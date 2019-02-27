package main

import (
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func TestInvoke(t *testing.T) {
	sccp := &DataBackSCC{}
	mockStub := shim.NewMockStub("tx", sccp)
	mockStub.MockTransactionStart("aaaa")
	res := mockStub.MockInvoke("TEST001", [][]byte{[]byte("func"), []byte("Hello world")})
	t.Log(res)
	mockStub.MockTransactionEnd("aaaa")
}
