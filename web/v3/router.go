package v2

import (
	"fmt"
	"strings"
)

// 用来支持路由树操作
type router struct {
	trees map[string]*node
}

func newRouter() router {
	return router{trees: map[string]*node{}}
}

type node struct {
	path string
	// 静态匹配
	children map[string]*node
	handler  HandleFunc
	// 通配符 * 匹配，任意匹配
	starChild *node
}

func (r *router) addRoute(method string, path string, handleFunc HandleFunc) {
	if path == "" {
		panic("web: 路由是空字符串")
	}
	if path[0] != '/' {
		panic("web: 路由必须以 / 开头")
	}
	if path != "/" && path[len(path)-1] == '/' {
		panic("web: 路由不能以 / 结尾")
	}
	root, ok := r.trees[method]
	if !ok {
		root = &node{path: "/"}
		r.trees[method] = root
	}
	if path == "/" {
		if root.handler != nil {
			panic("web: 路由冲突[/]")
		}
		root.handler = handleFunc
		return
	}

	path = path[1:]
	// 切割path
	segs := strings.Split(path, "/")
	for _, seg := range segs {
		if seg == "" {
			panic(fmt.Sprintf("web: 非法路由。不允许使用 //a/b, /a//b 之类的路由, [%s]", path))
		}
		root = root.childOrCreate(seg)
	}
	if root.handler != nil {
		panic(fmt.Sprintf("web: 路由冲突[%s]", path))
	}
	root.handler = handleFunc
}

func (r *router) findRoute(method string, path string) (*node, bool) {
	root, ok := r.trees[method]
	if !ok {
		return nil, false
	}
	if path == "/" {
		return root, true
	}
	segs := strings.Split(strings.Trim(path, "/"), "/")
	for _, seg := range segs {
		root, ok = root.childOf(seg)
		if !ok {
			return nil, false
		}
	}
	return root, true
}

func (n *node) childOf(path string) (*node, bool) {
	if n.children == nil {
		return n.starChild, n.starChild != nil
	}
	res, ok := n.children[path]
	if !ok {
		return n.starChild, n.starChild != nil
	}
	return res, ok
}

func (n *node) childOrCreate(seg string) *node {
	if seg == "*" {
		if n.starChild == nil {
			n.starChild = &node{path: seg}
		}
		return n.starChild
	}
	if n.children == nil {
		n.children = map[string]*node{}
	}
	res, ok := n.children[seg]
	if !ok {
		res = &node{path: seg}
		n.children[seg] = res
	}
	return res
}
