const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL;

const makeUrl = (path: string): string => `${backendUrl}${path}`;

// ###########################################################################

export async function fetchModels(orgId: string, accessToken: string) {
  const url = makeUrl(`org/${orgId}/model/all`);
  // console.log(url);
  const res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
  }).then((res) => res.json());
  // console.log(res.data);
  return res.Data;
}

export async function fetchModelReadme(
  orgId: string,
  modelName: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/model/${modelName}/readme/version`);
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

export async function writeModelReadme(
  orgId: string,
  modelName: string,
  content: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/model/${modelName}/readme`);
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

export async function fetchModel(
  orgId: string,
  projectId: string,
  modelId: string,
  accessToken: string
) {
  const url = makeUrl(
    `${orgId}/project/${projectId}/model/${modelId}/latest/details`
  );
  // console.log(url);
  const res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
  }).then((res) => res.json());
  return res.data;
}

export async function fetchModelVersions(
  orgId: string,
  modelId: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/model/${modelId}/branch/dev/version`);
  const res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
  }).then((res) => res.json());
  let versions: string[] = [];
  res.Data.forEach((version: any) => {
    // console.log(version.version,i)
    versions.push(version.version);
  });
  return versions;
}

export async function fetchModelMetrics(
  orgId: string,
  modelId: string,
  version: string,
  accessToken: string
) {
  const url = makeUrl(
    `org/${orgId}/model/${modelId}/branch/dev/version/${version}/log`
  );
  const res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
  }).then((res) => res.json());
  // console.log(res.Data[0].data)
  return res.Data[0].data;
}

export async function fetchModelParameters(
  orgId: string,
  projectId: string,
  modelId: string,
  version: string,
  accessToken: string
) {
  const url = makeUrl(
    `${orgId}/project/${projectId}/model/${modelId}/${version}/params`
  );
  const res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
  }).then((res) => res.json());
  return res.data;
}
