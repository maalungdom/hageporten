package handlers

import (
	"fmt"
	"net/http"
	"os"
)

func Filer(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir("uploads")
	if err != nil {
		http.Error(w, "Kunne ikkje lesa uploads/", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<html><head><meta charset="utf-8"><title>Filer</title></head><body><h1>📁 Opplasta filer</h1><ul>`)
	for _, f := range files {
		name := f.Name()
		fmt.Fprintf(w, `<li><a href="/uploads/%s">%s</a></li>`, name, name)
	}
	fmt.Fprint(w, `</ul><p><a href="/">← Tilbake</a></p></body></html>`)
}
