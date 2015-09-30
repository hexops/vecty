package model

type FilterState int

const (
	All FilterState = iota
	Active
	Completed
)

type Item struct {
	Title     string
	Completed bool
}
