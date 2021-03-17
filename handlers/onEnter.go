package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//OnEnter type
type OnEnter struct {
	l *log.Logger
}

//NewOnEnter method
func NewOnEnter(l *log.Logger) *OnEnter {
	return &OnEnter{l}
}

func (h *OnEnter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Welcome to REST API")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "Hello %s", string(d))
}
