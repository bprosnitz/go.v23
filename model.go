/*
Package v23 defines the Runtime interface of the public Vanadium API and its subdirectories define the entire Vanadium public API.

Once we reach a '1.0' version these public APIs will be stable over
an extended period and changes to them will be carefully managed to ensure backward compatibility. The same policy as used for go (http://golang.org/doc/go1compat) will be used for them.

The current release is 0.1 and although we will do our best to maintain
backwards compatibility we can't guarantee that until we reach the 1.0 milestone.

*/
package v23

import (
	"bytes"
	"fmt"
	"runtime"
	"sync"

	"v.io/v23/context"
	"v.io/v23/naming"
	"v.io/v23/naming/ns"
	"v.io/v23/rpc"
	"v.io/v23/security"
)

const (
	// LocalStop is the message received on WaitForStop when the stop was
	// initiated by the process itself.
	LocalStop = "localstop"
	// RemoteStop is the message received on WaitForStop when the stop was
	// initiated via an RPC call (AppCycle.Stop).
	RemoteStop            = "remotestop"
	UnhandledStopExitCode = 1
	ForceStopExitCode     = 1
)

// Task is streamed to channels registered using TrackTask to provide a sense of
// the progress of the application's shutdown sequence.  For a description of
// the fields, see the Task struct in the v23/services/mgmt/appcycle
// package, which it mirrors.
type Task struct {
	Progress, Goal int32
}

// AppCycle is the interface for managing the shutdown of a runtime
// remotely and locally. An appropriate instance of this is provided by
// the Profile to the runtime implementation which in turn arranges to
// serve it on an appropriate network address.
type AppCycle interface {
	// Stop causes all the channels returned by WaitForStop to return the
	// LocalStop message, to give the application a chance to shut down.
	// Stop does not block.  If any of the channels are not receiving,
	// the message is not sent on them.
	// If WaitForStop had never been called, Stop acts like ForceStop.
	Stop()

	// ForceStop causes the application to exit immediately with an error
	// code.
	ForceStop()

	// WaitForStop takes in a channel on which a stop event will be
	// conveyed.  The stop event is represented by a string identifying the
	// source of the event.  For example, when Stop is called locally, the
	// LocalStop message will be received on the channel.  If the channel is
	// not being received on, or is full, no message is sent on it.
	//
	// The channel is assumed to remain open while messages could be sent on
	// it.  The channel will be automatically closed during the call to
	// Cleanup.
	WaitForStop(chan<- string)

	// AdvanceGoal extends the goal value in the shutdown task tracker.
	// Non-positive delta is ignored.
	AdvanceGoal(delta int32)
	// AdvanceProgress advances the progress value in the shutdown task
	// tracker.  Non-positive delta is ignored.
	AdvanceProgress(delta int32)
	// TrackTask registers a channel to receive task updates (a Task will be
	// sent on the channel if either the goal or progress values of the
	// task have changed).  If the channel is not being received on, or is
	// full, no Task is sent on it.
	//
	// The channel is assumed to remain open while Tasks could be sent on
	// it.
	TrackTask(chan<- Task)

	// Remote returns an object to serve the remotely accessible AppCycle
	// interface (as defined in v23/services/mgmt/appcycle)
	Remote() interface{}
}

// Runtime is the interface that concrete Vanadium implementations must
// implement.  It will not be used directly by application builders.
// They will instead use the package level functions that mirror these
// factories.
// TODO(mattr): Some opts might not make sense now.  For example, the
// namespace is currently a ServerOpt, but it probably makes more sense
// to just use the current namespace in the context that is passed
// to NewServer.  The same for Profile and StreamManager.
type Runtime interface {

	// Init is a chance to initialize state in the runtime implementation
	// after the runtime has been registered in the v23 package.
	// Code that runs in this routine, unlike the code in the runtimes
	// constructor, can use the v23.Get/Set methods.
	Init(ctx *context.T) error

	// NewEndpoint returns an Endpoint by parsing the supplied endpoint
	// string as per the format described above. It can be used to test
	// a string to see if it's in valid endpoint format.
	//
	// NewEndpoint will accept strings both in the @ format described
	// above and in internet host:port format.
	//
	// All implementations of NewEndpoint should provide appropriate
	// defaults for any endpoint subfields not explicitly provided as
	// follows:
	// - a missing protocol will default to a protocol appropriate for the
	//   implementation hosting NewEndpoint
	// - a missing host:port will default to :0 - i.e. any port on all
	//   interfaces
	// - a missing routing id should default to the null routing id
	// - a missing codec version should default to AnyCodec
	// - a missing RPC version should default to the highest version
	//   supported by the runtime implementation hosting NewEndpoint
	NewEndpoint(ep string) (naming.Endpoint, error)

	// NewServer creates a new Server instance.
	//
	// It accepts at least the following options:
	// ServesMountTable and ServerBlessings.
	NewServer(ctx *context.T, opts ...rpc.ServerOpt) (rpc.Server, error)

	// SetNewStreamManager creates a new stream manager and context
	// with that StreamManager attached.
	SetNewStreamManager(ctx *context.T) (*context.T, error)

	// SetPrincipal attaches a principal to the returned context.
	SetPrincipal(ctx *context.T, principal security.Principal) (*context.T, error)

	// GetPrincipal returns the current Principal.
	GetPrincipal(ctx *context.T) security.Principal

	// SetNewClient creates a new Client instance and attaches it to a
	// new context.
	SetNewClient(ctx *context.T, opts ...rpc.ClientOpt) (*context.T, rpc.Client, error)

	// GetClient returns the current Client.
	GetClient(ctx *context.T) rpc.Client

	// SetNewNamespace creates a new Namespace and attaches it to the
	// returned context.
	SetNewNamespace(ctx *context.T, roots ...string) (*context.T, ns.Namespace, error)

	// GetNamespace returns the current namespace
	GetNamespace(ctx *context.T) ns.Namespace

	// GetAppCycle gets the current AppCycle.
	GetAppCycle(ctx *context.T) AppCycle

	// GetListenSpec gets the ListenSpec.
	GetListenSpec(ctx *context.T) rpc.ListenSpec

	// SetBackgroundContext creates a new context derived from the given context
	// with the given context set as the background context.
	SetBackgroundContext(ctx *context.T) *context.T

	// BackgroundContext retrieves a background context.  This context can
	// be used for general background activities.
	GetBackgroundContext(ctx *context.T) *context.T
}

// NewEndpoint returns an Endpoint by parsing the supplied endpoint
// string as per the format described above. It can be used to test
// a string to see if it's in valid endpoint format.
//
// NewEndpoint will accept strings both in the @ format described
// above and in internet host:port format.
//
// All implementations of NewEndpoint should provide appropriate
// defaults for any endpoint subfields not explicitly provided as
// follows:
// - a missing protocol will default to a protocol appropriate for the
//   implementation hosting NewEndpoint
// - a missing host:port will default to :0 - i.e. any port on all
//   interfaces
// - a missing routing id should default to the null routing id
// - a missing codec version should default to AnyCodec
// - a missing RPC version should default to the highest version
//   supported by the runtime implementation hosting NewEndpoint
func NewEndpoint(ep string) (naming.Endpoint, error) {
	return initState.currentRuntime().NewEndpoint(ep)
}

// NewServer creates a new Server instance.
//
// It accepts at least the following options:
// ServesMountTable and ServerBlessings.
//
// ServerBlessings defaults to v23.GetPrincipal(ctx).BlessingStore().Default().
// These Blessings are the server's Blessings for its lifetime.
func NewServer(ctx *context.T, opts ...rpc.ServerOpt) (rpc.Server, error) {
	return initState.currentRuntime().NewServer(ctx, opts...)
}

// SetNewStreamManager creates a new stream manager and context
// with that StreamManager attached.
func SetNewStreamManager(ctx *context.T) (*context.T, error) {
	return initState.currentRuntime().SetNewStreamManager(ctx)
}

// SetPrincipal attaches a principal to the returned context.
func SetPrincipal(ctx *context.T, principal security.Principal) (*context.T, error) {
	return initState.currentRuntime().SetPrincipal(ctx, principal)
}

// GetPrincipal returns the current Principal.
func GetPrincipal(ctx *context.T) security.Principal {
	return initState.currentRuntime().GetPrincipal(ctx)
}

// SetNewClient creates a new Client instance and attaches it to a
// new context.
func SetNewClient(ctx *context.T, opts ...rpc.ClientOpt) (*context.T, rpc.Client, error) {
	return initState.currentRuntime().SetNewClient(ctx, opts...)
}

// GetClient returns the current Client.
func GetClient(ctx *context.T) rpc.Client {
	return initState.currentRuntime().GetClient(ctx)
}

// SetNewNamespace creates a new Namespace and attaches it to the
// returned context.
func SetNewNamespace(ctx *context.T, roots ...string) (*context.T, ns.Namespace, error) {
	return initState.currentRuntime().SetNewNamespace(ctx, roots...)
}

// GetNamespace returns the current namespace.
func GetNamespace(ctx *context.T) ns.Namespace {
	return initState.currentRuntime().GetNamespace(ctx)
}

// GetAppCycle gets the current AppCycle.
func GetAppCycle(ctx *context.T) AppCycle {
	return initState.currentRuntime().GetAppCycle(ctx)
}

// GetListenSpec gets the current ListenSpec.
func GetListenSpec(ctx *context.T) rpc.ListenSpec {
	return initState.currentRuntime().GetListenSpec(ctx)
}

// SetBackgroundContext creates a new context derived from the given context
// with the given context set as the background context.
func SetBackgroundContext(ctx *context.T) *context.T {
	return initState.runtime.SetBackgroundContext(ctx)
}

// BackgroundContext retrieves a background context.  This context can
// be used for general background activities.
func GetBackgroundContext(ctx *context.T) *context.T {
	return initState.runtime.GetBackgroundContext(ctx)
}

var initState = &initStateData{}

type initStateData struct {
	mu           sync.RWMutex
	runtime      Runtime
	runtimeStack string
	profile      Profile
	profileStack string
}

func (i *initStateData) currentRuntime() Runtime {
	i.mu.RLock()
	defer i.mu.RUnlock()

	if i.runtimeStack == "" {
		panic(`Calling v23 method before initializing the runtime with Init().
You should call Init from your main or test function before calling
other v23 operations.`)
	}
	if i.runtime == nil {
		panic(`Calling v23 method during runtime initialization.  You cannot
call v23 methods until after the runtime has been constructed.  You may
be able to move the offending caller to the Runtime.Init() method of your
runtime implementation.`)
	}

	return i.runtime
}

// A profile represents the combination of hardware, operating system,
// compiler and libraries available to the application. The Profile
// creates a runtime implementation with the required hardware, operating system
// and library specific dependencies included.
//
// The implementations of the Profile are intended to capture all of
// the dependencies implied by that profile. For example, if a Profile requires
// a particular hardware specific library (say Bluetooth support), then the
// implementation of the Profile should include that dependency and
// the resulting runtime instance; the package implementing
// the Profile should expose the additional APIs needed to use the
// functionality.
//
// Profiles range from the generic to the very specific (e.g. "linux" or
// "my-sprinkler-controller-v2". Applications should, in general, use
// as generic a Profile as possbile.
//
// Profiles are registered using v23.RegisterProfileInit and
// subsequent registrations will panic. Packages that implement profiles will
// typically call RegisterProfileInit in their init functions so importing a
// profile will be sufficient to register it. Only one profile can be registered
// in any program, and subsequent registrations will panic.  Typically a
// program's main package will be the only place to import a profile.
//
// This scheme allows applications to use a pre-supplied Profile as well
// as for developers to create their own Profiles (to represent their
// hardware and software system). The Vanadium Build System, once fully
// developed will likely insert generated that uses one of the above schemes
// to configure profiles.
//
// At a minimum a Profile must do the following:
//   - Initialize a Runtime implementation (providing the flags to it)
//   - Return a Runtime implemenation, initial context, Shutdown func.
//
// See the v.io/x/ref/profiles package for a complete description of the
// precanned Profiles and how to use them.
type Profile func(ctx *context.T) (Runtime, *context.T, Shutdown, error)

// RegisterProfileInit register the specified Profile.
// It must be called before v23.Init; typically it will be called by an init
// function. It will panic if called more than once.
func RegisterProfileInit(f Profile) {
	// Skip 3 frames: runtime.Callers, getStack, RegisterProfileInit.
	stack := getStack(3)
	initState.mu.Lock()
	defer initState.mu.Unlock()
	if initState.profile != nil {
		format := `A profile has already been registered.
This is most likely because a library package is
importing a profile.  Look for imports of the form
'v.io/x/ref/profiles/...' and remove them.  Profiles should only be
imported in your main package.  Previous registration was from:
%s
Current registration is from:
%s
`
		panic(fmt.Sprintf(format, initState.profileStack, stack))
	}
	initState.profile = f
	initState.profileStack = stack
}

type Shutdown func()

func getStack(skip int) string {
	var buf bytes.Buffer
	stack := make([]uintptr, 16)
	stack = stack[:runtime.Callers(skip, stack)]
	for _, pc := range stack {
		fnc := runtime.FuncForPC(pc)
		file, line := fnc.FileLine(pc)
		fmt.Fprintf(&buf, "%s:%d: %s\n", file, line, fnc.Name())
	}
	return buf.String()
}

// Init should be called once for each vanadium executable, providing
// the setup of the initial context.T and a Shutdown function that can
// be used to clean up the runtime.  We allow calling Init multiple
// times (useful in tests), but only as long as you call the Shutdown
// returned previously before calling Init the second time.
func Init() (*context.T, Shutdown) {
	initState.mu.Lock()
	profile := initState.profile
	if initState.profile == nil {
		initState.mu.Unlock()
		panic("No profile has been registered nor specified. This is most" +
			" likely because your main package has not imported a profile")
	}

	// Skip 3 stack frames: runtime.Callers, getStack, Init
	stack := getStack(3)
	if initState.runtimeStack != "" {
		initState.mu.Unlock()
		format := `A runtime has already been initialized."
The previous initialization was from:
%s
This registration is from:
%s
`
		panic(fmt.Sprintf(format, initState.runtimeStack, stack))
	}
	initState.runtimeStack = stack
	initState.mu.Unlock()

	rootctx, rootcancel := context.RootContext()
	// Note we derive a second cancelable context here beyond the
	// rootctx.  This allows us to do shutdown in two steps.  First
	// we cancel this initial context to trigger cleanup of all
	// servers, clients, stream managers, etc.  Then after everything
	// is shut down we invoke rootcancel.  This allows the cleanup
	// to perform operations that require uncancelled contexts.
	ctx, cancel := context.WithCancel(rootctx)
	rt, ctx, shutdown, err := profile(ctx)
	if err != nil {
		cancel()
		rootcancel()
		panic(err)
	}

	initState.mu.Lock()
	initState.runtime = rt
	initState.mu.Unlock()

	vshutdown := func() {
		// Note we call our own cancel here to ensure that the
		// runtime/profile implementor has not attached anything to a
		// non-cancellable context.
		cancel()
		shutdown()
		rootcancel()

		initState.mu.Lock()
		initState.runtime = nil
		initState.runtimeStack = ""
		initState.mu.Unlock()
	}

	if err := rt.Init(ctx); err != nil {
		vshutdown()
		panic(err)
	}

	return ctx, vshutdown
}
