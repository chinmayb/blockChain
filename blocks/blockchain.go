package blocks

import (
	"encoding/json"

	log "go.uber.org/zap"
)

type BlockChainBuilder interface {
	Add() error
	IsValidNewBlock(newBlock *Block) bool
	Mine()
	Length() int
	Get() string
}

type BlockChain struct {
	Head *Block
}

func (b *BlockChain) IsValidNewBlock() bool {
	return false
}

func (b *BlockChain) Mine() {

}

func (b *BlockChain) Add() error {
	return nil
}

func (b *BlockChain) Get() string {
	if b.Head == nil {
		return ""
	}
	var a []byte
	cur := b.Head
	for cur != nil {
		b, _ := json.Marshal(cur)
		a = append(a, b...)
		cur = cur.Next
	}
	return string(a)
}

func (b *BlockChain) Length() int {
	if b.Head == nil {
		return 0
	}
	cur := b.Head
	count := 1
	for cur != nil {
		count++
		cur = cur.Next
	}
	if count != int(cur.Index)+1 {
		log.Any("length", "invalid")
	}
	return count
}
