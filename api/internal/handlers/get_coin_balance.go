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
	var params = api.CoinBalanceParams