package table

import (
	"strings"
)

type Generator interface {
	newTable(string) EncodingTable
}

type EncodingTable map[rune]string
type decodingTree struct {
	Value string
	Zero  *decodingTree
	One   *decodingTree
}

func (et EncodingTable) decodingTree() decodingTree {

	res := decodingTree{}

	for ch, code := range et {
		res.add(code, ch)
	}
	return res
}

func (dt *decodingTree) add(code string, value rune) {

	currNode := dt

	for _, ch := range code {

		switch ch {
		case '0':
			if currNode.Zero == nil {
				currNode.Zero = &decodingTree{}
			}
			currNode = currNode.Zero
		case '1':
			if currNode.One == nil {
				currNode.One = &decodingTree{}
			}
			currNode = currNode.One
		}
	}
	currNode.Value = string(value)
}

func (dt *decodingTree) Decode(str string) string {

	var buff strings.Builder
	currNode := dt

	for _, ch := range str {

		if currNode.Value != "" {
			buff.WriteString(currNode.Value)
			currNode = dt
		}

		switch ch {
		case '0':
			currNode = currNode.Zero
		case '1':
			currNode = currNode.One
		}

		if currNode.Value != "" {
			buff.WriteString(currNode.Value)
			currNode = dt
		}
	}

	return buff.String()
}
