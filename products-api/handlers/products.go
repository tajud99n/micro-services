package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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
	case http.MethodPost:
		p.addProduct(rw, r)
		break
	case http.MethodPut:
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.updateProduct(id, rw, r)
		break
	default:
		rw.WriteHeader(http.StatusNotFound)
	}
}

func (p *Product) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("get")
	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "error marshalling products", http.StatusInternalServerError)
		return
	}
}

func (p *Product) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("post")
	np := &data.Product{}

	err := np.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "error unmarshalling product", http.StatusBadRequest)
		return
	}
	data.AddProduct(np)
}

func (p *Product) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("put")
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "error unmarshalling product", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "server error", http.StatusInternalServerError)
		return
	}
}
