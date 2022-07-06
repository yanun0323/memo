package main

type Duck struct {
	QuarkBehavior QuarkBehavior
	FlyBehavior   FlyBehavior
}

func (d *Duck) Display() {

}

// Behavior interface
type QuarkBehavior interface {
	Quark()
}

// Behavior interface
type FlyBehavior interface {
	Fly()
}
