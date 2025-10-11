package server

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
)

type todo struct {
	UserID    int    `json:"userID"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var form = `
<h1>Todo #{{.ID}}</h1>
<div>{{printf "User %d" .UserID}}</div>
<div>{{printf "%s (completed: %t)" .Title .Completed}}</div>`

func handler(w http.ResponseWriter, r *http.Request) {
	const base = "https://jsonplaceholder.typicode.com"

	resp, err := http.Get(base + os.Args[1])

	if err != nil {
		http.Error(w, err.Error(), http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	var item todo

	// instead of doing a io.ReadAll, getting a slice of bytes
	// then taking that byte slice and passing it to
	// json.Unmarshal, can simply pass the response Body
	// to New Decoder and then Decode into the designated place all in one line
	err = json.NewDecoder(resp.Body).Decode(&item)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.New("mine")
	tmpl.Parse(form)
	tmpl.Execute(w, item)

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
