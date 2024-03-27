package main

import (
	"fmt"
  "github.com/alecthomas/kong"
)

const appName = "S1T9 Led Control by L2K"
const version = "v1.0.0"

var cli struct {
  Off OffCommand `cmd:"" help:"Switch Off Leds."`
  On OnCommand `cmd:"" help:"Switch On Leds."`
}

func main() {
	fmt.Printf("%s %s\n", appName, version)
  ctx := kong.Parse(&cli,
		kong.Name("s1t9ledcontrol"),
		kong.Description("A very simple utility to control S1/T9 Mini PC Led Indicator"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: false,
		}))
  err := ctx.Run()
  ctx.FatalIfErrorf(err)
}
