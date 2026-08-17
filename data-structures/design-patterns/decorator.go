package designpatterns

import "fmt"

type ProcessClass struct{}

func (process *ProcessClass) process() {
	fmt.Println("Process Class process")
}

type ProcessDecorator struct {
	processInstance *ProcessClass
}

func (decorator *ProcessDecorator) process() {
	if decorator.processInstance == nil {
		fmt.Println("Process Decorator process")
	}else {
		fmt.Printf("Process Decorator process and ")
		decorator.processInstance.process()
	}
}

func RunDecoratorPattern() {
	var process = &ProcessClass{}
	var decorator = &ProcessDecorator{}
	decorator.process()
	decorator.processInstance = process
	decorator.process()
}
