package handle

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func Mainhandle(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL.Path)
	//fmt.Fprint(w, "this is main")
	if r.URL.Path != "/main" {
		http.NotFound(w, r)
		return
	}
	data1 := map[string]interface{}{
		"text1": "Greet",
		"text2": "bye-bye",
	}
	tmpl, err := template.ParseFiles(path.Join("packhtml", "main.html"), path.Join("packhtml", "layout.html"))
	if err != nil {
		log.Print(err)
		http.Error(w, "ERROR", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, data1)
	if err != nil {
		log.Print(err)
		http.Error(w, "ERROR", http.StatusInternalServerError)
		return
	}
}
func Homehandle(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	data := map[string]interface{}{
		"text1": "Halo, Hola, Hai",
		"text2": "gracias, xie-xie, danke",
	}

	tmpl, err := template.ParseFiles(path.Join("packhtml", "portofolio.html"), path.Join("packhtml", "layout.html"))
	if err != nil {
		log.Print(err)
		http.Error(w, "ERROR", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Print(err)
		http.Error(w, "ERROR", http.StatusInternalServerError)
		return
	}
}
