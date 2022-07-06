package models

import (
	"encoding/json"
	block "go-grpc/commons/pb"
)

type Block struct {
	ID       string `json:"id,omitempty"`
	ParentID string `json:"parent_id,omitempty"`
}

var blocks = []Block{{
	ID:       "F1",
	ParentID: "C0",
}}

func GetBlockById(id string) Block {
	for _, b := range blocks {
		if b.ID == id {
			return b
		}
	}
	return Block{}
}

func TransformBlockDAOIntoResponse(blockDAO Block) *block.ResponseBlock {
	var responseBlock block.ResponseBlock
	blockBytes, err := json.Marshal(blockDAO)
	if err != nil {
		return nil
	}
	responseBlock.BlockJson = string(blockBytes)
	return &responseBlock
}

func TransformResponseIntoBlockDTO(response *block.ResponseBlock) Block {
	var resBlock Block
	err := json.Unmarshal([]byte(response.BlockJson), &resBlock)
	if err != nil {
		return Block{}
	}
	return resBlock
}
