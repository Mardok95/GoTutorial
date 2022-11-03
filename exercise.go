// This is a file that i use for exercise about random tasks.

package main

import (
	"fmt"
)

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	unit := map[string]int{}
	unit["quarter_of_a_dozen"] = 3
	unit["half_of_a_dozen"] = 6
	unit["dozen"] = 12
	unit["small_gross"] = 120
	unit["gross"] = 144
	unit["great_gross"] = 1728
	return unit

}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := map[string]int{}

	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, exists := units[unit]
	if exists == false {
		return false
	}
	firstValue := bill[item]
	bill[item] = value + firstValue
	return true

}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	value, exists := bill[item]
	value1, exists1 := units[unit]
	if exists == false || exists1 == false {
		return false
	}
	if value < 0 || value1 < 0 {
		return false
	}
	if value == 0 {
		delete(bill, item)
		return true
	}
	firstValue := bill[item]
	bill[item] = value + firstValue
	return true

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	value, exists := bill[item]
	if exists == false {
		return 0, false
	} else {
		return value, true
	}
}

func main() {
	//fmt.Println("Exercise")
	units := Units()
	fmt.Println(units)

	bill := NewBill()
	fmt.Println(bill)

	bill = NewBill()
	units = Units()
	ok := AddItem(bill, units, "carrot", "dozen")
	fmt.Println(ok)

	bill = NewBill()
	units = Units()
	ok = RemoveItem(bill, units, "carrot", "dozen")
	fmt.Println(ok)

	bill = map[string]int{"carrot": 12, "grapes": 3}
	qty, ok1 := GetItem(bill, "carrot")
	fmt.Println(qty)
	// Output: 12
	fmt.Println(ok1)
	// Output: true

}
