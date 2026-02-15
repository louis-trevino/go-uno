/*
 * BinTree.go
 * Custom Binary Tree
 *
 * Author: Louis Trevino
 * Copyright(C) Torino Consulting, 2020.
 *
 * Compiled and tested using Go version go1.13.8 windows/amd64
 */
package uno

import (
	"fmt"
)

type Comparable interface {
	CompareTo(obj Comparable) int
	GetNumericValue() float64
	GeStringValue() string
}

type BinNode struct {
	Data  Comparable
	Left  *BinNode
	Right *BinNode
}

type BinTree struct {
	Root *BinNode
}

type Floatie struct {
	Value float64
}

func (num Floatie) GetNumericValue() float64 {
	return num.Value
}

func (num Floatie) GeStringValue() string {
	var fmtVal string = fmt.Sprintf("%.2f", num.Value)
	return fmtVal
}

func (num Floatie) CompareTo(otherNum Comparable) int {
	if num.Value > otherNum.GetNumericValue() {
		return 1
	} else if num.Value < otherNum.GetNumericValue() {
		return -1
	} else {
		return 0
	}
}

func NewBinTree(initialValues []Comparable) *BinTree {
	binTree := &BinTree{}
	if initialValues == nil {
		return binTree
	}
	for _, iData := range initialValues {
		//fmt.Printf("Adding %v \n", iData)
		binTree.Insert(iData)
	}
	return binTree
}

func (inst *BinTree) InsertNode(iNode *BinNode, key Comparable) *BinNode {
	if iNode == nil {
		iNode = &BinNode{key, nil, nil}
		if inst.Root == nil {
			inst.Root = iNode
		}
		return iNode
	}
	if key.CompareTo(iNode.Data) < 0 {
		iNode.Left = inst.InsertNode(iNode.Left, key)
	} else if key.CompareTo(iNode.Data) > 0 {
		iNode.Right = inst.InsertNode(iNode.Right, key)
	}
	return iNode
}

func (inst *BinTree) Insert(data Comparable) *BinNode {
	var node *BinNode = inst.InsertNode(inst.Root, data)
	return node
}

func (inst *BinTree) Inorder() string {
	var sb string = ""
	inst.InorderNode(inst.Root, &sb)
	return sb
}

func (inst *BinTree) ToString() string {
	return inst.Inorder()
}

func (inst *BinTree) InorderNode(iNode *BinNode, sb *string) {
	if iNode != nil {
		inst.InorderNode(iNode.Left, sb)
		var sep string = ", "
		if len(*sb) == 0 {
			sep = ""
		}
		var fmtData string = iNode.Data.GeStringValue() // fmt.Sprintf("%.2f", iNode.Data)
		// fmt.Printf("+ %s \n", fmtData)
		*sb = fmt.Sprintf("%s%s%s", *sb, sep, fmtData)
		inst.InorderNode(iNode.Right, sb)
	}
}

func (inst *BinTree) SearchNode(data Comparable, iNode *BinNode) *BinNode {
	if iNode == nil {
		fmt.Println("iNode should be populated.")
		return nil
	}
	if data.CompareTo(iNode.Data) == 0 {
		return iNode
	}
	if data.CompareTo(iNode.Data) < 0 {
		return inst.SearchNode(data, iNode.Left)
	}
	if data.CompareTo(iNode.Data) > 0 {
		return inst.SearchNode(data, iNode.Right)
	}
	return nil
}

func (inst *BinTree) SearchNodeData(data Comparable) *BinNode {
	var binNode *BinNode = inst.SearchNode(data, inst.Root)
	return binNode
}

func GetFloatieArray(floatArr []float64) []Comparable {
	var compArr []Comparable = make([]Comparable, len(floatArr))
	for idx, val := range floatArr {
		floatie := Floatie{Value: val}
		compArr[idx] = Comparable(floatie)
	}
	return compArr
}

func BinTreeDemo() {
	fmt.Printf("\n* Demo of uno.BinTree (custom binary tree) \n")
	var initFloatValues = []float64{-5, 54, 36, 27, 20, 11, 7, 5}
	fmt.Printf("Initial non-sorted numeric values: %v \n", initFloatValues)
	var initValues []Comparable = GetFloatieArray(initFloatValues)
	binTree := NewBinTree(initValues)
	fmt.Printf("BinTree: %v \n", binTree.ToString())
	var searchData Comparable = Floatie{Value: 27}
	var found *BinNode = binTree.SearchNodeData(searchData)
	fmt.Printf("Searching: %s, found: %s \n", searchData.GeStringValue(), found.Data.GeStringValue())
}
