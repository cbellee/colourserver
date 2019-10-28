package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Colour         string
	EnvVars       []string
	Headers       []string
	Host          string
	RequestURI    string
	ClientAddress string
	HostName      string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	colour := r.URL.Path[len("/view/"):]
	envVars := make([]string, len(os.Environ()))
	headers := make([]string, len(r.Header))
	hostName, err := os.Hostname()
	if err!=nil {
		hostName = ""
	}

	for _, env := range os.Environ() {
		envVars = append(envVars, env)
	}

	for name, values := range r.Header {
		for _, value := range values {
			header := name + ": " + value
			headers = append(headers, header)
		}
	}

	renderTemplate(w, colour, &Page{Colour: colour, EnvVars: envVars, Headers: headers, ClientAddress: r.RemoteAddr, RequestURI: r.URL.RequestURI(), HostName: hostName})
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles("html/" + "main.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
