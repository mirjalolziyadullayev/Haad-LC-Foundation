package main 

import "fmt"

func main() {
	var alarm int
	alarm = 4


	if alarm == 5 {
		fmt.Println("Wake up!")
	} else if alarm < 5 {
		fmt.Println("Sleep enough.")
	} else {
		fmt.Println("wake up! Wake UP!")
	}
}