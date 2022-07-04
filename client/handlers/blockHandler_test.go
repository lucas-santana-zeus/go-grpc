package handlers

import (
	block "go-grpc/commons/pb"
	"testing"
)

var blockTest = &block.ResponseBlock{
	Id:       "F1",
	ParentId: "C0",
}

func TestGetBlockById(t *testing.T) {
	//	todo: test existent block
	//	todo: test nonexistent block

	t.Run("Getting existent block", func(t *testing.T) {

	})

	t.Run("Getting non existent block", func(t *testing.T) {

	})
}
