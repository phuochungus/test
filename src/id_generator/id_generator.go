package idGen

import "github.com/bwmarrin/snowflake"

var instance *snowflake.Node

func init() {
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	instance = node
}

func GenId() int {
	return int(instance.Generate())
}
