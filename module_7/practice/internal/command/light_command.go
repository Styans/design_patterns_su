package command

import (
	"smart-patterns-go/internal/device"
)

type LightOnCommand struct {
	light *device.Light
}

func NewLightOnCommand(light *device.Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

func (c *LightOnCommand) Undo() {
	c.light.Off()
}

func (c *LightOnCommand) Redo() {
	c.Execute()
}

type LightOffCommand struct {
	light *device.Light
}

func NewLightOffCommand(light *device.Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

func (c *LightOffCommand) Undo() {
	c.light.On()
}

func (c *LightOffCommand) Redo() {
	c.Execute()
}