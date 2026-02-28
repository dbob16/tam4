package prefixes

import (
	"errors"

	"git.dilangilluly.us/dbob16/tam4/apiserver/localdb"
	"git.dilangilluly.us/dbob16/tam4/apiserver/apikeys"
)

func GetAllPrefixes(api_key string) ([]Prefix, error) {
	db := localdb.NewSession()
	defer db.Close()

	authStatus := apikeys.CheckAPIKey(api_key)
	if !authStatus {
		return nil, errors.New("Invalid API key.")
	}

	results, err := db.Query("SELECT prefix, color, weight FROM prefixes ORDER BY weight, prefix")
	if err != nil {
		return nil, err
	}

	var constResults []Prefix

	for results.Next() {
		var curPrefix Prefix
		results.Scan(&curPrefix.Prefix, &curPrefix.Color, &curPrefix.Weight)
		constResults = append(constResults, curPrefix)
	}

	return constResults, nil
}

func GetOnePrefix(api_key string, prefix string) (Prefix, error) {
	db := localdb.NewSession()
	defer db.Close()

	authStatus := apikeys.CheckAPIKey(api_key)
	if !authStatus {
		return Prefix{}, errors.New("Invalid API key.")
	}

	var rtnPrefix Prefix

	err := db.QueryRow("SELECT prefix, color, weight FROM prefixes WHERE prefix = $1", prefix).Scan(&rtnPrefix.Prefix, &rtnPrefix.Color, &rtnPrefix.Weight)
	if err != nil {
		return Prefix{}, errors.New("Error finding Prefix.")
	}

	return rtnPrefix, nil
}

func PostOnePrefix(api_key string, prefix Prefix) (Prefix, error) {
	db := localdb.NewSession()
	defer db.Close()

	authStatus := apikeys.CheckAPIKey(api_key)
	if !authStatus {
		return Prefix{}, errors.New("Invalid API key.")
	}

	_, err := db.Exec(`INSERT INTO prefixes (prefix, color, weight) VALUES ($1, $2, $3) ON CONFLICT (prefix)
		DO UPDATE SET color = EXCLUDED.color, weight = EXCLUDED.weight`, prefix.Prefix, prefix.Color, prefix.Weight)
	if err != nil {
		return Prefix{}, errors.New("Error inserting prefix. " + err.Error())
	}

	return prefix, nil
}

func DeleteOnePrefix(api_key string, prefix string) (string, error) {
	db := localdb.NewSession()
	defer db.Close()

	authStatus := apikeys.CheckAPIKey(api_key)
	if !authStatus {
		return "", errors.New("Invalid API key.")
	}

	_, err := db.Exec("DELETE FROM prefixes WHERE prefix = $1", prefix)
	if err != nil {
		return "", err
	}

	return "Key deleted successfully.", nil
}
