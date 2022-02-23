package xstate

type Transition struct {
	Target   string
	Internal bool
	Actions  []Action
}
