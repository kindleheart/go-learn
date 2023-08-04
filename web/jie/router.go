package jie

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
	// /a/b/c 总路由
	route string
	// c //当前路由
	path string
	// 静态匹配
	children    map[string]*node
	handler     HandleFunc
	middlewares MiddlewaresChain
	// 通配符 * 匹配，任意匹配
	starChild *node
	// 参数匹配 :id
	paramChild *node
	isRoute    bool
}

func (r *router) addRoute(method, path string, handler HandleFunc, middlewares MiddlewaresChain) {
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
		if root.isRoute {
			panic("web: 路由冲突[/]")
		}
		root.handler = handler
		root.route = path
		root.isRoute = true
		root.middlewares = middlewares
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
	if root.isRoute {
		panic(fmt.Sprintf("web: 路由冲突[%s]", path))
	}
	root.handler = handler
	root.route = path
	root.isRoute = true
	root.middlewares = middlewares
}

func (n *node) childOrCreate(path string) *node {
	// 通配符
	if path == "*" {
		if n.paramChild != nil {
			panic(fmt.Sprintf("web: 非法路由，已有路径参数路由。不允许同时注册通配符路由和参数路由 [%s]", path))
		}
		if n.starChild == nil {
			n.starChild = &node{path: path}
		}
		return n.starChild
	}
	// 路由参数
	if path[0] == ':' {
		if n.starChild != nil {
			panic(fmt.Sprintf("web: 非法路由，已有通配符路由。不允许同时注册通配符路由和参数路由 [%s]", path))
		}
		if n.paramChild != nil {
			if n.paramChild.path != path {
				panic(fmt.Sprintf("web: 路由冲突，参数路由冲突，已有 %s，新注册 %s", n.paramChild.path, path))
			}
		} else {
			n.paramChild = &node{path: path}
		}
		return n.paramChild
	}
	if n.children == nil {
		n.children = map[string]*node{}
	}
	res, ok := n.children[path]
	if !ok {
		res = &node{path: path}
		n.children[path] = res
	}
	return res
}

func (r *router) findRoute(method string, path string) *nodeValue {
	root, ok := r.trees[method]
	if !ok {
		return nil
	}
	if path == "/" {
		if root == nil || !root.isRoute {
			return nil
		}
		return &nodeValue{
			handler:     root.handler,
			middlewares: root.middlewares,
			route:       root.route,
		}
	}
	segs := strings.Split(strings.Trim(path, "/"), "/")
	val := &nodeValue{}
	var starNode *node
	for _, seg := range segs {
		var matchParam bool
		var starChild *node
		root, starChild, matchParam = root.childOf(seg)
		if starNode == nil {
			starNode = starChild
		}
		if root == nil {
			if starNode != nil {
				root = starNode
				break
			}
			return nil
		}
		if matchParam {
			val.addValue(root.path[1:], seg)
		}
	}
	if !root.isRoute {
		return nil
	}
	val.handler = root.handler
	val.middlewares = root.middlewares
	val.route = root.route
	return val
}

// 第一个返回值是命中的节点
// 第二个返回值是通配符节点
// 第三个返回值是代表是否命中路由参数节点
func (n *node) childOf(path string) (*node, *node, bool) {
	if n.children == nil {
		if n.paramChild != nil {
			return n.paramChild, n.starChild, true
		}
		return n.starChild, n.starChild, false
	}
	res, ok := n.children[path]
	if !ok {
		if n.paramChild != nil {
			return n.paramChild, n.starChild, true
		}
		return n.starChild, n.starChild, false
	}
	return res, n.starChild, false
}

type nodeValue struct {
	handler     HandleFunc
	middlewares MiddlewaresChain
	route       string
	pathParams  map[string]string
}

func (m *nodeValue) addValue(key, value string) {
	if m.pathParams == nil {
		m.pathParams = map[string]string{}
	}
	m.pathParams[key] = value
}
