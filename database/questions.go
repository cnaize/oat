package database

import (
	"fmt"
	"time"

	"github.com/cnaize/oat/database/source"
	"github.com/cnaize/oat/database/source/sourcecsv"
	"github.com/cnaize/oat/database/source/sourcejson"
	"github.com/cnaize/oat/model"
)

// QuestionsDB represents database for questions
type Questions struct {
	source source.DataSource
}

func NewQuestions(sourceType SourceType) (*Questions, error) {
	var db Questions
	var err error
	switch sourceType {
	case SourceTypeCSV:
		if db.source, err = sourcecsv.NewDataSourceCSV(source.DataTypeQuestions); err != nil {
			return nil, err
		}
	case SourceTypeJSON:
		if db.source, err = sourcejson.NewDataSourceJSON(source.DataTypeQuestions); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("source type not supported: %v", sourceType)
	}
	return &db, nil
}

func (q *Questions) List() (model.QuestionList, *model.Error) {
	var res model.QuestionList
	list, err := q.source.List()
	if err != nil {
		return res, &model.Error{Text: err.Error()}
	}
	res.SetData(list)
	return res, nil
}

func (q *Questions) Create(item model.Question) *model.Error {
	item.CreatedAt = time.Now()
	if err := q.source.Create(item); err != nil {
		return &model.Error{Text: err.Error()}
	}
	return nil
}
