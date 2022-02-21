package ecc

import (
	"fmt"
	"math/rand"
)

// EncryptedMsg the encrypted message with ciphertext, nonce and ciphertext public key. // todo: auth tag
type EncryptedMsg struct {
	Ciphertext       string // the encrypted message
	Nonce            []byte // should have a length of 12
	CiphertextPubKey [2]int // the public key used for encryption
}

// InstanceECDH an implementation of a hybrid encryption scheme by using the ECDH (Elliptic Curve Diffie–Hellman) key exchange scheme
type InstanceECDH struct {
	Ep 			*Elliptic
	PrivateKey 	int
}

// NewInstanceECDH returns a pointer to a new InstanceECDH object.
//
// The input ep should already be assigned a generator pointer.
func NewInstanceECDH(ep *Elliptic) *InstanceECDH {
	return &InstanceECDH{Ep: ep}
}

// RandomlyPicksPrivateKey randomly picks a private key.
func (ecdh *InstanceECDH) RandomlyPicksPrivateKey() {
	ecdh.PrivateKey = rand.Intn(ecdh.Ep.P - 1) + 1		// 0 < k <= p
}

func (ecdh *InstanceECDH) PublicKey() (publicKey [2]int) {
	publicKey = Generate(ecdh.PrivateKey, ecdh.Ep)
	return
}

// generates and returns a shared ECC key and a ciphertext public key
// using a received public key.
//
// Use the sharedECCKey for symmetric encryption.
// Use the randomly generated ciphertextPubKey to calculate the decryption key later.
func (ecdh *InstanceECDH) calculateEncryptionKey(publicKey [2]int)  (sharedECCKey [2]int, ciphertextPubKey [2]int) {
	ciphertextPrivateKey := rand.Intn(ecdh.Ep.P - 1) + 1		// 0 < k <= p
	// shared ecc key = random k` * public key from the other party = random k` * G * private key of the other party,
	// since they agree on the same elliptic curve
	sharedECCKey = Calculate(publicKey, ciphertextPrivateKey, ecdh.Ep)
	ciphertextPubKey = Generate(ciphertextPrivateKey, ecdh.Ep)
	return
}

// returns the sharedECCKey which is used for the decryption.
func (ecdh *InstanceECDH) calculateDecryptionKey(ciphertextPubKey [2]int) (sharedECCKey [2]int) {
	// here, shared ecc key = random k` from the other party * G * private key of this party
	sharedECCKey = Calculate(ciphertextPubKey, ecdh.PrivateKey, ecdh.Ep)
	return
}

// Encrypt encrypts a message by a given public key and returns the ciphertext and the nonce.
func (ecdh *InstanceECDH) Encrypt(msg string, publicKey [2]int) (encryptedMsg *EncryptedMsg, err error) {
	sharedECCKey, ciphertextPubKey := ecdh.calculateEncryptionKey(publicKey)
	// the secret key is from the shared ecc key, so it is the same for both the sender and the receiver
	secretKey := PointTo256bit(StringifyPublicKey(sharedECCKey))
	ciphertext, nonce, err := EncryptGCM(msg, secretKey)
	encryptedMsg = &EncryptedMsg{
		Ciphertext:       ciphertext,
		Nonce:            nonce,
		CiphertextPubKey: ciphertextPubKey,
	}
	return
}

// Decrypt decrypts a message and returns the decrypted message.
func (ecdh *InstanceECDH) Decrypt(encryptedMsg *EncryptedMsg) (msg string, err error) {
	ciphertext, nonce, ciphertextPubKey := encryptedMsg.Ciphertext, encryptedMsg.Nonce, encryptedMsg.CiphertextPubKey
	sharedECCKey := ecdh.calculateDecryptionKey(ciphertextPubKey)
	secretKey := PointTo256bit(StringifyPublicKey(sharedECCKey))
	msg, err = DecryptGCM(ciphertext, nonce, secretKey)
	return
}

// StringifyPublicKey converts the key to string.
func StringifyPublicKey(key [2]int) string {
	return fmt.Sprintf("%d,%d", key[0], key[1])
}
