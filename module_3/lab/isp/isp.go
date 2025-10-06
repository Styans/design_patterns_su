package isp

import "fmt"

type Worker interface {
	Work()
}

type Eater interface {
	Eat()
}

type Sleeper interface {
	Sleep()
}

type HumanWorker struct{}

func (HumanWorker) Work()  { fmt.Println("Human is working...") }
func (HumanWorker) Eat()   { fmt.Println("Human is eating...") }
func (HumanWorker) Sleep() { fmt.Println("Human is sleeping...") }

type RobotWorker struct{}

func (RobotWorker) Work() { fmt.Println("Robot is working...") }
