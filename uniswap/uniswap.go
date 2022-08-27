package uniswap

import (
	"context"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/zombeer/go-uniswap-info/gen/IERC20"
	"github.com/zombeer/go-uniswap-info/gen/IUniswapV2Factory"
	"github.com/zombeer/go-uniswap-info/gen/IUniswapV2Pair"
)

type Reserves struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}

type Token struct {
	Adress common.Address
	Symbol string
	Name   string
}

type Pair struct {
	Adress   common.Address
	Reserves Reserves
	Token0   Token
	Token1   Token
	Price    *big.Float
}

func GetPairDetails(pair *IUniswapV2Pair.IUniswapV2Pair, blockNumber *big.Int, client *ethclient.Client) *Pair {
	result := Pair{}

	callOpts := &bind.CallOpts{
		Context:     context.Background(),
		BlockNumber: blockNumber,
	}

	r, _ := pair.GetReserves(callOpts)
	result.Reserves = r

	tkn0Address, err := pair.Token0(callOpts)
	if err != nil {
		log.Fatal("not able to get token0 address", err)
	}
	tkn1Address, err := pair.Token1(callOpts)
	if err != nil {
		log.Fatal("not able to get token1 address", err)
	}
	token0 := GetToken(tkn0Address, client)
	token1 := GetToken(tkn1Address, client)

	symbol0, _ := token0.Symbol(callOpts)
	name0, _ := token0.Name(callOpts)
	symbol1, _ := token1.Symbol(callOpts)
	name1, _ := token0.Name(callOpts)

	result.Token0 = Token{
		Adress: tkn0Address,
		Symbol: symbol0,
		Name:   name0,
	}

	result.Token1 = Token{
		Adress: tkn1Address,
		Name:   name1,
		Symbol: symbol1,
	}

	reserve0 := big.NewFloat(0).SetInt(result.Reserves.Reserve0)
	reserve1 := big.NewFloat(0).SetInt(result.Reserves.Reserve1)
	result.Price = big.NewFloat(0).Quo(reserve1, reserve0)
	return &result
}

func GetClient() *ethclient.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file", err)
	}
	clientUrl := os.Getenv("CLIENT_URL")
	client, err := ethclient.Dial(clientUrl)
	if err != nil {
		panic(err)
	}
	return client
}

func GetUniswapFactory(address string, client *ethclient.Client) *IUniswapV2Factory.IUniswapV2Factory {
	factoryAdress := common.HexToAddress(address)
	factory, err := IUniswapV2Factory.NewIUniswapV2Factory(factoryAdress, client)
	if err != nil {
		log.Fatal("not able to create factory instance", err)
	}
	return factory
}

func GetUniswapPair(address common.Address, client *ethclient.Client) *IUniswapV2Pair.IUniswapV2Pair {
	pair, err := IUniswapV2Pair.NewIUniswapV2Pair(address, client)
	if err != nil {
		log.Fatal("not able to get uniswap pair at address", address.String(), err)
	}
	return pair
}

func GetToken(adress common.Address, client *ethclient.Client) *IERC20.IERC20 {
	token, err := IERC20.NewIERC20(adress, client)
	if err != nil {
		log.Fatal("not able to get token at address", adress.String(), err)
	}
	return token
}
