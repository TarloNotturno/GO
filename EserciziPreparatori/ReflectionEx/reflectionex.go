package reflectionex

import (
	"fmt"
	"reflect"
)

type Pizza struct {
	price   int
	Request string
	recipe  Pizzadecorator
}

type Pizzadecorator interface {
	decorPizza() string
}

func (r *Pizza) decorPizza() string {
	if r.recipe != nil {

		//	return r.Request
		//} else {
		r.Request = r.recipe.decorPizza() + " " + r.Request
		r.recipe = nil
		//return r.recipe.decorPizza() + " " + r.Request
	}
	return r.Request
}

func SendOrder() *Pizza {
	baseRedPizza := Pizza{Request: "tomato pasta"}
	mozzarellaPizza := Pizza{Request: "mozzarella", recipe: &baseRedPizza}
	pepperoniPizza := Pizza{Request: "sausage", recipe: &mozzarellaPizza}
	wurstelPizza := Pizza{Request: "wurstel", recipe: &mozzarellaPizza}
	americanPizza := Pizza{Request: "french fries", recipe: &wurstelPizza}
	//fmt.Println(reflect.ValueOf(americanPizza.recipe).Elem().Field(1))
	fmt.Println("reflect su una variabile : reflect.ValueOf(22)->", reflect.ValueOf(22), "reflect.TypeOf(22)->", reflect.TypeOf(22))
	fmt.Println("reflect su una struct : reflect.ValueOf(americanPizza.recipe)->", reflect.ValueOf(americanPizza.recipe), "reflect.TypeOf(americanPizza.recipe)->", reflect.TypeOf(americanPizza.recipe))
	fmt.Println("reflect su una struct : reflect.ValueOf(americanPizza.recipe).Elem()->", reflect.ValueOf(americanPizza.recipe).Elem())
	pizzaType := reflect.ValueOf(americanPizza.recipe).Elem()
	for i := 0; i < pizzaType.NumField(); i++ {
		fmt.Println("element ", i, " type ", pizzaType.Field(i).Type(), " value", pizzaType.Field(i))
	}
	fmt.Println("americanPizza.decorPizza()", reflect.ValueOf(americanPizza.recipe))
	//app := reflect.ValueOf(americanPizza.recipe).Elem()
	//app.NumField()
	pepperoniPizza.decorPizza()
	fmt.Println(reflect.ValueOf(&pepperoniPizza).Elem().Field(1))
	fmt.Println(americanPizza.decorPizza())
	return &americanPizza
}
