package models

type Language struct {
	Id           int64  `json:"id"`
	Abbreviation string `json:"abbreviation,omitempty"`
	Name         string `json:"name"`
}

// NewLanguage is a language method that adds a new language into database
func (l *Language) NewLanguage() (int64, error) {

	qp, err := DbConn.Prepare(`INSERT INTO Languages(abbreviation, name) VALUES (?, ?)`)
	if err != nil {
		return -1, err
	}

	res, err := qp.Exec(l.Abbreviation, l.Name)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

// GetLanguageById is a function that gets the language from the database by id
func GetLanguageById(id int64) (*Language, error) {

	var lang Language

	qp, err := DbConn.Prepare(`SELECT * FROM Languages WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	rows, err := qp.Query(id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&lang.Id, &lang.Abbreviation, &lang.Name)
	}

	// Check lang exist
	if lang.Id == 0 {
		return nil, nil
	}

	return &lang, nil
}
