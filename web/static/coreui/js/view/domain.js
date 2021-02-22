
var client 

function run (){
    client = new $.RestClient('/api/');
    client.add('gatewaypackets');
    client.add('gateways');
}

function onLogClick(gatewayId) {    
    client.gatewaypackets.read(gatewayId).done(function(data) {

        $("#packetList").empty()
        data.forEach(element => {
            $("#packetList").append("<tr><td>"+element.ItemName+"</td><td>" + element.ItemValue +"</td></tr>"); 
        });
      });    
}

// function onRemoveClick(gatewayId) {    
//     client.gateways.del(gatewayId).done(function(data) {
//         alert('sss');
//     });
// }

//因为domain.js在jquery前加载，所以juery的加载判断是无法执行的
if (document.readyState !== 'loading') {
    run();
  } else {
    document.addEventListener('DOMContentLoaded', run);
}