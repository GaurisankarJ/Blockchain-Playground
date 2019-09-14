package main

// var lNode = `{"id":"ID","proof":"PROOF","signature":{"key":"KEY","time_stamp":"TIME_STAMP"}}`

// fmt.Printf("%T, %v", leaves, leaves)

// MERKLE TREE
// // Build leaves of Content to build tree
// var leaves []merkletree.Content

// for i := 0; i < 10; i++ {
// 	leaves = append(leaves, LeafNode{leaf: strconv.Itoa(i)})
// }

// // Create a new Merkle Tree from the leaves of Content
// tree, err := merkletree.NewTree(leaves)
// if err != nil {
// 	log.Fatal(err)
// }

// // Get the Merkle Root of the tree
// root := tree.MerkleRoot()
// log.Println(root)
// fmt.Printf("%T, %v", root, root)

// // Verify the entire tree (hashes for each node) is valid
// verifyTree, err := tree.VerifyTree()
// if err != nil {
// 	log.Fatal(err)
// }
// log.Println("Verify Tree: ", verifyTree)

// // Verify a specific content in in the tree
// var newLeaf = LeafNode{leaf: "11"}
// verifyContent, err := tree.VerifyContent(newLeaf)
// if err != nil {
// 	log.Fatal(err)
// }

// log.Println("Verify Content: ", verifyContent)

// // String representation
// log.Println(t)

// GENESIS
// var wg sync.WaitGroup
// wg.Add(1)
// // CREATE GENESIS BLOCK
// go func() {
// 	genesisBlock := Block{}
// 	genesisBlock.index = 0
// 	genesisBlock.blockHash = "blockHash"
// 	genesisBlock.prevBlockHash = "prevBlockHash"
// 	genesisBlock.timeStamp = time.Now().String()

// 	// CREATE MERKLE LEAF
// 	genesisBlock.merkleLeaf = make([]LeafNode, 10)
// 	for i := 0; i < 10; i++ {
// 		genesisBlock.merkleLeaf[i].leaf = lNode
// 	}

// 	genesisBlock.merkleRoot = "merkleRoot"

// 	spew.Dump(genesisBlock)

// 	Blockchain = append(Blockchain, genesisBlock)

// 	wg.Done()
// }()

// wg.Wait()
