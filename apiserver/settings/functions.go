package settings

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

func SettingsDirsMasher() SettingsDirs {
	var settingsdirs SettingsDirs = SettingsDirs{
		BaseDir: ".",
		ConfigDir: "./settings.json",
		CertsDir: "./certs",
	}
	if val, exists := os.LookupEnv("TAM4_CONFIG"); exists {
		settingsdirs.BaseDir = val
		settingsdirs.ConfigDir = strings.Join([]string{val, "/settings.json"}, "")
		settingsdirs.CertsDir = strings.Join([]string{val, "/certs"}, "")
	}
	return settingsdirs
}

func ReadSettings() SettingsStruct {
	var s SettingsStruct
	configDir := SettingsDirsMasher().ConfigDir
	_, err := os.Stat(configDir)
	if err != nil {
		if os.IsNotExist(err) {
			if val, exists := os.LookupEnv("TAM4_DB_HOST"); exists {
				s.DBCreds.DBHost = val
			} else {
				s.DBCreds.DBHost = "localhost"
			}
			if val, exists := os.LookupEnv("TAM4_DB_PORT"); exists {
				s.DBCreds.DBPort = val
			} else {
				s.DBCreds.DBPort = "5432"
			}
			if val, exists := os.LookupEnv("TAM4_DB_USER"); exists {
				s.DBCreds.DBUser = val
			} else {
				s.DBCreds.DBUser = "tam4"
			}
			if val, exists := os.LookupEnv("TAM4_DB_PASSWORD"); exists {
				s.DBCreds.DBPassword = val
			} else {
				s.DBCreds.DBPassword = "dbob16"
			}
			if val, exists := os.LookupEnv("TAM4_DB_DATABASE"); exists {
				s.DBCreds.DBDatabase = val
			} else {
				s.DBCreds.DBDatabase = "tam4"
			}
			if val, exists := os.LookupEnv("TAM4_API_PW"); exists {
				s.APIPW = val
			} else {
				s.APIPW = "dbob16"
			}

			fileToWrite, err := json.MarshalIndent(s, "", "  ")
			if err != nil {
				log.Fatalln("Error converting settings to JSON.")
			}

			err = os.WriteFile(configDir, fileToWrite, 0644)
			if err != nil {
				log.Fatalln("Error writing default config.")
			}
		} else {
			log.Fatalln("Unknown error for checking file existence.")
		}
	} else {
		readFile, err := os.ReadFile(configDir)
		if err != nil {
			log.Fatalln("Unknown error reading file.")
		}
		err = json.Unmarshal(readFile, &s)
	}
	return s
}
