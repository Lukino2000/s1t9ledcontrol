package main

type OffCommand struct {
  SerialPort string `arg:"" name:"serial-port" required help:"Serial port"`
}

func (cmd *OffCommand) Run() error {
  _, err := SendSerialData(cmd.SerialPort, 0x04, 0x05, 0x05)
	return err
}
