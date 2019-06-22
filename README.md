# workshop_led_blinking
Material for raspberry pi workshop

# Wiring
![Lighting an LED](doc/img/led-3v3.png)

# Code

## 1- Power your raspberry

You can achive it with connecting it to your pc trought the Micro USB Port of the raspberry pi

![power](doc/img/1-min.jpg)

## 2- Connect to your raspberry pi
Using putty if you're on windows, Ssh if you're on a linux based os
## 3- Download led binary file
If you're using Raspberry 0
```
curl -O https://raw.githubusercontent.com/ionoid-io-projects/workshop_led_blinking/master/bin/arm6/led
```

If you're using Raspberry 3 b
```
curl -O https://raw.githubusercontent.com/ionoid-io-projects/workshop_led_blinking/master/bin/arm7/led
```

## 4- execute binary to make led blinking
```
sudo ./led
```

## 5- Do you see the led Blink ?

# Ressources

[https://projects.raspberrypi.org/en/projects/physical-computing/4](https://projects.raspberrypi.org/en/projects/physical-computing/4)
