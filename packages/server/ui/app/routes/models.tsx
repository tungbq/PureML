import type { MetaFunction } from "@remix-run/node"
import { Meta, Outlet, useLoaderData } from "@remix-run/react"
import NavBar from "~/components/Navbar"
import { getSession } from "~/session"
import { fetchUserSettings } from "./api/auth.server"

export const meta: MetaFunction = () => ({
	charset: "utf-8",
	title: "Models | PureML",
	viewport: "width=device-width,initial-scale=1",
})

export async function loader({ request }: any) {
	const session = await getSession(request.headers.get("Cookie"))
	const accesstoken = session.get("accessToken")
	const profile = await fetchUserSettings(accesstoken)
	return profile
}

export default function ModelsLayout() {
	const prof = useLoaderData()
	return (
		<div>
			<head>
				<Meta />
			</head>
			{prof ? (
				<NavBar intent='loggedIn' user={prof[0].name.charAt(0).toUpperCase()} />
			) : (
				<NavBar intent='loggedOut' />
			)}
			<div className='px-12 pt-16 pb-12 h-full w-screen'>
				<Outlet />
			</div>
		</div>
	)
}
