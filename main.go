package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	templates  *template.Template
	currentDir string
	opDir      string
)

func main() {
	log.Println("Generating site")

	pages, _ := filepath.Glob(path.Join(currentDir, "pages", "*txt"))
	os.Mkdir(opDir, 0700)

	for _, page := range pages {
		log.Println("Rendering", page)
		content, _ := ioutil.ReadFile(page)
		basename := path.Base(page)
		filename := basename[:strings.LastIndex(basename, ".")]
		opFile, err := os.Create(path.Join(opDir, filename+".html"))
		if err != nil {
			log.Panic(err)
		}
		defer opFile.Close()
		templates.ExecuteTemplate(opFile, "app.html", Page{Content: string(content)})
	}

	log.Println("DONE")
}

type Page struct {
	Content string
}

func init() {
	templates = template.Must(template.ParseGlob(path.Join(currentDir, "layouts", "*html")))
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	opDir = path.Join(currentDir, "_site")
}
