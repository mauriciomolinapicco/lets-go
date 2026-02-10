package tools

import (
	log "github.com/sirupsen/logrus"
)

//types of data that the db will return (in django it would be a model)

type DatabaseDetails struct {
	AuthToken string
	Username string
}

type CoinDetails {
	Coins int64
	Username string
}

type DatabaseInterface interface { //define required methods
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetupDatabase() error
}