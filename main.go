package main

import (
	"encoding/json"
	"net/http"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func main() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	http.HandleFunc("/customers", GetAllCustomers)
	http.ListenAndServe(":8080", nil)
}
func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "John Doe", City: "New York", ZipCode: "10001"},
		{Name: "Jane Smith", City: "Los Angeles", ZipCode: "90001"},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
