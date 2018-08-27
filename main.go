package main


type Transactor interface{
	Transact() bool
}

type Transaction struct{
	Sender string
	Recipient string
	Amount int64
}

func (Transaction) Transact() bool {

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

func (b *Block) hash() string{

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

}