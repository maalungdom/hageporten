package main

import (
	"fmt"
	"net/http"

	"hageport/handlers" // NB: dette må matche modulnamn i go.mod
)

func main() {
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/opplasting", handlers.Upload)
	http.HandleFunc("/filer", handlers.Filer)
	http.HandleFunc("/slett", handlers.Slett)
	http.Handle("/uploads/", http.StripPrefix("/uploads/", http.FileServer(http.Dir("uploads"))))

	fmt.Println("🌿 Tenaren køyrer på http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("🚨 Feil ved oppstart:", err)
	}
}
