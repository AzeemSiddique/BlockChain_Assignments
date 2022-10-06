package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

var endhash=""
type Block struct {
	blockno int
	transact   string
	nonce   int
	pre_hash   string
	hash    string
}

func NewBlock(num int, transaction string, nonc int, pre_h string) *Block {
	s := new(Block)
	s.blockno = num
	s.transact = transaction
	s.nonce = nonc			//nonce
	s.pre_hash = pre_h		//previous hash
	non := strconv.Itoa(nonc)		//nonce to string
	var cal_hash string = transaction + non + pre_h		//calculate hash
	s.hash = CalculateHash(cal_hash)
	endhash=s.hash		//end hash for chain verification

	return s
}
func CalculateHash(ds string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(ds)))		//calculate hash using sha256
}
type BlocktList struct {
	list[]*Block		//list of blocks in blockchain
}

func (ls *BlocktList) Createblock (transaction string, nonc int)*Block{
	length_of_list := len(ls.list)		//length of list of blocks
	pre_hash := ""
	if length_of_list > 0 {
		pre_hash = ls.list[length_of_list-1].hash	//previous hash of last block
	}

	block := NewBlock( len(ls.list), transaction, nonc, pre_hash)		//create new block with previous hash
	ls.list = append(ls.list, block)		//append new block to list of blocks in blockchain
	return block
}
func (ls *BlocktList) Print(){
	length:=len(ls.list)
	i:=0
	for i < length {												//print all blocks in blockchain 
		fmt.Println("\n\n\tBlock Number = ", ls.list[i].blockno)
		fmt.Println("\tTransection  = ", ls.list[i].transact)
		fmt.Println("\tNonce        = ", ls.list[i].nonce)
		fmt.Println("\tPrev Hash    = ", ls.list[i].pre_hash)
		fmt.Println("\tHash         = ", ls.list[i].hash)
		i = i + 1
	}
}
func (ls *BlocktList) ChangeBlock(num int, transaction string) {
	size := len(ls.list)

	ls.list[num].transact = transaction		//change transaction of block number num

	for i:=num;i<size;i++{			//recalculate hash of all blocks after block number num 
		t:=ls.list[i].transact
		n:=ls.list[i].nonce
		prevh:=ls.list[i-1].hash

		non := strconv.Itoa(n)			//nonce to string 
		var cal_hash string = t + non + prevh		//calculate hash 
		ls.list[i].hash = CalculateHash(cal_hash)		//calculate hash using sha256
	}


}
func (ls *BlocktList) Chain_Verification() {
	i:=len(ls.list)-1
	if ls.list[i].hash==endhash {			//check if end hash is same as hash of last block 
		fmt.Println("\n\tThis BlockChain Is Valid!")
	} else {
		fmt.Println("\n\tThis BlockChain Is Not Valid! \n\tError detected in Block Number ", i)
	}

}

// func (ls *StudentList) Print(){
func main() {
	student := new( BlocktList)
	student.Createblock("This is transaction 1", 233)			//create block with transaction and nonce
	student.Createblock ("Transaction 2 is proccessing" , 432)
	student.Createblock ("Third no transaction" , 56)
	student.Print()

	student.ChangeBlock(1,"Change block")		//change transaction of block number 1
	student.Print()

	student.Chain_Verification()		//verify blockchain 
	
}