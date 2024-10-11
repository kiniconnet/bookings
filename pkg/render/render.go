package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/kiniconnet/bookings/pkg/config"
	"github.com/kiniconnet/bookings/pkg/models"
)

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) error {
	app = a
	return nil
}

/*  RenderTemplate renders templates using html/template */
func RenderTemplate(w http.ResponseWriter, t string, td *models.TemplateData) {

	var tc map[string]*template.Template

	/* Get the template cache from the app config */
	if app.UseCache{
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplate()

	}
	
	templ, ok := tc[t]
	if !ok {
		log.Fatal("error due to not accepting the mistake")
	}

	err := templ.Execute(w, td)
	if err != nil {
		log.Fatal()
	}
}

func CreateTemplate() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.html")
	if err != nil {
		return myCache, err
	}
 
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		mathces, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(mathces) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
