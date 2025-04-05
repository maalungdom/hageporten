
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "tenar/handlers"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "templates/index.html")
    })
    http.HandleFunc("/admin", handlers.AdminHandler)
    http.Handle("/opplasting", handlers.NewUploadHandler("uploads"))

    certExists := fileExists("cert.pem") && fileExists("key.pem")

    if certExists {
        fmt.Println("🔐 HTTPS-tenar startar på https://0.0.0.0:8443")
        go func() {
            log.Fatal(http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil))
        }()
    } else {
        fmt.Println("⚠️  Fant ikkje cert.pem / key.pem – hoppar over HTTPS")
    }

    fmt.Println("🌐 HTTP-tenar startar på http://0.0.0.0:8080")
    log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

func fileExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}
