# Introduction
Making led blink 

# How to
Compile led.go like this
```
go get github.com/stianeikeland/go-rpio
env GOOS=linux GOARCH=arm GOARM=6 go build led.go
```
Copy the generated file to your raspberry pi device and execute it with this command

```
env RPIO_PIN=18 BLINK_INTERVAL=1 ./led
```

# env variables used
RPIO_PIN = an integer to define rpi pin used in wiring, default 18
BLINK_INTERVAL = time interval in second, default 1

Happy blinking 