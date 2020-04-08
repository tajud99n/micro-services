package handlers

import (
	"log"
	"net/http"

	"github.com/tajud99n/nic/products-api/data"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		p.getProducts(rw, r)
		break
	default:
		rw.WriteHeader(http.StatusNotFound)
	}
}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "error marshalling product", http.StatusInternalServerError)
		return
	}
}
