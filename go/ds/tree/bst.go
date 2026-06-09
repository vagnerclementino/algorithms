package ds

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

type TreeNode struct {
	Value     int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

type BinarySearchTree struct {
	root *TreeNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{root: nil}
}

func (t *BinarySearchTree) Add(value int) {
	newNode := &TreeNode{Value: value}
	if t.root == nil {
		t.root = newNode
		return
	} else {
		addNode(t.root, newNode)
	}
}

func (t *BinarySearchTree) Search(value int) bool {
	if t == nil {
		return false
	}
	return search(t.root, value)
}

func (t *BinarySearchTree) MinValue() (int, error) {
	if t == nil {
		return math.MaxInt, fmt.Errorf("the tree is nil")
	}
	return minValue(t.root), nil
}

func (t *BinarySearchTree) MaxValue() (int, error) {
	if t == nil {
		return math.MinInt, fmt.Errorf("the tree is nil")
	}
	return maxValue(t.root), nil
}

func minValue(node *TreeNode) int {
	currentNode := node
	for currentNode.LeftNode != nil {
		currentNode = currentNode.LeftNode
	}
	return currentNode.Value
}

func maxValue(node *TreeNode) int {
	currentNode := node
	for currentNode.RightNode != nil {
		currentNode = currentNode.RightNode
	}
	return currentNode.Value
}

func search(node *TreeNode, value int) bool {
	if node == nil {
		return false
	}
	if value < node.Value {
		return search(node.LeftNode, value)
	}

	if value > node.Value {
		return search(node.RightNode, value)
	}
	return true
}

func addNode(rootNode, newNode *TreeNode) {
	if newNode.Value < rootNode.Value {
		if rootNode.LeftNode == nil {
			rootNode.LeftNode = newNode
		} else {
			addNode(rootNode.LeftNode, newNode)
		}
	} else {
		if rootNode.RightNode == nil {
			rootNode.RightNode = newNode
		} else {
			addNode(rootNode.RightNode, newNode)
		}
	}
}

func collectNodes(n *TreeNode) ([]*TreeNode, int) {
	nodes := make([]*TreeNode, 0)
	maxW := 0
	var dfs func(*TreeNode)
	dfs = func(nd *TreeNode) {
		if nd == nil {
			return
		}
		nodes = append(nodes, nd)
		w := len(fmt.Sprintf("%d", nd.Value))
		if w > maxW {
			maxW = w
		}
		dfs(nd.LeftNode)
		dfs(nd.RightNode)
	}
	dfs(n)
	return nodes, maxW
}

func assignInorder(n *TreeNode, xMap map[*TreeNode]int, depthMap map[*TreeNode]int, idx *int, depth int) {
	if n == nil {
		return
	}
	assignInorder(n.LeftNode, xMap, depthMap, idx, depth+1)
	xMap[n] = *idx
	depthMap[n] = depth
	(*idx)++
	assignInorder(n.RightNode, xMap, depthMap, idx, depth+1)
}

func (t *BinarySearchTree) String() string {
	if t.root == nil {
		return "<empty>"
	}
	// compute max width of node text and collect nodes (uses helper)
	_, maxW := collectNodes(t.root)
	if maxW == 0 {
		maxW = 1
	}
	// assign x positions using in-order traversal (compact)
	xMap := make(map[*TreeNode]int)
	depthMap := make(map[*TreeNode]int)
	idx := 0
	assignInorder(t.root, xMap, depthMap, &idx, 0)
	if idx == 0 {
		return "<empty>"
	}
	// layout parameters
	slotWidth := maxW + 1 // tight padding
	totalWidth := idx * slotWidth
	// find max depth
	maxDepth := 0
	for _, d := range depthMap {
		if d > maxDepth {
			maxDepth = d
		}
	}
	var buf bytes.Buffer
	// build lines per depth
	for depth := 0; depth <= maxDepth; depth++ {
		line := make([]rune, totalWidth)
		for i := range line {
			line[i] = ' '
		}
		// place node labels
		for n := range xMap {
			if depthMap[n] != depth {
				continue
			}
			x := xMap[n]*slotWidth + slotWidth/2
			ns := fmt.Sprintf("%d", n.Value)
			start := x - len(ns)/2
			for j, ch := range ns {
				pos := start + j
				if pos >= 0 && pos < totalWidth {
					line[pos] = ch
				}
			}
		}
		// trim trailing spaces to reduce horizontal width on small terminals
		sline := strings.TrimRight(string(line), " ")
		buf.WriteString(sline)
		buf.WriteByte('\n')
		// connector line (to next depth)
		if depth < maxDepth {
			conn := make([]rune, totalWidth)
			for i := range conn {
				conn[i] = ' '
			}
			for n := range xMap {
				if depthMap[n] != depth {
					continue
				}
				px := xMap[n]*slotWidth + slotWidth/2
				if n.LeftNode != nil {
					cx := xMap[n.LeftNode]*slotWidth + slotWidth/2
					mid := (px + cx) / 2
					if mid >= 0 && mid < totalWidth {
						conn[mid] = '/'
					}
				}
				if n.RightNode != nil {
					cx := xMap[n.RightNode]*slotWidth + slotWidth/2
					mid := (px + cx) / 2
					if mid >= 0 && mid < totalWidth {
						conn[mid] = '\\'
					}
				}
			}
			// trim connector line as well
			sconn := strings.TrimRight(string(conn), " ")
			buf.WriteString(sconn)
			buf.WriteByte('\n')
		}
	}
	return buf.String()
}
