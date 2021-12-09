package container

import (
	"bytes"
	"io/ioutil"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/taise-hub/shellgame/src/app/domain/model"
)

type ContainerRepository struct {
	ContainerHandler
}

func NewContainerRepository(handler ContainerHandler) *ContainerRepository {
	return &ContainerRepository {
			handler,
	}
} 

func (repo *ContainerRepository) Inspect(name string) error {
	err := repo.ContainerHandler.Inspect(name)
	return err
}

func (repo *ContainerRepository) Run(id string) error {
	return repo.ContainerHandler.Run(id)
}

func (repo *ContainerRepository) Create(name string) (id string, err error) {
	id, err = repo.ContainerHandler.Create(name)
	if err != nil {
		return
	}
	return id, err
}

func (repo *ContainerRepository) Remove(id string) error {
	return repo.ContainerHandler.Remove(id)
}

func (repo *ContainerRepository) Execute(cmd string, container string) (*model.ExecutionResult, error) {
	reader, err := repo.ContainerHandler.Execute(cmd, container)
	if err != nil {
		return nil, err
	}
	var stdoutBuf, stderrBuf bytes.Buffer
	if _, err = stdcopy.StdCopy(&stdoutBuf, &stderrBuf, reader); err != nil {
		return nil, err
	}
	stdout, err := ioutil.ReadAll(&stdoutBuf)
	if err != nil {
		return nil, err
	}
	stderr, err := ioutil.ReadAll(&stderrBuf)
	if err != nil {
		return nil, err
	}
	executionResult := new(model.ExecutionResult)
	executionResult.Command  = cmd
	executionResult.StdOut   = stdout
	executionResult.StdErr   = stderr

	return executionResult, nil
}