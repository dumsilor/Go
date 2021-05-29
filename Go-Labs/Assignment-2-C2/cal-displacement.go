/*

s = 1/2 at^2 + v0 t + s0

1. Create function called GenDisplayFn()
2. Function will take 3 float64 arguemt, initial displacement, velocity, accelartation
3. GenDisplayFn() will return a function
4. Returnting FUnction will take 1 argument, time
5. returning function return korbe displacement over given time

*/

package main

import "fmt"

func GenDisplayFn(initial_displacement, initial_velocity, accelartation float64) func(time float64) float64 {
	s0 := initial_displacement
	v0 := initial_velocity
	a := accelartation
	cal_displacement := func(t float64) float64 {
		time_square := t * t
		first_cal := a * time_square
		second_cal := 0.5 * first_cal
		third_cal := v0 * t
		displacement := second_cal + third_cal + s0
		return displacement

	}
	return cal_displacement
}

func take_values() (float64, float64, float64) {
	var initial_displacement float64
	var initial_velocity float64
	var accelartation float64

	fmt.Print("Please enter the initial Displacement: ")
	fmt.Scan(&initial_displacement)
	fmt.Print("Please enter the initial Displacement: ")
	fmt.Scan(&initial_velocity)
	fmt.Print("Please enter accelaration: ")
	fmt.Scan(&accelartation)
	return initial_displacement, initial_velocity, accelartation
}

func main() {
	var time float64
	id, iv, a := take_values()
	displacement_func := GenDisplayFn(iv, id, a)
	for {
		fmt.Print("Please enter time: ")
		fmt.Scan(&time)
		fmt.Print("Total Displacement: ", displacement_func(time), "\n")
		var cont string
		fmt.Print("Want to calculate more?(y/n)")
		fmt.Scan(&cont)
		if cont == "n" {
			break
		}
	}

}
