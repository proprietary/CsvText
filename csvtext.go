package csvtext

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"strings"
)

type CsvText struct {
	Csv string
	Arr []string
}

func makeArr(src string) []string {
	ret := strings.Split(src, ",")
	for i, v := range ret {
		ret[i] = strings.Trim(v, " ,")
	}
	return ret
}

func (c *CsvText) SetCsv(src string) *CsvText {
	c.Csv = src
	c.Arr = makeArr(src)
	return c
}

func (c *CsvText) SetArr(src []string) *CsvText {
	c.Arr = src
	b := new(bytes.Buffer)
	q := len(src) - 1
	for i, v := range src {
		b.WriteString(v)
		if i != q {
			b.WriteRune(',')
		}
	}
	c.Csv = b.String()
	return c
}

func (c *CsvText) Value() (driver.Value, error) {
	return c.Csv, nil
}

func (c *CsvText) Scan(src interface{}) error {
	s, ok := src.(string)
	if !ok {
		return fmt.Errorf("CsvText must be able to scan this as a string!")
	}
	c = c.SetCsv(s)
	return nil
}
