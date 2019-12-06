package table

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
)

// Fetches all the tables on a given page
func GetTables(url, selector string) (*Tables, error) {

	if selector == "" {
		selector = "table"
	}

	var tables []Table
	c := colly.NewCollector()
	c.OnHTML(selector, func(e *colly.HTMLElement) {
		header := getHeader(e)
		rows := getRows(e)
		table := NewTableFromHeaderAndRows(header, rows)
		tables = append(tables, table)
	})

	err := c.Visit(url)

	if err != nil {
		return nil, err
	}

	c.Wait()

	return &Tables{tables}, nil
}

type Tables struct {
	Tables []Table
}

func (t *Tables) Maps() [][]map[string]interface{} {
	var maps [][]map[string]interface{}
	for _, table := range t.Tables {
		tableMap := table.Map()
		maps = append(maps, tableMap)
	}
	return maps
}

func (t *Tables) JSON() ([]byte, error) {
	return json.Marshal(t.Maps())
}

func (t *Tables) PrintJSON() error {
	m := t.Maps()
	b, e := json.MarshalIndent(m, "", "\t")
	if e != nil {
		return e
	}
	fmt.Println(string(b))
	return nil
}

