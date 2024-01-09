package core

func UIntPtr(v int) *uint {
	c := uint(v)
	return &c
}
