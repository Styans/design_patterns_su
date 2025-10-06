package kiss

import "fmt"

type Calculator struct{}

func (c Calculator) Add(a, b int) {
	fmt.Println(a + b)
}

type Client struct{}

func (c Client) Execute() {
	fmt.Println("Doing something simple...")
}
