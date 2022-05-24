package models

type Title struct {
	Id       int64     `json:"id,omitempty"`
	Title    string    `json:"title"`
	Language *Language `json:"language"`
}

// NewTitle is a title method that adds a new title into database
func (t *Title) NewTitle() (int64, error) {

	lang, err := GetLanguageById(t.Language.Id)
	if err != nil {
		return -1, err
	}

	// if language doesn't exist
	if lang.Id == 0 {
		t.Language.Id = 0
		t.Language.Id, err = t.Language.NewLanguage()
	}

	qp, err := DbConn.Prepare(`INSERT INTO Titles(title, language) VALUES (?, ?)`)
	if err != nil {
		return -1, err
	}

	res, err := qp.Exec(t.Title, t.Language.Id)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

// GetTitleById is a function that gets the title from the database by id
func GetTitleById(id int64) (*Title, error) {

	var title Title
	var langId int64

	// Get title info from db
	qp, err := DbConn.Prepare(`SELECT * FROM Titles WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	rows, err := qp.Query(id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&title.Id, &title.Title, &langId)
	}

	// get language info of the title from db
	title.Language, err = GetLanguageById(langId)
	if err != nil {
		return nil, err
	}

	return &title, nil
}

// Exist is a title method that checks if title exist in the database
func (t *Title) Exist() (bool, error) {

	title, err := GetTitleById(t.Id)
	if err != nil {
		return false, err
	}

	if title.Id == 0 {
		return false, nil
	}

	return true, nil
}
