package main

import("fmt")

type Cache struct{

	key string
	
}


func (c *Cache) Set(key_ string){

	
	c.key = key_

}

func (c *Cache) Get()string{
if(c.key == ""){fmt.Println("Ключ пуст")}
	return c.key
}

func (c *Cache) Delete(){

	if(c.key != ""){
	 fmt.Println("Ключ удалён", c.key)
	 c.key = ""}else{fmt.Println("Ключ уже пуст")}
	
}


func main(){

var c Cache


c.Set("123456")
fmt.Println("Ключ: ", c.Get())
c.Delete()
fmt.Println(c.Get())
}