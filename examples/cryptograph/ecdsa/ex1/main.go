package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"log"
)

func main() {
	pvKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader) //elliptic.p224, elliptic.P384(), elliptic.P521()

	if err != nil {
		log.Println("ECDSA Keypair generation was Fail!")
	}

	pbKey := &pvKey.PublicKey

	fmt.Println("###### Keypair ######")
	fmt.Println("====== Private Key ======\n")
	fmt.Println(("Private Key: %x\n", pvKey.D))
	fmt.Println("========== Public Key(X, Y) =====\n")
	fmt.Println("X=")
	fmt.Println("Hex: X=#{pbKey.")
}
