// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nosql_test

import (
	"testing"

	"v.io/v23/context"
	wire "v.io/v23/services/syncbase/nosql"
	"v.io/v23/syncbase"
	"v.io/v23/syncbase/nosql"
	tu "v.io/v23/syncbase/testutil"
	"v.io/v23/verror"
)

// Tests schema checking logic within App.NoSQLDatabase() method.
// This test as following steps:
// 1) Call NoSQLDatabase() for a non existent db.
// 2) Create the database, and verify if Schema got stored properly.
// 3) Call EnforceSchema() to make sure that the method is no-op and is
//    able to read the schema from db.
// 4) Call NoSQLDatabase() on the same db to create a new handle with an
//    upgraded schema, call EnforceSchema() and check if SchemaUpgrader
//    is called and if the new schema is stored appropriately.
func TestSchemaCheck(t *testing.T) {
	ctx, sName, cleanup := tu.SetupOrDie(nil)
	defer cleanup()
	a := tu.CreateApp(t, ctx, syncbase.NewService(sName), "a")
	schema := tu.DefaultSchema(0)
	mockUpgrader := schema.Upgrader.(*tu.MockSchemaUpgrader)

	db1 := a.NoSQLDatabase("db1", schema)

	// Verify that calling Upgrade on a non existing database does not throw
	// errors.
	err := db1.EnforceSchema(ctx)
	if err != nil {
		t.Fatalf("db1.EnforceSchema() failed: %v", err)
	}
	if mockUpgrader.CallCount > 0 {
		t.Fatal("Call to upgrader was not expected.")
	}

	// Create db1, this step also stores the schema provided above
	if err := db1.Create(ctx, nil); err != nil {
		t.Fatalf("db1.Create() failed: %v", err)
	}
	// verify if schema was stored as part of create
	if _, err := getSchemaMetadata(ctx, db1.FullName()); err != nil {
		t.Fatalf("Failed to lookup schema after create: %v", err)
	}

	// Make redundant call to Upgrade to verify that it is a no-op
	if err := db1.EnforceSchema(ctx); err != nil {
		t.Fatalf("db1.EnforceSchema() failed: %v", err)
	}
	if mockUpgrader.CallCount > 0 {
		t.Fatal("Call to upgrader was not expected.")
	}

	// try to make a new database object for the same database but this time
	// with a new schema version
	schema.Metadata.Version = 1
	rule := wire.CrRule{"table1", "foo", "", wire.ResolverTypeLastWins}
	policy := wire.CrPolicy{
		Rules: []wire.CrRule{rule},
	}
	schema.Metadata.Policy = policy
	otherdb1 := a.NoSQLDatabase("db1", schema)
	if err := otherdb1.EnforceSchema(ctx); err != nil {
		t.Fatalf("otherdb1.EnforceSchema() failed: %v", err)
	}
	if mockUpgrader.CallCount != 1 {
		t.Fatalf("Unexpected number of calls to upgrader. Expected: %d, Actual: %d.", 1, mockUpgrader.CallCount)
	}

	// check if the contents of SchemaMetadata are correctly stored in the db.
	metadata, err3 := getSchemaMetadata(ctx, otherdb1.FullName())
	if err3 != nil {
		t.Fatalf("GetSchemaMetadata failed: %v", err3)
	}
	if metadata.Version != 1 {
		t.Fatalf("Unexpected version number: %d", metadata.Version)
	}
	if len(metadata.Policy.Rules) != 1 {
		t.Fatalf("Unexpected number of rules: %d", len(metadata.Policy.Rules))
	}
	if metadata.Policy.Rules[0] != rule {
		t.Fatalf("Unexpected number of rules: %d", len(metadata.Policy.Rules))
	}
}

func TestRPCSchemaCheckError(t *testing.T) {
	// Setup
	ctx, sName, cleanup := tu.SetupOrDie(nil)
	defer cleanup()
	a := tu.CreateApp(t, ctx, syncbase.NewService(sName), "a")
	schema := tu.DefaultSchema(0)

	// Create db1 with schema version 0 and add table1 and row1
	dbHandle1 := a.NoSQLDatabase("db1", schema)
	if err := dbHandle1.Create(ctx, nil); err != nil {
		t.Fatalf("db1.Create() failed: %v", err)
	}
	if err := dbHandle1.CreateTable(ctx, "table1", nil); err != nil {
		t.Fatalf("db1.CreateTable() failed: %v", err)
	}
	if err := dbHandle1.Table("table1").Put(ctx, "row1", "value1"); err != nil {
		t.Fatalf("table1.Put() failed: %v", err)
	}

	// Try writing to database db1 with a db handle with schema version 2
	schema2 := tu.DefaultSchema(2)
	dbHandle2 := a.NoSQLDatabase("db1", schema2)

	// verify write rpcs for Database
	if err := dbHandle2.CreateTable(ctx, "table1", nil); !isVersionMismatchErr(err) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(err))
	}
	if err := dbHandle2.DeleteTable(ctx, "table1"); !isVersionMismatchErr(err) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(err))
	}
	if err := dbHandle2.Destroy(ctx); !isVersionMismatchErr(err) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(err))
	}
	if _, err := dbHandle2.Exists(ctx); !isVersionMismatchErr(err) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(err))
	}
	if _, err := dbHandle2.BeginBatch(ctx, wire.BatchOptions{}); !isVersionMismatchErr(err) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(err))
	}
	if _, _, err := dbHandle2.Exec(ctx, ""); !isVersionMismatchErr(err) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(err))
	}

	// verify write rpcs for Table
	table := dbHandle2.Table("table1")
	if _, err := table.Exists(ctx); !isVersionMismatchErr(err) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(err))
	}
	if err := table.Delete(ctx, nosql.SingleRow("row1")); !isVersionMismatchErr(err) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(err))
	}
	stream := table.Scan(ctx, nosql.SingleRow("row1"))
	if stream.Advance() {
		t.Fatalf("Stream advanced unexpectedly")
	}
	if !isVersionMismatchErr(stream.Err()) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(stream.Err()))
	}
	if _, err := table.GetPermissions(ctx, "row1"); !isVersionMismatchErr(err) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(err))
	}
	if err := table.SetPermissions(ctx, nosql.Prefix("row"), nil); !isVersionMismatchErr(err) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(err))
	}
	if err := table.DeletePermissions(ctx, nosql.Prefix("row")); !isVersionMismatchErr(err) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(err))
	}

	// verify write rpcs for Row
	row := table.Row("row1")
	if _, err := row.Exists(ctx); !isVersionMismatchErr(err) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(err))
	}
	var str string
	if err := row.Get(ctx, &str); !isVersionMismatchErr(err) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(err))
	}
	if err := row.Put(ctx, "newValue"); !isVersionMismatchErr(err) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(err))
	}
	if err := row.Delete(ctx); !isVersionMismatchErr(err) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(err))
	}
}

func TestRPCSchemaCheckErrorForBatch(t *testing.T) {
	// Setup
	ctx, sName, cleanup := tu.SetupOrDie(nil)
	defer cleanup()
	a := tu.CreateApp(t, ctx, syncbase.NewService(sName), "a")
	schema := tu.DefaultSchema(0)

	// Create db1 with schema version 0 and add table1 and row1
	dbHandle1 := a.NoSQLDatabase("db1", schema)
	if err := dbHandle1.Create(ctx, nil); err != nil {
		t.Fatalf("db1.Create() failed: %v", err)
	}
	if err := dbHandle1.CreateTable(ctx, "table1", nil); err != nil {
		t.Fatalf("db1.CreateTable() failed: %v", err)
	}
	if err := dbHandle1.Table("table1").Put(ctx, "row1", "value1"); err != nil {
		t.Fatalf("table1.Put() failed: %v", err)
	}

	// Create three batches using dbHandle1
	batch1, batchErr1 := dbHandle1.BeginBatch(ctx, wire.BatchOptions{})
	if batchErr1 != nil {
		t.Fatalf("db1.BeginBatch() failed: %v", batchErr1)
	}
	batch1.Table("table1").Row("row1").Put(ctx, "newValue1")

	batch2, batchErr2 := dbHandle1.BeginBatch(ctx, wire.BatchOptions{})
	if batchErr2 != nil {
		t.Fatalf("db1.BeginBatch() failed: %v", batchErr2)
	}
	batch2.Table("table1").Row("row1").Put(ctx, "newValue2")

	batch3, batchErr3 := dbHandle1.BeginBatch(ctx, wire.BatchOptions{})
	if batchErr3 != nil {
		t.Fatalf("db1.BeginBatch() failed: %v", batchErr3)
	}

	// Upgrade schema version for underlying db using a different handle
	schema2 := tu.DefaultSchema(1)
	dbHandle2 := a.NoSQLDatabase("db1", schema2)
	dbHandle2.EnforceSchema(ctx)

	// Commit batch1, abort batch2, attempt writing a row using batch3.
	// Each of these operations should fail.
	if err := batch1.Commit(ctx); !isVersionMismatchErr(err) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(err))
	}
	if err := batch2.Abort(ctx); !isVersionMismatchErr(err) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(err))
	}
	if err := batch3.Table("table1").Row("row1").Put(ctx, "newValue3"); !isVersionMismatchErr(err) {
		t.Fatal("Expected ErrDatabaseVersionMismatch, found: " + toString(err))
	}

	// Verify that the value of row1 is the original value.
	var value string
	if err := dbHandle2.Table("table1").Get(ctx, "row1", &value); err != nil {
		t.Fatalf("table1.Get() failed: %v", err)
	}
}

func toString(err error) string {
	if err == nil {
		return "nil"
	}
	return string(verror.ErrorID(err)) + ": " + err.Error()
}

func isVersionMismatchErr(err error) bool {
	if err == nil {
		return false
	}
	return verror.ErrorID(err) == wire.ErrSchemaVersionMismatch.ID
}

func getSchemaMetadata(ctx *context.T, dbName string) (wire.SchemaMetadata, error) {
	return wire.DatabaseClient(dbName).GetSchemaMetadata(ctx)
}
