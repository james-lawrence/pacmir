package main

import (
	"github.com/alecthomas/kong"
)

type context struct {
	Config string
}

func main() {
	type CLI struct {
		Config string `required:"" default:"/etc/pacman.conf"`
		Daemon Daemon `cmd:"" help:"local mirror daemon" default:"1"`
		Mirror mirror `cmd:"" help:"hosted mirrior daemon"`
	}
	var (
		cli CLI
	)
	ctx := kong.Parse(&cli)
	ctx.FatalIfErrorf(
		ctx.Run(&context{Config: cli.Config}),
	)
}
