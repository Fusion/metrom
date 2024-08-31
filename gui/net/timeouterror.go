package net

type TimeoutError struct {
}

func (t *TimeoutError) Error() string {
	return "*"
}
