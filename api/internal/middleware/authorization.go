package middleware

import (
	"net/http"
	"errors"
	"github.com/mauriciomolinapicco/lets-go/api/api"
	"github.com/mauriciomolinapicco/lets-go/api/internal/tools"

	log "github.com/sirupsen/logrus"

)

var UnAuthorizedError = errors.New("invalid user or token")

func Authorization(next http.Handler) http.Handler {
	//        					(response writer, pointer to request)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		//instanciar puntero a db
		var database *tools.DatabaseInterface
		database, err = toolsNewDatabase()
		if err != nil {
			api.InternalErrorHandler(w, err)
			return
		}

		var loginDetails *tools.loginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
	} 
}

