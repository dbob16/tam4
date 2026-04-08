export const load = async ({ params, fetch }) => {
  const { prefix } = params;
  const prefixRes = await fetch(`/api/prefixes/${prefix}`);
  const prefixData = await prefixRes.json();
  return { prefix: prefixData };
}
