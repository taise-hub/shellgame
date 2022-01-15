package container

import (
	"bufio"
)

type ContainerHandler interface {
	Exists(string) bool
	Run(string) error
	Create(string, string) (string, error)
	Remove(string) error
	Execute(string, string) (*bufio.Reader, error)
}
