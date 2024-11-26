package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	sl "slices"
	str "strings"
)

type Item struct {
	ShortDescription string
	Price            string
}

type Receipt struct {
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Items        []Item `json:"items"`
	Total        string
}

func processRetailer(r string, p *int) {
	regex := regexp.MustCompile(`[a-zA-Z0-9]`)
	c := len(regex.FindAllString(r, -1))
	*p += c
}

func processTotal(t string, p *int) {
	_, c, _ := str.Cut(t, ".")
	if c == "00" {
		*p += 50
	}

	if sl.Contains([]string{"00", "25", "50", "75"}, c) {
		*p += 25
	}
}

func processItemCount(i []Item, p *int) {
	if len(i)%2 >= 0 {
		*p += (len(i) / 2) * 5
	}
}
func process(r Receipt) {
	p := 0
	processRetailer(r.Retailer, &p)
	processTotal(r.Total, &p)
	processItemCount(r.Items, &p)
	log.Println(r, p)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /receipts/{id}/points/", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprint(w, id)
	})

	mux.HandleFunc("POST /receipts/process/", func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		var rec Receipt
		err := decoder.Decode(&rec)
		if err != nil {
			panic(err)
		}
		process(rec)
	})

	http.ListenAndServe("localhost:8090", mux)
}
