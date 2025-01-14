package main

import (
	"encoding/json"
	"math"
	"net/http"
	"regexp"
	sl "slices"
	"strconv"
	str "strings"

	"github.com/google/uuid"
)

var receipts map[uuid.UUID]int

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

type PointsResponse struct {
	Points int `json:"points"`
}

type ProcessResponse struct {
	ID uuid.UUID `json:"id"`
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

func processItemsCount(i []Item, p *int) {
	if len(i)%2 >= 0 {
		*p += (len(i) / 2) * 5
	}
}

func processItem(i Item, p *int) {
	if len(str.TrimSpace(i.ShortDescription))%3 == 0 {
		pf, _ := strconv.ParseFloat(i.Price, 32)
		*p += int(math.Ceil((pf * 0.2)))
	}
}

func processItems(i []Item, p *int) {
	processItemsCount(i, p)
	for _, item := range i {
		processItem(item, p)
	}
}

func processPurchaseDate(d string, p *int) {
	date, _ := strconv.ParseInt(d[len(d)-2:], 10, 32)
	if date%2 != 0 {
		*p += 6
	}
}

func processPurchaseTime(t string, p *int) {
	hour, _ := strconv.ParseFloat(t[:2], 32)
	minute, _ := strconv.ParseFloat(t[3:], 32)
	time := hour + (minute / 100.00)

	if time > 14 && time < 16 {
		*p += 10
	}
}

func process(r Receipt) int {
	p := 0
	processRetailer(r.Retailer, &p)
	processTotal(r.Total, &p)
	processItems(r.Items, &p)
	processPurchaseDate(r.PurchaseDate, &p)
	processPurchaseTime(r.PurchaseTime, &p)
	return p
}

func main() {

	receipts = make(map[uuid.UUID]int)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /receipts/{id}/points/", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		uuid, err := uuid.Parse(id)
		if err != nil {
			panic(err)
		}
		res := PointsResponse{Points: receipts[uuid]}
		json.NewEncoder(w).Encode(res)
	})

	mux.HandleFunc("POST /receipts/process/", func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		var rec Receipt
		err := decoder.Decode(&rec)
		if err != nil {
			panic(err)
		}

		var id = uuid.New()
		receipts[id] = process(rec)

		res := ProcessResponse{ID: id}
		json.NewEncoder(w).Encode(res)
	})

	http.ListenAndServe("localhost:8090", mux)
}
