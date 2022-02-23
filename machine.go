package xstate

func NewMachine(cfg StateNodeConfig, context Data) StateNode {
	return NewStateNode(cfg, context)
}
