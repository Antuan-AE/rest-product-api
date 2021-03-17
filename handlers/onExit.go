package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

//OnExit type
type OnExit struct {
	l *log.Logger
}

//NewOnExit function to handle specific logging type
func NewOnExit(l *log.Logger) *OnExit {
	return &OnExit{l}
}

func (h *OnExit) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Ooops", http.StatusBadRequest)
		return
	}
	rw.Write([]byte("Exiting..." + string(d)))
	h.l.Println("Exiting REST API")
}
