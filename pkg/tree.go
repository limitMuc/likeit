package pkg

import (
	"fmt"
)

const (
	NEW_LINE       = "\n"
	INTERVAL_SPACE = "    "
	INTERNAL_NODE  = "├── "
	LAST_NODE      = "└── "
)

type Node struct {
	name      string
	number    int
	isLast    bool
	isDir     bool
	stdPrint  string
	absPath   string
	needSpace string
	level     int
	subNode   []Node
}

func (n *Node) GetName() string {
	return n.name
}

func (n *Node) GetNumber() int {
	return n.number
}

func (n *Node) GetIsLast() bool {
	return n.isLast
}

func (n *Node) GetStdPrint() string {
	return n.stdPrint
}

func (n *Node) GetAbsPath() string {
	return n.absPath
}

func (n *Node) GetSubNode() []Node {
	return n.subNode
}

func (n *Node) Print() {
	if n.level == 1 {
		fmt.Printf("%s <%d> %s", n.stdPrint, Total, NEW_LINE)
	} else if !n.isDir && len(n.subNode) == 0 {
		fmt.Printf("%s <%d> %s", n.stdPrint, n.number, NEW_LINE)
	} else {
		fmt.Printf("%s %s", n.stdPrint, NEW_LINE)
	}
	for _, node := range n.subNode {
		node.Print()
	}
}

func NewNode(name, stdPrint string, level int) *Node {
	return &Node{
		name:      name,
		number:    0,
		isLast:    false,
		stdPrint:  stdPrint,
		absPath:   name,
		needSpace: "",
		level:     level,
		subNode:   []Node{},
	}
}
