package handlers

import (
	"hello_world/data"
	"hello_world/utils"
	"log"
	"net/http"
)

var requestUtil utils.RequestUtil

type Product struct {
	l *log.Logger
}

//get new product object
func NewProduct(l *log.Logger) *Product {
	requestUtil = *utils.NewRequestUtil(l)
	return &Product{l: l}
}

//method required for interface http.ServeMux.Handler
func (p *Product) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProduct(rw, r)
		return
	}
	if r.Method == http.MethodPut {
		p.putProduct(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.postProduct(rw, r)
		return
	}
	http.Error(rw, "METHOD NOT ALLOWED", http.StatusBadRequest)
}

func (p *Product) postProduct(rw http.ResponseWriter, r *http.Request) {
	id, err := requestUtil.GetId(r)
	if err != nil {
		http.Error(rw, "INVALID URI", http.StatusBadRequest)
		return
	}
	_, _, err = data.FindProduct(id)
	if err == nil {
		http.Error(rw, "RESOURCE ALREADY EXISTS", http.StatusConflict)
		return
	}
	prod := &data.Product{}
	err = prod.FromJSON(r.Body)
	if err != nil {
		p.l.Fatal("Error ocurred while trying to convert json body to object")
		http.Error(rw, "INVALID URI", http.StatusBadRequest)
		return
	}
	data.AddProduct(prod)
}

func (p *Product) putProduct(rw http.ResponseWriter, r *http.Request) {
	id, err := requestUtil.GetId(r)
	if err != nil {
		http.Error(rw, "INVALID URI", http.StatusBadRequest)
		return
	}
	prod := &data.Product{}
	err = prod.FromJSON(r.Body)
	if err != nil {
		p.l.Fatal("Error ocurred while trying to convert json body to object")
		http.Error(rw, "INVALID URI", http.StatusBadRequest)
		return
	}
	err = data.UpdateProduct(id, prod)
	if err != nil {
		p.l.Fatal("Error ocurred while trying to update object")
		http.Error(rw, "INTERNAL SERVER ERROR", http.StatusInternalServerError)
		return
	}
}

func (p *Product) getProduct(rw http.ResponseWriter, r *http.Request) {
	id, err := requestUtil.GetId(r)
	if err != nil {
		http.Error(rw, "INVALID URI", http.StatusBadRequest)
		return
	}
	record, id, err := data.FindProduct(id)
	if err != nil {
		http.Error(rw, "RESOURCE NOT FOUND", http.StatusNotFound)
		return
	}
	err = record.ToJSON(rw)
	if err != nil {
		http.Error(rw, "UNABLE TO PARSE TO JSON", http.StatusInternalServerError)
		return
	}
}
