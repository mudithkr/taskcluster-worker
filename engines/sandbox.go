package engines

import "io"

// The Shell interface is not fully specified yet, but the idea is that it pipes
// data to a shell inside a Sandbox.
type Shell interface {
	StdinPipe() (io.WriteCloser, error)
	StdoutPipe() (io.ReadCloser, error)
	StderrPipe() (io.ReadCloser, error)
	Terminate() error
	Wait() (bool, error)
}

// The Sandbox interface represents an active sandbox.
//
// All methods on this interface must be thread-safe.
type Sandbox interface {
	// Wait for task execution and termination of all associated shells, and
	// return immediately if sandbox execution has finished.
	//
	// When this method returns all resources held by the Sandbox instance must
	// have been released or transferred to the ResultSet instance returned. If an
	// internal error occured resources may be freed and WaitForResult() may
	// return ErrNonFatalInternalError if the error didn't leak resources and we
	// don't expect the error to be persistent.
	//
	// When this method has returned any calls to Abort() or NewShell() should
	// return ErrSandboxTerminated. If Abort() is called before WaitForResult()
	// returns, WaitForResult() should return ErrSandboxAborted and release all
	// resources held.
	//
	// Notice that this method may be invoked more than once. In all cases it
	// should return the same value when it decides to return. In particular, it
	// must keep a reference to the ResultSet instance created and return the same
	// instance, so that any resources held aren't transferred to multiple
	// different ResultSet instances.
	//
	// Non-fatal errors: ErrNonFatalInternalError, ErrSandboxAborted.
	WaitForResult() (ResultSet, error)

	// NewShell creates a new Shell for interaction with the sandbox. This is
	// useful for debugging and other purposes.
	//
	// If the engine doesn't support interactive shells it may return
	// ErrFeatureNotSupported. This should not interrupt/abort the execution of
	// the task which should proceed as normal.
	//
	// If the WaitForResult() method has returned and the sandbox isn't running
	// anymore this method must return ErrSandboxTerminated, signaling that you
	// can't interact with the sandbox anymore.
	//
	// Non-fatal errors: ErrFeatureNotSupported, ErrSandboxTerminated.
	NewShell() (Shell, error)

	// Abort the sandbox, this means killing the task execution as well as all
	// associated shells and releasing all resources held.
	//
	// If called before the sandbox execution finished, then WaitForResult() must
	// return ErrSandboxAborted. If sandbox execution has finished when Abort() is
	// called Abort() should return ErrSandboxTerminated and not release any
	// resources as they should have been released by WaitForResult() or
	// transferred to the ResultSet instance returned.
	//
	// Non-fatal errors: ErrSandboxTerminated
	Abort() error
}