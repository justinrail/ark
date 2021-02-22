
var client 

function run (){
    client = new $.RestClient('/api/');
    client.add('corepointpackets');
    client.add('corepointstatepackets');
}
function onLogClick(gId,cId,pId) {    
    var qs = "?gatewayId=" +gId + "&coresourceId=" + cId + "&corepointId=" + pId
    //{gatewayId:gId},{coresourceId:cId},{corepointId:pId}
    //这个地方不直接传递对象而是手工拼接URL是这个类库不支持3个及以上参数的bug。但手工拼接会多出一个“/"需要在服务器特殊处理
    client.corepointpackets.read(qs).done(function(data) {

        $("#packetList").empty()
        data.forEach(element => {
            $("#packetList").append("<tr><td>"+element.ItemName+"</td><td>" + element.ItemValue +"</td></tr>"); 
        });
      });    

    client.corepointstatepackets.read(qs).done(function(data) {

        $("#stateList").empty()
        data.forEach(element => {
            $("#stateList").append("<tr><td>"+element.ItemName+"</td><td>" + element.ItemValue +"</td></tr>"); 
        });
      });    
}

//因为view.js在jquery前加载，所以juery的加载判断是无法执行的
if (document.readyState !== 'loading') {
    run();
  } else {
    document.addEventListener('DOMContentLoaded', run);
}