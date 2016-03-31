package main

import (
	"../driver"
	"fmt"
	//"time"
)

func main() {
	fisk := 1
	if driver.Elev_init() == false {
		fmt.Println("Noe fucket skjedde")
	} else if driver.Elev_init() == true {
		fmt.Println("Initialize successful")
	}

	//driver.Elev_set_button_lamp(2, 1, 1)
	driver.Elev_set_motor_direction(-1)
	for fisk == 1 {
		if driver.Elev_get_button_signal(driver.BUTTON_COMMAND1, 2) {
			driver.Elev_set_motor_direction(-1)
			if driver.Elev_get_floor_sensor_signal() == 2 {
				driver.Elev_set_motor_direction(0)
			}
		}
	}
}

//  FloorChan := make(chan int,1)
// ButtonChan :=make(chan Ressurs.MessageFromEvent,4)

// LightChan :=make(chan Ressurs.ButtonLigth,2)
//MotorChan :=make(chan string,1)

// StatusChan :=make(chan Ressurs.MessageFromEvent,1)

//FromNetwork :=make(chan Ressurs.MessageFromEvent,1)
//ToNetwork :=make(chan Ressurs.MessageFromEvent,1)

// go elev.DriverInit(ButtonChan, FloorChan,LightChan,MotorChan)

//go elevator.InitElevatorControl(FromNetwork,StatusChan,FloorChan,LightChan,MotorChan,ToNetwork)

//go network.MasterSlaveNetwork(StatusChan,ButtonChan,FromNetwork,ToNetwork)

//for{
//    time.Sleep(1e9)
//}
