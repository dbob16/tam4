package apikeys

import (
	"crypto/rand"
	"database/sql"
	"errors"

	"git.dilangilluly.us/dbob16/tam4/apiserver/localdb"
	"git.dilangilluly.us/dbob16/tam4/apiserver/settings"
)

func ListApikeys(api_pw string) ([]ApiKey, error) {
	db := localdb.NewSession()
	defer db.Close()
	config := settings.ReadSettings()

	if api_pw != config.APIPW {
		return nil, errors.New("Invalid Password.")
	}

	results, err := db.Query("SELECT api_key, description FROM api_keys")
	if err != nil {
		return nil, errors.New("Unable to fetch results.")
	}

	var api_keys []ApiKey

	for results.Next() {
		var apikey ApiKey
		if err := results.Scan(&apikey.ApiKey, &apikey.Description); err != nil {
			return nil, errors.New("Error parsing contents of a returned row.")
		}
		api_keys = append(api_keys, apikey)
	}

	return api_keys, nil
}

func CreateAPIkey(api_pw string, description string) (string, error) {
	db := localdb.NewSession()
	defer db.Close()
	config := settings.ReadSettings()

	if api_pw != config.APIPW {
		return "", errors.New("Invalid Password.")
	}

	var NewApiKey string

	for {
		NewApiKey = rand.Text()[0:16]
		err := db.QueryRow(`SELECT api_key FROM api_keys WHERE api_key = $1`, NewApiKey).Scan()
		if err != nil {
			if err == sql.ErrNoRows {
				break
			} else {
				return "", err
			}
		}
	}

	db.Exec("INSERT INTO api_keys (api_key, description) VALUES ($1, $2)", NewApiKey, description)

	return NewApiKey, nil
}

func DeleteAPIKey(api_pw string, key_to_del string) (string, error) {
	db := localdb.NewSession()
	defer db.Close()
	config := settings.ReadSettings()

	if api_pw != config.APIPW {
		return "", errors.New("Invalid Password.")
	}

	_, err := db.Exec("DELETE FROM api_keys WHERE api_key = $1", key_to_del)
	if err != nil {
		return "", err
	}

	return "Key deleted successfully.", nil
}

func CheckAPIKey(api_key string) (bool) {
	db := localdb.NewSession()
	defer db.Close()

	var out_key string
	err := db.QueryRow("SELECT api_key FROM api_keys WHERE api_key = $1", api_key).Scan(&out_key)
	if err != nil {
		return false
	}

	return true
}
