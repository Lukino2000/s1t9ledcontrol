package main

import (
  "fmt"
	"time"
	"go.bug.st/serial"
)

func SendSerialData(serialName string, mode byte, brightness byte, speed byte) (int, error) {
	crc := byte(uint16(mode + brightness + speed + 0xfa) & 0x00ff)
  data := []byte{0xfa, mode, brightness, speed, crc}

  port, err := serial.Open(serialName, &serial.Mode{BaudRate: 10000,})
  if err != nil {
    return -1, fmt.Errorf("Cannot open serial port: %s", err.Error())
  }

  for _, v := range data {
    _, err := port.Write([]byte{v})
    if err != nil {
      return -1, fmt.Errorf("Cannot send data to serial port: %s", err.Error())
    }
    time.Sleep(5 * time.Millisecond)
  }

  port.Close()

	return len(data), nil
}

