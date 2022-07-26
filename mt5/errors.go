package mt5

import "errors"

// Return code 8
var ErrUnauthorized = errors.New("not enough permissions")

// Return code 11
var ErrNoService = errors.New("service unavailable")

// Return code 13
var ErrNotFound = errors.New("not found")

// Return code 14
var ErrPartial = errors.New("partial error")

// Return code 15
var ErrShutdown = errors.New("shutdown in progress")

// Return code 16
var ErrCanceled = errors.New("operation canceled")
