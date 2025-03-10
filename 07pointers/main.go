package main


func main() {
	println("Pointers")

	// var ptr *int

	// println("Value of ptr: ", ptr)
	// println("Address of ptr: ", &ptr)

	mynumber := 23

	var ptr = &mynumber // refrensce memmey user & pointer

	println("Value of mynumber: ", ptr)
	println("Address of mynumber: ", *ptr)

	*ptr = *ptr + 2 //*ptr have actur valur ref 23
	println("Value of mynumber after increment: ", mynumber)


}
