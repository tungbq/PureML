import { redirect } from "@remix-run/node";

export default function Index() {
  return redirect("/models");
}
