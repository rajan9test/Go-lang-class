package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`this server is runnig`))
	})

	router.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(401)
        
		w.Write([]byte(`you are not authorised`))
		
	})

	s := &http.Server{
		Addr:              ":8080",
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	log.Println("the server runnning on the port 8080")

	err := s.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}
