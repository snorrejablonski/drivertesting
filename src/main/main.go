package main

import (
	def "driver" //Hvorfor må jeg ha def forran?!
	"fmt"
	"time"
)

func main() {
	fisk := 1
	if def.Elev_init() == false {
		fmt.Println("Noe fucket skjedde")
	} else if def.Elev_init() == true {
		fmt.Println("Initialize successful")
	}

	//def.Elev_set_button_lamp(2, 1, 1)

	for fisk == 1 {
		def.Elev_set_motor_direction(1)
		if def.Elev_get_button_signal(def.BUTTON_COMMAND1, 3) {

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
