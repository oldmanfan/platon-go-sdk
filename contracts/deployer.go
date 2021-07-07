package contracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/AlayaNetwork/Alaya-Go/accounts/abi/bind"
	"github.com/AlayaNetwork/Alaya-Go/crypto"
	"github.com/AlayaNetwork/Alaya-Go/ethclient"
	"math/big"
)

type Contract struct {
	Url        string
	PrivateKey string

	ContractAddr string
}

type Deployer interface {
	DoDeploy(opts *bind.TransactOpts, client bind.ContractBackend)
}

func (c *Contract) Deploy(deployer Deployer) error {
	client, err := ethclient.Dial(c.Url)
	if err != nil {
		return err
	}

	privateKey, err := crypto.HexToECDSA(c.PrivateKey)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	deployer.DoDeploy(auth, client)
	return nil
}
