package app

import "net/http"

func Start() {
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	http.HandleFunc("/customers", GetAllCustomers)
	http.ListenAndServe(":8080", nil)
}
