package blocks

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	log "go.uber.org/zap"
	"time"
)

const initData = "the genesis block \\m/"

type Blocker interface {
	IsValid() bool
}

type Block struct {
	Index        int32       `json:"index"`
	Timestamp    time.Time   `json:"timestamp"`
	Data         interface{} `json:"data,omitempty"`
	Hash         string      `json:"hash"`
	Next         *Block
	PreviousHash string `json:"previous_hash"`
}

// NewBlock ...
func NewBlock(index string, data interface{}, prevHash string, next *Block) *Block {
	return &Block{
		Index:        0,
		Timestamp:    time.Now(),
		Data:         initData,
		Hash:         calculateHash(index, prevHash, time.Now(), initData),
		Next:         next,
		PreviousHash: prevHash,
	}
}

// GetGenesisBlock ...
func GetGenesisBlock() *Block {
	return NewBlock(
		"0",
		initData,
		"",
		nil)
}

func calculateHash(index, prevHash string, timstamp time.Time, data interface{}) string {
	b := []byte(fmt.Sprintf("%s%s%v%v", index, prevHash, timstamp, data))
	return fmt.Sprintf("%s", sha256.Sum256(b))
}

func (b *Block) IsValid(prev *Block) bool {
	if prev == nil {
		return false
	}
	if prev.Index != b.Index-1 {
		log.Error(errors.New("invalid index"))
		return false
	}

	if prev.Hash != b.PreviousHash {
		log.Error(errors.New("invalid previous hash"))
		return false
	}
	if b.Hash != calculateHash(string(b.Index), b.PreviousHash, b.Timestamp, b.Data) {
		log.Error(errors.New("invalid current hash"))
		return false
	}
	return true
}
