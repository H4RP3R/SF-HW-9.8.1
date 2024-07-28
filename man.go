package main

import "fmt"

type Man struct {
	Name     string `json:"Name"`
	LastName string `json:"LastName"`
	Age      int    `json:"Age"`
	Gender   string `json:"Gender"`
	Crimes   int    `json:"Crimes"`
}

func (m Man) String() string {
	format := "[%s %s] age: %d, sex: %s, crimes: %d"
	return fmt.Sprintf(format, m.Name, m.LastName, m.Age, m.Gender, m.Crimes)
}
