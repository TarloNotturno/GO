package DP

type Prototyper interface {
	Prototype() Prototyper
}

func Copy(p Prototyper) Prototyper {
	copy := p.Prototype()
	return copy
}

func (s Smartphone) ReturnPrice() int {
	return s.Product_Price
}

func (s Smartphone) Prototype() Prototyper {
	var copy Smartphone
	copy.K_keyboard = s.K_keyboard
	copy.M_microphope = s.M_microphope
	copy.Product_Price = s.Product_Price
	copy.K_keyboard.keys = s.K_keyboard.keys
	return copy
}

func (s Shop) Prototype() Prototyper {
	var copy Shop
	copy.chairs = s.chairs
	copy.wardrobes = s.wardrobes
	return copy
}

func (s Shape) Prototype() Prototyper {
	var copy Shape
	copy.Shape_type = s.Shape_type
	copy.color = s.color
	return copy
}

type Shape_array []Shape

func (s Shape_array) Prototype() Prototyper {
	copy := make(Shape_array, len(s))
	for i, _ := range s {
		copy[i] = Copy(s[i]).(Shape)
	}
	return copy
}
