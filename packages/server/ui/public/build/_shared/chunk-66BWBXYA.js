import{a as T}from"/build/_shared/chunk-XZZCBT56.js";import{a as x}from"/build/_shared/chunk-P6UAVHBL.js";import{f as i,j as $}from"/build/_shared/chunk-6ERGQZP6.js";import{b as u}from"/build/_shared/chunk-CXVWUV7G.js";import{e as h}from"/build/_shared/chunk-ADMCF34Z.js";var e=h(u()),w=x("text-zinc-400 font-medium flex bg-slate-0 sticky z-10",{variants:{intent:{primaryModelTab:"pt-4 border-b-2 border-slate-100 top-44",primaryDatasetTab:"pt-4 border-b-2 border-slate-100 top-28",primarySettingTab:"pt-4 border-b-2 border-slate-100 top-28",modelTab:"pt-8 top-[16rem]",datasetTab:"pt-8 top-[11.7rem]",modelReviewTab:"pt-8 top-[16rem]",datasetReviewTab:"pt-8 top-[11.7rem]"},fullWidth:{true:"w-full"}},defaultVariants:{intent:"primaryModelTab",fullWidth:!0}});function y(s){return T(s?"text-white":"text-slate-600","flex justify-center items-center")}function N({intent:s,fullWidth:f,tab:t}){let k=$()[2].pathname,g=decodeURI(k.slice(1)),r=g.split("/")[1],n=g.split("/")[3],d=g.split("/")[3],c=[{id:"modelCard",name:"Model Card",hyperlink:`/org/${r}/models/${n}`},{id:"versions",name:"Versions",hyperlink:`/org/${r}/models/${n}/versions/metrics`},{id:"review",name:"Review",hyperlink:`/org/${r}/models/${n}/review`}],b=[{id:"datasetCard",name:"Dataset Card",hyperlink:`/org/${r}/datasets/${d}`},{id:"versions",name:"Versions",hyperlink:`/org/${r}/datasets/${d}/versions/datalineage`},{id:"review",name:"Review",hyperlink:`/org/${r}/datasets/${d}/review`}],v=[{id:"profile",name:"Profile",hyperlink:"/settings"},{id:"account",name:"Account",hyperlink:"/settings/account"},{id:"members",name:"Members",hyperlink:"/settings/members"}],o=[{id:"metrics",name:"Metrics",hyperlink:`/org/${r}/models/${n}/versions/metrics`},{id:"graphs",name:"Graphs",hyperlink:`/org/${r}/models/${n}/versions/graphs`}],l=[{id:"datalineage",name:"Data Lineage",hyperlink:`/org/${r}/datasets/${d}/versions/datalineage`}],p=[{id:"metrics",name:"Metrics",hyperlink:`/org/${r}/models/${n}/review/commit`}],m=[{id:"datalineage",name:"Data Lineage",hyperlink:`/org/${r}/datasets/${d}/review/commit`}];return(0,e.jsx)("div",{className:w({intent:s,fullWidth:f}),children:(0,e.jsx)("div",{className:"flex px-10",children:s==="primaryModelTab"||s==="primaryDatasetTab"||s==="primarySettingTab"?(0,e.jsx)(e.Fragment,{children:s==="primaryModelTab"?(0,e.jsx)(e.Fragment,{children:Object.keys(c).map(a=>(0,e.jsx)("div",{className:`${t===c[a].id?"text-blue-700 border-b-2 border-blue-700":"text-slate-600"} p-4`,children:(0,e.jsx)(i,{to:c[a].hyperlink,children:(0,e.jsx)("span",{children:c[a].name})})},a))}):(0,e.jsx)(e.Fragment,{children:s==="primaryDatasetTab"?(0,e.jsx)(e.Fragment,{children:Object.keys(b).map(a=>(0,e.jsx)("div",{className:`${t===b[a].id?"text-blue-700 border-b-2 border-blue-700":"text-slate-600"} p-4`,children:(0,e.jsx)(i,{to:b[a].hyperlink,children:(0,e.jsx)("span",{children:b[a].name})})},a))}):(0,e.jsx)(e.Fragment,{children:Object.keys(v).map(a=>(0,e.jsx)("div",{className:`${t===v[a].id?"text-blue-700 border-b-2 border-blue-700":"text-slate-600"} p-4`,children:(0,e.jsx)(i,{to:v[a].hyperlink,children:(0,e.jsx)("span",{children:v[a].name})})},a))})})}):(0,e.jsx)(e.Fragment,{children:s==="modelTab"||s==="datasetTab"?(0,e.jsx)("div",{className:"flex",children:s==="modelTab"?(0,e.jsx)(e.Fragment,{children:Object.keys(o).map(a=>(0,e.jsx)("div",{className:`${t===o[a].id?"bg-blue-700 rounded text-white":""} px-4 py-2`,children:(0,e.jsx)(i,{to:o[a].hyperlink,className:`${y(t===o[a].id)}`,children:(0,e.jsx)("span",{children:o[a].name})})},a))}):(0,e.jsx)(e.Fragment,{children:Object.keys(l).map(a=>(0,e.jsx)("div",{className:`${t===l[a].id?"bg-blue-700 rounded text-white":""} px-4 py-2`,children:(0,e.jsx)(i,{to:l[a].hyperlink,className:`${y(t===l[a].id)}`,children:(0,e.jsx)("span",{children:l[a].name})})},a))})}):(0,e.jsx)("div",{className:"flex",children:s==="modelReviewTab"?(0,e.jsx)(e.Fragment,{children:Object.keys(p).map(a=>(0,e.jsx)("div",{className:`${t===p[a].id?"bg-blue-700 rounded text-white":""} px-4 py-2`,children:(0,e.jsx)(i,{to:p[a].hyperlink,className:`${y(t===p[a].id)}`,children:(0,e.jsx)("span",{children:p[a].name})})},a))}):(0,e.jsx)(e.Fragment,{children:Object.keys(m).map(a=>(0,e.jsx)("div",{className:`${t===m[a].id?"bg-blue-700 rounded text-white":""} px-4 py-2`,children:(0,e.jsx)(i,{to:m[a].hyperlink,className:`${y(t===m[a].id)}`,children:(0,e.jsx)("span",{children:m[a].name})})},a))})})})})})}export{N as a};