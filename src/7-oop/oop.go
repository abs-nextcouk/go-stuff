package main

import (
	"fmt"
	"time"
)

type Budget struct {
	CampaignID string
	Balance    float64
	Expires    time.Time
}

// This is how to do "constructors"
func NewBudget(campaignID string, balance float64, expires time.Time) (*Budget, error) {
	if campaignID == "" {
		return nil, fmt.Errorf("empty campaignID")
	}
	if balance <= 0 {
		return nil, fmt.Errorf("balance must be > 0")
	}
	if expires.Before(time.Now()) {
		return nil, fmt.Errorf("bad expiration date")
	}

	b := Budget{
		CampaignID: campaignID,
		Balance:    balance,
		Expires:    expires,
	}
	return &b, nil
}

func (b Budget) TimeLeft() time.Duration {
	return b.Expires.Sub(time.Now().UTC())
}

func (b *Budget) Update(sum float64) {
	b.Balance += sum
}

type Shape interface {
	Area() float64
}

// generics
type Ordered interface {
	int | float64 | string
}

func main() {
	b1 := Budget{CampaignID: "aaaa", Balance: 111, Expires: time.Now()}
	fmt.Println(b1)
	fmt.Printf("%#v\n", b1)
	b1 = Budget{"bbbbb", 112, time.Now()}
	fmt.Println(b1)
	fmt.Printf("%#v\n", b1)
	b1 = Budget{Expires: time.Now()}
	fmt.Println(b1)
	fmt.Printf("%#v\n", b1)

	fmt.Println("================")
	b2 := Budget{Expires: time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b2.TimeLeft())

	fmt.Println("================")
	b3, err := NewBudget("pups", 32.1, time.Now().Add(7*24*time.Hour))
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("%#v\n", b3)
	}
}
