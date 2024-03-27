## S1/T9 Mini Pc Led Control

In releases you will find compiled version of this utility renamed with os postfix, feel free to download, rename and put it on your favorite folder

to switch off
```
Usage: s1t9ledcontrol off <serial-port> [flags]
Arguments:
  <serial-port>    Serial port

ex:
linux: s1t9ledcontrol off /dev/ttyUSB0
win: s1t9ledcontrol.exe off com3
```

to switch on
```
Usage: s1t9ledcontrol on <serial-port> [flags]
Arguments:
  <serial-port>    Serial port
Flags:
  --mode="auto"     Led mode
  --brightness=5    Brightness Level (1->5, 1 darker)
  --speed=5         Speed Level (1->5, 1 slower)

ex:
linux: s1t9ledcontrol on /dev/ttyUSB0 --mode auto --brightness 3 --speed 3
win: s1t9ledcontrol.exe on com3 --mode auto --brightness 3 --speed 3
```
