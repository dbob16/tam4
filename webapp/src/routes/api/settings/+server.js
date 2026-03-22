import { readConfig, writeConfig } from "$lib/server/settings";

export const GET = () => {
  const s = readConfig();
  return new Response(JSON.stringify(s))
}

export const POST = async ({ request }) => {
  const s = readConfig();
  const reqData = await request.json();
  const settingsData = { ...s, ...reqData };
  await writeConfig(settingsData);
  return new Response(JSON.stringify({detail: "Success."}))
}
