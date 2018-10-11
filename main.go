package main

import "fmt"

type Transactor interface{
	Transact() bool
}

type Transaction struct{
	Sender string
	Recipient string
	Amount int64
}

func (Transaction) Transact() bool {
return false
}

func NewTransaction(sender, recipient string, amount int64) Transactor {
	return &Transaction{
		sender, recipient, amount,
	}
}


type Block struct {
	index int32
	transaction Transaction
	proof int64
	previousHash string

}

func NewDefaultBlock() {

	//x := new([]string) // var a string
	y := make([]int, 0)
	var w [1]int
	x := new([]int)
	fmt.Println(x, y, w)
	if *x == nil {
		fmt.Println("x is  nil")
	}
	if y == nil{
		fmt.Println("y is  nil")
	}
	//if w == nil {
	//	fmt.Println("w is  nil", w)
	//}

}


type  BlockChainBuilder interface{
	NewTransaction() Transactor
	LasBlock() Block
	NewBlock() Block
}


type  BlockChain struct{
	chain []Block

}

func main(){
	NewDefaultBlock()
}