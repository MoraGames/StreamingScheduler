package models

type Episode struct {
	Id             int64   `json:"id,omitempty"`
	Series         *Series `json:"series"`
	Number         int     `json:"number,omitempty"`
	OriginalTitle  *Title  `json:"originalTitle,omitempty"`
	Plot           string  `json:"plot,omitempty"`
	AgeRestriction uint    `json:"ageRestriction,omitempty"`
}

// NewEpisode is a episode method that adds a new episode into database
func (e *Episode) NewEpisode() (int64, error) {

	// Check series exist
	exist, err := e.Series.Exist()
	if err != nil {
		return -1, err
	}

	if !exist {
		e.Series.Id = 0
		series, err := e.Series.NewSeries()
		if err != nil {
			return -1, err
		}
		e.Series.Id = series
	}

	// Check original title exist
	exist, err = e.OriginalTitle.Exist()
	if err != nil {
		return -1, err
	}

	if !exist {
		e.OriginalTitle.Id = 0
		title, err := e.OriginalTitle.NewTitle()
		if err != nil {
			return -1, err
		}
		e.OriginalTitle.Id = title
	}

	qp, err := DbConn.Prepare(`INSERT INTO Episodes(series, number, originalTitle, plot, ageRestriction) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return -1, err
	}

	res, err := qp.Exec(e.Series.Id, e.Number, e.OriginalTitle.Id, e.Plot, e.AgeRestriction)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

// GetEpisodeById is a function that gets the episode from the database by id
func GetEpisodeById(id int64) (*Episode, error) {

	var ep Episode
	var seriesId, originTitleId int64

	qp, err := DbConn.Prepare(`SELECT * FROM Episodes WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	rows, err := qp.Query(id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&ep.Id, &seriesId, &ep.Number, &originTitleId, &ep.Plot, &ep.AgeRestriction)
	}

	// Populate series
	ep.Series, err = GetSeriesById(seriesId)
	if err != nil {
		return nil, err
	}

	ep.OriginalTitle, err = GetTitleById(originTitleId)
	if err != nil {
		return nil, err
	}

	return &ep, nil
}

// Exist Check if the episode exist
func (ep *Episode) Exist() (bool, error) {
	episode, err := GetSeriesById(ep.Id)
	if err != nil {
		return false, err
	}

	if episode.Id == 0 {
		return false, nil
	}

	return true, nil
}
