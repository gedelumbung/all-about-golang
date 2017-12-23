package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Product struct {
	Id          string  `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	Price       int     `json:"price,omitempty"`
	Category    *Category `json:"category,omitempty"`
}
type Category struct {
	Title   string `json:"title,omitempty"`
	Slug    string `json:"slug,omitempty"`
}

var products []Product

func GetProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range products {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Product{})
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	products = append(products, product)
	json.NewEncoder(w).Encode(products)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range products {
		if item.Id == params["id"] {
			products = append(products[:index], products[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(products)
	}
}

func main() {
	router := mux.NewRouter()
	products = append(products, Product{Id: "1", Title: "Tipat Kuah", Price: 10000, Category: &Category{Title: "Balinese Food", Slug: "balinese-food"}})
	router.HandleFunc("/products", GetProducts).Methods("GET")
	router.HandleFunc("/products/{id}", GetProduct).Methods("GET")
	router.HandleFunc("/products", CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", DeleteProduct).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8800", router))
}
