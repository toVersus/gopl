package main

import (
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var dilbert Employee
	dilbert.Salary -= 5000 // demoted, for waiting too few lines of code

	position := &dilbert.Position
	*position = "Senior" + *position // promoted, for outsourcing to Elbonia

	var employeeOfTheMonth = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	(*employeeOfTheMonth).Position += " (proactive team player)" // same as the above
}
