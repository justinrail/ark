


function run (){
    mermaid.initialize({ theme: 'forest', startOnLoad:true }) 

    // var myChart = echarts.init(document.getElementById('graph'));

    // // 指定图表的配置项和数据
    // var option = {
    //     title: {
    //         text: 'topology DAG process flow'
    //     },
    //     tooltip: {
    //         trigger: 'item',
    //         triggerOn: 'mousemove'
    //     },
    //     series: [
    //         {
    //             type: 'sankey',
    //             data: [{
    //                 name: 'cog_spout'
    //             }, {
    //                 name: 'cog_state_updater'
    //             }, {
    //                 name: 'cog_ashbin'
    //             }, {
    //                 name: 'cor_spout'
    //             }, {
    //                 name: 'cor_state_updater'
    //             }, {
    //                 name: 'cor_ashbin'
    //             }, {
    //                 name: 'cov_ashbin'
    //             }, {
    //                 name: 'cov_spout'
    //             }, {
    //                 name: 'cov_data_updater'
    //             }, {
    //                 name: 'cov_state_normalizer'
    //             }, {
    //                 name: 'cov_state_updater'
    //             }, {
    //                 name: 'cov_coreactiveevent_updater'
    //             }, {
    //                 name: 'cov_hsicoreactiveevent_appender'
    //             }],
    //             links: [{
    //                 source: 'cov_spout',
    //                 target: 'cov_data_updater',
    //                 value:9
    //             },{
    //                 source: 'cov_data_updater',
    //                 target: 'cov_state_normalizer',
    //                 value:9
    //             },{
    //                 source: 'cov_state_normalizer',
    //                 target: 'cov_state_updater',
    //                 value:3
    //             },{
    //                 source: 'cov_state_updater',
    //                 target: 'cov_coreactiveevent_updater',
    //                 value:1
    //             },{
    //                 source: 'cov_coreactiveevent_updater',
    //                 target: 'cov_hsicoreactiveevent_appender',
    //                 value:0
    //             },{
    //                 source: 'cog_spout',
    //                 target: 'cog_state_updater',
    //                 value:1
    //             },{
    //                 source: 'cor_spout',
    //                 target: 'cor_state_updater',
    //                 value:1
    //             }],
    //             focusNodeAdjacency: 'allEdges',
    //             itemStyle: {
    //                 normal: {
    //                     borderWidth: 1,
    //                     borderColor: '#aaa'
    //                 }
    //             },
    //             lineStyle: {
    //                 normal: {
    //                     color: 'source',
    //                     curveness: 0.5
    //                 }
    //             }
    //         }
    //     ]
    // };

    // // 使用刚指定的配置项和数据显示图表。
    // myChart.setOption(option);
}


//因为domain.js在jquery前加载，所以juery的加载判断是无法执行的
if (document.readyState !== 'loading') {
    run();
  } else {
    document.addEventListener('DOMContentLoaded', run);
}