import { Meta, useLoaderData } from "@remix-run/react"
import { fetchPublicProfile, fetchUserSettings } from "~/routes/auth.server"
import Avatar from "~/components/ui/Avatar"
import Button from "~/components/ui/Button"
import { Edit2 } from "lucide-react"
import ProfileCard from "~/components/ProfileCard"
import NavBar from "~/components/Navbar"
import type { MetaFunction } from "@remix-run/node"
import Error from "~/Error404"
import { getSession } from "~/session"

export const meta: MetaFunction = () => ({
	charset: "utf-8",
	title: "Profile | PureML",
	viewport: "width=device-width,initial-scale=1",
})

export async function loader({ params, request }: any) {
	const session = await getSession(request.headers.get("Cookie"))
	const accesstoken = session.get("accessToken")
	const publicProfile = await fetchPublicProfile(params.username)
	const userProfile = await fetchUserSettings(accesstoken)
	if (!publicProfile) return null
	return { publicProfile, userProfile }
}

export default function UserProfile() {
	const userProfileData = useLoaderData()
	if (userProfileData) {
		return (
			<div>
				<head>
					<Meta />
				</head>
				{userProfileData.userProfile ? (
					<NavBar
						intent='loggedIn'
						user={userProfileData.userProfile[0].name.charAt(0).toUpperCase()}
					/>
				) : (
					<NavBar intent='loggedOut' />
				)}
				<div className='flex px-12 pt-24 pb-12 text-slate-800 font-medium'>
					<div className='h-full w-28 md:w-36 lg:w-56 2xl:w-96'>
						<div className='h-28 w-28 md:h-36 md:w-36 lg:w-56 lg:h-56 2xl:h-96 2xl:w-96 flex items-center justify-center text-md text-blue-600 bg-blue-200 rounded-lg'>
							<Avatar intent='profile'>
								{userProfileData.publicProfile[0]?.name
									.charAt(0)
									.toUpperCase() || "User"}
							</Avatar>
						</div>
						<div className='pt-6 font-semibold text-base text-slate-900'>
							{userProfileData.publicProfile[0]?.name || "Name"}
						</div>
						<div className='pb-6 text-base font-normal'>
							{userProfileData.publicProfile[0]?.email || "Email"}
						</div>
						<Button aria-label='follow' intent='primary' icon=''>
							Follow
						</Button>
						<div className='flex justify-between text-base pt-8'>
							<span>Bio</span>
							<Edit2 />
						</div>
						<div className='font-medium text-base text-slate-600'>
							{userProfileData.publicProfile[0]?.bio || "Add your bio"}
						</div>
						{/* <div className="text-base pt-8">Organizations</div>
            {userProfileData.publicProfile[0].orgs.length !== 0 ? (
              <div>
                {userProfileData.publicProfile[0].orgs.map((org: any) => (
                  <Button intent="org" icon="" key={org.name}>
                    {org.name}
                  </Button>
                ))}
              </div>
            ) : (
              "-"
            )} */}
					</div>
					<div className='pl-12 w-full'>
						<div className='pb-6'>Overview</div>
						<div className='flex w-full'>
							<div className='pr-4 w-full'>
								<ProfileCard
									title='Projects'
									count={
										userProfileData.publicProfile[0]?.number_of_projects || "0"
									}
								/>
							</div>
							<div className='pr-4 w-full'>
								<ProfileCard
									title='Models'
									count={
										userProfileData.publicProfile[0]?.number_of_models || "0"
									}
								/>
							</div>
							<div className='w-full'>
								<ProfileCard
									title='Datasets'
									count={
										userProfileData.publicProfile[0]?.number_of_datasets || "0"
									}
								/>
							</div>
						</div>
					</div>
				</div>
			</div>
		)
	}
	return (
		<div>
			<head>
				<Meta />
			</head>
			{userProfileData.userProfile ? (
				<NavBar
					intent='loggedIn'
					user={userProfileData.userProfile[0].name.charAt(0).toUpperCase()}
				/>
			) : (
				<NavBar intent='loggedOut' />
			)}
			<Error />
		</div>
	)
}
