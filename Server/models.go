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
	borrow(*tape) bool
}

func (s shop) list() []*tape {
	return shelf
}

func (s shop) borrow(btape *tape) bool {
	for _, stape := range shelf {
		if (*stape).Title == (*btape).Title && (*stape).Available {
			(*stape).Available = false
			return true
		}
	}
	return false
}
