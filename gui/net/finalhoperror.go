package net

type FinalHopError struct {
}

func (t *FinalHopError) Error() string {
	return "found final hop"
}
