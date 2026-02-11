package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mauriciomolinapicco/lets-go/api"
	"github.com/mauriciomolinapicco/lets-go/api/internal/tools"
	log "github.com/sirupsen/logrus"
	"github.com/gorilla/schema"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err = error

	err = decoder.Decode(&params, r.URL.Query()) //grab parameters on the url and set them to values on the struct

	if err != nil {
		log.Error(err)
		api.RequestErrorHandler(w, err)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	
	if err != nil {
		api.InternalErrorHandler(w)
		return
	}

	var tokenDetails *tools.TokenDetails
	tokenDetails = (*database).GetUserTokenDetails(params.Username)

	if tokenDetails == nil {
		log.Error(api.UnAuthorizedError)
		api.RequestErrorHandler(w, api.UnAuthorizedError)
		return
	}

	var response = api.CoinBalanceResponse{
		Code: http.StatusOK,
		Balance: (*tokenDetails).Coins,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
