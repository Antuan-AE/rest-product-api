package main

import (
	"log"
	"net/http"
	"os"

	"github.com/tony96314/rest-product-api/handlers"
)

func main() {

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	log.Println("Hello World")
	// 	d, err := ioutil.ReadAll(r.Body)
	// 	if err != nil {
	// 		http.Error(rw, "Ooops", http.StatusBadRequest)
	// 		// rw.WriteHeader(http.StatusBadRequest)
	// 		// rw.Write([]byte("Ooops"))
	// 		return
	// 	}

	// 	fmt.Fprintf(rw, "Hello %s", string(d))
	// })

	// http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
	// 	log.Println("Goodbye World")
	// })

	l := log.New(os.Stdout, "[rest-product-api] ", log.LstdFlags)

	enteringHand := handlers.NewOnEnter(l)
	exitHand := handlers.NewOnExit(l)

	servMux := http.NewServeMux()
	servMux.Handle("/", enteringHand)
	servMux.Handle("/exit", exitHand)

	http.ListenAndServe(":1522", servMux)
}
