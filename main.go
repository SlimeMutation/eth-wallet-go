package main

import (
	"fmt"

	address "github.com/SlimeMutation/eth-wallet-go/address"
)

func main() {
	addressInfo, error := address.CreateAddressFromPrivateKey()
	if error != nil {
		fmt.Println("CreateAddressFromPrivateKey error", error)
		return
	}
	fmt.Println("PublicKey: ", addressInfo.PublicKey)
	fmt.Println("PrivateKey: ", addressInfo.PrivateKey)
	fmt.Println("Address: ", addressInfo.Address)
}
