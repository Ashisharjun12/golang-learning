package main

import "fmt"

func main() {
	fmt.Println("slices starting...")

	var fruitlist = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	fmt.Printf("type of fruitlist: %T\n" , fruitlist)

	fmt.Println("fruitlist: ", fruitlist)
	fmt.Println("fruitlist: ", len(fruitlist))

	fruitlist = append(fruitlist, "mango " ,"bannan")

	fmt.Println("fruitlist: ", fruitlist)

	fruitlist = append(fruitlist[1:4])
	fmt.Println("fruitlist after slicing: ", fruitlist)

	//how to remove valuse based on index

	var course = []string{"react", "angular", "vue", "svelte", "javascript"}
	fmt.Println("course: ", course)
	course = append(course[:2], course[3:]...)
	fmt.Println("course after removing index 2: ", course)

	
}