export const load = async ({ params, fetch }) => {
  let res = await fetch(`/api/prefixes/${params.prefix}`);
  const prefix = await res.json();
  res = await fetch(`/api/reports/bybasket/${params.prefix}`);
  const reportData = await res.json();
  return { prefix, reportData }
}
