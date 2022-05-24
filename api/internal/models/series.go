package models

type Series struct {
	Id             int64   `json:"id"`
	OriginalTitle  *Title  `json:"originalTitle"`
	Plot           string  `json:"plot,omitempty"`
	Format         *Format `json:"format"`
	Favorites      int     `json:"favorites"`
	AgeRestriction uint    `json:"ageRestriction,omitempty"`
}

// NewSeries is a series method that adds a new series into database
func (s *Series) NewSeries() (int64, error) {

	// Check original title exist
	exist, err := s.OriginalTitle.Exist()
	if err != nil {
		return -1, err
	}

	if !exist {
		s.OriginalTitle.Id = 0
		title, err := s.OriginalTitle.NewTitle()
		if err != nil {
			return -1, err
		}
		s.OriginalTitle.Id = title
	}

	// Check format exist
	exist, err = s.Format.Exist()
	if err != nil {
		return -1, err
	}

	if !exist {
		s.Format.Id = 0
		format, err := s.Format.NewFormat()
		if err != nil {
			return -1, err
		}
		s.Format.Id = format
	}

	qp, err := DbConn.Prepare(`INSERT INTO Series(originalTitle, plot, format, favorites, ageRestriction) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return -1, err
	}

	res, err := qp.Exec(s.OriginalTitle.Id, s.Plot, s.Format.Id, s.Favorites, s.AgeRestriction)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

// GetSeriesById is a function that gets the series from the database by id
func GetSeriesById(id int64) (*Series, error) {

	var series Series

	qp, err := DbConn.Prepare(`SELECT * FROM Series WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	rows, err := qp.Query(id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&series.Id, series.OriginalTitle, series.Plot, series.Format, series.Favorites, series.AgeRestriction)
	}

	return &series, nil
}

// Exist Check if the series exist
func (s *Series) Exist() (bool, error) {
	series, err := GetSeriesById(s.Id)
	if err != nil {
		return false, err
	}

	if series.Id == 0 {
		return false, nil
	}

	return true, nil
}
