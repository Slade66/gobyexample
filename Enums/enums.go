package main

import "fmt"

type ServerState int

const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = [...]string{
	"idle", "connected", "error", "retrying",
}

func (s ServerState) String() string {
	if s < 0 || int(s) > len(stateName) {
		return "UNKNOW"
	}
	return stateName[s]
}

func main() {
	a := StateIdle
	fmt.Println(a)

	b := ServerState(-1)
	fmt.Println(b)

	fmt.Println(StateError == StateIdle)
	fmt.Println(StateIdle == StateIdle)
}
