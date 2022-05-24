package models

type Quality struct {
	Id         int64  `json:"id,omitempty"`
	Quality    string `json:"quality"`
	Resolution string `json:"resolution,omitempty"`
}

// NewQuality is a quaity method that adds a new quality into database
func (q *Quality) NewQuality() (int64, error) {

	qp, err := DbConn.Prepare(`INSERT INTO Qualities (Quality, Resolution) VALUES (?, ?)`)
	if err != nil {
		return -1, err
	}

	res, err := qp.Exec(q.Quality, q.Resolution)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

// GetQualityById is a function that gets the quality from the database by id
func GetQualityById(id int64) (*Quality, error) {

	var quality Quality

	qp, err := DbConn.Prepare(`SELECT * FROM Qualities WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	rows, err := qp.Query(id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&quality.Id, &quality.Quality, &quality.Resolution)
	}

	return &quality, nil
}

// Exist Check if the quality exist
func (q *Quality) Exist() (bool, error) {

	quality, err := GetQualityById(q.Id)
	if err != nil {
		return false, err
	}

	if quality.Id == 0 {
		return false, nil
	}

	return true, nil
}
