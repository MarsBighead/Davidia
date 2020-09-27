package util

import (
	"strings"

	"github.com/tealeg/xlsx"
)

// Student information structure
type Student struct {
	Name string
	ID   string
}

type Students struct {
	Header  map[string]int
	Student map[string]*Student
}

//Get student information for source
func (src *Source) Get(home string) (*Students, error) {
	f, err := xlsx.OpenFile(home + "/" + src.Name)
	if err != nil {
		return nil, err
	}
	ss := new(Students)
	ss.Student = make(map[string]*Student)
	for _, sheet := range f.Sheets {
		if len(sheet.Rows) < 10 {
			continue
		}
		for y, row := range sheet.Rows {
			if y == 0 {
				ss.getHeader(row)
			} else {
				ss.getData(row)
			}
		}

	}

	return ss, nil
}

//ExtractBpHeader  get header row value and coordinate

func (ss *Students) getHeader(row *xlsx.Row) {
	header := make(map[string]int)
	for x, cell := range row.Cells {
		val := strings.Trim(cell.String(), " ")
		if val == "姓名" || val == "学号" {
			header[val] = x
		}
	}
	ss.Header = header
}

//ExtractBpHeader  get header row value and coordinate

func (ss *Students) getData(row *xlsx.Row) {
	//header := make(map[string]int)
	s := new(Student)
	h := ss.Header
	s.Name = row.Cells[h["姓名"]].String()
	s.ID = row.Cells[h["学号"]].String()
	ss.Student[s.ID] = s
}
