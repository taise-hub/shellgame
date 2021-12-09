package container

import (
	"bytes"
	"io/ioutil"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/taise-hub/shellgame/src/app/domain/model"
)

type ConatainerRepository struct {
	ContainerHandler
}

func (repo *ConatainerRepository) Inspect(name string) error {
	err := repo.ContainerHandler.Inspect(name)
	return err
}

func (repo *ConatainerRepository) Run(id string) error {
	return repo.ContainerHandler.Run(id)
}

func (repo *ConatainerRepository) Create(name string) (id string, err error) {
	id, err = repo.ContainerHandler.Create(name)
	if err != nil {
		return
	}
	return id, err
}

func (repo *ConatainerRepository) Remove(id string) error {
	return repo.ContainerHandler.Remove(id)
}

func (repo *ConatainerRepository) Execute(cmd string, container string) (*model.ExecutionResult, error) {
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