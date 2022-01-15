package container

import (
	"bytes"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/taise-hub/shellgame/src/app/domain/model"
	"io/ioutil"
)

type ContainerRepository struct {
	ContainerHandler
}

func NewContainerRepository(handler ContainerHandler) *ContainerRepository {
	return &ContainerRepository{
		handler,
	}
}

func (repo *ContainerRepository) Exists(name string) bool {
	return repo.ContainerHandler.Exists(name)
}

func (repo *ContainerRepository) Run(id string) error {
	return repo.ContainerHandler.Run(id)
}

func (repo *ContainerRepository) Create(image string, name string) (id string, err error) {
	id, err = repo.ContainerHandler.Create(image, name)
	if err != nil {
		return
	}
	return id, err
}

func (repo *ContainerRepository) Remove(id string) error {
	return repo.ContainerHandler.Remove(id)
}

func (repo *ContainerRepository) Execute(cmd string, container string) (*model.CommandResult, error) {
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
	commandResult := new(model.CommandResult)
	commandResult.Command = cmd
	commandResult.StdOut = stdout
	commandResult.StdErr = stderr

	return commandResult, nil
}
