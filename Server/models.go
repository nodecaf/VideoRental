package main

type tape struct {
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Available bool   `json:"available"`
}

var shelf = []*tape{
	{"Hot Fuzz", 2007, true},
	{"Matrix", 1998, false},
	{"Minions", 2014, true},
}

type shop struct{}

type clerk interface {
	list() []*tape
}

func (s shop) list() []*tape {
	return shelf
}
