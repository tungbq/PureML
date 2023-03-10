import { createCookie, createCookieSessionStorage } from "@remix-run/node";

const sessionCookie = createCookie("__session", {
  sameSite: true,
  httpOnly: true,
});

export const { getSession, commitSession, destroySession } =
  createCookieSessionStorage({
    cookie: sessionCookie,
  });
