const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL;

const makeUrl = (path: string): string => `${backendUrl}${path}`;

// ###########################################################################

export async function fetchDatasets(orgId: string, accessToken: string) {
  const url = makeUrl(`org/${orgId}/dataset/all`);
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

export async function fetchDatasetBranch(
  orgId: string,
  datasetName: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/dataset/${datasetName}/branch`);
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

export async function fetchDatasetVersions(
  orgId: string,
  datasetName: string,
  // branchName: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/dataset/${datasetName}/branch/dev/version`);
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

export async function fetchDatasetReadme(
  orgId: string,
  datasetName: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/dataset/${datasetName}/readme/version`);
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

export async function writeDatasetReadme(
  orgId: string,
  datasetName: string,
  content: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/dataset/${datasetName}/readme`);
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
    body: JSON.stringify({
      content: content,
      file_type: "html",
    }),
  }).then((res) => res.json());
  return res;
}
