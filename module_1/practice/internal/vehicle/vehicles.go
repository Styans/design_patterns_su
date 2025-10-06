package vehicles

type Vehicle interface {
	GetBrand() string
	GetModel() string
	GetYear() int
	StartEngine() string
	StopEngine() string
}
