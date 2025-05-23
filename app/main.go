package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
)

type page struct {
	Colour        string
	Location 	  string
	EnvVars       []string
	Headers       []string
	Host          string
	RequestURI    string
	ClientAddress string
	HostName      string
	Version       string
}

var (
	colour       = GetEnvWithDefault("COLOUR","blue")
	version      = GetEnvWithDefault("VERSION","0.0.1")
	location	 = GetEnvWithDefault("LOCATION","local")
	validColours = []string{"red", "green", "blue", "yellow"}
)

func GetEnvWithDefault(envVarName string, defaultValue string) string {
	v, e := os.LookupEnv(envVarName)
	if e && v !="" {
		return v
	}

	os.Setenv(envVarName, defaultValue)
	return defaultValue
}

func contains(a []string, b string) bool {
	for _, s := range a {
		if b == s {
			return true
		}
	}
	return false
}

func viewHandler(w http.ResponseWriter, r *http.Request) {

	envVars := make([]string, len(os.Environ()))
	headers := make([]string, len(r.Header))
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "localhost"
	}

	envVars = append(envVars, os.Environ()...)
	sort.Strings(headers)
	renderTemplate(w, colour, &page{Location: location, Version: version, Colour: colour, EnvVars: envVars, Headers: headers, ClientAddress: r.RemoteAddr, RequestURI: r.URL.RequestURI(), HostName: hostName})
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *page) {
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
	flag.Parse()
	// validate colour flag
	if contains(validColours, colour) {
		http.HandleFunc("/", viewHandler)
		http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))
		log.Fatal(http.ListenAndServe(":80", nil))
	} else {
		fmt.Fprintln(os.Stderr, "missing colour option! ('red', 'green', 'blue', 'yellow')")
		os.Exit(127)
	}
}
