// custom type, error handling

package main

import (
	"fmt"
	"math"
	"os"
)

type Transaction struct {
	Symbol string
	Volume int
	Price  float64
	Buy    bool
}

func (t *Transaction) Value() float64 {
	value := float64(t.Volume) * t.Price
	valueRounded := math.Round(value * 1000) / 1000
	if t.Buy {
		value *= -1
	}
	return valueRounded
}

func NewTransaction(symbol string, volume int, price float64, buy bool) (*Transaction, error) {
	if symbol == "" {
		return nil, fmt.Errorf("Symbol missing.")
	}

	if volume <= 0 || price <= 0 {
		return nil, fmt.Errorf("Both volume and price needed")
	}

	transaction := &Transaction{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}

	return transaction, nil
}

func main() {
	transaction, err := NewTransaction("Google", 10, 314.95, true)
	if err != nil {
		fmt.Printf("error: cannot create transaction - %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Transaction created: %+v\n", transaction)
	fmt.Printf("Tranaction\"s value: %.3f\n", transaction.Value())
}
