package main
import (
    "fmt"
    "log"
    "net/http"
	
	"github.com/abdennour/go-to-do-app/router"
)
func main() {
    r := router.Router()
    fmt.Println("server starting on 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}