package main
import "fmt"

type ticketSystem struct {
	TicketNumber int
	EventName string
	Price float64
	Status bool
}
func (d ticketSystem) display() {
	fmt.Println("________________________________________\n")
	fmt.Println("Ticket number:", d.TicketNumber)
	fmt.Println("----------------------------")
	fmt.Println("Event name:", d.EventName)
	fmt.Println("----------------------------")	
	fmt.Println("Price:", d.Price)
	fmt.Println("----------------------------")
	fmt.Println("Status:", d.Status)
	fmt.Println("________________________________________\n")
}
func (s *ticketSystem) setStatus(in bool) {
	if in == true {
		s.Status = true
	} else {
		s.Status = false
	}
}
func main() {
	ticket1:= ticketSystem{
		TicketNumber: 1000,
		EventName: "Mr.Reeves",
		Price: 120,
		Status: false,
	}

	ticket1.display()
	ticket1.setStatus(true)
	ticket1.display()
}
