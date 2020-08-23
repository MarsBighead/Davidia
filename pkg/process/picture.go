package process

import (
	"Davidia/pkg/util"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"time"
)

//Report check report
type Report struct {
	Header []string
	Data   map[string]*data
	Date   string
}

type data struct {
	Date       *time.Time
	Name       string
	ID         string
	QRFilename string

	//status 1. submitted, -1 unsubmitted,
	Status  int
	Comment string
	//IsOk filename with right format of not
	IsOk bool
}

// Check student
func Check(qrCodePath string, ss map[string]*util.Student) (*Report, error) {

	fs, err := ioutil.ReadDir(qrCodePath)
	if err != nil {
		return nil, err
	}
	var n, rn int
	r := make(map[string]*data)
	for _, f := range fs {
		if f.IsDir() {
			continue
		} else {
			n++
			//fmt.Println("QR Pic:", f.Name())
			filename := f.Name()
			d := new(data)
			d.getQRFilename(filename)
			d.getID(filename)
			if _, ok := ss[d.ID]; ok {
				d.IsOk = ok
				rn++
			} else {
				d.ID = fmt.Sprintf("%d", n)
			}
			d.getName(filename)
			d.Status = 1
			r[d.ID] = d
			log.Printf("Name %s find with ID=%s at IsOK status %v.\n", d.Name, d.ID, d.IsOk)
		}

	}
	for id, v := range ss {
		if _, ok := r[id]; ok {
			continue
		} else {
			d := &data{
				Name:   v.Name,
				ID:     id,
				Status: -1,
			}
			log.Printf("Not find one's health QR Code for Name %s with ID=%s.\n", d.Name, d.ID)
			r[id] = d
		}
	}
	log.Printf("Right ID number is %d, total QR Code Pic is %d.\n", rn, n)
	return &Report{
		Data: r,
	}, nil
}

func (d *data) getName(s string) {
	d.Name = strings.Split(s, "-")[0]
}

func (d *data) getQRFilename(s string) {
	d.QRFilename = strings.Split(s, ".")[0]
}

func (d *data) getID(filename string) {

	re := regexp.MustCompile(`(2020\d{6})`)
	ss := re.FindStringSubmatch(filename)
	if len(ss) > 0 {
		d.ID = ss[0]
	} else {
		d.Status = -1
	}
}
