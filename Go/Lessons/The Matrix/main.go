package main 

import "fmt"

func main() {
	var get string
	fmt.Println("Choose your side Neo: Reality or The Matrix")
	fmt.Scan(&get)

	if get =="reality" {
		fmt.Println("\nWelcome to the reality! You are the person that the world need.\nAre you ready:")
		fmt.Println("Press y,Y for Yes, n,N for No")
		fmt.Scan(&get)

		
		if get =="y"{
			fmt.Println("Great")
			
		} else if get == "Y"{
			fmt.Println("Great")

		} else if get == "n"{
			fmt.Print("You should think about this again\n")
			main()

		} else if get =="N"{
			fmt.Print("You should think about this again\n")
			main()

		} else {
			fmt.Println("Incorrect answer")
			main()
		}

	} else if get == "matrix"{
		fmt.Println("You chose The Matrix, You chose the degradation.")
	} else {
		fmt.Println("You shouldn't leave without answering!\n")
		main()		
	}
}