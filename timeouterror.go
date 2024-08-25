package main

type TimeoutError struct {
}

func (t *TimeoutError) Error() string {
	return "*"
}
