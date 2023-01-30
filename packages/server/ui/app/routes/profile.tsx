import { redirect } from "@remix-run/node"
import { useLoaderData } from "@remix-run/react"
import { getSession } from "~/session"
import { fetchPublicProfile, fetchUserSettings } from "./auth.server"

export async function loader({ request }: any) {
	const session = await getSession(request.headers.get("Cookie"))
	const accesstoken = session.get("accessToken")
	const userSetting = await fetchUserSettings(accesstoken)
	const username = userSetting[0].handle
	const profile = await fetchPublicProfile(username)
	return redirect(`/${profile[0].handle}`, {})
}

export default function ProfileRedirection() {
	const prof = useLoaderData()
	return <></>
}
