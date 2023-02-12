import type { MetaFunction } from "@remix-run/node"
import { redirect } from "@remix-run/node"
import { Meta, useActionData, useNavigate } from "@remix-run/react"
import Button from "~/components/ui/Button"
import { getSession } from "~/session"
import base, { fetchUserSettings } from "./api/auth.server"

export const action = async ({ request }: any) => {
	const form = await request.formData()
	const comment = form.get("comment")
	const session = await getSession(request.headers.get("Cookie"))
	const accessToken = session.get("accessToken")
	const userDetails = await fetchUserSettings(accessToken)
	const email = userDetails[0].email
	// const submit = base("tbl7qXTTBN3Ln6KyE").create(
	// 	[
	// 		{
	// 			fields: {
	// 				Email: email,
	// 				Comment: comment,
	// 			},
	// 		},
	// 	],
	// 	(err: string) => {
	// 		if (err) {
	// 			console.error(err)
	// 		}
	// 	}
	// )
	return redirect("/models", {})
}

export const meta: MetaFunction = () => ({
	charset: "utf-8",
	title: "Contact Us | PureML",
	viewport: "width=device-width,initial-scale=1",
})

export default function Contact() {
	const navigate = useNavigate()
	const data = useActionData()
	return (
		<div className='flex h-screen justify-center items-center bg-zinc-800 opacity-60'>
			<head>
				<Meta />
			</head>
			<div className='bg-slate-0 p-4 rounded-lg w-[20rem]'>
				<div className='text-slate-800 font-medium pb-4'>Contact Us</div>
				<form method='post'>
					<label htmlFor='comment'>
						<textarea
							typeof='text'
							name='comment'
							required
							className='whitespace-pre-line w-full bg-transparent text-sm border border-slate-600 rounded-md h-full hover:border-blue-750 focus:outline-none focus:border-blue-750 max-h-[200px] p-4'
							placeholder='Add your query or feedback here...'
						/>
					</label>
					<div className='pt-12 grid justify-items-end w-full'>
						<div className='flex justify-between w-1/2'>
							<Button
								icon=''
								fullWidth={false}
								intent='secondary'
								onClick={() => {
									navigate("/models")
								}}>
								No
							</Button>
							<Button icon='' intent='danger' fullWidth={false} type='submit'>
								Submit
							</Button>
						</div>
					</div>
				</form>
			</div>
		</div>
	)
}
