const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL;

const makeUrl = (path: string): string => `${backendUrl}${path}`;

// ###########################################################################

// ########################### model details api ###########################

export async function fetchPublicModels(orgId: string, accessToken: string) {
  const url = makeUrl(`org/${orgId}/model/all`);
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

export async function fetchModels(orgId: string, accessToken: string) {
  const url = makeUrl(`org/${orgId}/model/all`);
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

export async function fetchModelByName(
  orgId: string,
  modelName: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/model/${modelName}`);
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

// ############################# model readme #############################

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
  return res.data;
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

// ######################### model branch details ###########################

export async function fetchModelBranch(
  orgId: string,
  modelName: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/model/${modelName}/branch`);
  const res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
  }).then((res) => res.json());
  if (res.data !== null) {
    let b: Object[] = [];
    res.data.forEach((branch: any) => {
      b.push({ value: branch.name, label: branch.name });
    });
    return b;
  }
  return res;
}

// ########################## model version details ###########################

export async function fetchOneModelVersion(
  orgId: string,
  modelId: string,
  branchName: string,
  version: string,
  accessToken: string
) {
  const url = makeUrl(
    `org/${orgId}/model/${modelId}/branch/${branchName}/version/${version}/log`
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

export async function fetchModelVersions(
  orgId: string,
  modelId: string,
  branchName: string,
  accessToken: string
) {
  const url = makeUrl(
    `org/${orgId}/model/${modelId}/branch/${branchName}/version?withLogs=true`
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
  return res.data;
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

// ######################### model review ###########################

export async function fetchModelReview(
  orgId: string,
  modelName: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/model/${modelName}/review`);
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

export async function submitModelReview(
  orgId: string,
  modelName: string,
  branch: string,
  version: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/model/${modelName}/review/create`);
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
    body: JSON.stringify({
      description: "string",
      from_branch: branch,
      from_branch_version: version,
      is_accepted: false,
      is_complete: false,
      title: "string",
      to_branch: "main",
    }),
  }).then((res) => res.json());
  return res;
}

export async function updateModelReview(
  orgId: string,
  modelName: string,
  reviewId: string,
  accepted: string,
  accessToken: string
) {
  const url = makeUrl(
    `org/${orgId}/model/${modelName}/review/${reviewId}/update`
  );
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
    body: JSON.stringify({
      description: "-",
      is_accepted: accepted === "true",
      is_complete: true,
      title: "-",
    }),
  }).then((res) => res.json());
  return res;
}
