package test

import (
	"fmt"
	"github.com/mcxxmc/simple-implementation-ecc/ecc"
	"math/rand"
	"testing"
	"time"
)

const loop = 100

func TestECDH1(t *testing.T) {
	rand.Seed(time.Now().Unix())
	alice, bob := prepareECDH(0, 7, 17, 15, 13)
	testECDHByLoop(alice, bob, loop, t)
}

func TestECDH2(t *testing.T) {
	rand.Seed(time.Now().Unix())
	alice, bob := prepareECDH(1, 1, 23, 3, 10)
	testECDHByLoop(alice, bob, loop, t)
}

func prepareECDH(a, b, p ,x, y int) (alice, bob *ecc.InstanceECDH) {
	ep := ecc.NewElliptic(a, b, p)
	ep.SetGeneratorPoint(x, y)
	alice = ecc.NewInstanceECDH(ep)
	bob = ecc.NewInstanceECDH(ep)
	return
}

func testECDHByLoop(alice, bob *ecc.InstanceECDH, loopNum int, t *testing.T) {
	//ecc.SetDebug(true)
	for i := 0; i < loopNum; i ++ {
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
