import{a as l}from"/build/_shared/chunk-CXVWUV7G.js";import{c}from"/build/_shared/chunk-ADMCF34Z.js";var d=c(f=>{"use strict";var u=l();function p(e,t){return e===t&&(e!==0||1/e===1/t)||e!==e&&t!==t}var v=typeof Object.is=="function"?Object.is:p,y=u.useState,S=u.useEffect,E=u.useLayoutEffect,m=u.useDebugValue;function h(e,t){var r=t(),o=y({inst:{value:r,getSnapshot:t}}),n=o[0].inst,s=o[1];return E(function(){n.value=r,n.getSnapshot=t,i(n)&&s({inst:n})},[e,r,t]),S(function(){return i(n)&&s({inst:n}),e(function(){i(n)&&s({inst:n})})},[e]),m(r),r}function i(e){var t=e.getSnapshot;e=e.value;try{var r=t();return!v(e,r)}catch{return!0}}function w(e,t){return t()}var x=typeof window>"u"||typeof window.document>"u"||typeof window.document.createElement>"u"?w:h;f.useSyncExternalStore=u.useSyncExternalStore!==void 0?u.useSyncExternalStore:x});var g=c((q,a)=>{"use strict";a.exports=d()});export{g as a};
