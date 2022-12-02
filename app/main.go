package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

type modifProduct struct {
	Id          int
	Name        string
	Description string
	Ammount     int
	Price       int
	Category    int
}
type Product struct {
	Id          int
	Name        string
	Description string
	Ammount     int
	Price       int
	Category    string
}
type Products struct {
	Products []Product
}

func main() {
	var err error
	db, err = sql.Open("postgres", "host=127.0.0.1 user=api password=123456 dbname=api sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Server started...")
	http.HandleFunc("/v1/products/modify/", modProduct)
	http.HandleFunc("/v1/products/", getProducts)
	log.Fatal(http.ListenAndServe(":80", nil))

}

func modProduct(w http.ResponseWriter, r *http.Request) {
	var g_product modifProduct

	if r.Method != "PUT" {
		http.Error(w, "Method not allowed", 405)
	} else {
		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&g_product)
		if err != nil {
			panic(err)
		}

		fmt.Println("# Insert Query")
		query := fmt.Sprintf("UPDATE products SET name = '%s', description = '%s', ammount = %d, price = %d, category = %d WHERE id = %d RETURNING id, name, description, ammount, price, category", g_product.Name, g_product.Description, g_product.Ammount, g_product.Price, g_product.Category, g_product.Id)
		rows := db.QueryRow(query)
		fmt.Println("# Insert QUERY: %s", rows)
		var id, ammount, price, category int
		var name, description, selcategory string
		a := rows.Scan(&id, &name, &description, &ammount, &price, &category)
		if a != nil {
			panic(a)
		}
		query2 := fmt.Sprintf("SELECT c.name from category c where id=%d", category)
		sel, err := db.Query(query2)
		if err != nil {
			panic(err)
		}
		for sel.Next() {
			err = sel.Scan(&selcategory)
			if err != nil {
				panic(err)
			}
			fmt.Fprintf(w, "{\"id\":%d,\"name\":%s,\"description\":%s,\"ammount\":%d,\"price\":%d,\"category\":%s}", id, name, description, ammount, price, selcategory)

		}
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
	} else {
		w_products := Products{}

		fmt.Println("# Querying")
		rows, err := db.Query("SELECT p.id, p.name, p.description, p.ammount, p.price, c.name as category from products p left join category c on p.category=c.id")
		if err != nil {
			panic(err)
		}
		for rows.Next() {
			w_product := Product{}
			err = rows.Scan(&w_product.Id, &w_product.Name, &w_product.Description, &w_product.Ammount, &w_product.Price, &w_product.Category)
			if err != nil {
				panic(err)
			}
			w_products.Products = append(w_products.Products, w_product)
		}
		json.NewEncoder(w).Encode(w_products)
	}

}

