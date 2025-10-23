package command

type ICommand interface {
	Execute()
	Undo()
}

type LightOnCommand struct {
	light *Light
}

func NewLightOnCommand(light *Light) *LightOnCommand { return &LightOnCommand{light: light} }

func (c *LightOnCommand) Execute() { c.light.On() }

func (c *LightOnCommand) Undo() { c.light.Off() }

type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand { return &LightOffCommand{light: light} }

func (c *LightOffCommand) Execute() { c.light.Off() }

func (c *LightOffCommand) Undo() { c.light.On() }

type TelevisionOnCommand struct {
	tv *Television
}

func NewTelevisionOnCommand(tv *Television) *TelevisionOnCommand { return &TelevisionOnCommand{tv: tv} }

func (c *TelevisionOnCommand) Execute() { c.tv.On() }

func (c *TelevisionOnCommand) Undo() { c.tv.Off() }

type TelevisionOffCommand struct {
	tv *Television
}

func NewTelevisionOffCommand(tv *Television) *TelevisionOffCommand { return &TelevisionOffCommand{tv: tv} }

func (c *TelevisionOffCommand) Execute() { c.tv.Off() }

func (c *TelevisionOffCommand) Undo() { c.tv.On() }

type TempUpCommand struct {
	thermostat *Thermostat
	degrees    int
	prevTemp   int
}

func NewTempUpCommand(thermostat *Thermostat, degrees int) *TempUpCommand {
	return &TempUpCommand{thermostat: thermostat, degrees: degrees}
}

func (c *TempUpCommand) Execute() {
	c.prevTemp = c.thermostat.GetTemp()
	c.thermostat.Up(c.degrees)
}

func (c *TempUpCommand) Undo() {
	c.thermostat.Down(c.degrees)
}

type TempDownCommand struct {
	thermostat *Thermostat
	degrees    int
	prevTemp   int
}

func NewTempDownCommand(thermostat *Thermostat, degrees int) *TempDownCommand {
	return &TempDownCommand{thermostat: thermostat, degrees: degrees}
}

func (c *TempDownCommand) Execute() {
	c.prevTemp = c.thermostat.GetTemp()
	c.thermostat.Down(c.degrees)
}

func (c *TempDownCommand) Undo() {
	c.thermostat.Up(c.degrees)
}

type MacroCommand struct {
	commands []ICommand
}

func NewMacroCommand(commands []ICommand) *MacroCommand {
	return &MacroCommand{commands: commands}
}

func (c *MacroCommand) Execute() {
	for _, cmd := range c.commands {
		cmd.Execute()
	}
}

func (c *MacroCommand) Undo() {
	for i := len(c.commands) - 1; i >= 0; i-- {
		c.commands[i].Undo()
	}
}