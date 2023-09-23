package day1

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	fmt.Println(binaryNumber(20))
}

func TestBinaryTreeDp(t *testing.T) {
	fmt.Println(binaryNumberDp(50))
}

func TestBracket(t *testing.T) {
	fmt.Println(bracket("(()())"))
	fmt.Println(bracket("())("))
	fmt.Println(bracket("()(("))
	fmt.Println(bracket(")"))
}
