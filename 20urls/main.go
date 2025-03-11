package main

import "fmt"

func main() {
	fmt.Println("studying urls..")

	changeOrderStatus(Cancelled)

	

}

type razorpay struct {

}

// func ( r razorpay) payment(amt float32){
// 	fmt.Printf("Razorpay payment of %f done.\n", amt)

// }


//enums  in go

// type OrderStatus int
type OrderStatus string

const (
	// Pending OrderStatus = iota
    // Processing
    // Shipped
    // Delivered
    // Cancelled
	Pending   OrderStatus= "Pending"
    Processing  = "Processing"
    Shipped  = "Shipped"
    Delivered  = "Delivered"
    Cancelled = "Cancelled"
   
)

func changeOrderStatus(status OrderStatus) {
	fmt.Println("chage order status: ", status)

}