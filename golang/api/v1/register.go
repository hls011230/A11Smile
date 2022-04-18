package v1

import (
	"A11Smile/db"
	"A11Smile/db/model"
	"encoding/hex"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func Register(user *model.User) error {
	addr, pk, _ := CreateAccount()
	user.BlockAddress = addr
	user.PrivateKey = pk

	cli := db.Get()
	cli.Table("users").Save(user)

	return nil
}

func CreateAccount() (addr string, pk string, err error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalln(err)
		return "", "", nil
	}

	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	log.Println("address: ", address)

	privateKey := hex.EncodeToString(key.D.Bytes())
	log.Println("privateKey: ", privateKey)
	return address, privateKey, nil
}
