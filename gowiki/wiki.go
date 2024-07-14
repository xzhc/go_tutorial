package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// page persistent storage(write)
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

// load pages(read)
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// viewHandler allow users to view a wiki page
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

// editHandler loads the page(or, if it dosen't exist, create an empty struct),
// and display an HTML form.
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w,
		"<h1>Editing %s</h1>"+
			"<form action=\"/save/%s\" method=\"POST\">"+
			"<textarea name=\"body\">%s</textarea><br>"+
			"<input type=\"submit\" value=\"Save\">"+
			"</form>",
		p.Title, p.Title, p.Body)
}
