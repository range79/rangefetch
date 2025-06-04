package model

type selection struct {
	name     string
	cursor   int
	selected map[int]struct{}
}
