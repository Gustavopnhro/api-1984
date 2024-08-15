package entity

import "gorm.io/gorm"

type Person struct {
	ID        ID       `json:"id"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Birthday  string   `json:"birthday"`
	District  string   `json:"district"`
	Gender    string   `json:"gender"`
	Job       string   `json:"job"`
	Hobbies   []string `json:"hobbies"`
	gorm.Model
}

func CreatePerson(first_name, last_name, birthday, district, gender, job string, hobbies []string) *Person {
	return &Person{
		ID:        NewID(),
		FirstName: first_name,
		LastName:  last_name,
		Birthday:  birthday,
		District:  district,
		Gender:    gender,
		Job:       job,
		Hobbies:   hobbies,
	}
}
