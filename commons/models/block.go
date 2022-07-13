package models

import (
	"encoding/json"
	"fmt"
	block "go-grpc/commons/pb"
	"time"
)

type Block struct {
	DataTimestamp     time.Time `json:"data_timestamp"`
	CreatedTimestamp  time.Time `json:"created_timestamp"`
	TemperatureInst   string    `json:"temperature_inst"`
	TemperatureMin    string    `json:"temperature_min"`
	TemperatureMax    string    `json:"temperature_max"`
	PrecipitationInst string    `json:"precipitation_inst"`
	PrecipitationMin  string    `json:"precipitation_min"`
	PrecipitationMax  string    `json:"precipitation_max"`
}

//var blocks = []Block{{
//	CreatedTimestamp: time.Now(),
//	DataTimestamp:    time.Now(),
//}}

//func GetBlockById(id string) Block {
//	//for _, b := range blocks {
//	//	if b.ID == id {
//	//		return b
//	//	}
//	//}
//	return Block{}
//}

func TransformBlockDAOIntoResponse(blockDAO Block) *block.ResponseBlock {
	var responseBlock block.ResponseBlock
	blockBytes, err := json.Marshal(blockDAO)
	if err != nil {
		fmt.Println("error marshalling block", err)
		return nil
	}
	responseBlock.BlockJson = string(blockBytes)
	return &responseBlock
}

func TransformResponseIntoBlockDTO(response *block.ResponseBlock) Block {
	var resBlock Block
	err := json.Unmarshal([]byte(response.BlockJson), &resBlock)
	if err != nil {
		fmt.Println("error unmarshalling block", err)
		return Block{}
	}
	return resBlock
}
