package web

import (
	"fmt"
	"net/http"
)

func StartServer() {
	http.Handle("/", http.FileServer(http.Dir("go/web/static")))
	fmt.Println("🌐 Web dashboard running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
