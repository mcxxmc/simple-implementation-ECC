package client

import (
	"math/rand"
	"simple-implementation-ECC/ecc"
)

// Msg the message for communication between clients.
type Msg struct {
	PublicKey [2]int			// the public key (not compressed)
}

// Client stands for a party in the ECDH key exchange.
type Client struct {
	Ep 			*ecc.Elliptic	// the elliptic curve
	K 			int				// the private key
	PublicKey 	[2]int			// the generated public key; it is not compressed
	Receive		chan *Msg		// the channel for receiving data
	Send		chan *Msg		// the channel for sending data
}

// NewClient returns a pointer to a new Client object.
//
// It only initializes the Elliptic curve of the client. Please use PickK() and GeneratePublicKey() for private
// and public key-pair.
func NewClient(ep *ecc.Elliptic, receive chan *Msg, send chan *Msg) *Client {
	return &Client{Ep: ep, Receive: receive, Send: send}
}

// PickK picks a valid k as the private key.
func (c *Client) PickK() {
	c.K = rand.Intn(c.Ep.P - 1) + 1		// k should not be 0
}

// GeneratePublicKey generates a public key. Should be called after PickK().
func (c *Client) GeneratePublicKey() {
	x, y := ecc.Generate(c.K, c.Ep)
	c.PublicKey = [2]int{x, y}
}

// SendPublicKey sends the public key to the channel.
func (c *Client) SendPublicKey() {
	c.Send <- &Msg{
		PublicKey: [2]int{c.PublicKey[0], c.PublicKey[1]},
	}
}

// CalculateShareKey calculate the shared key using the public key received from another party.
func (c *Client) CalculateShareKey() (int, int) {
	msg := <- c.Receive
	if msg == nil {
		panic("no public key received")
	}
	return ecc.Calculate(msg.PublicKey[0], msg.PublicKey[1], c.K, c.Ep)
}
