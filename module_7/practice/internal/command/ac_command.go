package command

import (
	"smart-patterns-go/internal/device"
)

type AcSetTempCommand struct {
	ac      *device.AirConditioner
	newTemp int
	oldTemp int 
}

func NewAcSetTempCommand(ac *device.AirConditioner, newTemp int) *AcSetTempCommand {
	return &AcSetTempCommand{ac: ac, newTemp: newTemp, oldTemp: ac.Temperature}
}

func (c *AcSetTempCommand) Execute() {
	c.oldTemp = c.ac.Temperature
	c.ac.SetTemperature(c.newTemp)
}

func (c *AcSetTempCommand) Undo() {
	c.ac.SetTemperature(c.oldTemp)
}

func (c *AcSetTempCommand) Redo() {
	c.Execute()
}