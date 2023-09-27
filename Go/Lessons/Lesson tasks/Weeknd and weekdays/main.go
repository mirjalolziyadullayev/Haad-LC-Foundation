package main  

import "fmt"

func main() {

	var get string 

	fmt. Println("enter the day:")
	fmt.Scan(&get)

	if get ==  "monday" || get == "tuesday" || get ==  "wednesday" || get ==  "thursday" || get == "Friday"  {
		fmt.Println("weekday")
	} else if get == "saturday" || get == "sunday" {
		fmt.Println("weeknd")
	} else {
		fmt.Println("unknown day")
	}

}