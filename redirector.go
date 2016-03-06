package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
)

type redirector struct {
	t *template.Template
}

func makeRedirector(templateStr string) (*redirector, error) {
	t, err := template.New("destination").Parse(templateStr)
	if err != nil {
		return nil, err
	}

	return &redirector{t}, nil
}

func (d *redirector) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	d.t.Execute(buf, r)

	dest := buf.String()
	if verbose {
		log.Printf("R: %s -> %s", r.URL, dest)
	}

	http.Redirect(w, r, dest, http.StatusFound)
}
