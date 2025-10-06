package isp

import "fmt"

type Printer interface {
	Print(content string)
}

type Scanner interface {
	Scan(content string)
}

type Faxer interface {
	Fax(content string)
}

type AllInOnePrinter struct{}

func (p AllInOnePrinter) Print(content string) {
	fmt.Println("Printing:", content)
}

func (p AllInOnePrinter) Scan(content string) {
	fmt.Println("Scanning:", content)
}

func (p AllInOnePrinter) Fax(content string) {
	fmt.Println("Faxing:", content)
}

type BasicPrinter struct{}

func (b BasicPrinter) Print(content string) {
	fmt.Println("Printing:", content)
}
