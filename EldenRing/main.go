package main

import (
	"fmt"

	"github.com/TarloNotturno/MyFunc"
)

func main() {
	var weaponCategory MyFunc.WeaponCategory
	//MyFunc.InitializeDefinition(&weaponCategory)
	//fmt.Println(weaponCategory.weaponType)
	var inventory MyFunc.Inventory
	//inventory.InitializeInv(&inventory)
	MyFunc.InitializeData(&inventory, &weaponCategory)
	fmt.Println(MyFunc.GetNumbRnd())
}
