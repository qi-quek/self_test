package main

func main() {

}

type Person struct {
	firstName,
	Lastname string
}

type Employee struct {
	Person
	EmployeeId string
}

type Customer struct {
	Person
	CustomerId  string
	PhoneNumber string
}

type CorporateModel interface {
	FullName() string
	ShowID() string
}

type MarketingModel interface {
	CorporateModel
	ShowPhoneNumber() string
}
