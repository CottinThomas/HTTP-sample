package myserver

import (
	//"fmt"
	"net/http"
	"sync"
	"text/template"
)

type MyHandler struct {
	sync.Mutex
	count int
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var count int

	h.Lock()
	h.count++
	count = h.count
	h.Unlock()
	t, err := template.ParseFiles("./template.html")
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(w, "template.html", count)
	if err != nil {
		panic(err)
	}
	//fmt.Fprintf(w, "Visitor count: %d.", count)
}