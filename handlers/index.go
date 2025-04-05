package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
		<!DOCTYPE html>
		<html lang="nn">
		<head>
			<meta charset="UTF-8">
			<title>🌿 Hageporten</title>
		</head>
		<body>
			<h1>🌱 Velkomen til hageporten</h1>
			<p><a href="/opplasting">→ Last opp ei fil</a></p>
			<p><a href="/filer">→ Sjå alle filer</a></p>
		</body>
		</html>
	`)
}
