package cryptocurrency

type Transactor interface {
	Transact() bool
}

type Transaction struct {
	Id        string
	Sender    string
	Recipient string
	Amount    int64
}

func (Transaction) Transact() bool {
	return false
}
