package main

import (
    "fmt"
)


type Room struct {
    ID   int 
    Cost int 
}


type ReservationManager struct {

    BookedRooms map[int]int
}

func (rm *ReservationManager) BookRoom(room Room, days int) {

    if rm.BookedRooms == nil {
        rm.BookedRooms = make(map[int]int)
    }
  
    rm.BookedRooms[room.ID] = days
    cost := days * room.Cost 
    fmt.Println("Номер ID:", room.ID," забронирован на ", days," дня со стоимостью ", cost )
}

func (rm *ReservationManager) IsRoomBooked(roomID int) {

    if _, ok := rm.BookedRooms[roomID]; ok {
       
        fmt.Println("Номер ID: ", roomID," забронирован")
    } else {
    
        fmt.Println("Номер ID:", roomID," свободен")
    }
}

func main() {
 
    r1 := Room{ID: 100, Cost: 2300}
    r2 := Room{ID: 900, Cost: 5300} 


    var rm ReservationManager
    rm.BookedRooms = make(map[int]int) 

 
    fmt.Println("--- Проверка статуса до бронирования ---")
    rm.IsRoomBooked(r1.ID)
    rm.IsRoomBooked(r2.ID)
    fmt.Println("----------------------------------------")


    rm.BookRoom(r1, 5) 
    rm.BookRoom(r2, 12) 

    fmt.Println("--- Проверка статуса после бронирования ---")
    rm.IsRoomBooked(r1.ID)
    rm.IsRoomBooked(r2.ID)
    fmt.Println("-------------------------------------------")

   

 
}