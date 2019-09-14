// JLL
// Land Records Management System

// {
// 	index			int
// 	blockHash 		string
// 	prevBlockHash	string
// 	timeStamp		string
// 	merkleTreeLeaf	string
// 	merkleTreeRoot	string
// }

// {
// 	id: 	STRING
// 	proof:	STRING
// 	signature: 	[{
// 		key:			STRING,
// 		time_stamp: 	STRING
// 	}]
// }

package main

import (
	"fmt"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type Block struct {
	index          int
	blockHash      string
	prevBlockHash  string
	timeStamp      string
	merkleTreeLeaf []LeafNode
	merkleTreeRoot string
}

type LeafNode struct {
	leaf string
}

var lNode = `{"id":"ID","proof":"PROOF","signature":{"key":"KEY","time_stamp":"TIME_STAMP"}}`

var Blockchain []Block

func main() {

	// // CREATE GENESIS BLOCK
	// go func() {
	// 	genesisBlock := Block{}
	// 	genesisBlock.index = 0
	// 	genesisBlock.blockHash = "blockHash"
	// 	genesisBlock.prevBlockHash = "prevBlockHash"
	// 	genesisBlock.timeStamp = time.Now().String()

	// 	// for i := 0; i < 5; i++ {
	// 	// 	genesisBlock.merkleTreeLeaf[i].leaf = lNode
	// 	// }

	// 	genesisBlock.merkleTreeRoot = "merkleTreeRoot"

	// 	spew.Dump(genesisBlock)

	// 	// Blockchain = append(Blockchain, genesisBlock)

	// 	fmt.Println(genesisBlock)
	// }()

	genesisBlock := Block{}
	genesisBlock.index = 0
	genesisBlock.blockHash = "blockHash"
	genesisBlock.prevBlockHash = "prevBlockHash"
	genesisBlock.timeStamp = time.Now().String()

	// for i := 0; i < 5; i++ {
	// 	genesisBlock.merkleTreeLeaf[0].leaf = lNode
	// }

	genesisBlock.merkleTreeRoot = "merkleTreeRoot"

	spew.Dump(genesisBlock)

	// Blockchain = append(Blockchain, genesisBlock)

	fmt.Println(1 << 10)
}
