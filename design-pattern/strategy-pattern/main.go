package main

type MyDuck struct {
	Duck
	QuarkBehavior QuarkBehavior
	FlyBehavior   FlyBehavior
}

// Duck Abstract
type Duck interface {
	Display()
}

// Behavior interface
type QuarkBehavior interface {
	Quark()
}

// Behavior interfacev
type FlyBehavior interface {
	Fly()
}
