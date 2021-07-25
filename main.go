package main

import (
	"encoding/json"
	"log"
)

type User struct {
	fullname string
	age      int
}

func JSONToUser(j string) User {
	var val struct {
		Fullname string `json:"fullname,omitempty"`
		Age      int    `json:"user_age"`
	}

	err := json.Unmarshal([]byte(j), &val)
	if err != nil {
		log.Fatal(err)
	}

	return User{
		fullname: val.Fullname,
		age:      val.Age,
	}
}

func (s *User) ToJSON() string {
	jsonD := struct {
		Fullname string `json:"fullname,omitempty"`
		Age      int    `json:"user_age"`
	}{
		Fullname: s.fullname,
		Age:      s.age,
	}

	b, err := json.Marshal(jsonD)
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}

type UserExported struct {
	Fullname string `json:"fullname"`
	Age      int    `json:"user_age"`
}

func main() {
	strJson := `{"fullname":"yuri","user_age":19}`

	// MANEIRA COM O MAP
	var mapping map[string]interface{}
	jsoErr := json.Unmarshal([]byte(strJson), &mapping)
	if jsoErr != nil {
		log.Fatal(jsoErr)
	}

	log.Print(mapping["fullname"])
	log.Print(mapping["user_age"])

	// MANEIRA COM STRUCT EXPORTAVEL
	var exported UserExported
	jsoErr = json.Unmarshal([]byte(strJson), &exported)
	if jsoErr != nil {
		log.Fatal(jsoErr)
	}

	log.Print(exported.Fullname)
	log.Print(exported.Age)

	// NÃO É POSSIVEL PASSAR O MAP PRA UMA JSON
	// FAZENDO STRUCT TO JSON
	userExport := UserExported{
		Fullname: "YURI",
		Age:      19,
	}

	b, jsoErr := json.Marshal(userExport)
	if jsoErr != nil {
		log.Fatal(jsoErr)
	}

	log.Print(string(b))

	// MANEIRAS DE CAMPOS NÃO EXPORTAVEIS
	user := JSONToUser(strJson)
	log.Print(user)

	userStruct := User{
		fullname: "yuri",
		age:      19,
	}
	log.Print(userStruct.ToJSON())

}
