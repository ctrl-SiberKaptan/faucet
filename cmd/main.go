package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/obscuronet/faucet/faucet"
)

// local execution: PORT=80 go run . --nodeHost 127.0.0.1 --pk 0x8dfb8083da6275ae3e4f41e3e8a8c19d028d32c9247e24530933782f2a05035b
func main() {
	cfg := parseCLIArgs()

	if cfg.PK == "" {
		panic("no key loaded")
	}
	nodeAddr := fmt.Sprintf("%s:%d", cfg.Host, cfg.WSPort)
	key, err := crypto.HexToECDSA(cfg.PK[2:])
	if err != nil {
		panic(err)
	}

	f, err := faucet.NewFaucet(nodeAddr, cfg.ChainID, key)
	if err != nil {
		panic(err)
	}
	server := faucet.NewWebServer(f)
	server.Start()
}
