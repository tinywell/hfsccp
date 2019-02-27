package main

import (
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
	args := stub.GetStringArgs()
	for _, s := range args {
		logger.Infof("BackUp String: '%s'\n", s)
		filePath := filepath.Join(backpath, backfile)
		f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logger.Error(err)
		}
		_, err = f.WriteString(s + "\n")
		if err != nil {
			logger.Error(err)
		}
	}
	return shim.Success(nil)
}

func checkExist() {
	filePath := filepath.Join(backpath, backfile)
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsExist(err) {
			logger.Debug("Exist")
			return
		}
	} else {
		return
	}
	logger.Infof("Not Exist, Create Dir %s", backpath)
	err = os.MkdirAll(backpath, 666)
	if err != nil {
		logger.Error(err)
	}
	logger.Infof("Create File %s", filePath)
	_, err = os.Create(filePath)
	if err != nil {
		logger.Error(err)
	}
}

func main() {}
