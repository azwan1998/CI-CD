package appModel

type PersonModel interface {
	GetByEmailAndPassword(email string, password string) (Person, error)
	GetAll() ([]Person, error)
	Add(Person) (Person, error)
	Edit(int, Person) (Person, error)
	IsActive(int, Person) (Person, error)
	GetByEmail(email string) (Person, error)
}
