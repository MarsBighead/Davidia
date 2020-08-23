package process

import (
	"github.com/tealeg/xlsx"
)

// Output report filename
func (r *Report) Output(filename string) error {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("QRCode Status")
	if err != nil {
		return err
	}
	r.addTitle(sheet)
	r.addData(sheet)
	return file.Save(filename)

}

//addTitle for table in Excel sheet
func (r *Report) addTitle(sheet *xlsx.Sheet) {
	row := sheet.AddRow()
	for _, header := range r.Header {
		cell := row.AddCell()
		cell.Value = header
	}
}

//addData for table in Excel sheet
func (r *Report) addData(sheet *xlsx.Sheet) {

	for _, v := range r.Data {
		row := sheet.AddRow()
		row.AddCell().SetString(r.Date)
		row.AddCell().SetString(v.Name)
		row.AddCell().SetString(v.ID)

		if v.Status == 1 {
			row.AddCell().SetString("Submitted")
		} else {
			row.AddCell().SetString("Unsubmitted")
		}
		if v.IsOk {
			row.AddCell().SetString("Ok")
		} else {
			row.AddCell().SetString("Not Ok")
		}
		row.AddCell().SetString(v.QRFilename)
	}
}
