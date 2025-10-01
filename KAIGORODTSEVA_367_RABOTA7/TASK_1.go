package main

import("fmt")

type BankAccount struct{

	accountNumber int64
	holderName string
balance float64

}


func (b *BankAccount) Deposit (amount float64){


b.balance += amount
 fmt.Println("На счёт положена сумма " , amount)

}

func (b *BankAccount) Withdraw (amount float64) string{

	var err string

	if(b.balance > 0){
b.balance -= amount
 err = ""
 fmt.Println("Со счёта снята сумма " , amount)
	}else{
	  err = "Баланс пуст, невозможно ничего снять"
	}
return err

}

func (b *BankAccount) GetBalance()float64{
return b.balance
}

func (b *BankAccount) GetName()string{
return b.holderName
}
func (b *BankAccount) GetNumber()int64{
return b.accountNumber
}


func main(){

var bankAccount BankAccount

bankAccount.balance = 1000
bankAccount.holderName = "Tan"
bankAccount.accountNumber = 12234

fmt.Println("Аккаунт с именем", bankAccount.GetName(),"счётом номер", bankAccount.GetNumber() )

bankAccount.Deposit(200)
fmt.Println("Баланс кошелька",bankAccount.GetBalance())
bankAccount.Withdraw(100)
fmt.Println("Баланс кошелька",bankAccount.GetBalance())
fmt.Println(bankAccount.Withdraw(1100))
fmt.Println("Баланс кошелька",bankAccount.GetBalance())
fmt.Println(bankAccount.Withdraw(1100))

}