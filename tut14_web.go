package main

import (
	"fmt"
)

func tut14() {
	grades := make(map[string]float32)
	grades["Timmy"] = 42
	grades["Jess"] = 92
	grades["Sam"] = 67

	fmt.Println(grades)

	TimsGrade := grades["Timmy"]
	fmt.Println(TimsGrade)

	delete(grades, "Timmy")
	fmt.Println(grades)

	for key, value := range grades {
		fmt.Printf("%s got a score %v\n", key, value)
	}
}
