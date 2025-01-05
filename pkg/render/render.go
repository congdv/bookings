package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/congdv/bookings/pkg/config"
	"github.com/congdv/bookings/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func AddDefaultData(templateData *models.TemplateData) *models.TemplateData {
	return templateData
}

func NewTemplate(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, templateData *models.TemplateData) {
	// create a template cache
	var templateCache map[string]*template.Template
	if app.UseCache {
		templateCache = app.TemplateCache

	} else {
		templateCache, _ = CreateTemplateCache()

	}

	// fmt.Printf("+%v", templateCache)
	// get requested template from cache
	template, ok := templateCache[tmpl]

	if !ok {
		log.Fatal("Error when getting cache from template cache")
	}

	buf := new(bytes.Buffer)

	td := AddDefaultData(templateData)
	err := template.Execute(buf, td)

	if err != nil {
		log.Println(err)
	}
	// render the template

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	log.Println("Creating template cache")
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files name *.page.html
	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = templateSet
	}

	return myCache, nil
}

// Map to store template
// var templateCache = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, templateName string) {
// 	var tmpl *template.Template
// 	var err error

// 	// check to see if we already have the template in our cache
// 	_, inMap := templateCache[templateName]
// 	if !inMap {
// 		// need to create the template
// 		err = createTemplateCache(templateName)
// 		if err != nil {
// 			log.Println("error: ", err)
// 		}
// 	} else {
// 		// we have the template in the cache

// 		log.Println("Using cached template")
// 	}

// 	tmpl = templateCache[templateName]
// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println("error: ", err)
// 	}
// }

// func createTemplateCache(templateName string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s.html", templateName),
// 		"./templates/base.layout.html",
// 	}

// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	log.Println("Creating template cache", templateName)
// 	templateCache[templateName] = tmpl

// 	return nil
// }
