import { Meta, Outlet, useLoaderData } from "@remix-run/react"
import NavBar from "~/components/Navbar"
import { getSession } from "~/session"
import { fetchUserSettings } from "./auth.server"
import type { MetaFunction } from "@remix-run/node"

export const meta: MetaFunction = () => ({
	charset: "utf-8",
	title: "Settings | PureML",
	viewport: "width=device-width,initial-scale=1",
})

export async function loader({ request }: any) {
	const session = await getSession(request.headers.get("Cookie"))
	const accesstoken = session.get("accessToken")
	const profile = await fetchUserSettings(accesstoken)
	return profile
}

export default function SettingsLayout() {
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
			<div className='pt-20 pb-12 h-full w-screen'>
				<div className='px-12 flex justify-between font-medium text-slate-800 text-base'>
					Settings
				</div>
				<Outlet />
			</div>
		</div>
	)
}
