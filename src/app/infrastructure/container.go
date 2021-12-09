package infrastructure

import (
	"bufio"
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)
type ContainerHandler struct {
	Client *client.Client
}

func NewContainerHandler() *ContainerHandler {
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	handler := new(ContainerHandler)
	handler.Client = cli
	return handler
}


func (h *ContainerHandler) Inspect(name string) error {
	_, err := h.Client.ContainerInspect(context.Background(), name)
	if err != nil {
		return nil
	}
	return err
}

// run container
func (h *ContainerHandler) Run(id string) error {
	if err := h.Client.ContainerStart(context.Background(), id, types.ContainerStartOptions{}); err != nil {
		return err
	}

	return nil
}

// create container
func (h *ContainerHandler) Create(name string) (id string, err error) {
	cc := &container.Config{
		Image: "shellgame",
		Tty:   true,
	}
	hc := &container.HostConfig{
		AutoRemove: true,
	}
	body, err := h.Client.ContainerCreate(context.Background(), cc, hc, nil, nil, name)
	if err != nil {
		return
	}
	id = body.ID
	return 
}

// remove container
func (h *ContainerHandler) Remove(id string) error {
	option := types.ContainerRemoveOptions {
		RemoveVolumes: true,
		RemoveLinks:   true,
		Force:         true,
	}
	return h.Client.ContainerRemove(context.Background(), id, option)
}

// execute command on container
func (h *ContainerHandler) Execute(cmd string, container string)(*bufio.Reader, error) {
	cmds := []string{"/bin/bash", "-c", cmd}
	ec := &types.ExecConfig{
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          cmds,
	}
	conf, err := h.Client.ContainerExecCreate(context.Background(), container, *ec)
	if err != nil {
		return nil, err
	}
	resp, err := h.Client.ContainerExecAttach(context.Background(), conf.ID, types.ExecStartCheck{})
	if err != nil {
		return nil, err
	}
	return resp.Reader, nil
}