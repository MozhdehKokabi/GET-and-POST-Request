package main

import (
	"fmt"
	"log"
	"net/http"
)

func abc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "forms.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "parseForm() err: %v\n", err)
			return
		}

		fmt.Println(w, "post form website r.postform = %v\n", r.PostForm)
		name := r.FormValue("name")
		address := r.FormValue("address")

		fmt.Fprintf(w, "name=%s\n", name)
		fmt.Fprintf(w, "name=%s\n", address)

	default:
		fmt.Fprint(w, "only get and post")
	}
}
func main() {
	http.HandleFunc("/", abc)

	fmt.Printf("starting server got testing\n")
	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatal(err)

	}

}
