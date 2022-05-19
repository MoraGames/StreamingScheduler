package utils

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gorilla/securecookie"
	"github.com/segmentio/ksuid"
	"net/http"
	"path/filepath"
	"regexp"
	"text/template"
)

type ParamsInfo struct {
	Key      string
	Required bool
}

//GetParams ritorna i parametri inviati tramite metodo GET dell'HTTP request.
func GetParams(params []ParamsInfo, r *http.Request) (map[string]interface{}, error) {

	values := make(map[string]interface{})

	for _, param := range params {

		fmt.Println(param)

		keys, err := r.URL.Query()[param.Key]

		fmt.Println(err)

		if (!err || len(keys[0]) < 1) && param.Required == true {
			return nil, errors.New(fmt.Sprintf("Url Param \"%s\" is missing", param.Key))
		}

		if err == false || len(keys[0]) < 1 {
			continue
		}

		values[param.Key] = keys[0]
	}

	return values, nil

}

//Restituisce l'IP del client che ha effettuato la richiesta.
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

//PrintErr ritorna un errore al client impostando a 400 lo status code della risposta HTTP.
func PrintErr(w http.ResponseWriter, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(fmt.Sprintf("{\"code\": 400, \"msg\": \"%s\"}\n", err)))
}

func IsEmailValid(email string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(email) < 3 && len(email) > 254 {
		return false
	}
	return emailRegex.MatchString(email)
}

//PrintInternalErr imposta a 500 lo status code della risposta HTTP.
func PrintInternalErr(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("{\"code\": 500, \"msg\": \"Internal Server Error\"}\n"))
}

func GenerateID() string {
	return ksuid.New().String()
}

func SetCookies(w http.ResponseWriter, token, tokenIss, cookieKey string) error {

	var s = securecookie.New([]byte(cookieKey), nil)

	value := map[string]string{
		"RefreshToken": token,
	}
	if encoded, err := s.Encode(tokenIss, value); err == nil {
		cookie := &http.Cookie{
			Name:     tokenIss,
			Value:    encoded,
			Secure:   false,
			HttpOnly: false,
			Path:     "/",
		}
		fmt.Println("COOKIE", cookie)
		http.SetCookie(w, cookie)
	} else {
		return err
	}

	return nil
}

func ParseTemplate(tmpl string, data interface{}) (string, error) {

	templatePath, err := filepath.Abs(tmpl)
	if err != nil {
		return "", errors.New("invalid template name")
	}

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)

	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
