package generate

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
)

func init() {
	node, _ = snowflake.NewNode(time.Now().UnixNano() % 1024)
}

func SnowFlask() int64 {
	return node.Generate().Int64()
}
