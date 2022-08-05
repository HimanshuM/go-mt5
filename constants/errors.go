package constants

import "errors"

// ErrUnauthorized = Return code 8
var ErrUnauthorized = errors.New("not enough permissions")

// ErrNoService = Return code 11
var ErrNoService = errors.New("service unavailable")

// ErrNotFound = Return code 13
var ErrNotFound = errors.New("not found")

// ErrPartial = Return code 14
var ErrPartial = errors.New("partial error")

// ErrShutdown = Return code 15
var ErrShutdown = errors.New("shutdown in progress")

// ErrCanceled = Return code 16
var ErrCanceled = errors.New("operation canceled")
