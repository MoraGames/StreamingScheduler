package models

type Resource struct {
	Id        int64     `json:"id"`
	Url       string    `json:"url"`
	Language  *Language `json:"language,omitempty"`
	Subtitles *Language `json:"subtitles,omitempty"`
	Quality   *Quality  `json:"quality,omitempty"`
	Episode   *Episode  `json:"episode"`
}

// NewResource is a series method that adds a new resource into database
func (r *Resource) NewResource() (int64, error) {

	var sub interface{}
	var qual interface{}

	// Check original language exist
	exist, err := r.Language.Exist()
	if err != nil {
		return -1, err
	}

	if !exist {
		r.Language.Id = 0
		lang, err := r.Language.NewLanguage()
		if err != nil {
			return -1, err
		}
		r.Language.Id = lang
	}

	// Check original subtitle exist
	if r.Subtitles != nil {
		exist, err = r.Subtitles.Exist()
		if err != nil {
			return -1, err
		}

		if !exist {
			r.Language.Id = 0
			sub, err := r.Subtitles.NewLanguage()
			if err != nil {
				return -1, err
			}
			r.Subtitles.Id = sub
		}

		sub = r.Subtitles.Id

	} else {
		sub = nil
	}

	// Check quality exist
	if r.Quality != nil {
		exist, err = r.Quality.Exist()
		if err != nil {
			return -1, err
		}

		if !exist {
			r.Quality.Id = 0
			quality, err := r.Quality.NewQuality()
			if err != nil {
				return -1, err
			}
			r.Quality.Id = quality
		}

		qual = r.Quality.Id

	} else {
		qual = nil
	}

	// Check episode exist
	exist, err = r.Episode.Exist()
	if err != nil {
		return -1, err
	}

	if !exist {
		r.Episode.Id = 0
		episode, err := r.Episode.NewEpisode()
		if err != nil {
			return -1, err
		}
		r.Episode.Id = episode
	}

	qp, err := DbConn.Prepare(`INSERT INTO Resources(url, language, subtitles, quality, episode) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return -1, err
	}

	res, err := qp.Exec(r.Url, r.Language.Id, sub, qual, r.Episode.Id)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

// GetResourceById is a function that gets the resource from the database by id
func GetResourceById(id int64) (*Resource, error) {

	var resource Resource
	var langId, subId, qualId, epId int64

	qp, err := DbConn.Prepare(`SELECT * FROM Resources WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	rows, err := qp.Query(id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&resource.Id, &resource.Url, &langId, &subId, &qualId, &epId)
	}

	// populate language
	resource.Language, err = GetLanguageById(langId)
	if err != nil {
		return nil, err
	}

	// populate subtitles
	resource.Subtitles, err = GetLanguageById(subId)
	if err != nil {
		return nil, err
	}

	// populate quality
	resource.Quality, err = GetQualityById(qualId)
	if err != nil {
		return nil, err
	}

	// populate episode
	resource.Episode, err = GetEpisodeById(epId)
	if err != nil {
		return nil, err
	}

	return &resource, nil
}

// Exist Check if the resources exist
func (r *Resource) Exist() (bool, error) {
	resource, err := GetResourceById(r.Id)
	if err != nil {
		return false, err
	}

	if resource.Id == 0 {
		return false, nil
	}

	return true, nil
}
