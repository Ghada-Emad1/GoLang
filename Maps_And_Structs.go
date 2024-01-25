package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}
type Animal struct {
	Name   string
	origin string
}
type Bird struct {
	Animal   // it is a compostion relation ship we don't say that Bird is Animal instead we say Bird has animal
	speedkpH float32
	canFly   bool
}
type Tags struct{
	Name string `required max:"100"`
	origin string
}
type Teacher struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Domains []string `json:"domains"`

}
type User struct{
	username string `ankit:"required,min=5,max=12"`
	email string `ankita:"required"`
	Test string
}

func main() {
	//initalization --->1
	statePopulation := map[string]int{
		"Cailforina": 123456,
		"Taxes":      987654,
		"Hollywood":  2345189,
	}
	fmt.Println(statePopulation)
	fmt.Println(statePopulation["Taxes"])

	// ---> 2
	statePopulation1 := make(map[string]int)
	statePopulation1["hawai"] = 1234321
	fmt.Println(statePopulation1["hawai"])

	fmt.Println(statePopulation["egypt"]) // if the request key doesn;t exit , it returns 0
	fmt.Println(len(statePopulation))     // length of map

	//To delete key from the map
	fmt.Println(statePopulation)
	delete(statePopulation, "Taxes")
	fmt.Println(statePopulation)

	pop, ok := statePopulation["ohio"]
	fmt.Println(pop) // the value stored under the key ohio which it is not exit , so it will return 0
	fmt.Println(ok)  // is a bool that is true if the key exist else it is false

	pop1, ok1 := statePopulation["Cailforina"]
	fmt.Println(pop1)
	fmt.Println(ok1)

	//To test without retrieveing the value
	_, ok2 := statePopulation["Taxes"]
	fmt.Println(ok2)

	// map type is reference types like slice and pointer
	sp := statePopulation // the change will be in two maps
	delete(sp, "Hollywood")
	sp["Taxes"] = 1000000
	fmt.Printf("StatePopulation map %v\n", statePopulation)
	fmt.Printf("The reference Map %v\n", sp)

	//Structs : is a collection of type
	aDoctor := Doctor{
		number:     3,
		actorName:  "john pertwee",
		companions: []string{"lisa", "omar"},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.number)
	fmt.Println(aDoctor.actorName)

	//Another declaration
	aDoctor1 := struct{ name string }{name: "john"}
	fmt.Println(aDoctor1)

	//Struct is not reference type so change in one don't change in another one
	anotherDoctor := aDoctor
	anotherDoctor.number = 4
	fmt.Println(anotherDoctor.number)
	fmt.Println(aDoctor.number)

	//if you want the two struct have same value you can use pointer
	anotherDoctor1 := &aDoctor
	anotherDoctor1.actorName = "ahmed ali"
	fmt.Println(anotherDoctor1.actorName)
	fmt.Println(aDoctor.actorName)

	//Go doesn't support traditional oop
	//it uses a model that's similar to inheritance called composition
	// it suports composition through what is called embedding

	b := Bird{}
	//Bird have the instance of Animal
	b.Name = "Bird"
	b.origin = "Egyptian"
	b.speedkpH = 45
	b.canFly = false
	fmt.Println(b.Animal)
	fmt.Println(b.Name)

	//Literal Syntax
	c := Bird{Animal: Animal{Name: "ricky", origin: "canda"}, speedkpH: 23, canFly: false}
	fmt.Println(c)
	fmt.Println(c.Animal)

	//using tag
	t:=reflect.TypeOf(Tags{})
	field,_:=t.FieldByName("Name")
	fmt.Println(field.Tag)

	teach:=Teacher{}
	st:=reflect.TypeOf(teach)
	for i:=0;i<st.NumField();i++{
		field:=st.Field(i)
		fmt.Println(field.Tag.Get("json"))
	}

  
	tag:=reflect.TypeOf(User{})
	field1,_:=tag.FieldByName("username")
	fmt.Println(field1.Tag)
	

}
