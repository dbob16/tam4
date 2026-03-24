import { readConfig } from "$lib/server/settings";

export const load = () => {
  const s = readConfig();
  const currentKey = s.remote_key;
  return { currentKey }
}
