package settings

type SettingsStruct struct {
	DBCreds struct {
		DBHost string `json:"hostname"`
		DBPort string `json:"port"`
		DBUser string `json:"user"`
		DBPassword string `json:"password"`
		DBDatabase string `json:"database"`
	} `json:"db_creds"`
	APIPW string `json:"api_pw"`
}

type SettingsDirs struct {
	BaseDir string `json:"base_dir"`
	ConfigDir string `json:"config_dir"`
	CertsDir string `json:"certs_dir"`
}
