package main

import (
	"go.uber.org/zap"

	"github.com/chinmayb/blockChain/blocks"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	log := logger.Sugar()
	log.Debug("initializing..")

	// this will be the head
	_ = blocks.GetGenesisBlock()
}
