// https://go.dev/doc/articles/wiki/
package main

import (
	"os"
	"log"
	"fmt"
    "net/http"
	"io/ioutil"
	"html/template"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func getText(country string, quantity int) string {
	response, err := http.Get("https://gutendex.com/books?languages="+country)

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }
	// var responseData map[string]string
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(responseData))
	return "nothing"
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    // title := r.URL.Path[len("/edit/"):]
    // p, err := loadPage(title)
    // if err != nil {
    //     p = &Page{Title: title}
    // }
    t, _ := template.ParseFiles("./data/pt_600_translated.html")
    t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	getText("pt", 2)
    log.Fatal(http.ListenAndServe(":8080", nil))
}