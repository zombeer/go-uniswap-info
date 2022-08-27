package main

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/zombeer/go-uniswap-info/uniswap"
)

const UNISWAP_FACTORY_ADDRESS = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"

func main() {
	log.Println("Hello WEB3 world!")

	client := uniswap.GetClient()
	factory := uniswap.GetUniswapFactory(UNISWAP_FACTORY_ADDRESS, client)

	currentBlockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal("not able to fetch block number", err)
	}

	callOpts := bind.CallOpts{
		Context:     context.Background(),
		BlockNumber: big.NewInt(int64(currentBlockNumber)),
	}
	pairsCount, err := factory.AllPairsLength(&callOpts)
	if err != nil {
		log.Fatal("not able to fetch pairs count", err)
	}
	log.Println("Pairs count:", pairsCount.String())

	for i := 1; i < 30; i++ {
		pairAddress, err := factory.AllPairs(&callOpts, big.NewInt(int64(i)))
		if err != nil {
			log.Fatal("not able to get pair address", i, err)
		}
		pair := uniswap.GetUniswapPair(pairAddress, client)

		pairDetails := uniswap.GetPairDetails(pair, big.NewInt(int64(currentBlockNumber)), client)
		log.Printf("%v [%v-%v], price: %v",
			i,
			pairDetails.Token0.Symbol,
			pairDetails.Token1.Symbol,
			pairDetails.Price,
		)
	}
}
