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
	r := Receipt{Retailer: "abcdef123456!@#$%^&"}
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

func TestProcessItemsCountOne(t *testing.T) {
	i1 := Item{ShortDescription: "Item 1", Price: "1.00"}
	r := Receipt{Items: []Item{i1}}
	p := 0

	processItemsCount(r.Items, &p)

	if p != 0 {
		t.Fatalf("expected 0, got %v", p)
	}
}
func TestProcessItemsCountTwo(t *testing.T) {
	i1 := Item{ShortDescription: "Item 1", Price: "1.00"}
	i2 := Item{ShortDescription: "Item 2", Price: "2.00"}
	r := Receipt{Items: []Item{i1, i2}}
	p := 0

	processItemsCount(r.Items, &p)

	if p != 5 {
		t.Fatalf("expected 5, got %v", p)
	}
}
func TestProcessItemsCountThree(t *testing.T) {
	i1 := Item{ShortDescription: "Item 1", Price: "1.00"}
	i2 := Item{ShortDescription: "Item 2", Price: "2.00"}
	i3 := Item{ShortDescription: "Item 3", Price: "3.00"}
	r := Receipt{Items: []Item{i1, i2, i3}}
	p := 0

	processItemsCount(r.Items, &p)

	if p != 5 {
		t.Fatalf("expected 5, got %v", p)
	}
}

func TestProcessItemsCountFour(t *testing.T) {
	i1 := Item{ShortDescription: "Item 1", Price: "1.00"}
	i2 := Item{ShortDescription: "Item 2", Price: "2.00"}
	i3 := Item{ShortDescription: "Item 3", Price: "3.00"}
	i4 := Item{ShortDescription: "Item 4", Price: "4.00"}
	r := Receipt{Items: []Item{i1, i2, i3, i4}}
	p := 0

	processItemsCount(r.Items, &p)

	if p != 10 {
		t.Fatalf("expected 10, got %v", p)
	}
}

func TestProcessItemMultiple3(t *testing.T) {
	i := Item{ShortDescription: "123456 89", Price: "12.25"}
	p := 0
	processItem(i, &p)

	if p != 3 {
		t.Fatalf("expected 3, got %v", p)
	}
}

func TestProcessItemNotMultiple3(t *testing.T) {
	i := Item{ShortDescription: "123456 8", Price: "12.25"}
	p := 0
	processItem(i, &p)

	if p != 0 {
		t.Fatalf("expected 0, got %v", p)
	}
}

func TestProcessItems(t *testing.T) {
	i1 := Item{ShortDescription: "123456", Price: "10.00"}   //2
	i2 := Item{ShortDescription: "12345", Price: "2.00"}     //0
	i3 := Item{ShortDescription: "123456789", Price: "6.00"} //1.2 => 2
	i4 := Item{ShortDescription: "12", Price: "4.00"}
	r := Receipt{Items: []Item{i1, i2, i3, i4}}
	p := 0

	processItems(r.Items, &p)

	if p != 14 {
		t.Fatalf("expected 14, got %v", p)
	}

}

func TestProcessPurchaseDateOdd(t *testing.T) {
	r := Receipt{PurchaseDate: "2024-01-11"}
	p := 0

	processPurchaseDate(r.PurchaseDate, &p)

	if p != 6 {
		t.Fatalf("expected 6, got %v", p)
	}
}

func TestProcessPurchaseDateEven(t *testing.T) {
	r := Receipt{PurchaseDate: "2024-01-02"}
	p := 0

	processPurchaseDate(r.PurchaseDate, &p)

	if p != 0 {
		t.Fatalf("expected 0, got %v", p)
	}
}
func TestProcessPurchaseTimeIn(t *testing.T) {
	r := Receipt{PurchaseTime: "14:01"}
	p := 0

	processPurchaseTime(r.PurchaseTime, &p)

	if p != 10 {
		t.Fatalf("expected 10, got %v", p)
	}
}

func TestProcessPurchaseTimeBefore(t *testing.T) {
	r := Receipt{PurchaseTime: "13:59"}
	p := 0

	processPurchaseTime(r.PurchaseTime, &p)

	if p != 0 {
		t.Fatalf("expected 0, got %v", p)
	}
}

func TestProcessPurchaseTimeAfter(t *testing.T) {
	r := Receipt{PurchaseTime: "16:01"}
	p := 0

	processPurchaseTime(r.PurchaseTime, &p)

	if p != 0 {
		t.Fatalf("expected 0, got %v", p)
	}
}
