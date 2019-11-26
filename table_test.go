package table

import (
	_ "github.com/stretchr/testify/assert"
	"testing"
)

const url1 = "https://www.sec.gov/Archives/edgar/data/320193/000032019319000119/0000320193-19-000119-index.htm"
const url2 = "https://www.sec.gov/cgi-bin/browse-edgar?company=Google&match=&filenum=&State=&Country=&SIC=&myowner=exclude&action=getcompany"
const url3 = "http://eoddata.com/stocklist/NYSE/B.htm"

/*
func TestTable(t *testing.T) {
	table := Table{}
	table.Get(url2, 0)
	table.Print()
	table.PrintJSON()
}
*/

func TestTables(t *testing.T) {
	tables, _ := GetTables(url3, "table")
	for _, table := range tables {
		table.PrintJSON()
	}
}