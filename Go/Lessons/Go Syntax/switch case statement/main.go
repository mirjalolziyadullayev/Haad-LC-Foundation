package main

import "fmt"

func main() {
	//switch case with integer
	var daysInt int = 6

	switch daysInt {
	case 1:
		fmt.Println("Monday")
	case 2: 
		fmt.Println("Tuesday")
	case 3: 
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6: 
		fmt.Println("Sunday")
	case 7:
		fmt.Println("Saturday")
	default :
		fmt.Println("Unknown day")
	} 

	switch daysInt {
	case 6,7:
		fmt.Println("Weekend")
	}

	// switch case with string
	var daysString string = "Friday"
	switch daysString {
	case 
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday":

		fmt.Println("Weekday")
	case 
		"Sunday", 
		"Saturday":

		fmt.Println("Weekend")
	default :
		fmt.Println("Unknown day")
	}
}