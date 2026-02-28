import { env } from "$env/dynamic/private";

export const load = async ({ fetch }) => {
  const adminPw = env.ADMIN_PW || ""
  let conRes = "";

  const res = await fetch("/api");
  if (!res.ok) {
    conRes = `Server error: [${res.status}] ${res.statusText}`
  } else {
    conRes = await res.json();
  }

  return {admin_pw: adminPw, conn_res: conRes}
}
