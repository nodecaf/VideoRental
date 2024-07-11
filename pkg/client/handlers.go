package client

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

type tape struct {
	Title     string `json:"title"`
	Year      int    `json:"year"`
	Available bool   `json:"available"`
}
type clientClerk struct {
	server string
}

func (c clientClerk) tapeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//Get tape name from form
	tn := r.FormValue("tape")
	sb := r.FormValue("submit")
	var mt = tape{tn, 2000, true} //we only look at the name in the api for now
	je, _ := json.Marshal(mt)
	http.Post("http://"+c.server+":8080/"+sb, "application/json", bytes.NewBuffer(je))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func (c clientClerk) mainHandler(w http.ResponseWriter, r *http.Request) {
	br, err := http.Get("http://" + c.server + ":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	in, _ := io.ReadAll(br.Body)
	var pj []tape
	err = json.Unmarshal(in, &pj)
	if err != nil {
		log.Fatal(err.Error())
	}
	//Create template
	t, err := template.New("list").ParseGlob("pkg/client/templates/*")
	if err != nil {
		log.Fatal((err.Error()))
	}
	w.WriteHeader(http.StatusOK)
	t.ExecuteTemplate(w, "index.gohtml", pj)
	w.Write([]byte("nodecaf was here"))
	//w.Write(in)
	//io.Copy(w, br.Body)
}
