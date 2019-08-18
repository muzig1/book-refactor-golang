package main

/*
手法： Replace Method With Method Object

介绍：
	自己对此手法的理解，重写一个例子
	一个回收衣服和回收鞋子的仓库例子

*/

type typ int

const (
	ClothesTyp typ = iota
	ShoeTyp
)

func main() {
	var (
		clothes_price    = 100 // 衣服价格
		clothes_quantity = 8   // 衣服质量 - 满分10分
		clothes_size     = 178 // 衣服大小

		shoe_price    = 500 // 鞋子价格
		shoe_quantity = 7   // 鞋子质量 - 满分10分
		shoe_size     = 39  // 鞋子码数
	)

	// 优化前：
	// 计算衣服的回收价格
	_ = func(price, quantity, size int, t typ) int {
		var aimQuantity int
		// 计算回收价格 - 价格随意计算的，似乎没有什么严谨性
		switch t {
		case ClothesTyp:
			if quantity < 4 {
				aimQuantity = 4
			}
			return price * (aimQuantity - 2) * size
		case ShoeTyp:
			if quantity < 2 {
				aimQuantity = 2
			}
			return price * (aimQuantity - 1) * size
		}
		return -1
	}(clothes_price, clothes_quantity, clothes_size, ClothesTyp)

	// 优化后:
	// 将函数做成函数对象，还可以在其他地方复用,当然此处比较好理解的函数抽象成函数对象，可能比较复杂的函数，就很难以具体的对象来抽象
	// 需要多加理解，以及抽象
	_ = NewRepositoryGetPrice(&Shoe{
		price:    shoe_price,
		quantity: shoe_quantity,
		size:     shoe_size,
		t:        ShoeTyp,
	}, ShoeTyp)
}

func NewRepositoryGetPrice(src interface{}, t typ) int {
	switch t {
	case ClothesTyp:
		if s, ok := src.(*Clothes); ok {
			return s.getPrice()
		}
	case ShoeTyp:
		if s, ok := src.(*Shoe); ok {
			return s.getPrice()
		}
	}
	return -1
}

type (
	Clothes struct {
		price, quantity, size int
		t                     typ
	}

	Shoe struct {
		price, quantity, size int
		t                     typ
	}
)

func (c *Clothes) getPrice() int {
	var aimQuantity int
	if c.quantity < 4 {
		aimQuantity = 4
	}
	return c.price * (aimQuantity - 2) * c.size
}

func (s *Shoe) getPrice() int {
	var aimQuantity int
	if s.quantity < 2 {
		aimQuantity = 2
	}
	return s.price * (aimQuantity - 1) * s.size
}
