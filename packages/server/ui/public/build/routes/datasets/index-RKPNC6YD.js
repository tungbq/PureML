import{a as n}from"/build/_shared/chunk-D2BMCLLG.js";import{a as o}from"/build/_shared/chunk-TP56N7TT.js";import"/build/_shared/chunk-GERQPFJA.js";import"/build/_shared/chunk-4MN4GN7E.js";import"/build/_shared/chunk-P6UAVHBL.js";import{f as d,k as r}from"/build/_shared/chunk-6ERGQZP6.js";import"/build/_shared/chunk-WNC7G6UM.js";import"/build/_shared/chunk-LQHMM3AA.js";import{b as i}from"/build/_shared/chunk-CXVWUV7G.js";import{c,e as s}from"/build/_shared/chunk-ADMCF34Z.js";var l=c((u,g)=>{g.exports={}});var m=s(l());var t=s(i());function p(){let a=r();return(0,t.jsxs)("div",{id:"datasets",children:[(0,t.jsx)("div",{className:"flex justify-between font-medium text-slate-800 text-base pt-6",children:"Datasets"}),a?(0,t.jsx)(t.Fragment,{children:a.datasets[0].length!==0?(0,t.jsx)("div",{className:"pt-6 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 gap-8 min-w-72",children:a.datasets.map(e=>(0,t.jsx)(d,{to:`/org/${a.orgId}/datasets/${e.name}`,children:(0,t.jsx)(n,{intent:"datasetCard",name:e.name,description:`Updated by ${e.updated_by.handle}`,tag2:e.created_by.handle},e.updated_at)},e.id))},"0"):(0,t.jsx)(o,{})}):"All public datasets shown here"]})}export{p as default};