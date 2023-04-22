package park

import (
	"encoding/json"

	"fmt"
)

type Links struct {
	Type LinksType `json:"type"`
	Link string    `json:"link"`
}

func (e Links) String() string {
	res, _ := json.Marshal(e)
	return string(res)
}

type LinksCollection []Links

func LinksFromJSON(data json.RawMessage) LinksCollection {
	entity := []Links{}
	if err := json.Unmarshal(data, &entity); err != nil {
		fmt.Println("unmarshal error", err)
	}
	return entity
}

func (e Links) LinksToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error", err)
	}
	return res
}

func (e LinksCollection) LinksToJSON() json.RawMessage {
	res, err := json.Marshal(e)
	if err != nil {
		fmt.Println("marshal error", err)
	}
	return res
}
