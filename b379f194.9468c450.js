(window.webpackJsonp=window.webpackJsonp||[]).push([[33],{101:function(e,t,n){"use strict";n.r(t),n.d(t,"frontMatter",(function(){return a})),n.d(t,"metadata",(function(){return c})),n.d(t,"toc",(function(){return s})),n.d(t,"default",(function(){return u}));var r=n(3),i=n(7),o=(n(0),n(121)),a={},c={unversionedId:"external/architecture-and-design/overview",id:"external/architecture-and-design/overview",isDocsHomePage:!1,title:"overview",description:"\x3c!--",source:"@site/docs/external/architecture-and-design/overview.md",slug:"/external/architecture-and-design/overview",permalink:"/weaver-dlt-interoperability/docs/external/architecture-and-design/overview",editUrl:"https://github.com/hyperledger-labs/weaver-dlt-interoperability/edit/master/docs/external/architecture-and-design/overview.md",version:"current",sidebar:"Documentation",previous:{title:"legacy-integration",permalink:"/weaver-dlt-interoperability/docs/external/user-stories/legacy-integration"},next:{title:"relay",permalink:"/weaver-dlt-interoperability/docs/external/architecture-and-design/relay"}},s=[{value:"title: Overview",id:"title-overview",children:[]},{value:"Network",id:"network",children:[]},{value:"Relay",id:"relay",children:[]},{value:"Synchronous vs Asynchronous message communication",id:"synchronous-vs-asynchronous-message-communication",children:[]},{value:"Message vs connection oriented communication",id:"message-vs-connection-oriented-communication",children:[]}],l={toc:s};function u(e){var t=e.components,a=Object(i.a)(e,["components"]);return Object(o.b)("wrapper",Object(r.a)({},l,a,{components:t,mdxType:"MDXLayout"}),Object(o.b)("hr",null),Object(o.b)("p",null,"id: overview"),Object(o.b)("h2",{id:"title-overview"},"title: Overview"),Object(o.b)("p",null,"The below diagram shows a high level architecture diagram of the Weaver framework."),Object(o.b)("p",null,Object(o.b)("img",{src:n(201).default})),Object(o.b)("h2",{id:"network"},"Network"),Object(o.b)("p",null,"The networks in the system can be made up of various heterogenious technologies, including Hyperledger Fabric and Corda. Each network in the system needs to contain an interoperability (IOP) module that enables them to communicate with the relays."),Object(o.b)("h2",{id:"relay"},"Relay"),Object(o.b)("p",null,"The relays act as a conduit to facilitate communication of protocols between networks (e.g. data transfer, asset exchange etc). The roles of the relays are described in more detail in ",Object(o.b)("a",{parentName:"p",href:"/weaver-dlt-interoperability/docs/external/architecture-and-design/relay"},"relay"),"."),Object(o.b)("h1",{id:"design-decisions"},"Design Decisions"),Object(o.b)("p",null,"The high level design decisions that were made for the system are outlined here."),Object(o.b)("h2",{id:"synchronous-vs-asynchronous-message-communication"},"Synchronous vs Asynchronous message communication"),Object(o.b)("p",null,"We decided to go with an asynchronous message architecture. The primary reason for this is because requests can take an arbitary amount of time to respond, it is not practical for a synchronous message to wait that long for a reply. For example, obtaining a 12 block confirmation on the Bitcoin network can take about 2 hours."),Object(o.b)("h2",{id:"message-vs-connection-oriented-communication"},"Message vs connection oriented communication"),Object(o.b)("p",null,"We decided to go with a message oriented architecture. The primary reason for this is because it makes the system more fault tolerant. With a message oriented architecture the requester and responder don't need to be alive at the same time. For example, if the requestor crashes while the responder is processing the request, the communication is not interrupted since the responder will just send a message when it has finished processing the request. The design choice also enables the systen to be made more fault tolerant in the future by implementing message queues between components in the system."))}u.isMDXComponent=!0},121:function(e,t,n){"use strict";n.d(t,"a",(function(){return d})),n.d(t,"b",(function(){return p}));var r=n(0),i=n.n(r);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function c(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,r,i=function(e,t){if(null==e)return{};var n,r,i={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var l=i.a.createContext({}),u=function(e){var t=i.a.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):c(c({},t),e)),n},d=function(e){var t=u(e.components);return i.a.createElement(l.Provider,{value:t},e.children)},m={inlineCode:"code",wrapper:function(e){var t=e.children;return i.a.createElement(i.a.Fragment,{},t)}},h=i.a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,o=e.originalType,a=e.parentName,l=s(e,["components","mdxType","originalType","parentName"]),d=u(n),h=r,p=d["".concat(a,".").concat(h)]||d[h]||m[h]||o;return n?i.a.createElement(p,c(c({ref:t},l),{},{components:n})):i.a.createElement(p,c({ref:t},l))}));function p(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var o=n.length,a=new Array(o);a[0]=h;var c={};for(var s in t)hasOwnProperty.call(t,s)&&(c[s]=t[s]);c.originalType=e,c.mdxType="string"==typeof e?e:r,a[1]=c;for(var l=2;l<o;l++)a[l]=n[l];return i.a.createElement.apply(null,a)}return i.a.createElement.apply(null,n)}h.displayName="MDXCreateElement"},201:function(e,t,n){"use strict";n.r(t),t.default=n.p+"assets/images/architecture_overview-7b33025d4ef4bce5161426ac77ec512a.png"}}]);