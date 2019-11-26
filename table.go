package table

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/olekukonko/tablewriter"
	"os"
)

type Table struct {
	Headers []string
	Rows [][]string
}

func NewTableFromHeaderAndRows(header []string, rows [][]string) Table {
	table := Table{Headers:header, Rows:rows}

	// If there are fewer headers than columns,
	// fill in the blanks
	nHeaders := len(table.Headers)
	nColumns := len(table.Rows[0])
	fmt.Println(nHeaders, nColumns)
	difference := nColumns - nHeaders

	for i := 0; i < difference; i ++ {
		table.Headers = append(table.Headers, "___")
	}
	return table
}

func getHeader(h *colly.HTMLElement) []string {
	var headers []string
	h.ForEach("tr th", func(_ int, el *colly.HTMLElement) {
		headers = append(headers, el.Text)
	})
	fmt.Println(headers)
	return headers
}

func getRows(rowElements *colly.HTMLElement) [][]string {
	var result [][]string
	rowElements.ForEach("tr", func(i int, rowElement *colly.HTMLElement) {
		var row [] string
		rowElement.ForEach("td", func(j int, columnElement *colly.HTMLElement) {
			row = append(row, columnElement.Text)
		})
		if len(row) > 0 {
			result = append(result, row)
		}
	})
	return result
}

func (t *Table) Map() []map[string]string {
	var result []map[string]string
	for _, row := range t.Rows {
		var rowMap = make(map[string]string)
		for j, value:= range row {
			name := t.Headers[j]
			rowMap[name] = value
		}
		result = append(result, rowMap)
	}
	return result
}

func (t *Table) JSON() ([]byte, error) {
	m := t.Map()
	return json.Marshal(m)
}

func (t *Table) Print() {
	twriter := tablewriter.NewWriter(os.Stdout)
	twriter.SetHeader(t.Headers)
	for _, v := range t.Rows {
		twriter.Append(v)
	}
	twriter.Render()
}

func (t *Table) PrintJSON() error {
	m := t.Map()
	b, e := json.MarshalIndent(m, "", "\t")
	if e != nil {
		return e
	}
	fmt.Println(string(b))
	return nil
}