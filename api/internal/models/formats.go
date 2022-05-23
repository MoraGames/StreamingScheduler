package models

type Format struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
}

// NewFormat is a format method that adds a new format into database
func (f *Format) NewFormat() (int64, error) {

	qp, err := DbConn.Prepare(`INSERT INTO Formats(type) VALUES (?)`)
	if err != nil {
		return -1, err
	}

	res, err := qp.Exec(f.Type)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

// GetFormatById is a function that gets the format from the database by id
func GetFormatById(id int64) (*Format, error) {

	var f Format

	qp, err := DbConn.Prepare(`SELECT * FROM Formats WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	rows, err := qp.Query(id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&f.Id, &f.Type)
	}

	// Check lang exist
	if f.Id == 0 {
		return nil, nil
	}

	return &f, nil
}
