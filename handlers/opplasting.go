package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Handter opplasting
func Upload(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		file, header, err := r.FormFile("fil")
		if err != nil {
			http.Error(w, "Feil ved opplasting", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		os.MkdirAll("uploads", os.ModePerm)
		out, err := os.Create("uploads/" + header.Filename)
		if err != nil {
			http.Error(w, "Feil ved lagring", http.StatusInternalServerError)
			return
		}
		defer out.Close()
		io.Copy(out, file)

		fmt.Fprintf(w, "✅ Fil lasta opp: %s", header.Filename)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html lang="nn">
		<head><meta charset="UTF-8"><title>Opplasting</title></head>
		<body>
			<h1>🌿 Last opp ei fil</h1>
			<form method="POST" enctype="multipart/form-data">
				<input type="file" name="fil">
				<input type="submit" value="Last opp">
			</form>
		</body>
		</html>
	`)
}
