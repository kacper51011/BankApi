package app

import (
	"encoding/xml"
	"net/http"
)

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "John Doe", City: "New York", ZipCode: "10001"},
		{Name: "Jane Smith", City: "Los Angeles", ZipCode: "90001"},
	}

	w.Header().Set("Content-Type", "application/xml")

	response := Customers{
		Customers: customers,
	}

	xml.NewEncoder(w).Encode(response)
}
