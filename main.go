package main

import "github.com/nandanai/finalexam/customer"

func main() {
	r := customer.MainAPI()
	r.Run(":2019")
}
