package shell

import (
	"StackDB/internal/utils"
)

type commandNode struct {
	uuid  string
	cmd   string
	child *commandNode
	args  []string
}

func newCommandNode() *commandNode {
	return &commandNode{
		uuid: utils.GetUuid(),
	}
}

func (cn *commandNode) walkDown() *commandNode {
	return cn.child
}

func (cn *commandNode) execute() error {
	return nil
}
