package people

import "fmt"

type person struct {
	name    string
	address string
	phone   string
}

func (prs *person) Print(name, address, phone string) {
	pers := person{
		name:    name,
		address: address,
		phone:   phone,
	}
	fmt.Sprintf("%s\n%s\n%s\n", pers.name, pers.address, pers.phone)
}

func (prs *person) Greet() string {
	return fmt.Sprintf("Hai %s\n", prs.name)
}

func (prs *person) UpdateGreet() *person {
	return &person{
		name:    prs.name,
		address: prs.address,
		phone:   prs.phone,
	}
}
