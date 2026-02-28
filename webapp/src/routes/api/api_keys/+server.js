import { error } from '@sveltejs/kit';
import { URLSearchParams } from 'url';

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

  return new Response(JSON.stringify(await res.json()))
}

export const DELETE = async ({ url }) => {
  const addr = url.searchParams.get('addr');
  const api_pw = url.searchParams.get('api_pw')
  const key_to_del = url.searchParams.get('key_to_del');

  const searchStr = new URLSearchParams({ api_pw: api_pw, key_to_delete: key_to_del }).toString()

  const res = await fetch(`${addr}/api/api_keys?${searchStr}`, {
    method: 'DELETE'
  });

  if (!res.ok) {
    throw error(res.status, res.statusText)
  }

  console.log(`[${res.status}] ${JSON.stringify(await res.json())}`)

  return new Response()
}
