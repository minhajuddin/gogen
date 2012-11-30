package main

import (
	"github.com/russross/blackfriday"
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

func render(page string) {
	log.Println("Rendering", page)
	content, _ := ioutil.ReadFile(page)
	content = blackfriday.MarkdownBasic(content)
	basename := path.Base(page)
	filename := basename[:strings.LastIndex(basename, ".")]
	opFile, err := os.Create(path.Join(opDir, filename+".html"))
	if err != nil {
		log.Panic(err)
	}
	defer opFile.Close()
	templates.ExecuteTemplate(opFile, "app.html", Page{Content: template.HTML(content)})
}

func main() {
	//runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("Generating site")

	pages, _ := filepath.Glob(path.Join(currentDir, "pages", "*md"))
	log.Println("Rendering", pages)
	os.Mkdir(opDir, 0700)

	for _, page := range pages {
		render(page)
	}

	log.Println("DONE")
}

type Page struct {
	Content template.HTML
}

func init() {
	templates = template.Must(template.ParseGlob(path.Join(currentDir, "layouts", "*html")))
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	opDir = path.Join(currentDir, "_site")
}
