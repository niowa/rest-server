package ethereum

import (
	"github.com/onrik/ethrpc"
	"log"
	"fmt"
)

func ConnectToEthereum()  {
	client := ethrpc.New("http://127.0.0.1:8545")

	version, err := client.Web3ClientVersion()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(version)
}
