package main

import (
	"fmt"
	"github.com/AlayaNetwork/Alaya-Go/accounts/abi/bind"
	"log"
	"platon-go-sdk/contracts"
	"platon-go-sdk/examples/store"
)

const AlayaEndpoint = "http://172.16.64.132:6789"
const privateKey = "ed72066fa30607420635be56785595ccf935675a890bef7c808afc1537f52281"

type StoreContract struct{}

func (sc StoreContract) DoDeploy(opts *bind.TransactOpts, client bind.ContractBackend) {
	input := "1.0"
	address, tx, instance, err := store.DeployStore(opts, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0

	_ = instance
}

func main() {
	store := StoreContract{}

	contract := contracts.Contract{
		Url:          AlayaEndpoint,
		PrivateKey:   privateKey,
		ContractAddr: "",
	}

	err := contract.Deploy(store)

	fmt.Println("deploy error : ", err)
}
