const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL;

const makeUrl = (path: string): string => `${backendUrl}${path}`;

// ###########################################################################

export async function fetchOrgDetails(orgId: string, accessToken: string) {
  const url = makeUrl(`org/id/${orgId}`);
  const res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
  }).then((res) => res.json());
  return res.Data;
}
