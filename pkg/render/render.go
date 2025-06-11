package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/jaysg1thub/bookings/pkg/config"
	"github.com/jaysg1thub/bookings/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td

}

// takes in a "RW" & "template" we want to parse & writes to Brwoser window:
// RenderTemplate renders templages using html/template;
// Note:	we change func name to UPPERCASE "R" so func is now visible from outside pkg
func RenderTemplate(w http.ResponseWriter, html string, td *models.TemplateData) {

	var tc map[string]*template.Template

	if app.UseCache {

		// get the template cache from the app config
		tc = app.TemplateCache

	} else {
		tc, _ = CreateTemplateCache()
	}

	// // get the template cache from the app config
	// tc := app.TemplateCache

	// can now get rid of this block since "tc" gets it's value directly above
	//tc, err := CreateTemplateCache()
	//if err != nil {
	//	log.Fatal(err)
	//}

	// get requested template from cache
	t, ok := tc[html]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// look for any layouts that exist in that directory:
		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil

}
