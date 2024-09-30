package DP

import "fmt"

type Shop struct {
	chairs    []Chair
	wardrobes []Wardrobe
}

func (s *Shop) Create(chairs []Chair, wardrobes []Wardrobe) {
	s.chairs = chairs
	s.wardrobes = wardrobes
}

func (s Shop) ListWearhouse() {
	fmt.Printf("List of product available:\n  >>>>  %v:", s.chairs[0].product_type)
	for i, chair := range s.chairs {
		fmt.Printf("%v. weight:%v    price %v\n", i+1, chair.weight, chair.price)
	}
	fmt.Printf("  >>>>  %v:\n", s.wardrobes[0].product_type)
	for i, wardrobes := range s.wardrobes {
		fmt.Printf("%v. capacity:%v    price %v\n", i+1, wardrobes.capacity, wardrobes.price)
	}

}

type Factorer interface {
	Create(int)
}

type Wardrobe struct {
	product_type string
	capacity     int
	price        int
}

func (a *Wardrobe) Create(price int) {
	a.product_type = "Wardrobe"
	a.capacity = 30
	a.price = price
}

type Chair struct {
	product_type string
	weight       int
	price        int
}

func (a *Chair) Create(price int) {
	a.product_type = "Chair"
	a.weight = 30
	a.price = price
}
