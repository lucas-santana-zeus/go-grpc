package models

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
