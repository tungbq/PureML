const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL;

const makeUrl = (path: string): string => `${backendUrl}${path}`;

// ###########################################################################

// ########################### org details api ###########################

export async function fetchAllOrgs(accessToken: string) {
  const url = makeUrl(`org`);
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
  return res.data;
}

// ########################### create org ###########################

export async function fetchCreateOrg(
  handle: string,
  description: string,
  accessToken: string,
  avatar?: any
) {
  const url = makeUrl(`org/create`);
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
    body: JSON.stringify({
      handle,
      description,
      avatar: "",
    }),
  }).then((res) => res.json());
  return res;
}

// ########################### join org ###########################

export async function fetchJoinOrg(code: string, accessToken: string) {
  const url = makeUrl(`org/join`);
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
    body: JSON.stringify({ join_code: code }),
  });
  return res;
}

// ########################### update org ###########################

export async function updateOrg(
  description: string,
  name: string,
  orgId: string,
  accessToken: string,
  avatar?: string
) {
  const url = makeUrl(`org/${orgId}/update`);
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
    body: JSON.stringify({
      avatar: "",
      description,
      name,
    }),
  }).then((res) => res.json());
  return res;
}

export async function updateOrgAddMember(
  orgId: string,
  email: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/add`);
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
    body: JSON.stringify({
      orgId,
      email,
    }),
  }).then((res) => res.json());
  return res;
}

export async function updateOrgRemoveMember(
  orgId: string,
  email: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/remove`);
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
    body: JSON.stringify({
      orgId,
      email,
    }),
  }).then((res) => res.json());
  return res;
}

export async function updateOrgChangeRole(
  orgId: string,
  email: string,
  role: string,
  accessToken: string
) {
  const url = makeUrl(`org/${orgId}/role`);
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
    body: JSON.stringify({
      orgId,
      email,
      role: role,
    }),
  }).then((res) => res.json());
  return res;
}
