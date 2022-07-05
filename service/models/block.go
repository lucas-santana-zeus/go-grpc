package models

import block "go-grpc/commons/pb"

var blocks = []block.ResponseBlock{{
	Id:       "F1",
	ParentId: "C0",
}}

func GetBlockById(id string) *block.ResponseBlock {
	for i := 0; i < len(blocks); i++ {
		if blocks[i].GetId() == id {
			return &blocks[i]
		}
	}
	return nil
}
