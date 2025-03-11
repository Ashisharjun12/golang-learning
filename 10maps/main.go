package main

import "fmt"

func main() {

	fmt.Println("study maps")
	languaguages := make(map[string]string)

	languaguages["go"] = "Go Lang"
	languaguages["java"] = "Java"
	languaguages["python"] = "Python"
	languaguages["javascript"] = "JavaScript"

	fmt.Println("languaguages: ", languaguages)

	fmt.Println("languaguages[go]: ", languaguages["go"])

	//delete the valuse
	delete(languaguages, "java")
	fmt.Println("languaguages: ", languaguages)


	//loops
	for key , value := range languaguages{
		fmt.Println(key, " : ", value)
	}

}