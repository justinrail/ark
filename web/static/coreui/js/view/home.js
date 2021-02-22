
function run (){
    /* eslint-disable no-magic-numbers */
// Disable the on-canvas tooltip
Chart.defaults.global.pointHitDetectionRadius = 1;
Chart.defaults.global.tooltips.enabled = false;
Chart.defaults.global.tooltips.mode = 'index';
Chart.defaults.global.tooltips.position = 'nearest';
//Chart.defaults.global.tooltips.custom = CustomTooltips; // eslint-disable-next-line no-unused-vars
var ctxCPU = document.getElementById("cpuline");

var cardChart1 = new Chart(ctxCPU, {
  type: 'line',
  data: {
    labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
    datasets: [{
      label: 'My First dataset',
      backgroundColor: getStyle('--primary'),
      borderColor: 'rgba(255,255,255,.55)',
      data: [65, 59, 84, 84, 51, 55, 40]
    }]
  },
  options: {
    maintainAspectRatio: false,
    legend: {
      display: false
    },
    scales: {
      xAxes: [{
        gridLines: {
          color: 'transparent',
          zeroLineColor: 'transparent'
        },
        ticks: {
          fontSize: 2,
          fontColor: 'transparent'
        }
      }],
      yAxes: [{
        display: false,
        ticks: {
          display: false,
          min: 35,
          max: 89
        }
      }]
    },
    elements: {
      line: {
        borderWidth: 1
      },
      point: {
        radius: 4,
        hitRadius: 10,
        hoverRadius: 4
      }
    }
  }
});

var ctxMEM = document.getElementById("membar");
var cardChart4 = new Chart(ctxMEM, {
    type: 'bar',
    data: {
      labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December', 'January', 'February', 'March', 'April'],
      datasets: [{
        label: 'My First dataset',
        backgroundColor: 'rgba(255,255,255,.2)',
        borderColor: 'rgba(255,255,255,.55)',
        data: [78, 81, 80, 45, 34, 12, 40, 85, 65, 23, 12, 98, 34, 84, 67, 82]
      }]
    },
    options: {
      maintainAspectRatio: false,
      legend: {
        display: false
      },
      scales: {
        xAxes: [{
          display: false,
          barPercentage: 0.6
        }],
        yAxes: [{
          display: false
        }]
      }
    }
  });

  var ctxMain = document.getElementById("main-chart");
  var mainChart = new Chart(ctxMain, {
    type: 'line',
    data: {
      labels: ['M', 'T', 'W', 'T', 'F', 'S', 'S', 'M', 'T', 'W', 'T', 'F', 'S', 'S', 'M', 'T', 'W', 'T', 'F', 'S', 'S', 'M', 'T', 'W', 'T', 'F', 'S', 'S'],
      datasets: [{
        label: 'My First dataset',
        backgroundColor: hexToRgba(getStyle('--info'), 10),
        borderColor: getStyle('--info'),
        pointHoverBackgroundColor: '#fff',
        borderWidth: 2,
        data: [165, 180, 70, 69, 77, 57, 125, 165, 172, 91, 173, 138, 155, 89, 50, 161, 65, 163, 160, 103, 114, 185, 125, 196, 183, 64, 137, 95, 112, 175]
      }, {
        label: 'My Second dataset',
        backgroundColor: 'transparent',
        borderColor: getStyle('--success'),
        pointHoverBackgroundColor: '#fff',
        borderWidth: 2,
        data: [92, 97, 80, 100, 86, 97, 83, 98, 87, 98, 93, 83, 87, 98, 96, 84, 91, 97, 88, 86, 94, 86, 95, 91, 98, 91, 92, 80, 83, 82]
      }, {
        label: 'My Third dataset',
        backgroundColor: 'transparent',
        borderColor: getStyle('--danger'),
        pointHoverBackgroundColor: '#fff',
        borderWidth: 1,
        borderDash: [8, 5],
        data: [65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65, 65]
      }]
    },
    options: {
      maintainAspectRatio: false,
      legend: {
        display: false
      },
      scales: {
        xAxes: [{
          gridLines: {
            drawOnChartArea: false
          }
        }],
        yAxes: [{
          ticks: {
            beginAtZero: true,
            maxTicksLimit: 5,
            stepSize: Math.ceil(250 / 5),
            max: 250
          }
        }]
      },
      elements: {
        point: {
          radius: 0,
          hitRadius: 10,
          hoverRadius: 4,
          hoverBorderWidth: 3
        }
      }
    }
  });

    //  // 基于准备好的dom，初始化echarts实例
    //  var cpuChart = echarts.init(document.getElementById('c'),'light');

    //  // 指定图表的配置项和数据
    //  var option = {
    //      tooltip: {

    //      },
    //      xAxis: {
    //         show: false,
    //          data: ["14:20","14:25","14:30","14:35","14:40","14:45"]
    //      },
    //      yAxis: {
    //          show: false
    //      },
    //      series: [{
    //         itemStyle: {
    //             // 点的颜色。
    //             color: '#99d2dd'
    //         },
    //          type: 'line',
    //          smooth:true,
    //          name: 'cpu ',
    //          data: [54.5, 20, 36, 10, 10, 20]
    //      }]
    //  };

    //  // 使用刚指定的配置项和数据显示图表。
    //  cpuChart.setOption(option);
    //  cpuChart.resize();
}


//因为domain.js在jquery前加载，所以juery的加载判断是无法执行的
if (document.readyState !== 'loading') {
    run();
  } else {
    document.addEventListener('DOMContentLoaded', run);
}