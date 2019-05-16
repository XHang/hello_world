/**
实现一个二叉树结构
功能
1. 可以将数组根据顺序录入二叉树
2. 可以遍历二叉树,从左到右
*/
package main

import (
	"fmt"
)

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func main() {
	var t  Tree
	arr:=[]int{3,5,1,6,8,1,4,3,1,5}
	addValues(&t,arr)
	fmt.Println(Traversing(&t,""))
	//

}

func addValues(node *Tree, arr []int) {
	for _, i := range arr {
		addValue(node, i)
	}
}

/**
将一个值添加到二叉树里
*/
func addValue(node *Tree, value int) *Tree {

	if node == nil {
		//构建这个类型,并且把这个类型的指针分配给node
		node = new(Tree)
		//存放结构体的变量本身就是指针,所以不用对这个指针*运算取值...
		node.Value = value
	} else {
		if node.Value == 0 || node.Value == value {
			node.Value = value
		} else if node.Value > value {
			node.Left = addValue(node.Left, value)
		} else {
			node.Right = addValue(node.Right, value)
		}
	}
	return node
}

/**
遍历二叉树,然后build字符串
*/
func Traversing(tree *Tree, str string) string {
	if tree != nil {
		str = Traversing(tree.Left, str)
		//go语言数字转字符串
		str += fmt.Sprintf("%d", tree.Value)
		str = Traversing(tree.Right, str)
	}
	return str
}
