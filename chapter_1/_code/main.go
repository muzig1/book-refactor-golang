package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	// 构造用户
	cus := NewCustomer("xiaoming")

	// 构造电影
	var (
		Names = []string{"Roman", "Cat", "Hacker"}
	)
	m1 := NewNewRelease(Names[0])
	m2 := NewChildrenMovie(Names[1])
	m3 := NewRegular(Names[2])
	cus.AddRental(NewRental(m1, int(rand.Int31n(10)+1)))
	cus.AddRental(NewRental(m2, int(rand.Int31n(10)+1)))
	cus.AddRental(NewRental(m3, int(rand.Int31n(10)+1)))

	// 输出信息
	fmt.Println(cus.Statement())
}

// --------- 构造区域 ---------

func NewCustomer(name string) *Customer {
	return &Customer{
		Name:    name,
		Rentals: make([]*Rental, 0),
	}
}

func NewRental(mov IMovie, days int) *Rental {
	return &Rental{
		Movie:      mov,
		DaysRented: days,
	}
}

func NewChildrenMovie(title string) IMovie {
	return &ChildrenMovie{
		Movie: newMovie(title, Children),
	}
}

func NewRegular(title string) IMovie {
	return &RegularMovie{
		Movie: newMovie(title, Regular),
	}
}

func NewNewRelease(title string) IMovie {
	return &NewReleaseMovie{
		Movie: newMovie(title, NewRelease),
	}
}

func newMovie(title string, passCode MovieType) *Movie {
	return &Movie{
		Title:     title,
		PriceCode: passCode,
	}
}

// --------- 用户 ---------

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
		expression += "\t" + r.Movie.GetTitle() + "\t" + strconv.Itoa(int(r.AmountFor())) + "\n"
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

// --------- 数据层 ---------

type (
	// Rental .
	Rental struct {
		Movie      IMovie
		DaysRented int
	}

	MovieType uint

	IMovie interface {
		GetTitle() string
		GetPriceCode() MovieType
		Amount(int) float32
		Points(int) int
	}

	// Movie .
	Movie struct {
		Title     string
		PriceCode MovieType
	}

	ChildrenMovie struct {
		*Movie
	}

	RegularMovie struct {
		*Movie
	}

	NewReleaseMovie struct {
		*Movie
	}
)

func (nm *NewReleaseMovie) GetTitle() string {
	return nm.Title
}

func (nm *NewReleaseMovie) GetPriceCode() MovieType {
	return nm.PriceCode
}

func (rm *RegularMovie) GetTitle() string {
	return rm.Title
}

func (rm *RegularMovie) GetPriceCode() MovieType {
	return rm.PriceCode
}

func (cm *ChildrenMovie) GetTitle() string {
	return cm.Title
}

func (cm *ChildrenMovie) GetPriceCode() MovieType {
	return cm.PriceCode
}

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

func (cm *ChildrenMovie) Amount(days int) (thisAmount float32) {
	thisAmount += 2
	if days > 2 {
		thisAmount += float32(days-2) * 1.5
	}
	return
}

func (cm *RegularMovie) Amount(days int) (thisAmount float32) {
	thisAmount += float32(days) * 3
	return
}

func (cm *NewReleaseMovie) Amount(days int) (thisAmount float32) {
	thisAmount += 1.5
	if days > 3 {
		thisAmount += float32(days-3) * 1.5
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
