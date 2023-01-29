import { Link, useLoaderData } from "@remix-run/react"
import Card from "~/components/Card"
import { getSession } from "~/session"
import { fetchModels } from "../models.server"
import EmptyModel from "./EmptyModel"

export type model = {
	id: string
	name: string
	updated_at: string
	uuid: string
	created_by: string
	updated_by: string
}

export async function loader({ request }: any) {
	const session = await getSession(request.headers.get("Cookie"))
	const orgId = session.get("orgId")
	const models: model[] = await fetchModels(
		session.get("orgId"),
		session.get("accessToken")
	)
	return { models, orgId }
}

export default function Index() {
	const modelData = useLoaderData()
	return (
		<div id='models'>
			<div className='flex justify-between font-medium text-slate-800 text-base pt-6'>
				Models
			</div>
			{modelData ? (
				<>
					{modelData.models[0].length !== 0 ? (
						<div className='pt-6 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 gap-8 min-w-72'>
							{modelData.models.map((model: any) => (
								<Link
									to={`/org/${modelData.orgId}/models/${model.name}`}
									key={model.id}>
									<Card
										intent='modelCard'
										name={model.name}
										description={`Updated by ${model.updated_by.handle || "-"}`}
										// tag1={model.tag1}
										tag2={model.created_by.handle}
									/>
								</Link>
							))}
						</div>
					) : (
						<EmptyModel />
					)}
				</>
			) : (
				"All public models shown here"
			)}
		</div>
	)
}
