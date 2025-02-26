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

func isGeneral(param string) bool {
	return strings.HasPrefix(param, ":")
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

	result := dfs(t, params)
	if result == nil {
		return nil
	}

	return result.handler
}

func dfs(node *TreeNode, params []string) *TreeNode {
	currentParam := params[0]
	isLastParam := len(params) == 1

	for _, child := range node.children {
		if isLastParam {
			if isGeneral(child.param) {
				return child
			}

			if child.param == currentParam {
				return child
			}
			continue
		}

		if !isGeneral(child.param) && child.param != currentParam {
			continue
		}

		result := dfs(child, params[1:])
		if result != nil {
			return result
		}
	}
	return nil
}
