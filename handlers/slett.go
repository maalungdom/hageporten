package handlers

import (
	"net/http"
	"os"
)

func Slett(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Berre POST er lov", http.StatusMethodNotAllowed)
		return
	}
	fil := r.FormValue("fil")
	if fil == "" {
		http.Error(w, "Manglar filnamn", http.StatusBadRequest)
		return
	}
	err := os.Remove("uploads/" + fil)
	if err != nil {
		http.Error(w, "Klarte ikkje sletta "+fil, http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/filer", http.StatusSeeOther)
}
