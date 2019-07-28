package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
)

func main() {
	// 构造用户
	cus := NewCustomer("xiaoming")

	// 构造电影
	var (
		Names = []string{"Roman", "Cat", "Hacker"}
		typs  = []MovieType{NewRelease, Children, Regular}
	)
	for i, name := range Names {
		m := NewMovie(name, typs[i])
		cus.AddRental(NewRental(m, int(rand.Int31n(10)+1)))
	}

	// 输出信息
	fmt.Println(cus.Statement())
}

// NewCustomer .
func NewCustomer(name string) *Customer {
	return &Customer{
		Name:    name,
		Rentals: make([]*Rental, 0),
	}
}

func NewRental(mov *Movie, days int) *Rental {
	return &Rental{
		Movie:      mov,
		DaysRented: days,
	}
}

func NewMovie(title string, passCode MovieType) *Movie {
	return &Movie{
		Title:     title,
		PriceCode: passCode,
	}
}

// Customer .
type Customer struct {
	Name    string
	Rentals []*Rental
}

// AddRental .
func (c *Customer) AddRental(rental *Rental) {
	c.Rentals = append(c.Rentals, rental)
}

// Statement .
func (c *Customer) Statement() string {
	var (
		expression string // 顾客描述信息
	)
	expression += "Rental Record for " + c.Name + "\n"

	for _, r := range c.Rentals {
		// 增加用户描述
		expression += "\t" + r.Movie.Title + "\t" + strconv.Itoa(int(r.AmountFor())) + "\n"
	}

	// 总结
	expression += "Amount owed is " + strconv.Itoa(c.GetTotalAmount()) + "\n"
	expression += "You earned " + strconv.Itoa(c.GetPoints()) + " frequent renter points"
	return expression
}

func (c *Customer) GetTotalAmount() (totalAmount int) {
	for _, r := range c.Rentals {
		totalAmount += int(r.AmountFor())
	}
	return
}

func (c *Customer) GetPoints() (points int) {
	for _, r := range c.Rentals {
		points += r.CalcPoints()
	}
	return
}

type (
	// Rental .
	Rental struct {
		Movie      *Movie
		DaysRented int
	}

	MovieType uint

	// Movie .
	Movie struct {
		Title     string
		PriceCode MovieType
	}
)

const (
	Children MovieType = iota + 1
	Regular
	NewRelease
)

// AmountFor 计算价格
func (r *Rental) AmountFor() (thisAmount float32) {
	return r.Movie.Amount(r.DaysRented)
}

func (r *Rental) CalcPoints() (points int) {
	return r.Movie.Points(r.DaysRented)
}

func (m *Movie) Amount(days int) (thisAmount float32) {
	switch m.PriceCode {
	case Children:
		thisAmount += 2
		if days > 2 {
			thisAmount += float32(days-2) * 1.5
		}
	case Regular:
		thisAmount += float32(days) * 3
	case NewRelease:
		thisAmount += 1.5
		if days > 3 {
			thisAmount += float32(days-3) * 1.5
		}
	default:
		log.Printf("error: not find MovieType%v", m.PriceCode)
	}
	return
}

func (m *Movie) Points(days int) (points int) {
	if m.PriceCode == NewRelease && days > 1 {
		return 2
	} else {
		return 1
	}
}
