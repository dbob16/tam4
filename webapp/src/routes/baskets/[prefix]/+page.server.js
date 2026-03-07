export const load = async ({ fetch, params }) => {
  const prefixRes = await fetch(`/api/prefixes/${params.prefix}`);
  const prefixJson = await prefixRes.json();
  return {prefix: prefixJson.prefix, color: prefixJson.color}
}
