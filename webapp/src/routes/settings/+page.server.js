import { readConfig } from "$lib/server/settings";

export const load = () => {
  const c = readConfig();
  return { settings: c }
}
