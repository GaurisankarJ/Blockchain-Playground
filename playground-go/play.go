package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/joho/godotenv"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	Hash      string
	PrevHash  string
}

var Blockchain []Block

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Data + block.PrevHash
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, Data string) (Block, error) {

	var newBlock Block

	t := time.Now()

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = Data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock, nil
}

func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if calculateHash(newBlock) != newBlock.Hash {
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

	io.WriteString(conn, "Enter new Data:")

	scanner := bufio.NewScanner(conn)

	// Take in Data from stdin and add it to the blockchain after conducting necessary validation
	go func() {
		for scanner.Scan() {
			Data := scanner.Text()
			// if err != nil {
			// 	log.Printf("%v not a string: %v", scanner.Text(), err)
			// 	continue
			// }

			newBlock, err := generateBlock(Blockchain[len(Blockchain)-1], Data)
			if err != nil {
				log.Println(err)
				continue
			}

			if isBlockValid(newBlock, Blockchain[len(Blockchain)-1]) {
				newBlockchain := append(Blockchain, newBlock)
				replaceChain(newBlockchain)
			}

			bcServer <- Blockchain
			io.WriteString(conn, "\nEnter new Data:")
		}
	}()

	// Simulate receiving broadcast
	go func() {
		for {
			time.Sleep(30 * time.Second)
			output, err := json.Marshal(Blockchain)
			if err != nil {
				log.Fatal(err)
			}
			io.WriteString(conn, string(output))
		}
	}()

	for _ = range bcServer {
		spew.Dump(Blockchain)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	bcServer = make(chan []Block)

	// Create genesis block
	go func() {
		t := time.Now()
		genesisBlock := Block{0, t.String(), "", "", ""}
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
