package services

type lifecycle int

const (
	Transient lifecycle = iota
	Singleto
	Scoped
)
