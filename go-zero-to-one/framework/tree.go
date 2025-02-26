package framework

import (
	"strings"
)

type TreeNode struct {
	children []*TreeNode
	handler  func(ctx *MyContext)
	param    string
	parent   *TreeNode
}

func Constructor() *TreeNode {
	return &TreeNode{
		children: []*TreeNode{},
		param:    "",
	}
}

func isGeneral(param string) bool {
	return strings.HasPrefix(param, ":")
}

func (t *TreeNode) Insert(path string, handler func(ctx *MyContext)) {
	node := t
	params := strings.Split(path, "/")
	for _, param := range params {
		child := node.findChild(param)
		if child == nil {
			child = &TreeNode{
				children: []*TreeNode{},
				param:    param,
				parent:   node,
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

func (t *TreeNode) Search(path string) *TreeNode {
	params := strings.Split(path, "/")

	return dfs(t, params)
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

func (t *TreeNode) ParseParams(path string) map[string]string {
	node := t
	path = strings.TrimSuffix(path, "/")
	paramArr := strings.Split(path, "/")

	paramDicts := make(map[string]string)
	for i := len(paramArr) - 1; i >= 0; i-- {
		if isGeneral(node.param) {
			paramDicts[node.param] = paramArr[i]
		}
		node = node.parent
	}
	return paramDicts
}
