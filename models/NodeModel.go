package models

// NodeRead 节点信息
func NodeRead(nodeID int) Node {
	node := Node{NodeID: nodeID}
	O.Read(&node)

	return node
}

// AllNode 所有节点
func AllNode() interface{} {
	var node []Node
	O.QueryTable(new(Node)).All(&node)

	return node
}
