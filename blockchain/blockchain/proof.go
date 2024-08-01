package Blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"
)

// Proof of Work
// The ideia is secure the blockchain by forcing the network to do work to add a block to the chain.
// The work is just computational power.
// So, if someone wants to tamper the blockchain by altering a block, this person will have to work each block after the tampered one.
// Making it hard to create non valid blocks.
// Meanwhile, it's kind of easy to validate a block. Much less computational power as well.

// "The work must be hard to do, but proving the work must be relative easy."

// In a real Blockchain, an algorithm would slowly increment this in a large period of time.
// Account the number of minors in the network.
// Account the power of computers in general.
// We need to evolve the Difficulty in order for the time to work a block stay the same.
const Difficulty = 18

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	// fmt.Printf("Target before left shifting: %d\n", target)
	target.Lsh(target, uint(256-Difficulty))
	// fmt.Printf("Target after left shifting: %d\n", target)

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) InitData(nounce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nounce)),
			ToHex(int64(Difficulty)),
		},
		[]byte{},
	)

	return data
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)

	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)

	hash := sha256.Sum256(data)

	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Println()

	return nonce, hash[:]
}
