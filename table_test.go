package table

import (
	_ "github.com/stretchr/testify/assert"
	"testing"
)

const url = "https://www.sec.gov/Archives/edgar/data/320193/000032019319000119/0000320193-19-000119-index.htm"

func TestTable(t *testing.T) {
	table := Table{}
	table.Get(url, 1)
	table.Print()
	table.PrintJSON()
}
