import{a as f,b as m,h as u}from"/build/_shared/chunk-4BJBRKM7.js";import{a as b}from"/build/_shared/chunk-P6UAVHBL.js";import{a as y,b as $}from"/build/_shared/chunk-CXVWUV7G.js";import{e as l}from"/build/_shared/chunk-ADMCF34Z.js";var e=l(y());var p="Avatar",[x,F]=m(p),[S,L]=x(p),_=(0,e.forwardRef)((t,r)=>{let{__scopeAvatar:o,...s}=t,[d,i]=(0,e.useState)("idle");return(0,e.createElement)(S,{scope:o,imageLoadingStatus:d,onImageLoadingStatusChange:i},(0,e.createElement)(u.span,f({},s,{ref:r})))});var k="AvatarFallback",R=(0,e.forwardRef)((t,r)=>{let{__scopeAvatar:o,delayMs:s,...d}=t,i=L(k,o),[g,A]=(0,e.useState)(s===void 0);return(0,e.useEffect)(()=>{if(s!==void 0){let N=window.setTimeout(()=>A(!0),s);return()=>window.clearTimeout(N)}},[s]),g&&i.imageLoadingStatus!=="loaded"?(0,e.createElement)(u.span,f({},d,{ref:r})):null});var c=_;var n=R;var a=l($()),v=b("flex items-center px-3 py-2 font-medium focus:outline-none",{variants:{intent:{primary:"w-6 h-6 flex items-center justify-center text-md text-blue-600 bg-blue-200 rounded-full capitalize",profile:"text-black h-9 rounded justify-center capitalize",org:"bg-brand-100 text-black h-9 rounded-full justify-center capitalize"},fullWidth:{true:"w-full"}},defaultVariants:{intent:"primary",fullWidth:!0}});function h({intent:t,fullWidth:r,children:o}){return(0,a.jsx)("div",{children:t==="primary"?(0,a.jsx)("div",{className:"h-full",children:(0,a.jsx)(c,{children:(0,a.jsx)(n,{className:v({intent:t,fullWidth:r}),children:o})})}):(0,a.jsx)(a.Fragment,{children:t==="profile"?(0,a.jsx)("div",{className:"px-1",children:(0,a.jsx)(c,{children:(0,a.jsx)(n,{className:v({intent:t,fullWidth:r}),children:o})})}):(0,a.jsx)("div",{className:"h-full",children:(0,a.jsx)(c,{children:(0,a.jsx)(n,{className:v({intent:t,fullWidth:r}),children:o})})})})})}export{h as a};
