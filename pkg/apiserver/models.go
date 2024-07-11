package apiserver

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
	borrowit(*tape) bool
	returnit(*tape) bool
}

func (s shop) list() []*tape {
	return shelf
}

func (s shop) borrowit(btape *tape) bool {
	for _, stape := range shelf {
		if (*stape).Title == (*btape).Title && (*stape).Available {
			(*stape).Available = false
			return true
		}
	}
	return false
}
func (s shop) returnit(btape *tape) bool {
	for _, stape := range shelf {
		if (*stape).Title == (*btape).Title && (*stape).Available {
			(*stape).Available = true
			return true
		}
	}
	return false
}
