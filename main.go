package main

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	uniinfo "github.com/zombeer/go-uniswap-info/uni-info"
)

const UNISWAP_FACTORY_ADDRESS = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"

func main() {
	log.Println("Hello WEB3 world!")

	client := uniinfo.GetClient()
	factory := uniinfo.GetUniswapFactory(UNISWAP_FACTORY_ADDRESS, client)

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
		pair := uniinfo.GetUniswapPair(pairAddress, client)

		reserves, _ := pair.GetReserves(&callOpts)

		tkn0Address, err := pair.Token0(&callOpts)
		if err != nil {
			log.Fatal("not able to get token0 address", err)
		}
		tkn1Address, err := pair.Token1(&callOpts)
		if err != nil {
			log.Fatal("not able to get token1 address", err)
		}
		token0 := uniinfo.GetToken(tkn0Address, client)
		token1 := uniinfo.GetToken(tkn1Address, client)

		token0Name, err := token0.Symbol(&callOpts)
		if err != nil {
			log.Fatal("error getting token name", err)
		}
		token1Name, err := token1.Symbol(&callOpts)
		if err != nil {
			log.Fatal("error getting token name", err)
		}
		reserve0 := big.NewFloat(0).SetInt(reserves.Reserve0)
		reserve1 := big.NewFloat(0).SetInt(reserves.Reserve1)
		pairPrice := big.NewFloat(0).Quo(reserve0, reserve1)
		log.Printf("%v %v-%v reserves: %v, %v, price: %v",
			i,
			token0Name,
			token1Name,
			reserves.Reserve0.String(),
			reserves.Reserve1.String(),
			pairPrice.String(),
		)
	}
}
