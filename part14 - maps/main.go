package main

import "fmt"

func main() {
	Grades := make(map[string]float64)

	Grades["Timmy"] = 42
	Grades["Jess"] = 92
	Grades["Sam"] = 67

	fmt.Println(Grades)

	TimsGrade := Grades["Timmy"]
	fmt.Println(TimsGrade)

	delete(Grades, "Timmy")
	fmt.Println(Grades)

	for k, v := range Grades {
		fmt.Println(k, ":", v)
	}
}
