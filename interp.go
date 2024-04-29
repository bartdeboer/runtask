package main

import (
	"github.com/bartdeboer/runtask/yaegilib"
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
	"github.com/traefik/yaegi/stdlib/unrestricted"
)

func newInterpreter() *interp.Interpreter {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)
	i.Use(unrestricted.Symbols)
	i.Use(yaegilib.Symbols)
	return i
}
