package models

import block "go-grpc/commons/pb"

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
	responseBlock.Id = blockDAO.ID
	responseBlock.ParentId = blockDAO.ParentID
	return &responseBlock
}

func TransformResponseIntoBlockDTO(response *block.ResponseBlock) Block {
	return Block{
		ID:       response.GetId(),
		ParentID: response.GetParentId(),
	}
}
