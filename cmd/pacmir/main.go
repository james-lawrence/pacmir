package main

import (
	"github.com/alecthomas/kong"
)

// CmdContext ...
type CmdContext struct {
	Config string
}

func main() {
	type CLI struct {
		Config string `required:"" default:"/etc/pacman.conf"`
		Daemon Daemon `cmd:"" help:"local mirror daemon" default:"1"`
		Mirror Mirror `cmd:"" help:"hosted mirrior daemon"`
		Spike  Spike  `cmd:"" help:"spike"`
	}

	var (
		cli CLI
	)

	ctx := kong.Parse(&cli)
	ctx.FatalIfErrorf(
		ctx.Run(&CmdContext{Config: cli.Config}),
	)
}
