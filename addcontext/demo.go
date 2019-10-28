package main

import (
	"context"
)

type Foo interface {
	FooA(i int)
	FooB(j int)
	FooC(ctx context.Context)
}

type Bar interface {
	BarA(i int)
	BarB()
	BarC()
}
