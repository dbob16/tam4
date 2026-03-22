import json
import os
import pathlib

data_path = pathlib.Path(os.getenv("TAM4_DATA_PATH", "."))

class Config:
    def __init__(self):
        self.config_path = data_path / "config.json"
    def read_config(self):
        """Reads the config file and returns the config data.

        Also will create a default config file if the config file doesn't exist."""
        if self.config_path.exists():
            with open(self.config_path, "r") as f:
                conf_data = json.loads(f.read())
        else:
            conf_data = {
                "db": {
                    "host": os.getenv("TAM4_DB_HOST", "localhost"),
                    "port": os.getenv("TAM4_DB_PORT", "5432"),
                    "user": os.getenv("TAM4_DB_USER", "tam4"),
                    "password": os.getenv("TAM4_DB_PASSWORD", "dbob16"),
                    "database": os.getenv("TAM4_DB_DATABASE", "tam4")
                },
                "api_pw": os.getenv("TAM4_API_PW", "dbob16")
            }
            with open(self.config_path, "w") as f:
                json.dump(conf_data, f, indent="  ")
        return conf_data
