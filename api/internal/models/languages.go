package models

type Language struct {
	Id           int    `json:"id"`
	Abbreviation string `json:"abbreviation,omitempty"`
	Name         string `json:"name"`
}

// NewLanguage is a function that adds a new language into database
func NewLanguage(language Language) (int64, error) {

	qp, err := DbConn.Prepare(`INSERT INTO LANGUAGES(abbreviation, name) VALUES (?, ?)`)
	if err != nil {
		return -1, err
	}

	res, err := qp.Exec(language.Abbreviation, language.Name)
	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

// GetLanguageById is a function that gets the language from the database by id
func (l *Language) GetLanguageById(id int64) (*Language, error) {

	var lang Language

	qp, err := DbConn.Prepare(`SELECT * FROM Languages WHERE id = ?`)
	if err != nil {
		return nil, err
	}

	rows, err := qp.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&lang)
	}

	return &lang, nil
}
