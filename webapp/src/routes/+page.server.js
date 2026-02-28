import { env } from "$env/dynamic/private";

export const load = () => {
  const adminPw = env.ADMIN_PW || ""
  return {admin_pw: adminPw}
}
