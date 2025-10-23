package command

type Command interface {
	Execute()
	Undo()
}

type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(l *Light) *LightOnCommand { return &LightOnCommand{light: l} }
func (c *LightOnCommand) Execute()               { c.light.TurnOn() }
func (c *LightOnCommand) Undo()                  { c.light.TurnOff() }

type DoorOpenCommand struct {
	door *Door
}

func NewDoorOpenCommand(d *Door) *DoorOpenCommand { return &DoorOpenCommand{door: d} }
func (c *DoorOpenCommand) Execute()                { c.door.Open() }
func (c *DoorOpenCommand) Undo()                   { c.door.Close() }

type TempIncreaseCommand struct {
	thermostat *Thermostat
}

func NewTempIncreaseCommand(t *Thermostat) *TempIncreaseCommand { return &TempIncreaseCommand{thermostat: t} }
func (c *TempIncreaseCommand) Execute()                        { c.thermostat.IncreaseTemp() }
func (c *TempIncreaseCommand) Undo()                           { c.thermostat.DecreaseTemp() }

type TempDecreaseCommand struct {
	thermostat *Thermostat
}

func NewTempDecreaseCommand(t *Thermostat) *TempDecreaseCommand { return &TempDecreaseCommand{thermostat: t} }
func (c *TempDecreaseCommand) Execute()                        { c.thermostat.DecreaseTemp() }
func (c *TempDecreaseCommand) Undo()                           { c.thermostat.IncreaseTemp() }