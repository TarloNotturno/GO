package main

import (
	"MyRnd"
	"fmt"
)

type WeaponValue struct {
	weapon_type      string
	scalingVal       string
	damage           uint8
	damage_reduction uint8
	damage_type      string
	faith_scaling    string
	int_scaling      string
	dex_scaling      string
	str_scaling      string
}

type WeaponCategory struct {
	weaponType map[uint8]string
	scalingVal map[string]float32
}

func InitializeDefinition(weaponCategory *WeaponCategory) {
	weaponType := map[uint8]string{
		1: "mele",
		2: "shield",
		3: "staff",
		4: "chime",
	}
	scalingVal := map[string]float32{
		"S": 3,
		"A": 2,
		"B": 1.5,
		"C": 1.25,
		"D": 1,
		"E": 0,
	}
	weaponCategory.weaponType = weaponType
	weaponCategory.scalingVal = scalingVal
}

type Inventory struct {
	weapons []WeaponValue
}

func (v WeaponValue) New(droppedWeapon *WeaponValue) {

	droppedWeapon.weapon_type = "mele"
	droppedWeapon.scalingVal = "A"
	droppedWeapon.damage = 22
	droppedWeapon.damage_reduction = 0
	droppedWeapon.damage_type = "raw"
	droppedWeapon.faith_scaling = "E"
	droppedWeapon.int_scaling = "E"
	droppedWeapon.dex_scaling = "B"
	droppedWeapon.str_scaling = "C"

}

func main() {
	var weaponCategory WeaponCategory
	InitializeDefinition(&weaponCategory)
	fmt.Println(weaponCategory.weaponType)
	var inventory Inventory
	var droppedWeapon WeaponValue
	droppedWeapon.New(&droppedWeapon)
	inventory.weapons = append(inventory.weapons, droppedWeapon)
	fmt.Println(inventory.weapons)
	fmt.Println(MyRnd.GetNumbRnd())
}
