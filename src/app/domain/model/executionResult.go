package model

type CommandResult struct {
	Command string
	StdOut  []byte
	StdErr  []byte
}
