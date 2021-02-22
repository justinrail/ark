
var client 

function run (){
    client = new $.RestClient('/api/');
    client.add('coresourcepackets');
}

function onLogClick(gId,cId) {    
    client.coresourcepackets.read({ gatewayId:gId }, { coresourceId:cId }).done(function(data) {

        $("#packetList").empty()
        data.forEach(element => {
            $("#packetList").append("<tr><td>"+element.ItemName+"</td><td>" + element.ItemValue +"</td></tr>"); 
        });
      });    
}

//因为domain.js在jquery前加载，所以juery的加载判断是无法执行的
if (document.readyState !== 'loading') {
    run();
  } else {
    document.addEventListener('DOMContentLoaded', run);
}