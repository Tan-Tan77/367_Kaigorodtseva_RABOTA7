package main

import (
	"fmt"
)

type Customer struct {
	Name string
}

type OrderItem struct {
	Name  string
	Price float64
}

type Order struct {
	Orders    map[Customer][]OrderItem
	Condition string
}

func (o *Order) TotalCost(c Customer) float64 {
	sum := 0.0
	if orderSlice, ok := o.Orders[c]; ok {

		for _, item := range orderSlice {
			sum += item.Price
		}
	} else {

		fmt.Printf("У посетителя %s нет заказов\n", c.Name)
	}
	return sum
}

func (o *Order) UpdateStatus(status string) {
	o.Condition = status
	fmt.Printf("Статус заказа изменён на \"%s\"\n", status)
}

func (o *Order) AddProduct(item OrderItem, c Customer) {

	currentItems, ok := o.Orders[c]
	if !ok {

		currentItems = []OrderItem{}
	}
	currentItems = append(currentItems, item)

	o.Orders[c] = currentItems
	fmt.Println("Товар ", item.Name, " добавлен к заказу посетителя", c.Name)
}

func (o *Order) DeleteProduct(itemToDelete OrderItem, c Customer) {

	if orderSlice, ok := o.Orders[c]; ok {

		var updatedOrderSlice []OrderItem
		deleted := false

		for _, item := range orderSlice {

			if item.Name != itemToDelete.Name || item.Price != itemToDelete.Price {
				updatedOrderSlice = append(updatedOrderSlice, item)
			} else {
				deleted = true
			}
		}

		if deleted {
			o.Orders[c] = updatedOrderSlice
			fmt.Println("Товар ", itemToDelete.Name, " удалён из заказа", c.Name)
		} else {
			fmt.Println("Товар ", itemToDelete.Name, " не найден для", c.Name)
		}
	} else {
		fmt.Println("У посетителя ", c.Name, " нет заказов для удаления товара")
	}
}

func main() {

	c := Customer{Name: "Tim"}
	item1 := OrderItem{Name: "choco", Price: 130.7}
	item2 := OrderItem{Name: "milk", Price: 110.0}

	var order1 Order

	order1.Orders = make(map[Customer][]OrderItem)

	order1.AddProduct(item1, c)

	fmt.Printf("Общая стоимость для %s: %.2f\n", c.Name, order1.TotalCost(c))

	order1.AddProduct(item2, c)
	fmt.Printf("Общая стоимость для %s после добавления %s: %.2f\n", c.Name, item2.Name, order1.TotalCost(c))

	order1.DeleteProduct(item1, c)
	fmt.Printf("Общая стоимость для %s после удаления %s: %.2f\n", c.Name, item1.Name, order1.TotalCost(c))

	item3 := OrderItem{Name: "water", Price: 50.0}
	order1.DeleteProduct(item3, c)

	order1.UpdateStatus("delivered")
}
