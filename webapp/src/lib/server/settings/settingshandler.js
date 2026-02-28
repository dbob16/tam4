import { env } from "$env/dynamic/private";
import { readFileSync, writeFileSync, existsSync } from "fs";

const settingsPath = env.CONFIG_DIR || "." + "/settings.json"

export const readSettings = () => {
  if (existsSync(settingsPath)) {
    const settingsData = readFileSync(settingsPath);
    const jsonData = JSON.parse(settingsData);
    return jsonData
  } else {
    const stockConfig = {
      display_theme: "light",
      remote_tls: false,
      remote_server: "",
      remote_key: ""
    };
    const writeData = JSON.stringify(stockConfig, null, 2);
    writeFileSync(settingsPath, writeData);
    return stockConfig
  }
}

export const writeSettings = (fileData) => {
  const writeData = JSON.stringify(fileData, null, 2);
  writeFileSync(settingsPath, writeData);
  return fileData
}

export const getRemote = () => {
  const config = readSettings();
  const prefix = config.remote_tls ? "https://" : "http://";
  const port = config.remote_tls ? ":8443" : ":8000";
  const connStr = prefix + config.remote_server + port;
  return {conn_str: connStr, api_key: config.remote_key}
}
