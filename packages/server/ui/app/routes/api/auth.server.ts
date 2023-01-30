// ############################## contactus api ##############################
import Airtable from "airtable";

const base = new Airtable({
  endpointUrl: "https://api.airtable.com",
  apiKey: process.env.NEXT_PUBLIC_AIRTABLE_API_KEY,
}).base("appAR7Cxhflh7YVe9");

export default base;

// ###########################################################################
const backendUrl = process.env.NEXT_PUBLIC_BACKEND_URL;

const makeUrl = (path: string): string => `${backendUrl}${path}`;
// ###########################################################################

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
  bio: string,
  avatar?: string
) {
  const url = makeUrl(`user/signup`);
  const res = await fetch(url, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
    },
    body: new URLSearchParams({ name, username, email, password, bio }),
  }).then((res) => res.json());
  // console.log("res=", res);
  return res;
}

export async function fetchUserOrg(accessToken: string) {
  const url = makeUrl(`org/`);
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

export async function fetchUserSettings(accessToken: string) {
  const url = makeUrl(`user/profile`);
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

export async function fetchPublicProfile(username: string) {
  const url = makeUrl(`user/profile/${username}`);
  const res = await fetch(url, {
    method: "GET",
    headers: {
      "Content-Type": "application/json",
      Accept: "application / json",
    },
  }).then((res) => res.json());
  return res.Data;
}
