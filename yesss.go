package assignment01bca

import (
	"crypto/sha256"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Block struct {
	transactions []string
	data         string
	nonce        int
	prev_hash    string
	hash         string
}

type BlockChain struct {
	block_chain  []*Block
	genesisBlock Block
}

func (chain *BlockChain) NewBlock(dat string, nonce_val int) *Block {
	// *TODO*
	prev_h := chain.block_chain[len(chain.block_chain)-1].hash
	hsh := CalculateHash(dat + strconv.Itoa(nonce_val) + prev_h)
	curr := Block{
		data:      dat,
		nonce:     nonce_val,
		prev_hash: prev_h,
		hash:      hsh}

	chain.block_chain = append(chain.block_chain, &curr)
	return &curr
}

func (b *Block) MineBlock(target int) (int, string) {
	//	- Find Nonce value for the block
	//	- Target should be number of trailing zeros in 256 bit output string
	nonce := 0

	for nonce < math.MaxInt64 {
		hsh := CalculateHash(b.data + strconv.Itoa(nonce) + b.prev_hash)
		pref := strings.Repeat("0", target)
		if hsh[:target] == pref {
			return nonce, hsh
		} else {
			nonce += 1
		}
	}
	return nonce, "Unsuccessful Mining"
}

func (chain BlockChain) DisplayBlocks() {
	// *TODO*
	for j, i := range chain.block_chain {
		fmt.Printf("\n____________________")
		fmt.Printf("\n|_____Block No %x_____|", j)
		fmt.Printf("\n| Block Hash %x |", i.hash)
		fmt.Printf("\n| Previous Hash %s |", i.prev_hash)
		fmt.Printf("\n| Nonce %x |", i.nonce)
		fmt.Printf("\n| Data: %s |\n", i.data)
		fmt.Printf("|__________________|\n\n")
	}
}

func (b Block) DisplayMerkelTree() {
	//	- Print all transaction in nice format. Showing transactions and hashes
	for j, i := range b.transactions {
		if j == 0 {
			fmt.Printf("\n|-%s", i)
		} else if j < 3 {
			fmt.Printf("\n   |-%s", i)
		} else if j < 5 {
			fmt.Printf("\n        |-%s", i)
		}
	}
}

func (chain BlockChain) ChangeBlock(i int, d string) {
	//	- Change one or more transactions of the block
	if i >= len(chain.block_chain) {
		println("Index Not In Range")
	} else {
		fmt.Printf("Transaction of block : %s\n", chain.block_chain[i].data)

		chain.block_chain[i].data = d
		fmt.Println("Transaction Updated")
	}

}

func (chain BlockChain) VerifyChain() {
	//	- Verify blockchain for changes.
	//	- Changes to transactions in Merkel Tree

	for j, i := range chain.block_chain {
		hsh := CalculateHash(i.data + strconv.Itoa(i.nonce) + i.prev_hash)
		if i.data == "genesis" {
			continue
		}
		if hsh != i.hash {
			fmt.Printf("\nBlock %x was changed.\n", j)
		}
	}

}

func CalculateHash(data string) string {
	//	- Calculate Hash of transaction or block
	//	- if size of transaction very large, Merkle-Damgard Transform

	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))

}

func createBlockChain() *BlockChain {
	genesis := Block{
		transactions: []string{"genesis"},
		data:         "genesis",
		prev_hash:    "0",
		nonce:        0,
		hash:         fmt.Sprintf("%x", sha256.Sum256([]byte("Genesis"))),
	}
	fmt.Println("Genesis Block Created")
	return &BlockChain{
		block_chain:  []*Block{&genesis},
		genesisBlock: genesis}

}
