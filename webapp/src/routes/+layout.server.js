import { readSettings } from "$lib/server/settings/settingshandler";

export const load = () => {
  return {settings: readSettings()}
}
