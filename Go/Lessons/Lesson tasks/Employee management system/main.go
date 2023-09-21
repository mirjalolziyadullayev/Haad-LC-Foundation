package main
import "fmt"

type employeeSystem struct {
	Name string
	Age int
	Salary float32
}
func (d employeeSystem) display() {
	fmt.Println("____________________________\n")
	fmt.Println("Name: ", d.Name)
	fmt.Println("-------------------------")
	fmt.Println("Age: ", d.Age)
	fmt.Println("-------------------------")
	fmt.Println("Salary: ", d.Salary)
	fmt.Println("____________________________\n")
}
func (s *employeeSystem) chngSalary(in float32) {
	if in <= 0 {
		fmt.Println("0 va 0 dan kam miqdor kiritib bolmaydi")
		return
	}
	s.Salary += s.Salary * in /100
}
func main() {
	worker1:= employeeSystem{
		Name: "Mr.Reeves",
		Age: 34,
		Salary: 10000,
	}
	worker1.display()
	worker1.chngSalary(10)
	worker1.display()
}