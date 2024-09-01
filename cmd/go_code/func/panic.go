package main

import "fmt"

func main() {
	//defer fmt.Println("宕机后要做的事情1")
	//defer fmt.Println("宕机后要做的事情2")
	// panic("宕机")
	/*
		宕机后要做的事情2
		宕机后要做的事情1
		panic: 宕机
	*/
	doDBOperations()

}
func connectToDB() {
	fmt.Println("ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB() //function called here with defer
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return //terminate the program
	// deferred function executed here just before actually returning, even if
	// there is a return or abnormal termination before
}
