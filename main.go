package main

import (
	"log"
	"html/template"
	"os"
)

var (
	templates = template.Must(template.ParseGlob("./layouts/*html"))
)

func main(){
	log.Println("Generating site")
	templates.ExecuteTemplate(os.Stdout, "app.html", Page{Content: "Awesome page"})
	log.Println("DONE")
}

type Page struct {
	Content string
}
