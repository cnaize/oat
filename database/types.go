package database

type SourceType string

var (
	SourceTypeCSV  SourceType = "csv"
	SourceTypeJSON SourceType = "json"
)

type DB struct {
	Questions *Questions
}
