import{a as f}from"/build/_shared/chunk-I2RIK3CX.js";import{a as h}from"/build/_shared/chunk-FNDXDYOH.js";import{a as v}from"/build/_shared/chunk-Z2L4KIN7.js";import{a as u}from"/build/_shared/chunk-XZZCBT56.js";import{b as m,e as c,g as p}from"/build/_shared/chunk-4MN4GN7E.js";import{a as d}from"/build/_shared/chunk-5TQH6BYT.js";import{a as n}from"/build/_shared/chunk-P6UAVHBL.js";import{a as i,j as l}from"/build/_shared/chunk-6ERGQZP6.js";import{b as o}from"/build/_shared/chunk-CXVWUV7G.js";import{e as r}from"/build/_shared/chunk-ADMCF34Z.js";var e=r(o());function x(t){return u(t?" text-blue-700 ":" text-slate-600 "," hover:text-blue-700 flex justify-center items-center px-5 cursor-pointer ")}var y=n("fixed z-20 h-18 px-12 py-4 w-full bg-slate-0 flex justify-between text-sm font-medium border-b-2 border-slate-200",{variants:{intent:{loggedIn:"",loggedOut:""},fullWidth:{true:"w-full"}},defaultVariants:{intent:"loggedOut",fullWidth:!0}});function b({intent:t,fullWidth:g,user:N}){let s=i(),a=l()[1].pathname;return(0,e.jsx)(e.Fragment,{children:(0,e.jsxs)("div",{className:y({intent:t,fullWidth:g}),children:[(0,e.jsxs)("div",{className:"flex",children:[(0,e.jsx)("a",{href:"/models",className:"flex items-center justify-center pr-8",children:(0,e.jsx)("img",{src:"/LogoWText.svg",alt:"Logo",width:"140",height:"96"})}),(0,e.jsx)(v,{intent:"search",placeholder:"Search models, datasets, users...",fullWidth:!1})]}),(0,e.jsxs)("div",{className:"flex justify-center items-center",children:[(0,e.jsxs)("div",{onClick:()=>{s("/models")},className:`${x(a==="/models")}`,children:[(0,e.jsx)(m,{className:"w-4 h-4"}),(0,e.jsx)("span",{className:"pl-2",children:"Models"})]}),(0,e.jsxs)("div",{onClick:()=>{s("/datasets")},className:`${x(a==="/datasets")}`,children:[(0,e.jsx)(c,{className:"w-4 h-4"}),(0,e.jsx)("span",{className:"pl-2",children:"Datasets"})]}),(0,e.jsxs)("a",{href:"https://docs.pureml.com",className:"flex justify-center items-center cursor-pointer px-5 hover:text-blue-700 border-r-2 border-slate-slate-200 font-medium text-slate-600",children:[(0,e.jsx)(p,{className:"w-4 h-4"}),(0,e.jsx)("span",{className:"pl-2",children:"Docs"})]}),t==="loggedOut"?(0,e.jsxs)(e.Fragment,{children:[(0,e.jsx)("div",{className:"w-full flex justify-center items-center px-5",children:(0,e.jsx)("a",{href:"/auth/signin",children:"Sign in"})}),(0,e.jsx)(d,{intent:"primary",icon:"",children:"Sign up"})]}):(0,e.jsx)("div",{className:"w-full flex justify-center items-center px-5",children:(0,e.jsx)(h,{intent:"primary",children:(0,e.jsx)(f,{children:N})})})]})]})})}export{b as a};
