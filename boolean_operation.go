package main

import "fmt"

func main() {
	var finalScore = 90
	var attendance = 81

	var passedFinalScore bool = finalScore > 80
	var passedAttendance bool = attendance > 80

	var passed bool = passedFinalScore && passedAttendance

	fmt.Println(passed)
}
