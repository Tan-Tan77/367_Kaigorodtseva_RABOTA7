package main

import("fmt"
"crypto/sha256"
"encoding/hex"
)


type User struct{

Username string
Email string
Password string

	
}

func (u *User) SetPassword(password string)string{

	passwordBytes := []byte(password)


 hashArray := sha256.Sum256(passwordBytes)

 
 hashString := hex.EncodeToString(hashArray[:])

 u.Password = hashString
return hashString

}

func (u *User) VerifyPassword(password string) bool {

	if(u.Password == password){
		return true
}else{return false}


}



func main(){

	var user1 User
	user1.Username = "Tan"
	user1.Email ="some@mail.com"
user1.Password ="1234end"


fmt.Println( "Пароль изменён на" ,user1.SetPassword("1234end"))

	if(user1.VerifyPassword("622573f090d5885289166b7ea6ca26b3c00ab2157b4a8a247847abc4b5abaf5e") == true){
		fmt.Println("Вход в аккаунт", user1.Username , user1.Email)

	}else{
			fmt.Println("Неверный пароль")
	}





}
