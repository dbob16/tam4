import { writeSettings } from "$lib/server/settings/settingshandler";

export const POST = async ({ request }) => {
  const postBody = await request.json();
  await writeSettings(postBody);
  return new Response(JSON.stringify("Settings saved successfully."))
}
