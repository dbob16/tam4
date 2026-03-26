export const load = async ({ fetch }) => {
  const prefixRes = await fetch("/api/prefixes");
  if (!prefixRes.ok) {
    return {prefixes: []}
  } else {
    const prefixData = await prefixRes.json();
    return {prefixes: [...prefixData]}
  }
}
