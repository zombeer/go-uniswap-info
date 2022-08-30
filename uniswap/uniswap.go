package uniswap

import (
	"context"
	"fmt"
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

const UNISWAP_FACTORY_ADDRESS = "0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
const PAIR_ADRESS = "0xA478c2975Ab1Ea89e8196811F51A7B7Ade33eB11"

func GetCallOptsAtBlock(bn int64) bind.CallOpts {
	return bind.CallOpts{
		Context:     context.Background(),
		BlockNumber: big.NewInt(bn),
	}
}

type Price struct {
	BlockNumber uint64  `json:"block_number"`
	Value       float64 `json:"value"`
}

type TokenInfo struct {
	Address string `json:"address"`
	Symbol  string `json:"symbol"`
}

type PairInfo struct {
	Address string    `json:"pair_address"`
	Symbols string    `json:"symbols"`
	Token0  TokenInfo `json:"token0"`
	Token1  TokenInfo `json:"token1"`
	Prices  []Price   `json:"prices"`
}

// Getting pair info by its hex address string representation
func GetPairInfo(address string) PairInfo {
	result := PairInfo{
		Address: address,
	}
	client := GetClient()
	currentBn, _ := client.BlockNumber(context.Background())
	callOpts := GetCallOptsAtBlock(int64(currentBn))
	pairAddress := common.HexToAddress(address)

	pair := GetPair(pairAddress, client)
	tkn0Address, _ := pair.Token0(&callOpts)
	tkn1Address, _ := pair.Token1(&callOpts)

	tkn0 := GetToken(tkn0Address, client)
	tkn1 := GetToken(tkn1Address, client)

	tkn0Symbol, _ := tkn0.Symbol(&callOpts)
	tkn1Symbol, _ := tkn1.Symbol(&callOpts)

	result.Symbols = fmt.Sprintf("%s-%s", tkn0Symbol, tkn1Symbol)
	result.Token0 = TokenInfo{Address: tkn0Address.String(), Symbol: tkn0Symbol}
	result.Token1 = TokenInfo{Address: tkn1Address.String(), Symbol: tkn1Symbol}
	return result
}

func GetPairPrices(pairAddress string) {}

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

func GetFactory(address string, client *ethclient.Client) *IUniswapV2Factory.IUniswapV2Factory {
	factoryAdress := common.HexToAddress(address)
	factory, err := IUniswapV2Factory.NewIUniswapV2Factory(factoryAdress, client)
	if err != nil {
		log.Fatal("not able to create factory instance", err)
	}
	return factory
}

func GetPair(address common.Address, client *ethclient.Client) *IUniswapV2Pair.IUniswapV2Pair {
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

// GetPairPriceAtBlock
