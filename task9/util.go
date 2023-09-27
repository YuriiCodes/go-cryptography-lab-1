package task9

import (
	"crypto/rand"
	"math/big"
)

type PublicKey struct {
	P, G, H *big.Int
}

type PrivateKey struct {
	PublicKey
	X *big.Int
}

// GenerateKeys generates a public and private key pair.
func GenerateKeys(bits int) (*PublicKey, *PrivateKey, error) {
	// Generate a prime number p
	p, err := rand.Prime(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}

	// Generate a generator g
	g := big.NewInt(2)

	// Generate private key x
	x, err := rand.Int(rand.Reader, p)
	if err != nil {
		return nil, nil, err
	}

	// Compute h = g^x mod p
	h := new(big.Int).Exp(g, x, p)

	return &PublicKey{p, g, h}, &PrivateKey{PublicKey{p, g, h}, x}, nil
}

// Encrypt encrypts the given message using the provided public key.
func (pub *PublicKey) Encrypt(m *big.Int) (*big.Int, *big.Int) {
	// Generate a random y from [1, p-1]
	y, _ := rand.Int(rand.Reader, new(big.Int).Sub(pub.P, big.NewInt(1)))
	y.Add(y, big.NewInt(1))

	// Compute s = h^y mod p
	s := new(big.Int).Exp(pub.H, y, pub.P)

	// Compute c1 = g^y mod p
	c1 := new(big.Int).Exp(pub.G, y, pub.P)

	// Compute c2 = m * s mod p
	c2 := new(big.Int).Mod(new(big.Int).Mul(m, s), pub.P)

	return c1, c2
}

// Decrypt decrypts the given ciphertext using the provided private key.
func (priv *PrivateKey) Decrypt(c1, c2 *big.Int) *big.Int {
	// Compute s = c1^x mod p
	s := new(big.Int).Exp(c1, priv.X, priv.P)

	// Compute sInv = s^-1 mod p
	sInv := new(big.Int).ModInverse(s, priv.P)

	// Compute m = c2 * sInv mod p
	m := new(big.Int).Mod(new(big.Int).Mul(c2, sInv), priv.P)

	return m
}

//func main() {
//	// Generate public and private keys
//	pubKey, privKey, err := GenerateKeys(2048)
//	if err != nil {
//		fmt.Println("Error generating keys:", err)
//		return
//	}
//
//	// The message to be encrypted
//	message := big.NewInt(123456)
//
//	fmt.Println("Original Message :", message)
//
//	// Encrypt the message
//	c1, c2 := pubKey.Encrypt(message)
//
//	fmt.Println("Encrypted Message C1:", c1)
//	fmt.Println("Encrypted Message C2:", c2)
//
//	// Decrypt the message
//	decryptedMessage := privKey.Decrypt(c1, c2)
//
//	fmt.Println("Decrypted Message  :", decryptedMessage)
//}
