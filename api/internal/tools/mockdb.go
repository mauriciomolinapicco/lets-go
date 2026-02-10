package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"user1": {AuthToken: "token1", Username: "user1"},
	"user2": {AuthToken: "token2", Username: "user2"},
}

var mockCoinDetails = map[string]CoinDetails{
	"user1": {Coins: 100, Username: "user1"},
	"user2": {Coins: 200, Username: "user2"},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(50 * time.Millisecond) //simular delay
	
	var clientData = LoginDetails{}
	clientData, ok := modkLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	//no setup needed for mock db
	return nil
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(50 * time.Millisecond) //simular delay

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}
