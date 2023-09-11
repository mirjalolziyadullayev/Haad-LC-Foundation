package main
import "fmt"

func main() {
	var get int = 0
	fmt.Println("Foydalanuvchi son kiritsin: ")
	fmt.Scan(&get)

	if get < 0 {
		fmt.Println("Manfiy")
	} else if get == 0{
		fmt.Println("0 qiymati musbat ham manfiy ham emas")
	} else {
		fmt.Println("Musbat")
	}
}