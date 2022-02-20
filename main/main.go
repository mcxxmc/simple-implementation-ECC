package main

import (
	"fmt"
	"simple-implementation-ECC/client"
	"simple-implementation-ECC/ecc"
)

const defaultChannelSize = 8

func main() {
	fmt.Println("Start simulation.")
	// use channel as the "Internet"
	Alice2Bob := make(chan *client.Msg, defaultChannelSize)
	Bob2Alice := make(chan *client.Msg, defaultChannelSize)

	// Alice and Bob both agree on the same elliptic curve
	ep := ecc.SampleElliptic()

	Alice := client.NewClient(ep, Bob2Alice, Alice2Bob)
	Bob := client.NewClient(ep, Alice2Bob, Bob2Alice)

	Alice.PickK()
	Alice.GeneratePublicKey()
	fmt.Println("Alice private key: ", Alice.K, ", public key: ", Alice.PublicKey)

	Bob.PickK()
	Bob.GeneratePublicKey()
	fmt.Println("Bob private key: ", Bob.K, ", public key: ", Bob.PublicKey)

	Alice.SendPublicKey()
	xb, yb := Bob.CalculateShareKey()
	fmt.Println("Bob shared key: ", xb, ", ", yb)

	Bob.SendPublicKey()
	xa, ya := Alice.CalculateShareKey()
	fmt.Println("Alice share key: ", xa, ", ", ya)

	if xa != xb || ya != yb {
		fmt.Println("ERROR: shared key not equal!")
		return
	}
}
