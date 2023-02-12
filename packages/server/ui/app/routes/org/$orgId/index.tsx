import type { MetaFunction } from "@remix-run/node"
import { Link, Meta, useLoaderData } from "@remix-run/react"
import { ArrowLeft } from "lucide-react"
import Card from "~/components/Card"
import AvatarIcon from "~/components/ui/Avatar"
import Error from "~/Error404"
import { fetchDatasets } from "~/routes/api/datasets.server"
import { fetchModels } from "~/routes/api/models.server"
import { fetchOrgDetails } from "~/routes/api/org.server"
import EmptyDataset from "~/routes/datasets/EmptyDataset"
import EmptyModel from "~/routes/models/EmptyModel"
import { getSession } from "~/session"

export const meta: MetaFunction = () => ({
	charset: "utf-8",
	title: "Organization Details | PureML",
	viewport: "width=device-width,initial-scale=1",
})

export type org = {
	uuid: string
	id: string
	name: string
}

export type model = {
	id: string
	name: string
	project_name: string
	project_id: string
	updated_at: string
	created_by: string
	updated_by_name: string
}

export type dataset = {
	id: string
	name: string
	project_name: string
	updated_at: string
	created_by: string
	project_id: string
	updated_by_name: string
}

export async function loader({ params, request }: any) {
	const session = await getSession(request.headers.get("Cookie"))
	// console.log('session=', session.get('accessToken'));
	const orgDetails: org[] = await fetchOrgDetails(
		params.orgId,
		session.get("accessToken")
	)
	if (!orgDetails) {
		return null
	}
	const modelDetails: model[] = await fetchModels(
		orgDetails[0].uuid,
		session.get("accessToken")
	)
	const datasetDetails: dataset[] = await fetchDatasets(
		orgDetails[0].uuid,
		session.get("accessToken")
	)
	return { orgDetails, modelDetails, datasetDetails }
}

export default function OrgIndex() {
	const orgData = useLoaderData<model[]>()
	if (orgData === null) {
		return <Error />
	}
	return (
		<div className='flex h-full'>
			<head>
				<Meta />
			</head>
			<div className='w-4/5 pt-6 px-12'>
				<Link
					to='/models'
					className='flex text-sm font-medium text-slate-600 pb-6'>
					<ArrowLeft /> Go back
				</Link>
				<div className='p-6 bg-brand-50 flex items-center rounded-lg'>
					<AvatarIcon intent='org'>
						{orgData.orgDetails[0].name.charAt(0)}
					</AvatarIcon>
					<div className='pl-8'>
						<div className='text-sm font-medium text-slate-800'>
							{orgData.orgDetails[0].name}
						</div>
						<div className='text-xs text-slate-600'>
							{orgData.orgDetails[0].description}
						</div>
					</div>
				</div>
				<div>
					<div className='flex justify-between font-medium text-slate-800 text-base pt-6'>
						Models
					</div>
					{orgData ? (
						<>
							{orgData.modelDetails.length !== 0 ? (
								<div className='pt-6 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 gap-8 min-w-72'>
									{orgData.modelDetails.map((model: any) => (
										<Link
											to={`/org/${orgData.orgDetails[0].uuid}/models/${model.name}`}
											key={model.id}>
											<Card
												intent='modelCard'
												name={model.name}
												description={`Updated by ${model.updated_by.handle}`}
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
				<div>
					<div className='flex justify-between font-medium text-slate-800 text-base pt-6'>
						Datasets
					</div>
					{orgData ? (
						<>
							{orgData.datasetDetails.length !== 0 ? (
								<div className='pt-6 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 gap-8 min-w-72'>
									{orgData.datasetDetails.map((dataset: any) => (
										<Link
											to={`/org/${orgData.orgDetails[0].uuid}/datasets/${dataset.name}`}
											key={dataset.id}>
											<Card
												intent='datasetCard'
												key={dataset.updated_at}
												name={dataset.name}
												description={`Updated by ${dataset.updated_by.handle}`}
												// tag1={dataset.tag1}
												tag2={dataset.created_by.handle}
											/>
										</Link>
									))}
								</div>
							) : (
								<EmptyDataset />
							)}
						</>
					) : (
						"All public datasets shown here"
					)}
				</div>
			</div>
			<div className='w-1/5 border-l-2 border-slate-200'>
				<div className='text-slate-900 font-medium text-base px-8 py-6'>
					Activity
				</div>
			</div>
		</div>
	)
}
