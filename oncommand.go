package main

import (
  "errors"
)

type OnCommand struct {
  SerialPort string `arg:"" name:"serial-port" required help:"Serial port"`
	Mode string `long:"mode" enum:"auto,rainbow,breathing,cycle" default:"auto" help:"Led mode (auto,rainbow,breathing,cycle)"`
  Brightness byte `long:"brightness" default:"5" help:"Brightness Level (1->5, 1 darker)"`
  Speed byte `long:"speed" default:"5" help:"Speed Level (1->5, 1 slower)"`
}

func (cmd *OnCommand) Run() error {
  if (cmd.Brightness < 1 || cmd.Brightness > 5) {
    return errors.New("Brightness must be in 1-5 range")
  }
  if (cmd.Speed < 1 || cmd.Speed > 5) {
    return errors.New("Speed must be in 1-5 range")
  }

  var brightness byte = 6 - cmd.Brightness
  var speed byte = 6 - cmd.Speed
  var mode byte

  switch cmd.Mode {
    case "auto":
      mode = 0x05
    case "rainbow":
      mode = 0x01
    case "breathing":
      mode = 0x02
    case "cycle":
      mode = 0x03
  }

  _, err := SendSerialData(cmd.SerialPort, mode, brightness, speed)
	return err
}
