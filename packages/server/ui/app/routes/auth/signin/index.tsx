import type { ErrorBoundaryComponent } from "@remix-run/node"
import { json, redirect } from "@remix-run/node"
import { useActionData } from "@remix-run/react"
import toast from "react-hot-toast"
import Button from "~/components/ui/Button"
import Input from "~/components/ui/Input"
import Link from "~/components/ui/Link"
import { fetchSignIn, fetchUserOrg } from "~/routes/auth.server"
import { commitSession, getSession } from "~/session"

export const ErrorBoundary: ErrorBoundaryComponent = ({ error }) => {
	return SignIn(error.message)
}

export const action = async ({ request }: any) => {
	const form = await request.formData()
	// const username = form.get('username');
	const email = form.get("email")
	const password = form.get("password")
	const data = await fetchSignIn(email, password)
	const session = await getSession(request.headers.get("Cookie"))
	if (data.Message === "User logged in") {
		const accessToken = data.Data[0].accessToken
		session.set("accessToken", accessToken)
		const org = await fetchUserOrg(accessToken)
		session.set("orgId", org[0].org.uuid)
		session.set("orgName", org[0].org.name)
		toast.success("Here is your toast.")
		return redirect("/models", {
			headers: { "Set-Cookie": await commitSession(session) },
		})
	} else if (data.Message === "User not found") {
		return json({ message: "User not found" })
	} else if (data.Message === "Invalid username") {
		return json({ message: "Invalid username" })
	} else if (data.Message === "Invalid credentials") {
		return json({ message: "Invalid credentials" })
	} else {
		return json({ message: "Something went Wrong!" })
	}
}

export default function SignIn(err: string) {
	const data = useActionData()
	let errorComp: JSX.Element = <div></div>
	if (err.length > 0) {
		errorComp = (
			<p className='errorBox'>
				There was an error with your data: <i className='errorMsg'>{err}</i>
			</p>
		)
	}
	return (
		<div className='md:w-3/5 bg-slate-0 md:flex md:flex-col justify-center items-center md:pt-0 text-white'>
			{errorComp}
			<div className='md:max-w-[450px] w-96 text-center'>
				<h2 className='font-semibold text-2xl text-slate-800 mb-12'>
					Sign In to PureML
				</h2>
				<form method='post' className='text-slate-400 flex flex-col text-left'>
					<p className='text-red-500'>{data ? data.message : null}</p>
					{/* <div className='pb-6'>
            <label htmlFor='username' className='text-base pb-1'>
              Username
              <Input
                intent='primary'
                // onChange={(e) => setEmail(e.target.value)}
                type='text'
                name='username'
                placeholder='Enter email ID...'
                aria-label='emailid'
                data-testid='email-input1'
                required
              />
            </label>
          </div> */}
					<div className='pb-6'>
						<label htmlFor='email' className='text-base pb-1'>
							Email ID
							<Input
								intent='primary'
								// onChange={(e) => setEmail(e.target.value)}
								type='email'
								name='email'
								placeholder='Enter email ID...'
								aria-label='emailid'
								data-testid='email-input1'
								required
							/>
						</label>
					</div>
					<div className='pb-6'>
						<label htmlFor='password' className='text-base pb-1'>
							Password
							<Input
								intent='primary'
								required
								// onChange={(e) => setPassword(e.target.value)}
								type='password'
								name='password'
								placeholder='Enter password...'
								aria-label='password'
								data-testid='password-input1'
							/>
						</label>
					</div>
					<Button intent='primary' icon='' type='submit'>
						Sign in
					</Button>
				</form>
				<div className='flex items-center text-slate-600 space-x-3 justify-center mt-6'>
					<Link intent='secondary' hyperlink='/auth/forgot_password'>
						Forgot Password?
					</Link>
					<p>|</p>
					<div className='flex items-center space-x-1'>
						<span className='text-sm'>Dont have an account?</span>
						<Link intent='secondary' hyperlink='/auth/signup'>
							Sign Up
						</Link>
					</div>
				</div>
			</div>
		</div>
	)
}
