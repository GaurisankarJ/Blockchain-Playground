// JLL
// Land Records Management System

// {
// 	index			int
// 	blockHash 		string
// 	prevBlockHash	string
// 	timeStamp		string
// 	merkleLeaf	string
// 	merkleRoot	string
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
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"github.com/cbergoon/merkletree"
	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

type Block struct {
	index         int
	blockHash     string
	prevBlockHash string
	timeStamp     string
	merkleLeaf    []LeafNode
	merkleRoot    string
}

type LeafNode struct {
	leaf string
}

var Blockchain []Block

// ############################################################################
// ############################################################################
// merkletree
// ############################################################################
// CalculateHash hashes the value of a LeafNode
func (leaf LeafNode) CalculateHash() ([]byte, error) {
	hash := sha256.New()
	if _, err := hash.Write([]byte(leaf.leaf)); err != nil {
		return nil, err
	}

	return hash.Sum(nil), nil
}

// Equals tests for equality of two Contents
func (leaf LeafNode) Equals(other merkletree.Content) (bool, error) {
	return leaf.leaf == other.(LeafNode).leaf, nil
}

// ############################################################################
// ############################################################################

func calculateBlockHash(block Block) string {
	plainText := string(block.index) + block.prevBlockHash + block.timeStamp + block.merkleRoot
	hash := sha256.New()
	hash.Write([]byte(plainText))
	hashed := hash.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, leafNodes []LeafNode) (Block, error) {
	var newBlock Block

	newBlock.index = oldBlock.index + 1
	newBlock.prevBlockHash = oldBlock.blockHash
	newBlock.timeStamp = time.Now().String()
	newBlock.merkleLeaf = leafNodes

	// MERKLE ROOT
	// Build leaves of Content to build tree
	var leaves []merkletree.Content

	for i := 0; i < 5; i++ {
		leaves = append(leaves, LeafNode{leaf: leafNodes[i].leaf})
	}

	// Create a new Merkle Tree from the leaves of Content
	tree, err := merkletree.NewTree(leaves)
	if err != nil {
		log.Fatal(err)
	}

	// Get the Merkle Root of the tree
	root := tree.MerkleRoot()
	log.Println(root)

	newBlock.merkleRoot = string(root)
	newBlock.blockHash = calculateBlockHash(newBlock)

	return newBlock, nil
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.index+1 != newBlock.index {
		return false
	}

	if oldBlock.blockHash != newBlock.prevBlockHash {
		return false
	}

	if calculateBlockHash(newBlock) != newBlock.blockHash {
		return false
	}

	return true
}

func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

var bcServer chan []Block

func handleConn(conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn, "Enter 5 space seperated LEAVES: ")

	scanner := bufio.NewScanner(conn)

	// Take input from stdin and add it to the blockchain after validation
	go func() {
		for scanner.Scan() {
			var leaves []LeafNode

			input := scanner.Text()
			leafNodes := strings.Split(input, " ")

			if len(leafNodes) != 5 {
				continue
			}

			for i := 0; i < 5; i++ {
				leaves = append(leaves, LeafNode{leaf: leafNodes[i]})
				fmt.Println(leaves)
			}

			newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], leaves)
			if err != nil {
				log.Println(err)
				continue
			}

			if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
				newBlockchain := append(Blockchain, newBlock)
				replaceChain(newBlockchain)
			}

			bcServer <- Blockchain
			io.WriteString(conn, "Enter 5 space seperated LEAVES: ")
		}
	}()

	// // Simulate receiving broadcast
	// go func() {
	// 	for {
	// 		time.Sleep(30 * time.Second)
	// 		output, err := json.Marshal(Blockchain)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		io.WriteString(conn, string(output))
	// 	}
	// }()

	// for _ = range bcServer {
	// 	spew.Dump(Blockchain)
	// }
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	bcServer = make(chan []Block)

	// CREATE GENESIS BLOCK
	go func() {
		genesisBlock := Block{}
		genesisBlock.index = 0
		genesisBlock.blockHash = "blockHash"
		genesisBlock.prevBlockHash = "prevBlockHash"
		genesisBlock.timeStamp = time.Now().String()
		// CREATE MERKLE LEAF
		genesisBlock.merkleLeaf = make([]LeafNode, 10)
		genesisBlock.merkleRoot = "merkleRoot"

		spew.Dump(genesisBlock)

		Blockchain = append(Blockchain, genesisBlock)
	}()

	// Start TCP and serve TCP server
	server, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConn(conn)
	}
}
