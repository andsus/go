package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

//PrivateKey generates random key between 1 and p -- prime number
func PrivateKey(p *big.Int) *big.Int {
	start := big.NewInt(2)
	private, _ := rand.Int(rand.Reader, new(big.Int).Sub(p, start))
	return private.Add(private, start)
}

//PublicKey generates public key
func PublicKey(private, p *big.Int, g int64) *big.Int {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

//NewPair generates pair of keys
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

//SecretKey func
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
