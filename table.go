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
	Rows [][]RowValue
}

// Returns rows as only the text values
func (t *Table) rowText() [][]string {

	var result [][]string
	for _, row := range t.Rows {
		var stringColumns []string
		for _, c := range row {
			stringColumns = append(stringColumns, c.Value)
		}
		result = append(result, stringColumns)
	}
	return result
}

type RowValue struct {
	Value string `json:"value"`
	Link  string `json:"link"`
}

func NewTableFromHeaderAndRows(header []string, rows [][]RowValue) Table {
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
	return headers
}

func getRows(rowElements *colly.HTMLElement) [][]RowValue {
	var result [][]RowValue
	rowElements.ForEach("tr", func(i int, rowElement *colly.HTMLElement) {
		var row [] RowValue
		rowElement.ForEach("td", func(j int, columnElement *colly.HTMLElement) {
			contentValue := columnElement.Text
			link := columnElement.ChildAttr("a", "href")
			row = append(row, RowValue{
				contentValue,
				link,
			})
		})
		if len(row) > 0 {
			result = append(result, row)
		}
	})
	return result
}

func (t *Table) Map() []map[string]interface{} {
	var result []map[string]interface{}
	for _, row := range t.Rows {
		var rowMap = make(map[string]interface{})
		for j, value:= range row {
			name := t.Headers[j]
			contentMap := make(map[string]string)
			contentMap["value"] = value.Value
			contentMap["link"] = value.Link
			rowMap[name] = contentMap
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
	for _, v := range t.rowText() {
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