package main

type SurferError struct {
	What string
}

func (error SurferError) Error() string {
	return error.What
}

func newSurferError(message string) SurferError {
	return SurferError{message}
}
