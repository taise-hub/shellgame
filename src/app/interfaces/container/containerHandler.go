package container

import (
	"bufio"
)

type ContainerHandler interface {
	Inspect(string) error
	Run(string) error
	Create(string) (string, error)
	Remove(string) (error)
	Execute(string, string) (*bufio.Reader, error)
}
