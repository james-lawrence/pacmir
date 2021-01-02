package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
)

type mirror struct {
	HTTPBind string `default:"localhost:4000" help:"HTTP address to bind the mirror"`
}

func (t mirror) Run(ctx *context) error {
	log.Println("mirror", spew.Sdump(t), spew.Sdump(ctx))
	return nil
}
