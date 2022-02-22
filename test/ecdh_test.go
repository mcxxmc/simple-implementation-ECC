package test

import (
	"fmt"
	"github.com/mcxxmc/simple-implementation-ECC/ecc"
	"math/rand"
	"testing"
	"time"
)

const loop = 100

func TestECDH(t *testing.T) {
	ep := ecc.SampleElliptic()
	ep.SetGeneratorPoint(15, 13)
	rand.Seed(time.Now().Unix())
	alice := ecc.NewInstanceECDH(ep)
	bob := ecc.NewInstanceECDH(ep)

	//ecc.SetDebug(true)

	for i := 0; i < loop; i ++ {
		alice.RandomlyPicksPrivateKey()
		bob.RandomlyPicksPrivateKey()
		publicKeyFromAlice := alice.PublicKey()
		publicKeyFromBob := bob.PublicKey()
		msg := "this is a message."
		encrypted, err := alice.Encrypt(msg, publicKeyFromBob)

		info := func() {
			fmt.Println("loop number: ", i)
			fmt.Println("alice private key: ", alice.PrivateKey)
			fmt.Println("bob private key: ", bob.PrivateKey)
			fmt.Println("public key from alice: ", publicKeyFromAlice)
			fmt.Println("public key from bob: ", publicKeyFromBob)
			if encrypted != nil {
				fmt.Println("ciphertext: ", encrypted.Ciphertext)
				fmt.Println("nonce: ", encrypted.Nonce)
				fmt.Println("ciphertext public key: ", encrypted.CiphertextPubKey)
			}
		}

		if err != nil {
			info()
			t.Error(err)
			return
		}
		decrypted, err := bob.Decrypt(encrypted)
		if err != nil {
			info()
			t.Error(err)
			return
		}
		if decrypted != msg {
			info()
			fmt.Println(decrypted)
			t.Error("wrong decryption result")
			return
		}
	}

	//ecc.SetDebug(false)
}
