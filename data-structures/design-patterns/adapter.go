package designpatterns

import "fmt"

type IProcess interface {
	process()
}

type Adapter struct {
	adaptee Adaptee
}

func (adapter Adapter) process() {
	fmt.Println("Adapter process");
	adapter.adaptee.convert()
}

type Adaptee struct {
	adapterType int
}

func (adaptee Adaptee) convert() {
	fmt.Println("Adaptee convert method")
}

//ANOTHER EXAMPLE

type Printer interface {
    Print([]byte)
}

type LegacyPrinter struct{}

func (l LegacyPrinter) PrintString(s string) {
    fmt.Println(s)
}

type LegacyAdapter struct {
    legacy LegacyPrinter
}

func (a LegacyAdapter) Print(data []byte) {
    a.legacy.PrintString(string(data))
}



// MAIN FUNCTION
func RunAdapterPattern () {
	var processor IProcess = Adapter{}
	processor.process()
}