import{a as $}from"/build/_shared/chunk-CXVWUV7G.js";import{e as k}from"/build/_shared/chunk-ADMCF34Z.js";var D=k($(),1),j=k($(),1),b=k($(),1);var V={data:""},Y=e=>typeof window=="object"?((e?e.querySelector("#_goober"):window._goober)||Object.assign((e||document.head).appendChild(document.createElement("style")),{innerHTML:" ",id:"_goober"})).firstChild:e||V;var Z=/(?:([\u0080-\uFFFF\w-%@]+) *:? *([^{;]+?);|([^;}{]*?) *{)|(}\s*)/g,q=/\/\*[^]*?\*\/|  +/g,H=/\n+/g,v=(e,t)=>{let a="",o="",s="";for(let r in e){let n=e[r];r[0]=="@"?r[1]=="i"?a=r+" "+n+";":o+=r[1]=="f"?v(n,r):r+"{"+v(n,r[1]=="k"?"":t)+"}":typeof n=="object"?o+=v(n,t?t.replace(/([^,])+/g,i=>r.replace(/(^:.*)|([^,])+/g,l=>/&/.test(l)?l.replace(/&/g,i):i?i+" "+l:l)):r):n!=null&&(r=/^--/.test(r)?r:r.replace(/[A-Z]/g,"-$&").toLowerCase(),s+=v.p?v.p(r,n):r+":"+n+";")}return a+(t&&s?t+"{"+s+"}":s)+o},h={},L=e=>{if(typeof e=="object"){let t="";for(let a in e)t+=a+L(e[a]);return t}return e},G=(e,t,a,o,s)=>{let r=L(e),n=h[r]||(h[r]=(l=>{let d=0,c=11;for(;d<l.length;)c=101*c+l.charCodeAt(d++)>>>0;return"go"+c})(r));if(!h[n]){let l=r!==e?e:(d=>{let c,y,f=[{}];for(;c=Z.exec(d.replace(q,""));)c[4]?f.shift():c[3]?(y=c[3].replace(H," ").trim(),f.unshift(f[0][y]=f[0][y]||{})):f[0][c[1]]=c[2].replace(H," ").trim();return f[0]})(e);h[n]=v(s?{["@keyframes "+n]:l}:l,a?"":"."+n)}let i=a&&h.g?h.g:null;return a&&(h.g=h[n]),((l,d,c,y)=>{y?d.data=d.data.replace(y,l):d.data.indexOf(l)===-1&&(d.data=c?l+d.data:d.data+l)})(h[n],t,o,i),n},J=(e,t,a)=>e.reduce((o,s,r)=>{let n=t[r];if(n&&n.call){let i=n(a),l=i&&i.props&&i.props.className||/^go/.test(i)&&i;n=l?"."+l:i&&typeof i=="object"?i.props?"":v(i,""):i===!1?"":i}return o+s+(n??"")},"");function O(e){let t=this||{},a=e.call?e(t.p):e;return G(a.unshift?a.raw?J(a,[].slice.call(arguments,1),t.p):a.reduce((o,s)=>Object.assign(o,s&&s.call?s(t.p):s),{}):a,Y(t.target),t.g,t.o,t.k)}var _,M,S,Ae=O.bind({g:1}),p=O.bind({k:1});function B(e,t,a,o){v.p=t,_=e,M=a,S=o}function u(e,t){let a=this||{};return function(){let o=arguments;function s(r,n){let i=Object.assign({},r),l=i.className||s.className;a.p=Object.assign({theme:M&&M()},i),a.o=/ *go\d+/.test(l),i.className=O.apply(a,o)+(l?" "+l:""),t&&(i.ref=n);let d=e;return e[0]&&(d=i.as||e,delete i.as),S&&d[0]&&S(i),_(d,i)}return t?t(s):s}}var w=k($(),1);var x=k($(),1),K=e=>typeof e=="function",A=(e,t)=>K(e)?e(t):e,Q=(()=>{let e=0;return()=>(++e).toString()})(),R=(()=>{let e;return()=>{if(e===void 0&&typeof window<"u"){let t=matchMedia("(prefers-reduced-motion: reduce)");e=!t||t.matches}return e}})(),W=20,T=new Map,X=1e3,U=e=>{if(T.has(e))return;let t=setTimeout(()=>{T.delete(e),E({type:4,toastId:e})},X);T.set(e,t)},ee=e=>{let t=T.get(e);t&&clearTimeout(t)},F=(e,t)=>{switch(t.type){case 0:return{...e,toasts:[t.toast,...e.toasts].slice(0,W)};case 1:return t.toast.id&&ee(t.toast.id),{...e,toasts:e.toasts.map(r=>r.id===t.toast.id?{...r,...t.toast}:r)};case 2:let{toast:a}=t;return e.toasts.find(r=>r.id===a.id)?F(e,{type:1,toast:a}):F(e,{type:0,toast:a});case 3:let{toastId:o}=t;return o?U(o):e.toasts.forEach(r=>{U(r.id)}),{...e,toasts:e.toasts.map(r=>r.id===o||o===void 0?{...r,visible:!1}:r)};case 4:return t.toastId===void 0?{...e,toasts:[]}:{...e,toasts:e.toasts.filter(r=>r.id!==t.toastId)};case 5:return{...e,pausedAt:t.time};case 6:let s=t.time-(e.pausedAt||0);return{...e,pausedAt:void 0,toasts:e.toasts.map(r=>({...r,pauseDuration:r.pauseDuration+s}))}}},N=[],z={toasts:[],pausedAt:void 0},E=e=>{z=F(z,e),N.forEach(t=>{t(z)})},te={blank:4e3,error:4e3,success:2e3,loading:1/0,custom:4e3},ae=(e={})=>{let[t,a]=(0,D.useState)(z);(0,D.useEffect)(()=>(N.push(a),()=>{let s=N.indexOf(a);s>-1&&N.splice(s,1)}),[t]);let o=t.toasts.map(s=>{var r,n;return{...e,...e[s.type],...s,duration:s.duration||((r=e[s.type])==null?void 0:r.duration)||e?.duration||te[s.type],style:{...e.style,...(n=e[s.type])==null?void 0:n.style,...s.style}}});return{...t,toasts:o}},re=(e,t="blank",a)=>({createdAt:Date.now(),visible:!0,type:t,ariaProps:{role:"status","aria-live":"polite"},message:e,pauseDuration:0,...a,id:a?.id||Q()}),I=e=>(t,a)=>{let o=re(t,e,a);return E({type:2,toast:o}),o.id},m=(e,t)=>I("blank")(e,t);m.error=I("error");m.success=I("success");m.loading=I("loading");m.custom=I("custom");m.dismiss=e=>{E({type:3,toastId:e})};m.remove=e=>E({type:4,toastId:e});m.promise=(e,t,a)=>{let o=m.loading(t.loading,{...a,...a?.loading});return e.then(s=>(m.success(A(t.success,s),{id:o,...a,...a?.success}),s)).catch(s=>{m.error(A(t.error,s),{id:o,...a,...a?.error})}),e};var se=(e,t)=>{E({type:1,toast:{id:e,height:t}})},oe=()=>{E({type:5,time:Date.now()})},ie=e=>{let{toasts:t,pausedAt:a}=ae(e);(0,j.useEffect)(()=>{if(a)return;let r=Date.now(),n=t.map(i=>{if(i.duration===1/0)return;let l=(i.duration||0)+i.pauseDuration-(r-i.createdAt);if(l<0){i.visible&&m.dismiss(i.id);return}return setTimeout(()=>m.dismiss(i.id),l)});return()=>{n.forEach(i=>i&&clearTimeout(i))}},[t,a]);let o=(0,j.useCallback)(()=>{a&&E({type:6,time:Date.now()})},[a]),s=(0,j.useCallback)((r,n)=>{let{reverseOrder:i=!1,gutter:l=8,defaultPosition:d}=n||{},c=t.filter(g=>(g.position||d)===(r.position||d)&&g.height),y=c.findIndex(g=>g.id===r.id),f=c.filter((g,P)=>P<y&&g.visible).length;return c.filter(g=>g.visible).slice(...i?[f+1]:[0,f]).reduce((g,P)=>g+(P.height||0)+l,0)},[t]);return{toasts:t,handlers:{updateHeight:se,startPause:oe,endPause:o,calculateOffset:s}}},ne=p`
from {
  transform: scale(0) rotate(45deg);
	opacity: 0;
}
to {
 transform: scale(1) rotate(45deg);
  opacity: 1;
}`,le=p`
from {
  transform: scale(0);
  opacity: 0;
}
to {
  transform: scale(1);
  opacity: 1;
}`,de=p`
from {
  transform: scale(0) rotate(90deg);
	opacity: 0;
}
to {
  transform: scale(1) rotate(90deg);
	opacity: 1;
}`,ce=u("div")`
  width: 20px;
  opacity: 0;
  height: 20px;
  border-radius: 10px;
  background: ${e=>e.primary||"#ff4b4b"};
  position: relative;
  transform: rotate(45deg);

  animation: ${ne} 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275)
    forwards;
  animation-delay: 100ms;

  &:after,
  &:before {
    content: '';
    animation: ${le} 0.15s ease-out forwards;
    animation-delay: 150ms;
    position: absolute;
    border-radius: 3px;
    opacity: 0;
    background: ${e=>e.secondary||"#fff"};
    bottom: 9px;
    left: 4px;
    height: 2px;
    width: 12px;
  }

  &:before {
    animation: ${de} 0.15s ease-out forwards;
    animation-delay: 180ms;
    transform: rotate(90deg);
  }
`,pe=p`
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
`,ue=u("div")`
  width: 12px;
  height: 12px;
  box-sizing: border-box;
  border: 2px solid;
  border-radius: 100%;
  border-color: ${e=>e.secondary||"#e0e0e0"};
  border-right-color: ${e=>e.primary||"#616161"};
  animation: ${pe} 1s linear infinite;
`,me=p`
from {
  transform: scale(0) rotate(45deg);
	opacity: 0;
}
to {
  transform: scale(1) rotate(45deg);
	opacity: 1;
}`,fe=p`
0% {
	height: 0;
	width: 0;
	opacity: 0;
}
40% {
  height: 0;
	width: 6px;
	opacity: 1;
}
100% {
  opacity: 1;
  height: 10px;
}`,ye=u("div")`
  width: 20px;
  opacity: 0;
  height: 20px;
  border-radius: 10px;
  background: ${e=>e.primary||"#61d345"};
  position: relative;
  transform: rotate(45deg);

  animation: ${me} 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275)
    forwards;
  animation-delay: 100ms;
  &:after {
    content: '';
    box-sizing: border-box;
    animation: ${fe} 0.2s ease-out forwards;
    opacity: 0;
    animation-delay: 200ms;
    position: absolute;
    border-right: 2px solid;
    border-bottom: 2px solid;
    border-color: ${e=>e.secondary||"#fff"};
    bottom: 6px;
    left: 6px;
    height: 10px;
    width: 6px;
  }
`,ge=u("div")`
  position: absolute;
`,he=u("div")`
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
  min-width: 20px;
  min-height: 20px;
`,be=p`
from {
  transform: scale(0.6);
  opacity: 0.4;
}
to {
  transform: scale(1);
  opacity: 1;
}`,ve=u("div")`
  position: relative;
  transform: scale(0.6);
  opacity: 0.4;
  min-width: 20px;
  animation: ${be} 0.3s 0.12s cubic-bezier(0.175, 0.885, 0.32, 1.275)
    forwards;
`,xe=({toast:e})=>{let{icon:t,type:a,iconTheme:o}=e;return t!==void 0?typeof t=="string"?w.createElement(ve,null,t):t:a==="blank"?null:w.createElement(he,null,w.createElement(ue,{...o}),a!=="loading"&&w.createElement(ge,null,a==="error"?w.createElement(ce,{...o}):w.createElement(ye,{...o})))},we=e=>`
0% {transform: translate3d(0,${e*-200}%,0) scale(.6); opacity:.5;}
100% {transform: translate3d(0,0,0) scale(1); opacity:1;}
`,Ee=e=>`
0% {transform: translate3d(0,0,-1px) scale(1); opacity:1;}
100% {transform: translate3d(0,${e*-150}%,-1px) scale(.6); opacity:0;}
`,ke="0%{opacity:0;} 100%{opacity:1;}",$e="0%{opacity:1;} 100%{opacity:0;}",Oe=u("div")`
  display: flex;
  align-items: center;
  background: #fff;
  color: #363636;
  line-height: 1.3;
  will-change: transform;
  box-shadow: 0 3px 10px rgba(0, 0, 0, 0.1), 0 3px 3px rgba(0, 0, 0, 0.05);
  max-width: 350px;
  pointer-events: auto;
  padding: 8px 10px;
  border-radius: 8px;
`,je=u("div")`
  display: flex;
  justify-content: center;
  margin: 4px 10px;
  color: inherit;
  flex: 1 1 auto;
  white-space: pre-line;
`,Ie=(e,t)=>{let a=e.includes("top")?1:-1,[o,s]=R()?[ke,$e]:[we(a),Ee(a)];return{animation:t?`${p(o)} 0.35s cubic-bezier(.21,1.02,.73,1) forwards`:`${p(s)} 0.4s forwards cubic-bezier(.06,.71,.55,1)`}},Ce=b.memo(({toast:e,position:t,style:a,children:o})=>{let s=e.height?Ie(e.position||t||"top-center",e.visible):{opacity:0},r=b.createElement(xe,{toast:e}),n=b.createElement(je,{...e.ariaProps},A(e.message,e));return b.createElement(Oe,{className:e.className,style:{...s,...a,...e.style}},typeof o=="function"?o({icon:r,message:n}):b.createElement(b.Fragment,null,r,n))});B(x.createElement);var Te=({id:e,className:t,style:a,onHeightUpdate:o,children:s})=>{let r=x.useCallback(n=>{if(n){let i=()=>{let l=n.getBoundingClientRect().height;o(e,l)};i(),new MutationObserver(i).observe(n,{subtree:!0,childList:!0,characterData:!0})}},[e,o]);return x.createElement("div",{ref:r,className:t,style:a},s)},Ne=(e,t)=>{let a=e.includes("top"),o=a?{top:0}:{bottom:0},s=e.includes("center")?{justifyContent:"center"}:e.includes("right")?{justifyContent:"flex-end"}:{};return{left:0,right:0,display:"flex",position:"absolute",transition:R()?void 0:"all 230ms cubic-bezier(.21,1.02,.73,1)",transform:`translateY(${t*(a?1:-1)}px)`,...o,...s}},ze=O`
  z-index: 9999;
  > * {
    pointer-events: auto;
  }
`,C=16,_e=({reverseOrder:e,position:t="top-center",toastOptions:a,gutter:o,children:s,containerStyle:r,containerClassName:n})=>{let{toasts:i,handlers:l}=ie(a);return x.createElement("div",{style:{position:"fixed",zIndex:9999,top:C,left:C,right:C,bottom:C,pointerEvents:"none",...r},className:n,onMouseEnter:l.startPause,onMouseLeave:l.endPause},i.map(d=>{let c=d.position||t,y=l.calculateOffset(d,{reverseOrder:e,gutter:o,defaultPosition:t}),f=Ne(c,y);return x.createElement(Te,{id:d.id,key:d.id,onHeightUpdate:l.updateHeight,className:d.visible?ze:"",style:f},d.type==="custom"?A(d.message,d):s?s(d):x.createElement(Ce,{toast:d,position:c}))}))};export{_e as a};
