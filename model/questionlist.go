package model

import (
	"github.com/mitchellh/mapstructure"
)

// QuestionList model
type QuestionList struct {
	Data []Question `json:"data"`
}

func (ql *QuestionList) SetData(data []interface{}) {
	ql.Data = make([]Question, len(data))
	for i, d := range data {
		switch d.(type) {
		case map[string]interface{}:
			mapstructure.Decode(d, &ql.Data[i])
		default:
			ql.Data[i] = d.(Question)
		}
	}
}
