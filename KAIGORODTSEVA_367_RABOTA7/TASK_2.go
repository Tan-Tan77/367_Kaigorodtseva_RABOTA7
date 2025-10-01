package main

import (
	"fmt"
)


type Product struct {
	ID       int     
	Name     string  
	Price    float64 
	Quantity int     
}


type Inventory struct {

	products map[int]Product
}


func NewInventory() *Inventory {
	return &Inventory{
		products: make(map[int]Product), 
	}
}


func (i *Inventory) AddProduct(p Product) {

	if _, ok := i.products[p.ID]; ok {
	
		i.products[p.ID] = p
		fmt.Println("Товар ", p.Name,"обновлён на складе. Новое количество: ", p.Quantity)
	} else {
		
		i.products[p.ID] = p
		fmt.Println("Товар ", p.Name," добавлен на склад.")
	}
}


func (i *Inventory) WriteOff(productID int, quantity int) error {

    if product, ok := i.products[productID]; ok {
       
        if product.Quantity >= quantity {
            product.Quantity -= quantity       
            i.products[productID] = product    
            fmt.Printf("Товар \"%s\" (ID: %d) списан со склада в количестве %d. Остаток: %d.\n", product.Name, productID, quantity, product.Quantity)
            return nil 
        } else {
          
            err := fmt.Errorf("недостаточно товара \"%s\" (ID: %d) на складе для списания. В наличии: %d, требуется: %d", product.Name, productID, product.Quantity, quantity)
            fmt.Println(err)
            return err
        }
    } else {
  
        err := fmt.Errorf("товар с ID %d не найден на складе, невозможно списать", productID)
        fmt.Println(err)
        return err 
    }
}

func (i *Inventory) RemoveProduct(productID int) error {

	if _, ok := i.products[productID]; ok {
		delete(i.products, productID) 
		fmt.Println("Товар с ID ", productID," удалён со склада.")
		return nil 
	} else {
	
		err := fmt.Errorf("товар с ID %d не найден на складе, невозможно удалить", productID)
		fmt.Println(err)
		return err 
	}
}


func (i *Inventory) GetTotalValue() float64 {
	var sum float64 = 0.0 

	for _, product := range i.products {
	
		sum += product.Price * float64(product.Quantity)
	}
	return sum 
}


func (i *Inventory) GetProductInfo(productID int) (Product, error) {
	if product, ok := i.products[productID]; ok {
		return product, nil 
	} else {
		err := fmt.Errorf("товар с ID %d не найден на складе", productID)
		return Product{}, err 
	}
}

func main() {
	
	inventory1 := NewInventory()

	product1 := Product{ID: 101, Name: "Пицца", Price: 269.9, Quantity: 20}
	product2 := Product{ID: 102, Name: "Шоколад", Price: 69.9, Quantity: 19}
	product3 := Product{ID: 103, Name: "Орехи", Price: 36.4, Quantity: 10}
	product4 := Product{ID: 104, Name: "Вода", Price: 30.0, Quantity: 50}


	fmt.Println("--- Добавление товаров ---")
	inventory1.AddProduct(product1)
	inventory1.AddProduct(product2)
	inventory1.AddProduct(product3)
	inventory1.AddProduct(product4)

	product1Updated := Product{ID: 101, Name: "Пицца Пепперони", Price: 280.0, Quantity: 15}
	inventory1.AddProduct(product1Updated)
	fmt.Println("--------------------------")


	fmt.Println("--- Получение информации о товаре ---")
	info, err := inventory1.GetProductInfo(102)
	if err == nil {
		fmt.Printf("Информация о товаре (ID: %d): Название - \"%s\", Цена - %.2f, Количество - %d\n", info.ID, info.Name, info.Price, info.Quantity)
	}
	_, err = inventory1.GetProductInfo(999) 
	fmt.Println("-------------------------------------")

	fmt.Println("--- Списание товаров ---")
	err = inventory1.WriteOff(101, 5) 

	err = inventory1.WriteOff(102, 25) 
	err = inventory1.WriteOff(999, 10)
	fmt.Println("------------------------")

	fmt.Println("\n--- Удаление товаров ---")
	
	err = inventory1.RemoveProduct(103) 



	fmt.Println("--- Общая стоимость склада ---")
	fmt.Println("Общая стоимость всех товаров на складе: ", inventory1.GetTotalValue())
	fmt.Println("------------------------------")

	
}
