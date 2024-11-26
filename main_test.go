package main

import (
	"testing"
)

func TestProcessRetailer(t *testing.T) {
	r := Receipt{Retailer: "abcdef123456"}
	p := 0
	processRetailer(r.Retailer, &p)

	if p != 12 {
		t.Fatalf("expected 12, got %v", p)
	}
}

func TestProcessRetailerNonAlpha(t *testing.T) {
	r := Receipt{Retailer: "abcdef123456@#$#@$"}
	p := 0
	processRetailer(r.Retailer, &p)

	if p != 12 {
		t.Fatalf("expected 12, got %v", p)
	}
}

func TestProcessTotalRound(t *testing.T) {
	r := Receipt{Total: "2.00"}
	p := 0

	processTotal(r.Total, &p)

	if p != 75 {
		t.Fatalf("expected 75, got %v", p)
	}
}

func TestProcessTotalMultiple(t *testing.T) {
	r := Receipt{Total: "2.25"}
	p := 0

	processTotal(r.Total, &p)

	if p != 25 {
		t.Fatalf("expected 25, got %v", p)
	}
}

func TestProcessTotalNotRoundNotMultiple(t *testing.T) {
	r := Receipt{Total: "2.65"}
	p := 0

	processTotal(r.Total, &p)

	if p != 0 {
		t.Fatalf("expected 0, got %v", p)
	}
}

func TestProcessItemCountOne(t *testing.T) {
	i1 := Item{ShortDescription: "Item 1", Price: "1.00"}
	r := Receipt{Items: []Item{i1}}
	p := 0

	processItemCount(r.Items, &p)

	if p != 0 {
		t.Fatalf("expected 0, got %v", p)
	}
}
func TestProcessItemCountTwo(t *testing.T) {
	i1 := Item{ShortDescription: "Item 1", Price: "1.00"}
	i2 := Item{ShortDescription: "Item 2", Price: "2.00"}
	r := Receipt{Items: []Item{i1, i2}}
	p := 0

	processItemCount(r.Items, &p)

	if p != 5 {
		t.Fatalf("expected 5, got %v", p)
	}
}
func TestProcessItemCountThree(t *testing.T) {
	i1 := Item{ShortDescription: "Item 1", Price: "1.00"}
	i2 := Item{ShortDescription: "Item 2", Price: "2.00"}
	i3 := Item{ShortDescription: "Item 3", Price: "3.00"}
	r := Receipt{Items: []Item{i1, i2, i3}}
	p := 0

	processItemCount(r.Items, &p)

	if p != 5 {
		t.Fatalf("expected 5, got %v", p)
	}
}

func TestProcessItemCountFour(t *testing.T) {
	i1 := Item{ShortDescription: "Item 1", Price: "1.00"}
	i2 := Item{ShortDescription: "Item 2", Price: "2.00"}
	i3 := Item{ShortDescription: "Item 3", Price: "3.00"}
	i4 := Item{ShortDescription: "Item 4", Price: "4.00"}
	r := Receipt{Items: []Item{i1, i2, i3, i4}}
	p := 0

	processItemCount(r.Items, &p)

	if p != 10 {
		t.Fatalf("expected 10, got %v", p)
	}
}
