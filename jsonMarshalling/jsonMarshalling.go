package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	 Speed int
	 Name string
}

func main() {
	myCar := Car{220, "Peugeot 207 CC Roland Garros"}

	myJsonCar, _ := json.Marshal(myCar) // Converts the object to JSON

	fmt.Println(string(myJsonCar))

	unmarshallingError := json.Unmarshal(myJsonCar, &myCar) // Converts JSON to an object

	if unmarshallingError == nil {
		fmt.Println(myCar.Name, myCar.Speed)
	} else {
		fmt.Println(unmarshallingError)
	}

}
