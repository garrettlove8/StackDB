package database

import "time"

type Data struct {
	Uuid  string            `json:"uuid"`
	CTime string            `json:"cTime"`
	MTime string            `json:"mTime"`
	Body  map[string][]byte `json:"body"`
}

func (d *Data) Edit() (*Data, error) {
	data := Data{
		Uuid:  d.Uuid,
		MTime: time.Now().String(),
		Body:  d.Body,
	}
	return &data, nil
}

func (d *Data) Delete() error {
	return nil
}
