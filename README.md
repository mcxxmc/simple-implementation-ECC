# simple-implementation-ecc


 A simple implementation for Elliptic Curve Cryptography. 
 

It implements the algorithm of 2D points over a finite field
(e.g., addition, doubling, multiplication and inverse),
the elliptic curve (in the form: y^2 = x^3 + ax + b), and the ECDH
(Elliptic Curve Diffie–Hellman) key exchange algorithm, including 
test cases for verification.

Generally speaking, it consists of the following 3 parts: 
the calculations over finite field, the elliptic curve and the 
ECDH key exchange.


I learnt many of concepts regarding ECC and ECDH here:


https://cryptobook.nakov.com/asymmetric-key-ciphers/elliptic-curve-cryptography-ecc


### Calculations Over Finite Field
This is in the "galois" package as the finite field is also called "galois field".


The implementation are inspired by the online discussions here:


https://math.stackexchange.com/questions/67171/calculating-the-modular-multiplicative-inverse-without-all-those-strange-looking/67190#67190


https://math.stackexchange.com/questions/1340484/addition-of-points-on-elliptic-curves-over-a-finite-field


Note that, different from my common practice, the "Point" object is passed as a struct
instead of a pointer. This will cause higher memory cost, but it simulates the ECC 
process in a better way, as public keys are distributed across the net as real copies 
(a stream of 0s and 1s) instead of a shared reference (usually 8-byte) to a single being.


### The Elliptic Curve


"elliptic.go" defines a struct for the elliptic curve. Most of its functions are 
based on "galois-field.go" in the "galois" package.


### ECDH Key Exchange


Implements the encryption and decryption process of the ECDH key exchange, using
the GCM cipher.


Including files "ecdh.go" and "gcm.go".


### Testing


The testing is based on the curve y^2 = x^3 + 0x + 7 (mod 17), with a generator
point (15, 13). This is a very small sample used only for testing. Details of this curve
can be found in the first link. 


Another testcase is y^2 = x^3 + 1 x + 1 (mod 23), with a generator point (3, 10),
which is slightly bigger.


Example of how to use my codes to simulate the ECDH key exchange process can be
found in the test file "ecdh_test.go".


### TODO


1. Better coding in "galois".
2. Currently, my codes assume all integers are at most 64 bits and there is 
no overflow during the process; however, in the real world most keys are 256+
bits. Since my codes are just experimental and not intended for production use,
this may not be a concern; however, consider an update in the future.
3. Implement ECIES (Elliptic Curve Integrated Encryption Scheme).
4. Introduce auth tag for the encryption / decryption process to fully simulate
ECDH and ECIES.


