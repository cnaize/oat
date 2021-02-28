package sourcejson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/cnaize/oat/database/source"
)

type DataSourceJSON struct {
	path string
	data []interface{}
}

func NewDataSourceJSON(dataType source.DataType) (*DataSourceJSON, error) {
	ds := DataSourceJSON{
		data: []interface{}{},
	}
	switch dataType {
	case source.DataTypeQuestions:
		ds.path = filepath.FromSlash("./database/source/sourcejson/questions.json")
	default:
		return nil, fmt.Errorf("data type not supported: %v", dataType)
	}

	file, err := os.Open(ds.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &ds.data); err != nil {
		return nil, err
	}

	return &ds, nil
}

func (ds *DataSourceJSON) List() ([]interface{}, error) {
	return ds.data, nil
}

func (ds *DataSourceJSON) Create(item interface{}) error {
	if item == nil {
		return fmt.Errorf("empty data")
	}
	ds.data = append(ds.data, item)
	data, err := json.Marshal(&ds.data)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(ds.path, data, 0755)
}
