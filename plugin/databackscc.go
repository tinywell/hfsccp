package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

const (
	backpath = "/data/backup"
	backfile = "backup.txt"
)

var logger = shim.NewLogger("DataBackSCC")

// New returns an implementation of the chaincode interface
func New() shim.Chaincode {
	return &DataBackSCC{}
}

// DataBackSCC for scc plugin test
type DataBackSCC struct{}

// Init implements the chaincode shim interface
func (s *DataBackSCC) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("Init Success")
	return shim.Success(nil)
}

// Invoke implements the chaincode shim interface
func (s *DataBackSCC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("Invoke")
	checkExist()
	_, args := stub.GetFunctionAndParameters()
	for _, s := range args {
		logger.Infof("Back String: '%s'\n", s)
		filePath := filepath.Join(backpath, backfile)
		err := ioutil.WriteFile(filePath, []byte(s), os.ModeAppend)
		if err != nil {
			fmt.Println(err)
		}
	}
	return shim.Success(nil)
}

func checkExist() {
	filePath := filepath.Join(backpath, backfile)
	_, err := os.Open(filePath)
	if os.IsExist(err) {
		return
	}
	os.MkdirAll(backpath, 666)
}

func main() {}
