import Airtable from "airtable";

// ############################## contactus api ##############################
const base = new Airtable({
  endpointUrl: "https://api.airtable.com",
  apiKey: process.env.NEXT_PUBLIC_AIRTABLE_API_KEY,
}).base("appAR7Cxhflh7YVe9");

export default base;

// ###########################################################################
const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL;

const makeUrl = (path: string): string => `${backendUrl}${path}`;
// ###########################################################################

// ############################ authentication api ###########################

export async function fetchSignIn(
  email: string,
  password: string,
  username?: string
) {
  const url = makeUrl(`user/login`);
  const res = await fetch(url, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
  }).then((res) => res.json());
  return res;
}

export async function fetchSignUp(
  name: string,
  username: string,
  email: string,
  password: string,
  bio?: string,
  avatar?: string
) {
  const url = makeUrl(`user/signup`);
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
    },
    body: JSON.stringify({
      name,
      handle: username,
      email,
      password,
      bio: "",
      avatar: "",
    }),
  }).then((res) => res.json());
  return res;
}

export async function fetchVerifyEmail(token: string | undefined) {
  const url = makeUrl(`user/verify-email`);
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
    },
    body: JSON.stringify({
      token: token,
    }),
  }).then((res) => res.json());
  return res;
}

export async function fetchForgotPassword(email: string) {
  const url = makeUrl(`user/forgot-password`);
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
    },
    body: JSON.stringify({
      email,
    }),
  }).then((res) => res.json());
  return res;
}

export async function fetchVerifyResetPassword(accessToken: string) {
  const url = makeUrl(`user/verify-reset-password`);
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
    },
    body: JSON.stringify({
      token: accessToken,
    }),
  }).then((res) => res.json());
  return res;
}

export async function fetchResetPassword(
  new_password: string,
  old_password: string,
  accessToken: string | undefined
) {
  const url = makeUrl(`user/reset-password`);
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
    },
    body: JSON.stringify({
      new_password: new_password,
      old_password: old_password,
      token: accessToken,
    }),
  }).then((res) => res.json());
  return res;
}

// ############################# user details api ############################

export async function fetchUserSettings(accessToken: string) {
  const url = makeUrl(`user/profile`);
  const res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
  }).then((res) => (res.ok ? res.json() : null));
  return res?.data;
}

export async function fetchPublicProfile(username: string) {
  const url = makeUrl(`user/profile/${username}`);
  const res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
    },
  }).then((res) => res.json());
  return res.data;
}

// ######################### update user details api #############################

export async function updateProfile(
  name: string,
  bio: string,
  accessToken: string,
  avatar?: string
) {
  const url = makeUrl(`user/profile`);
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
      Authorization: `Bearer ${accessToken}`,
    },
    body: JSON.stringify({
      name,
      bio,
      avatar: "",
    }),
  }).then((res) => res.json());
  return res;
}
