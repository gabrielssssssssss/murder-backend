package entity

type ResultEntity struct {
	Person PersonEntity
}

type PersonEntity struct {
	Civility  string
	FirstName string
	LastName  string
	Birthdate string
	Email     string
	Mobile    string
	Zipcode   string
	City      string
	Street    string
}
