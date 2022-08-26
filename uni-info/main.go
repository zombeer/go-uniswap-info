package uniinfo

import (
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"github.com/zombeer/go-uniswap-info/gen/IERC20"
	"github.com/zombeer/go-uniswap-info/gen/IUniswapV2Factory"
)

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

func GetToken(adress common.Address, client *ethclient.Client) *IERC20.IERC20 {
	token, err := IERC20.NewIERC20(adress, client)
	if err != nil {
		log.Fatal("not able to get token at address", adress.String(), err)
	}
	return token
}
