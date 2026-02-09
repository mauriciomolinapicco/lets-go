package api

import (
	"encoding/json"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	//generalmente 200
	Code int
	Balance int64
}

type Error struct {
	Code int
	Message string //error message
}