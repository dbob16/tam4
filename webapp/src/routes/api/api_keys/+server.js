import { error } from '@sveltejs/kit';

export const GET = async ({ url }) => {
  const addr = url.searchParams.get('addr');
  const api_pw = url.searchParams.get('api_pw');

  const res = await fetch(`${addr}/api/api_keys?api_pw=${api_pw}`);
  if (!res.ok) {
    throw error(res.status, res.statusText)
  }

  const jsonData = await res.json();

  return new Response(JSON.stringify(jsonData))
}

export const POST = async ({ request }) => {
  const { addr, api_pw, desc } = await request.json();

  const res = await fetch(`${addr}/api/api_keys`, {
    method: 'POST',
    body: JSON.stringify({
      api_pw: api_pw,
      description: desc
    }),
    headers: {'Content-Type': 'application/json'}
  })

  if (!res.ok) {
    throw error(res.status, res.statusText)
  }

  return new Response()
}
