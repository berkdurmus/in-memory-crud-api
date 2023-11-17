// main.go
package InMemoryCRUDAPI

import (
	"log"
	"net/http"
)

func main() {
	store := NewStore()
	http.HandleFunc("/create", store.createItemHandler)
	http.HandleFunc("/get", store.getItemHandler)
	http.HandleFunc("/update", store.updateItemHandler)
	http.HandleFunc("/delete", store.deleteItemHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
