# Workshop LED blinking
Blinking an LED is one of the first projects a beginner will do to learn electronics, it's no different from learning Raspberry Pi, to do so we will learn in this workshop how to use the Raspberry Pi GPIO pin to blink an LED.

# Wiring

![Lighting an LED](doc/img/led-3v3.jpg)

|LED                  |  Raspberry Pi  |
|---------------------|----------------|
| Ground 		      | GND            |
| Vcc / + 		      | GPIO18         |

![wiring](doc/img/gpio.png)

# Code

## 1- Power your raspberry

You can achive it with connecting it to your pc trought the Micro USB Port of the raspberry pi

![power](doc/img/1-min.jpg)

## 2- Connect to your raspberry pi
Using putty if you're on windows, Ssh if you're on a linux based os
Follow the following instruction if you dont know how to connect to raspberry pi
[Connect to Raspberry Pi using Putty](https://github.com/ionoid-io-projects/workshop/blob/master/doc/od-iot-raspbian-rpi-zero-windows.md#5-first-boot)

## 3- Download led binary file

Assuming you're connected with... copy and past this command
If you're using Raspberry zero
```
curl -O https://raw.githubusercontent.com/ionoid-io-projects/workshop_led_blinking/master/bin/arm6/led
```

If you're using Raspberry 3 b
```
curl -O https://raw.githubusercontent.com/ionoid-io-projects/workshop_led_blinking/master/bin/arm7/led
```
## make it executable
```
chmod +x led
```

## 4- execute binary to make led blink
```
./led
```

## How to stop the program
To quit or stop the program click on **Ctrl+C**

# Ressources


