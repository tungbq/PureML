import{a as n}from"/build/_shared/chunk-XUGTD4NM.js";import{a}from"/build/_shared/chunk-D2BMCLLG.js";import"/build/_shared/chunk-GERQPFJA.js";import"/build/_shared/chunk-4MN4GN7E.js";import"/build/_shared/chunk-P6UAVHBL.js";import{f as r,k as s}from"/build/_shared/chunk-6ERGQZP6.js";import"/build/_shared/chunk-WNC7G6UM.js";import"/build/_shared/chunk-LQHMM3AA.js";import{b as i}from"/build/_shared/chunk-CXVWUV7G.js";import{c,e as o}from"/build/_shared/chunk-ADMCF34Z.js";var m=c((f,l)=>{l.exports={}});var p=o(m());var e=o(i());function g(){let d=s();return(0,e.jsxs)("div",{id:"models",children:[(0,e.jsx)("div",{className:"flex justify-between font-medium text-slate-800 text-base pt-6",children:"Models"}),d?(0,e.jsx)(e.Fragment,{children:d.models[0].length!==0?(0,e.jsx)("div",{className:"pt-6 grid justify-center grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-6 4xl:grid-cols-10 gap-8 min-w-72",children:d.models.map(t=>(0,e.jsx)(r,{to:`/org/${d.orgId}/models/${t.name}`,children:(0,e.jsx)(a,{intent:"modelCard",name:t.name,description:`Updated by ${t.updated_by.handle||"-"}`,tag2:t.created_by.handle})},t.id))}):(0,e.jsx)(n,{})}):"All public models shown here"]})}export{g as default};