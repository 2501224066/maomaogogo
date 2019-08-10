package models

import (
	_ "github.com/go-sql-driver/mysql"
)

func AllNode() interface{} {
	var node []Node
	O.QueryTable(new(Node)).All(&node)

	return node
}
