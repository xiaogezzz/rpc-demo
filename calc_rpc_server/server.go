package main

import (
	"calc"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type CalcService struct{}

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

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":1234", nil)

	// 使用 curl 模拟调用
	// 
	// curl localhost:1234/jsonrpc -X POST \
	// --data '{"method":"CalcService.GetOperators","params":[{}],"id":0}'
	// 
	// curl localhost:1234/jsonrpc -X POST \
	// --data '{"method":"CalcService.CalcTwoNumber","params":[{"Number1":5,"Number2":2,"Operator":"/"}],"id":0}'
}
