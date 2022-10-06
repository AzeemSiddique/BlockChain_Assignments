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
	s.nonce = nonc
	s.pre_hash = pre_h
	non := strconv.Itoa(nonc)
	var cal_hash string = transaction + non + pre_h
	s.hash = CalculateHash(cal_hash)
	endhash=s.hash

	return s
}
func CalculateHash(ds string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(ds)))
}
type BlocktList struct {
	list[]*Block
}

func (ls *BlocktList) Createblock (transaction string, nonc int)*Block{
	length_of_list := len(ls.list)
	pre_hash := ""
	if length_of_list > 0 {
		pre_hash = ls.list[length_of_list-1].hash
	}

	block := NewBlock( len(ls.list), transaction, nonc, pre_hash)
	ls.list = append(ls.list, block)
	return block
}
func (ls *BlocktList) Print(){
	length:=len(ls.list)
	i:=0
	for i < length {
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

	ls.list[num].transact = transaction

	for i:=num;i<size;i++{	
		t:=ls.list[i].transact
		n:=ls.list[i].nonce
		prevh:=ls.list[i-1].hash

		non := strconv.Itoa(n)
		var cal_hash string = t + non + prevh
		ls.list[i].hash = CalculateHash(cal_hash)
	}


}
func (ls *BlocktList) Chain_Verification() {
	i:=len(ls.list)-1
	if ls.list[i].hash==endhash {
		fmt.Println("\n\tChain Is Valid!")
	} else {
		fmt.Println("\n\tChain Is Not Valid!\n\tError in Block Number ", i)
	}

}

// func (ls *StudentList) Print(){
func main() {
	student := new( BlocktList)
	student.Createblock("Haji Ibrar 2 rr", 233)
	student.Createblock ("Mola Jutt2 ss" , 432)
	student.Createblock ("gd ss ss" , 56)
	student.Print()

	student.ChangeBlock(1,"AHAHAHAHAH 2 PhukYuu")
	student.Print()

	student.Chain_Verification()
	
}