package addresss

import (
	"fmt"
	"testing"
)

/*
9199b2e172d34219d73b4fbd9f766d8a1602f8c87d3436f2df75f9fc4f2a2107
04e2be8cce5770d9a21f77f0b2d5456b43b07bf4bfa714fb9c6d840809b584fa2a5677b19b3d326a20d4c0c5cd99e228b943f6c0f79e757c66449fee1d3d92d9b7
0xB570A993f916B72fE98ac6199178d098910f1364
*/
func TestCreateAddressFromPrivateKey(t *testing.T) {
	address, err := CreateAddressFromPrivateKey()
	if err != nil {
		return
	}
	fmt.Println(address.PrivateKey, address.PublicKey, address.Address)
}

// 0xB570A993f916B72fE98ac6199178d098910f1364
func TestPubkeyToAddress(t *testing.T) {
	address, err := PubkeyToAddress("04e2be8cce5770d9a21f77f0b2d5456b43b07bf4bfa714fb9c6d840809b584fa2a5677b19b3d326a20d4c0c5cd99e228b943f6c0f79e757c66449fee1d3d92d9b7")
	if err != nil {
		return
	}
	fmt.Println(address)
}
