package model

type Item struct {
	Title     string
	Completed bool
}

type FilterState int

const (
	All FilterState = iota
	Active
	Completed
)
