import { ErrorBoundaryComponent, json, redirect } from "@remix-run/node"
import { useActionData } from "@remix-run/react"
import toast from "react-hot-toast"
import Button from "~/components/ui/Button"
import Input from "~/components/ui/Input"
import Link from "~/components/ui/Link"
import { fetchSignUp } from "~/routes/api/auth.server"
import { commitSession, getSession } from "~/session"

export const ErrorBoundary: ErrorBoundaryComponent = ({ error }) => {
	return SignUp(error.message)
}

export const action = async ({ request }: any) => {
	const form = await request.formData()
	console.log("form=", form)
	const name = form.get("name")
	const username = form.get("username")
	const email = form.get("email")
	const password = form.get("password")
	const bio = form.get("bio")
	console.log(name, username, email, password, bio)
	const data = await fetchSignUp(name, username, email, password, bio)
	console.log("data=", data)
	if (data.Message === "User created") {
		toast.success("Here is your toast.")
		json({ message: "Account created successfully" })
		return redirect("/auth/signin", {})
	} else if (data.Message === "User already exists") {
		return json({ message: "User not found" })
	} else {
		return json({ message: "Something went Wrong!" })
	}
}

export default function SignUp(err: string) {
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
					Sign up to PureML
				</h2>
				<form method='post' className='text-slate-400 flex flex-col text-left'>
					<div className='pb-6'>
						<label htmlFor='name' className='text-base pb-1'>
							Name
							<Input
								intent='primary'
								// onChange={(e) => setName(e.target.value)}
								type='text'
								name='name'
								placeholder='Enter name...'
								aria-label='name'
								data-testid='name-input'
								required
							/>
						</label>
					</div>
					<div className='pb-6'>
						<label htmlFor='username' className='text-base pb-1'>
							Username
							<Input
								intent='primary'
								// onChange={(e) => setName(e.target.value)}
								type='text'
								name='username'
								placeholder='Enter your username...'
								aria-label='username'
								data-testid='username-input'
								required
							/>
						</label>
					</div>
					<div className='pb-6'>
						<label htmlFor='email' className='text-base pb-1'>
							Email ID
							<Input
								intent='primary'
								required
								// onChange={(e) => setEmail(e.target.value)}
								type='email'
								name='email'
								placeholder='Enter email ID...'
								aria-label='emalid'
								data-testid='email-input2'
							/>
						</label>
					</div>
					<div className='pb-6'>
						<label htmlFor='password' className='text-base pb-1'>
							Password
							<Input
								intent='primary'
								// onChange={(e) => setPassword(e.target.value)}
								type='password'
								name='password'
								required
								placeholder='Enter password...'
								aria-label='password'
								data-testid='password-input2'
							/>
						</label>
					</div>
					<div className='pb-6'>
						<label htmlFor='bio' className='text-base pb-1'>
							Short Bio
							<Input
								intent='primary'
								// onChange={(e) => setName(e.target.value)}
								type='text'
								name='bio'
								placeholder='Enter your short bio...'
								aria-label='bio'
								data-testid='bio-input'
								required
							/>
						</label>
					</div>
					<Button intent='primary' icon=''>
						Sign Up
					</Button>
				</form>
				<div className='flex items-center text-slate-600 space-x-2 justify-center mt-6'>
					<Link intent='secondary' hyperlink='/forgot_password'>
						Forgot Password?
					</Link>
					<p>|</p>
					<div className='flex items-center space-x-1'>
						<span className='text-sm'>Already have an account?</span>
						<Link intent='secondary' hyperlink='/auth/signin'>
							Sign In
						</Link>
					</div>
				</div>
			</div>
		</div>
	)
}
