package framework

import (
	"net/http"
	"strings"
)

type TreeNode struct {
	children []*TreeNode
	handler  func(w http.ResponseWriter, r *http.Request)
	param    string
}

func Constructor() TreeNode {
	return TreeNode{
		children: []*TreeNode{},
		param:    "",
	}
}

func (t *TreeNode) Insert(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	node := t
	params := strings.Split(path, "/")
	for _, param := range params {
		child := node.findChild(param)
		if child == nil {
			child = &TreeNode{
				children: []*TreeNode{},
				param:    param,
			}
			node.children = append(node.children, child)
		}
		node = child
	}

	node.handler = handler
}

func (t *TreeNode) findChild(param string) *TreeNode {
	for _, child := range t.children {
		if child.param == param {
			return child
		}
	}
	return nil
}

func (t *TreeNode) Search(path string) func(w http.ResponseWriter, r *http.Request) {
	params := strings.Split(path, "/")
	node := t
	for _, param := range params {
		child := t.findChild(param)
		if child == nil {
			return nil
		}
		node = child
	}
	return node.handler
}
