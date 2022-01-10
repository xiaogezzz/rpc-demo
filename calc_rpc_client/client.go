package main

import (
	"calc"
	"log"
)

func main() {
	client, err := calc.DialCalcService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Err Dial Client:", err)
	}

	// var opers []string
	// err = client.GetOperators(struct{}{}, &opers)
	// if err != nil {
	// 	log.Println("Err Get Operators:", err)
	// }
	// log.Println(opers)

	testAdd := calc.Calc{
		Number1:  5,
		Number2:  2,
		Operator: "/",
	}
	var result float64
	err = client.CalcTwoNumber(testAdd, &result)
	if err != nil {
		log.Println("Err Calculate:", err)
	}
	log.Println(result)
}
