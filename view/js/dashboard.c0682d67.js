webpackJsonp([10],{"11aa":function(t,e,a){"use strict";var i=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"qcontent"},[a("div",{staticClass:"qcontent-header hidden-xs"},[t._v("服务器概览 "),a("button",{staticClass:"btn btn-xs btn-default full-btn",attrs:{type:"button",title:"点击大屏展示"},on:{click:function(e){e.preventDefault(),t.fullscreen(e)}}},[a("i",{staticClass:"fa fa-arrows-alt"})])]),t._v(" "),a("div",{staticClass:"container-fluid",style:"height:"+this.pageHeight+"px;min-height:500px;"},[a("div",{staticClass:"col-xs-12 col-md-6"},[a("div",{staticClass:"view-dashboard view-left"},[a("div",{staticClass:"panel"},[a("div",{staticClass:"panel-body"},[a("div",{staticClass:"col-md-12 panel-title"},[t._v("\r\n                            通道信息\r\n                        ")]),t._v(" "),a("div",{staticClass:"col-xs-6 auth-view",staticStyle:{"text-align":"center"}},[a("el-progress",{attrs:{color:"#337ab7",type:"circle",height:this.authViewHeight,width:this.authViewHeight,percentage:0==t.authData.DeviceTotal?0:parseFloat((t.authData.DeviceOnline/t.authData.DeviceTotal*100).toFixed(2))}}),t._v(" "),a("div",{staticClass:"auth-title"},[t._v("在线数:"),a("span",{staticClass:"auth-num"},[t._v(t._s(t.authData.DeviceOnline))]),a("br"),t._v("总设备:"),a("span",{staticClass:"auth-num"},[t._v(t._s(t.authData.DeviceTotal))])])],1),t._v(" "),a("div",{staticClass:"col-xs-6 auth-view",staticStyle:{"text-align":"center"}},[a("el-progress",{attrs:{color:"#337ab7",type:"circle",height:this.authViewHeight,width:this.authViewHeight,percentage:0==t.authData.ChannelTotal?0:parseFloat((t.authData.ChannelOnline/t.authData.ChannelTotal*100).toFixed(2))}}),t._v(" "),a("div",{staticClass:"auth-title"},[t._v("在线数:"),a("span",{staticClass:"auth-num"},[t._v(t._s(t.authData.ChannelOnline))]),a("br"),t._v("总通道:"),a("span",{staticClass:"auth-num"},[t._v(t._s(t.authData.ChannelTotal))])])],1)])])]),t._v(" "),a("div",{staticClass:"view-dashboard  view-left view-split"},[a("div",{staticClass:"panel"},[a("div",{staticClass:"panel-body"},[a("ve-bar",{ref:"storeChart",attrs:{height:"100%",colors:t.chartColors,data:t.storeData,settings:t.storeSettings,"legend-position":"bottom",title:{text:"存储使用",left:"center"}}})],1)])])]),t._v(" "),a("div",{staticClass:"col-xs-12 col-md-6"},[a("div",{staticClass:"view-dashboard view-right view-split-m"},[a("div",{staticClass:"panel"},[a("div",{staticClass:"panel-body"},[a("ve-line",{ref:"cpuChart",attrs:{height:"100%",colors:t.chartColors,data:t.cpuData,settings:t.memSettings,"legend-visible":!1,title:{text:"CPU使用",left:"center"}}})],1)])]),t._v(" "),a("div",{staticClass:"view-dashboard view-right view-split"},[a("div",{staticClass:"panel"},[a("div",{staticClass:"panel-body"},[a("ve-line",{ref:"memChart",attrs:{height:"100%",colors:t.chartColors,data:t.memData,settings:t.memSettings,"legend-visible":!1,title:{text:"内存使用",left:"center"}}})],1)])])])])])},s=[],n={render:i,staticRenderFns:s};e.a=n},ettG:function(t,e,a){var i=a("pAFs");"string"==typeof i&&(i=[[t.i,i,""]]),i.locals&&(t.exports=i.locals);a("FIqI")("3825c506",i,!0,{})},pAFs:function(t,e,a){e=t.exports=a("UTlt")(!1),e.push([t.i,".qcontent .fullscreen[data-v-2b92d112]{position:fixed;left:0;top:0;right:0;bottom:0;background-color:#ecf0f5}.qcontent .full-btn[data-v-2b92d112]{margin-left:2px;margin-top:-3px}.qcontent .col-md-6[data-v-2b92d112]{padding-left:0;padding-right:0;height:100%}.qcontent .panel-body[data-v-2b92d112],.qcontent .panel[data-v-2b92d112]{height:100%}.qcontent .panel-title[data-v-2b92d112]{text-align:center;font-weight:700;font-size:18px}.qcontent .container-fluid[data-v-2b92d112]{padding:10px 10px 20px}.qcontent[data-v-2b92d112]{margin:-16px -15px}.view-dashboard[data-v-2b92d112]{height:50%}.view-left[data-v-2b92d112]{margin-right:5px}.view-right[data-v-2b92d112]{margin-left:5px}.view-split[data-v-2b92d112]{margin-top:10px}.auth-live[data-v-2b92d112],.auth-vlive[data-v-2b92d112]{margin-top:6%}.auth-num[data-v-2b92d112]{margin-top:20%;font-size:14px;font-weight:700;color:#337ab7}.auth-view[data-v-2b92d112]{min-height:200px;padding-top:78px}.auth-title[data-v-2b92d112]{font-size:12px;color:#1d3b55}.auth-promt-div[data-v-2b92d112]{margin-top:10%}@media (max-width:992px){.col-md-6[data-v-2b92d112]{width:100%}.container-fluid[data-v-2b92d112]{height:100%!important}.view-dashboard[data-v-2b92d112]{height:320px}.view-left[data-v-2b92d112]{margin-right:0}.view-right[data-v-2b92d112]{margin-left:0}.view-split-m[data-v-2b92d112],.view-split[data-v-2b92d112]{margin-top:10px}.auth-live[data-v-2b92d112],.auth-vlive[data-v-2b92d112]{margin-top:0}.auth-num[data-v-2b92d112]{margin-top:20%;font-size:14px;font-weight:700;color:#337ab7}.auth-view[data-v-2b92d112]{height:200px!important;padding-top:78px!important}.qcontent .container-fluid[data-v-2b92d112]{padding-bottom:10px}}",""])},uBsg:function(t,e,a){"use strict";(function(t){function i(t){return t&&t.__esModule?t:{default:t}}Object.defineProperty(e,"__esModule",{value:!0});var s=Object.assign||function(t){for(var e=1;e<arguments.length;e++){var a=arguments[e];for(var i in a)Object.prototype.hasOwnProperty.call(a,i)&&(t[i]=a[i])}return t},n=a("Jz72"),o=i(n),r=a("6ROu"),d=(i(r),a("HVJf")),l=a("lRRz"),h=i(l);o.default.use(h.default),e.default={computed:s({},(0,d.mapState)(["serverInfo","userInfo"])),components:{},data:function(){return{chartColors:["#337ab7","#7FFFD4"],pageWidth:0,pageHeight:0,protocol:location.protocol,authViewHeight:80,timer:0,vods:{},duration:30,memData:{columns:["time","use"],rows:[]},memSettings:{area:!0,xAxisType:"time",yAxisType:["percent"],min:[0],max:[1],labelMap:{use:"使用"}},cpuData:{columns:["time","use"],rows:[]},vodData:{columns:["type","count"],rows:[]},vodSettings:{label:{show:!0,formatter:"{b}: {c}"},roseType:"area"},storeData:{columns:["Name","Used","FreeSpace"],rows:[]},storeSettings:{dimension:["Name"],metrics:["Used","FreeSpace"],stack:{store:["Used","FreeSpace"]},legendName:{Used:"已使用(G)",FreeSpace:"剩余(G)"},labelMap:{Used:"已使用(G)",FreeSpace:"剩余(G)"}},authData:{ChannelOnline:0,ChannelTotal:0,ChannelCount:0,DeviceOnline:0,DeviceTotal:0},bandwidthData:{columns:["time","use"],rows:[]},bandwidthSettings:{area:!0,xAxisType:"time",labelMap:{use:"使用(Mbps)"}}}},mounted:function(){var e=this;this.top(),this.timer1=setInterval(function(){e.top()},2e3),this.store(),this.timer2=setInterval(function(){e.store()},5e3),this.timer3=setTimeout(function(){e.resizeCharts()},1e3),t(window).on("resize",this.resize),t(document).on("expanded.pushMenu",this.resizeCharts),t(document).on("collapsed.pushMenu",this.resizeCharts),this.fixAuthView()},created:function(){this.initHeight()},beforeDestroy:function(){this.timer1&&(clearInterval(this.timer1),this.timer1=0),this.timer2&&(clearInterval(this.timer2),this.timer2=0),this.timer3&&(clearTimeout(this.timer3),this.timer3=0),t(window).off("resize",this.resize),t(document).off("expanded.pushMenu",this.resizeCharts),t(document).off("collapsed.pushMenu",this.resizeCharts)},methods:{fullscreen:function(){var t=this;this.$fullscreen.enter(this.$el.querySelector(".container-fluid"),{wrap:!1,callback:function(){t.resize(),t.resizeCharts()}})},top:function(){var e=this;t.get("/api/v1/dashboard/top").then(function(t){var a=t.data;e.memData.rows=a.memData,e.cpuData.rows=a.cpuData})},store:function(){var e=this;t.get("/api/v1/dashboard/store").then(function(t){var a=t.data;e.storeData.rows=a}),t.get("/api/v1/dashboard/auth").then(function(t){var a=t.data;e.authData=a})},isIE:function(){return!!(window.ActiveXObject||"ActiveXObject"in window)},initHeight:function(){this.pageWidth=window.innerWidth,this.pageHeight=window.innerHeight,"number"!=typeof this.pageWidth&&("CSS1Compat"==document.compatMode?(this.pageWidth=document.documentElement.clientWidth,this.pageHeight=document.documentElement.clientHeight):(this.pageWidth=document.body.clientWidth,this.pageHeight=document.body.clientHeight)),this.pageHeight=this.pageHeight-140},resizeCharts:function(){this.$refs.storeChart.resize(),this.$refs.cpuChart.resize(),this.$refs.memChart.resize()},resize:function(){this.initHeight(),this.fixAuthView()},fixAuthView:function(){this.pageHeight>=600?t(".auth-view").css("padding-top",this.pageHeight/10+"px"):t(".auth-view").css("padding-top","20px")}}}}).call(e,a("0iPh"))},"w+5F":function(t,e,a){"use strict";function i(t){a("ettG")}Object.defineProperty(e,"__esModule",{value:!0});var s=a("uBsg"),n=a.n(s);for(var o in s)["default","default"].indexOf(o)<0&&function(t){a.d(e,t,function(){return s[t]})}(o);var r=a("11aa"),d=a("C7Lr"),l=i,h=d(n.a,r.a,!1,l,"data-v-2b92d112",null);e.default=h.exports}});