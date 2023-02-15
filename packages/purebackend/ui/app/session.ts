import { createCookie, createCookieSessionStorage } from "@remix-run/node";

const sessionCookie = createCookie("__session", {
  secrets: ["r3m1xr0ck5"],
  sameSite: true,
  httpOnly: true,
});

export const { getSession, commitSession, destroySession } =
  createCookieSessionStorage({
    cookie: sessionCookie,
  });
