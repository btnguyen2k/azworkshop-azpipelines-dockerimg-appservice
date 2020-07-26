package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")

		respHtml := fmt.Sprintf(`Hello, %q<br>Now is: %s`, html.EscapeString(r.URL.Path), time.Now().Format(time.RFC3339))

		respHtml += "Env:<br>"
		respHtml += "<pre>"
		for _, e := range os.Environ() {
			respHtml += e
		}
		respHtml += "</pre>"

		fmt.Fprintf(w, respHtml)
	})

	listenPort := 8080
	log.Printf("Running application on port %d\n", listenPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", listenPort), nil))
}
