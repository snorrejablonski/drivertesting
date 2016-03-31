package driver

//inspirasjon fra mortenfyhn, fiks matrix-imputs i funksjonene slik at index'ene stemmer.

import (
	"fmt"
)

const N_FLOORS = 4
const N_BUTTONS = 3

type ELEV_BUTTON_TYPE int

const (
	CALL_UP      = 0
	CALL_DOWN    = 1
	CALL_COMMAND = 2
)

type ELEV_MOTOR_DIR int

const (
	DIR_UP   = 1
	DIR_DOWN = -1
	DIR_STOP = 0
)

var lamp_channel_matrix = [N_FLOORS][N_BUTTONS]int{
	{LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
	{LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
	{LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
	{LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
}

var button_channel_matrix = [N_FLOORS][N_BUTTONS]int{
	{BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
	{BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
	{BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
	{BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4},
}

func Elev_init() bool {
	if Io_init() == false {
		return false
	}
	for f := 0; f < N_FLOORS; f++ {
		for b := 0; b < N_BUTTONS; b++ {
			Elev_set_button_lamp(b, f, 0)
		}
	}
	return true
}

func Elev_set_motor_direction(dirn ELEV_MOTOR_DIR) {
	if dirn == 0 {
		Io_write_analog(MOTOR, 0)
	} else if dirn > 0 {
		Io_clear_bit(MOTORDIR)
		Io_write_analog(MOTOR, 2800)
	} else if dirn < 0 {
		Io_set_bit(MOTORDIR)
		Io_write_analog(MOTOR, 2800)
	}
}

func Elev_set_button_lamp(button int, floor int, value int) {
	if value != 0 {
		Io_set_bit(lamp_channel_matrix[floor][button])
	} else {
		Io_clear_bit(lamp_channel_matrix[floor][button])
	}
}

func Elev_set_floor_indicator(floor int) {
	// Binary encoding. One light must always be on.
	if floor&0x02 > 0 {
		Io_set_bit(LIGHT_FLOOR_IND1)
	} else {
		Io_clear_bit(LIGHT_FLOOR_IND1)
	}

	if floor&0x01 > 0 {
		Io_set_bit(LIGHT_FLOOR_IND2)
	} else {
		Io_clear_bit(LIGHT_FLOOR_IND2)
	}
}

func Elev_set_door_open_lamp(value int) {
	if value != 0 {
		Io_set_bit(LIGHT_DOOR_OPEN)
	} else {
		Io_clear_bit(LIGHT_DOOR_OPEN)
	}
}

func Elev_set_stop_lamp(value int) {
	if value != 0 {
		Io_set_bit(LIGHT_STOP)
	} else {
		Io_clear_bit(LIGHT_STOP)
	}
}

func Elev_get_button_signal(button int, floor int) bool {
	fmt.Printf("button %i, floor %i\n", button, floor)
	if Io_read_bit(button_channel_matrix[floor][button]) != 0 {
		return true
	} else {
		return false
	}
}

func Elev_get_floor_sensor_signal() int {
	if Io_read_bit(SENSOR_FLOOR1) != 0 {
		return 0
	} else if Io_read_bit(SENSOR_FLOOR2) != 0 {
		return 1
	} else if Io_read_bit(SENSOR_FLOOR3) != 0 {
		return 2
	} else if Io_read_bit(SENSOR_FLOOR4) != 0 {
		return 3
	} else {
		return -1
	}
}

func Elev_get_stop_signal() int {
	return Io_read_bit(STOP)
}

func Elev_get_obstruction_signal() int {
	return Io_read_bit(OBSTRUCTION)
}
