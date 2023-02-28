import{a as N}from"/build/_shared/chunk-GA5CE6QK.js";import{a as p}from"/build/_shared/chunk-CXVWUV7G.js";import{e as d}from"/build/_shared/chunk-ADMCF34Z.js";function u(){return u=Object.assign?Object.assign.bind():function(e){for(var t=1;t<arguments.length;t++){var r=arguments[t];for(var n in r)Object.prototype.hasOwnProperty.call(r,n)&&(e[n]=r[n])}return e},u.apply(this,arguments)}var i=d(p());function H(e,t=[]){let r=[];function n(a,c){let f=(0,i.createContext)(c),l=r.length;r=[...r,c];function h(x){let{scope:$,children:E,...v}=x,g=$?.[e][l]||f,q=(0,i.useMemo)(()=>v,Object.values(v));return(0,i.createElement)(g.Provider,{value:q},E)}function w(x,$){let E=$?.[e][l]||f,v=(0,i.useContext)(E);if(v)return v;if(c!==void 0)return c;throw new Error(`\`${x}\` must be used within \`${a}\``)}return h.displayName=a+"Provider",[h,w]}let s=()=>{let a=r.map(c=>(0,i.createContext)(c));return function(f){let l=f?.[e]||a;return(0,i.useMemo)(()=>({[`__scope${e}`]:{...f,[e]:l}}),[f,l])}};return s.scopeName=e,[n,F(s,...t)]}function F(...e){let t=e[0];if(e.length===1)return t;let r=()=>{let n=e.map(s=>({useScope:s(),scopeName:s.scopeName}));return function(a){let c=n.reduce((f,{useScope:l,scopeName:h})=>{let x=l(a)[`__scope${h}`];return{...f,...x}},{});return(0,i.useMemo)(()=>({[`__scope${t.scopeName}`]:c}),[c])}};return r.scopeName=t.scopeName,r}var m=d(p());function M(e){let t=(0,m.useRef)(e);return(0,m.useEffect)(()=>{t.current=e}),(0,m.useMemo)(()=>(...r)=>{var n;return(n=t.current)===null||n===void 0?void 0:n.call(t,...r)},[])}var j=d(p()),B=Boolean(globalThis?.document)?j.useLayoutEffect:()=>{};var P=d(p());function I(e,t){typeof e=="function"?e(t):e!=null&&(e.current=t)}function C(...e){return t=>e.forEach(r=>I(r,t))}function T(...e){return(0,P.useCallback)(C(...e),e)}var b=d(p()),R=d(N());var o=d(p());var y=(0,o.forwardRef)((e,t)=>{let{children:r,...n}=e,s=o.Children.toArray(r),a=s.find(_);if(a){let c=a.props.children,f=s.map(l=>l===a?o.Children.count(c)>1?o.Children.only(null):(0,o.isValidElement)(c)?c.props.children:null:l);return(0,o.createElement)(S,u({},n,{ref:t}),(0,o.isValidElement)(c)?(0,o.cloneElement)(c,void 0,f):null)}return(0,o.createElement)(S,u({},n,{ref:t}),r)});y.displayName="Slot";var S=(0,o.forwardRef)((e,t)=>{let{children:r,...n}=e;return(0,o.isValidElement)(r)?(0,o.cloneElement)(r,{...k(n,r.props),ref:C(t,r.ref)}):o.Children.count(r)>1?o.Children.only(null):null});S.displayName="SlotClone";var X=({children:e})=>(0,o.createElement)(o.Fragment,null,e);function _(e){return(0,o.isValidElement)(e)&&e.type===X}function k(e,t){let r={...t};for(let n in t){let s=e[n],a=t[n];/^on[A-Z]/.test(n)?s&&a?r[n]=(...f)=>{a(...f),s(...f)}:s&&(r[n]=s):n==="style"?r[n]={...s,...a}:n==="className"&&(r[n]=[s,a].filter(Boolean).join(" "))}return{...e,...r}}var A=["a","button","div","h2","h3","img","label","li","nav","ol","p","span","svg","ul"],Q=A.reduce((e,t)=>{let r=(0,b.forwardRef)((n,s)=>{let{asChild:a,...c}=n,f=a?y:t;return(0,b.useEffect)(()=>{window[Symbol.for("radix-ui")]=!0},[]),(0,b.createElement)(f,u({},c,{ref:s}))});return r.displayName=`Primitive.${t}`,{...e,[t]:r}},{});function U(e,t){e&&(0,R.flushSync)(()=>e.dispatchEvent(t))}export{u as a,H as b,M as c,B as d,C as e,T as f,y as g,Q as h,U as i};