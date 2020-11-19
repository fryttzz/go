package main

import (
	"fryttzz/go-alura/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8080", nil)
}
