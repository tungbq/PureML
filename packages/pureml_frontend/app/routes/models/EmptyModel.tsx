export default function EmptyModel() {
  return (
    <div className="px-12 pt-2 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 min-w-72">
      <div className="rounded-lg border-2 border-slate-200 px-6 py-4">
        <div className="font-medium text-sm pb-6">There are no models yet</div>
        <div className="rounded-lg h-2 bg-slate-200 w-1/3" />
        <div className="pt-2"></div>
        <div className="rounded-lg h-2 bg-slate-200 w-2/3" />
      </div>
    </div>
  );
}
