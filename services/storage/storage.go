package storage

import "errors"

// ErrNotFound returned in case when entity was not found.
var ErrNotFound = errors.New("entity was not found")

// ErrAlreadyExist returned in case when entity exists
var ErrAlreadyExist = errors.New("record already exists")
