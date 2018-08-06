package entity

import "errors"

// ErrNotFound not found
var ErrNotFound = errors.New("Not found")

// ErrCannotBeDeleted entity cannot be deleted
var ErrCannotBeDeleted = errors.New("Cannot Be Deleted")
