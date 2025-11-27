package virtualbox

import "errors"

var (
	ErrVMNotFound      = errors.New("vm not found")
	ErrVMAlreadyLocked = errors.New("vm is already locked by a session")
	ErrVMNotRunning    = errors.New("vm is not running")
	ErrUnknownError    = errors.New("unknown error")
)
