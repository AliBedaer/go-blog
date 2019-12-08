package main

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {
	validate = validator.New()
	router := routes()

	http.ListenAndServe(":8000", router)

}
