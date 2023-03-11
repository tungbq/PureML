import { Link, useMatches } from "@remix-run/react";

function Breadcrumbs() {
  const matches = useMatches();
  const pathname = matches[2].pathname;
  const url = decodeURI(pathname.slice(1)).split("/");
  const urlitems = url.filter(function (val, idx) {
    if ((idx + 1) % 2 == 0) return val;
  });
  return (
    <ul className="flex pt-6">
      {urlitems.map((item, index) => (
        <li key={item}>
          <Link
            to={`/${url.slice(0, index + 2).join("/")}`}
            className={
              item === urlitems[urlitems.length - 1]
                ? "font-medium text-slate-600 w-full"
                : "text-slate-400 hover:text-slate-600 w-full"
            }
          >
            {item}
          </Link>
          {index !== url.length - 1 && (
            <span className="text-slate-400 mx-1">/</span>
          )}
        </li>
      ))}
    </ul>
  );
}

export default Breadcrumbs;
