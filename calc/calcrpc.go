package calc

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Calc struct {
	Number1  float64
	Number2  float64
	Operator string
}

const ServiceName = "CalcService"

type ServiceInterface interface {
	// CalcTwoNumber 对两个数进行运算
	CalcTwoNumber(request Calc, reply *float64) error
	// GetOperators 获取所有支持的运算
	GetOperators(request struct{}, reply *[]string) error
}

func RegisterCalcService(svc ServiceInterface) error {
	return rpc.RegisterName(ServiceName, svc)
}

type CalcClient struct {
	*rpc.Client
}

var _ ServiceInterface = (*CalcClient)(nil)

func DialCalcService(network, address string) (*CalcClient, error) {
	conn, err := net.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &CalcClient{
		rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn)),
	}, nil
}

func (c *CalcClient) CalcTwoNumber(request Calc, reply *float64) error {
	return c.Client.Call(ServiceName+".CalcTwoNumber", request, reply)
}

func (c *CalcClient) GetOperators(request struct{}, reply *[]string) error {
	return c.Client.Call(ServiceName+".GetOperators", request, reply)
}
