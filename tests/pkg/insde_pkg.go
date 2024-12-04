package pkg

type person struct {
	name string
	age  int
}

func InitPerson(name string, age int) *person {
	return &person{name, age}
}

type Creature struct {
	p     person
	exist bool
}

var (
	C1 = Creature{person{"ly", 111}, false}
)

func (c Creature) GetPerson() person {
	return c.p
}

func EditPerson(p *person, name string, age int) {
	p.name = name
	p.age = age
}
