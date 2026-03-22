import { readFileSync, writeFileSync, existsSync } from "fs";
import { env } from "$env/dynamic/private";
import { randomUUID } from "crypto";

export const readConfig = () => {
  const { TAM4_DATA_DIR } = env;
  const configPath = TAM4_DATA_DIR || "." + "/settings.json";
  if (existsSync(configPath)) {
    const settingsData = readFileSync(configPath, "utf8")
    return JSON.parse(settingsData)
  } else {
    const settingsData = {
      tam4_unique_id: randomUUID(),
      remote_server: env.TAM4_REMOTE_SERVER || "",
      remote_port: env.TAM4_REMOTE_PORT || "8000",
      remote_tls: env.TAM4_TLS || false,
      default_pref: env.TAM4_DEFAULT_PREF || "CALL"
    }
    writeFileSync(configPath, JSON.stringify(settingsData, null, 2), { encoding: "utf8" })
    return settingsData
  }
}

export const writeConfig = (settingsData) => {
  const { TAM4_DATA_DIR } = env;
  const configPath = TAM4_DATA_DIR || "." + "/settings.json";
  writeFileSync(configPath, JSON.stringify(settingsData, null, 2), { encoding: "utf8" });
}
