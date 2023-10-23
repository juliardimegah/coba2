package main

import (
	"Proyek3/controller"
	"Proyek3/login"
	"Proyek3/signup"
	"net/http"
)

func main() {
	controller.Auth()
	// Menghubungkan rute HTTP dari package login
	// Mendaftarkan rute HTTP dari package login
	login.RegisterLoginRoutes()
	// Mendaftarkan rute HTTP dari package signup
	http.HandleFunc("/signup", signup.SignupHandler)

	// Melayani form login
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "templates/login.html")
		} else {
			http.Error(w, "Metode tidak diizinkancoy", http.StatusMethodNotAllowed)
		}
	})

	// Mulai server HTTP
	http.ListenAndServe(":9000", nil)
}
