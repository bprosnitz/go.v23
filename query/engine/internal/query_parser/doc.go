// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package query_parser is a parser to parse a simplified select statement (a la SQL) for the
// Vanadium key value store (a.k.a., syncbase).
//
// The select is of the form:
//
// <query_specification> ::=
//   <select_clause> <from_clause> [<where_clause>] [<limit_offset_clause>]
//
// <select_clause> ::= SELECT <selector> [{<comma><selector>}...]
//
// <from_clause> ::= FROM <table>
//
// <where_clause> ::= WHERE <expression>
//
// <limit_offset_clause> ::=
//   LIMIT <int_literal> [OFFSET <int_literal>]
//   | OFFSET <int_literal> [LIMIT <int_literal>]
//
// <selector> ::= <column> [AS <string_literal>]
//
// <column> ::=
//   k
//   | v[<period><field>]
//   | <function>
//
// <field> ::= <segment>[{<period><segment>}...]
//
// <segment> ::= <identifier>[<keys>]
//
// <keys> ::= <key>...
//
// <key> ::= <left_bracket> <operand> <right_bracket>
//
// <function> ::= <identifier><left_paren>[<operand>[{<comma><operand>}...]<right_paren>
//
// <table> ::= <identifier>
//
// <expression> ::=
//   <left_paren> <expression> <right_paren>
//   | <logical_expression>
//   | <binary_expression>
//
// <logical_expression> ::=
//   <expression> <logical_op> <expression>
//
// <logical_op> ::=
//   AND
//   | OR
//
// <binary_expression> ::=
//   <operand> <binary_op> <operand>
//   | v[<period><field>] IS [NOT] NIL
//
// <operand> ::=
//   k
//   | v[<period><field>]
//   | <literal>
//   | <function>
//
// <binary_op> ::=
//   =
//   | EQUAL
//   | <>
//   | NOT EQUAL
//   | LIKE
//   | NOT LIKE
//   | <
//   | <=
//   | >=
//   | >
//
// <literal> ::=
//   <string_literal>
//   | <bool_literal>
//   | <int_literal>
//   | <float_literal>
//
// Example:
// select v.Foo.Far, v.Baz[2] from Foobarbaz where Type(v) like "%.Customer" and (v.Foo = 42 and v.Bar not like "abc%) or (k >= "100" and  k < "200")
package query_parser