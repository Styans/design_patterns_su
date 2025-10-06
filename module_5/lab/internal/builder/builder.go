package builder

type Computer struct {
	CPU     string
	RAM     string
	Storage string
	GPU     string
	OS      string
	PSU     string
	Cooling string
}

func (c Computer) String() string {
	return "CPU: " + c.CPU + "; RAM: " + c.RAM + "; Storage: " + c.Storage + "; GPU: " + c.GPU + "; OS: " + c.OS + "; PSU: " + c.PSU + "; Cooling: " + c.Cooling
}

type IComputerBuilder interface {
	SetCPU()
	SetRAM()
	SetStorage()
	SetGPU()
	SetOS()
	SetPSU()
	SetCooling()
	GetComputer() Computer
}

type OfficeComputerBuilder struct {
	computer Computer
}

func NewOfficeComputerBuilder() *OfficeComputerBuilder {
	return &OfficeComputerBuilder{}
}

func (b *OfficeComputerBuilder) SetCPU() {
	b.computer.CPU = "Intel i3"
}

func (b *OfficeComputerBuilder) SetRAM() {
	b.computer.RAM = "8GB"
}

func (b *OfficeComputerBuilder) SetStorage() {
	b.computer.Storage = "1TB HDD"
}

func (b *OfficeComputerBuilder) SetGPU() {
	b.computer.GPU = "Integrated"
}

func (b *OfficeComputerBuilder) SetOS() {
	b.computer.OS = "Windows 10"
}

func (b *OfficeComputerBuilder) SetPSU() {
	b.computer.PSU = "400W"
}

func (b *OfficeComputerBuilder) SetCooling() {
	b.computer.Cooling = "Air"
}

func (b *OfficeComputerBuilder) GetComputer() Computer {
	return b.computer
}

type GamingComputerBuilder struct {
	computer Computer
}

func NewGamingComputerBuilder() *GamingComputerBuilder {
	return &GamingComputerBuilder{}
}

func (b *GamingComputerBuilder) SetCPU() {
	b.computer.CPU = "Intel i9"
}

func (b *GamingComputerBuilder) SetRAM() {
	b.computer.RAM = "32GB"
}

func (b *GamingComputerBuilder) SetStorage() {
	b.computer.Storage = "1TB SSD"
}

func (b *GamingComputerBuilder) SetGPU() {
	b.computer.GPU = "NVIDIA RTX 4080"
}

func (b *GamingComputerBuilder) SetOS() {
	b.computer.OS = "Windows 11"
}

func (b *GamingComputerBuilder) SetPSU() {
	b.computer.PSU = "850W"
}

func (b *GamingComputerBuilder) SetCooling() {
	b.computer.Cooling = "Liquid"
}

func (b *GamingComputerBuilder) GetComputer() Computer {
	return b.computer
}

type ComputerDirector struct {
	builder IComputerBuilder
}

func NewComputerDirector(b IComputerBuilder) ComputerDirector {
	return ComputerDirector{builder: b}
}

func (d ComputerDirector) ConstructComputer() Computer {
	d.builder.SetCPU()
	d.builder.SetRAM()
	d.builder.SetStorage()
	d.builder.SetGPU()
	d.builder.SetOS()
	d.builder.SetPSU()
	d.builder.SetCooling()
	return d.builder.GetComputer()
}
