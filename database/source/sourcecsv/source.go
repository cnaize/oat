package sourcecsv

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
	"github.com/jszwec/csvutil"

	"github.com/cnaize/oat/database/source"
)

type DataSourceCSV struct {
	path string
	data []interface{}
}

func NewDataSourceCSV(dataType source.DataType) (*DataSourceCSV, error) {
	ds := DataSourceCSV{
		data: []interface{}{},
	}
	switch dataType {
	case source.DataTypeQuestions:
		ds.path = filepath.FromSlash("./database/source/sourcecsv/questions.csv")
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
	if err := csvutil.Unmarshal(data, &ds.data); err != nil {
		return nil, err
	}

	return &ds, nil
}

func (ds *DataSourceCSV) List() ([]interface{}, error) {
	return ds.data, nil
}

func (ds *DataSourceCSV) Create(item interface{}) error {
	if item == nil {
		return fmt.Errorf("empty data")
	}
	ds.data = append(ds.data, item)
	file, err := os.Create(ds.path)
	if err != nil {
		return err
	}
	defer file.Close()
	if err := gocsv.MarshalFile(ds.data, file); err != nil {
		return err
	}
	return nil
}
