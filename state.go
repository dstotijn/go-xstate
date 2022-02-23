package xstate

import "context"

type StateNodeConfig struct {
	ID      string
	Initial string
	Type    StateNodeType
	States  map[string]StateNodeConfig
	On      map[string][]Transition
	Invoke  []InvokeConfig
	Always  []Transition
}

type StateNode struct {
	ID      string
	Initial string
	Type    StateNodeType
	Context Data
	States  map[string]StateNode
	On      map[string][]Transition
	Meta    map[string]interface{}

	config StateNodeConfig
}

type StateNodeType string

type State struct {
	Value   StateValue
	Context Data
	Event   Event
	Actions []Action
	Meta    Meta
	Done    bool
}

type Meta map[string]interface{}

type StateValue map[string]StateValue

func NewStateNode(cfg StateNodeConfig, context Data) StateNode {
	sn := StateNode{
		ID:      cfg.ID,
		Initial: cfg.Initial,
		Type:    cfg.Type,
		Context: context,
		On:      map[string][]Transition{},

		config: cfg,
	}

	sn.States = make(map[string]StateNode, len(cfg.States))

	for key, stateCfg := range cfg.States {
		sn.States[key] = NewStateNode(stateCfg, nil)
	}

	return sn
}

func (sn StateNode) Transition(ctx context.Context, state State, event Event) State {
	panic("TODO")
}

func MatchesState(parent, child StateValue) bool {
	for key, parentValue := range parent {
		childValue, ok := child[key]
		if !ok {
			return false
		}
		if parentValue != nil {
			if !MatchesState(parentValue, childValue) {
				return false
			}
		}
	}

	return true
}
