package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

type Node struct {
	value  int
	left   *Node
	right  *Node
	parent *Node
}

func NewNode(value int, parent *Node) *Node {
	return &Node{
		value:  value,
		parent: parent,
	}
}

func swapWithLeftChild(p, v *Node, nodeMap map[int]*Node) {
	if v.right != nil {
		v.right.parent = p
	}
	if p.right != nil {
		p.right.parent = v
	}

	p.value, v.value = v.value, p.value
	p.right, v.right = v.right, p.right
	nodeMap[v.value] = v
	nodeMap[p.value] = p
}

func swapWithRightChild(p, v *Node, nodeMap map[int]*Node) {
	if v.left != nil {
		v.left.parent = p
	}
	if p.left != nil {
		p.left.parent = v
	}

	p.value, v.value = v.value, p.value
	p.left, v.left = v.left, p.left
	nodeMap[v.value] = v
	nodeMap[p.value] = p
}

func swap(p, v *Node, nodeMap map[int]*Node) {
	if p.left == v {
		swapWithLeftChild(p, v, nodeMap)
	} else {
		swapWithRightChild(p, v, nodeMap)
	}
}

func buildTree(n int) (root *Node, nodeMap map[int]*Node) {
	root = NewNode(1, nil)
	queue := list.New()
	queue.PushBack(root)
	currentNodeId := 2
	nodeMap = make(map[int]*Node, 0)
	nodeMap[1] = root

	for queue.Len() > 0 && currentNodeId <= n {
		parent := queue.Front().Value.(*Node)
		queue.Remove(queue.Front())

		if currentNodeId <= n {
			parent.left = NewNode(currentNodeId, parent)
			queue.PushBack(parent.left)
			nodeMap[currentNodeId] = parent.left
			currentNodeId++
		}
		if currentNodeId <= n {
			parent.right = NewNode(currentNodeId, parent)
			queue.PushBack(parent.right)
			nodeMap[currentNodeId] = parent.right
			currentNodeId++
		}
	}

	return root, nodeMap
}

func lvr(root *Node) {
	if root == nil {
		return
	}

	stack := make([]*Node, 0)
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		fmt.Printf("%d ", curr.value)
		curr = curr.right
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanWords)
	scan.Scan()
	N, err := strconv.Atoi(scan.Text())
	if err != nil {
		panic(err)
	}

	// Пропускаем число изменений
	scan.Scan()

	root, nodeMap := buildTree(N)

	for scan.Scan() {
		nodeId, err := strconv.Atoi(scan.Text())
		if err != nil {
			panic(err)
		}

		node := nodeMap[nodeId]
		if node == root {
			continue
		}

		swap(node.parent, node, nodeMap)
	}

	lvr(root)
}
