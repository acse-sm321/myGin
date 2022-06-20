package myGin

import "strings"

type node struct {
	pattern  string  // matching route e.g. /p/:lang
	part     string  // part of the route e.g. :lang
	children []*node // children nodes are stored in a pointer list
	isWild   bool    // is the matching accurate, : or * in part will be true
}

// find the first matching node for insertion
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// find all the matching node for search
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

/* Insert new node to tier tree */
func (n *node) insert(pattern string, parts []string, height int) {
	// if matched to len(parts) level,copy the pattern then return
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part) // search child that matches the part
	// if no child matches the searching part, create a new node
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	// if exists child matching the searching part, insert recursively to the correct location
	child.insert(pattern, parts, height+1)
}

/* Search existing node in the Trie tree */
func (n *node) search(parts []string, height int) *node {
	// if searched to bottom of trie tree and has prefix "*", return node
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part) // find all children matching the search parts and return it

	// do the search recursively until all matching children are found
	for _, child := range children {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
