package main

import (
	"calc"
	"log"
	"net"
	"net/rpc"
)

type CalcService struct {}

func (c *CalcService) CalcTwoNumber(request calc.Calc, reply *float64) error {
	oper, err := CreateOperation(request.Operator)
	if err != nil {
		return err
	}
	*reply = oper(request.Number1, request.Number2)
	return nil
}

func (c *CalcService) GetOperators(request struct{}, reply *[]string) error {
	*reply = make([]string, 0)
	for k := range Operators {
		*reply = append(*reply, k)
	}
	return nil
}

func main() {
	calc.RegisterCalcService(new(CalcService))
	// rpc.HandleHTTP()
	// http.ListenAndServe(":1234", nil)

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for	{
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}
}