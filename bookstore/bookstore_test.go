package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title: "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()
	b :=  bookstore.Book{
		Title: "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
	var want int = 1
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies 
	if want != got {
		t.Errorf("want %d copies after buying 1 copy from a stock of 1, got %d", want, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
	}
	_, err :=bookstore.Buy(b)
	if err == nil {
		t.Error("want error buying book when zero copies left, but got nil")
	}
}