package controllers

import (
	"errors"
	"fmt"
	"net/http"
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

//PrintInternalErr imposta a 500 lo status code della risposta HTTP.
func PrintInternalErr(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("{\"code\": 500, \"msg\": \"Internal Server Error\"}\n"))
}
