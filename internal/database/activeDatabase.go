package database

type ActiveDatabase struct {
	Id          string
	Name        string
	Type        string
	ColStrings  []string
	Collections []Collection
}

func (ad *ActiveDatabase) Load() (*ActiveDatabase, error) {
	data, err := readDbFile(ad.Name)
	if err != nil {
		return nil, err
	}

	db, err := decodeDbFile(data)
	if err != nil {
		return nil, err
	}

	ad.Id = db.Id
	ad.Type = db.Type

	for _, v := range db.Collections {
		data, _ := readColFile(db.Name, v)
		col, _ := decodeColFile(data)
		ad.Collections = append(ad.Collections, *col)
	}

	return ad, nil
}
