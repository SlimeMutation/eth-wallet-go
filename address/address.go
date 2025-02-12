package addresss

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type EthAddress struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	Address    string `json:"address"`
}

func CreateAddressFromPrivateKey() (*EthAddress, error) {
	privateKey, error := crypto.GenerateKey()
	if error != nil {
		return nil, error
	}
	address := &EthAddress{
		PrivateKey: hex.EncodeToString(crypto.FromECDSA(privateKey)),
		PublicKey:  hex.EncodeToString(crypto.FromECDSAPub(&privateKey.PublicKey)),
		Address:    crypto.PubkeyToAddress(privateKey.PublicKey).String(),
	}
	return address, nil
}

func PubkeyToAddress(publicKeyHex string) (string, error) {
	publicKeyBytes, error := hex.DecodeString(publicKeyHex)
	if error != nil {
		return "", error
	}
	return common.BytesToAddress(crypto.Keccak256(publicKeyBytes[1:])[12:]).String(), nil
}
