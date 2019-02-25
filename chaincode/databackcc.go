package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("DataBackCC")

// New returns an implementation of the chaincode interface
func New() shim.Chaincode {
	return &DataBackCC{}
}

// DataBackCC for scc plugin test
type DataBackCC struct{}

// Init implements the chaincode shim interface
func (s *DataBackCC) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("Init Success")
	return shim.Success(nil)
}

// Invoke implements the chaincode shim interface
func (s *DataBackCC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("Invoke")
	rsp := stub.InvokeChaincode("databackscc", [][]byte{[]byte(""), []byte(stub.GetTxID())}, stub.GetChannelID())
	return shim.Success(rsp.GetPayload())
}

func main() {
	err := shim.Start(&DataBackCC{})
	if err != nil {
		logger.Error(err)
	}
}
