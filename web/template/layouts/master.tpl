<!DOCTYPE html>
<!--
* CoreUI - Free Bootstrap Admin Template
* @version v2.1.11
-->

<html lang="en">
  <head>
    <base href="./">
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, shrink-to-fit=no">
    <meta name="description" content="CoreUI - Open Source Bootstrap Admin Template">
    <meta name="author" content="Łukasz Holeczek">
    <meta name="keyword" content="Bootstrap,Admin,Template,Open,Source,jQuery,CSS,HTML,RWD,Dashboard">
    <title>ark</title>
    <!-- Icons-->
    <link href="/vendors/@coreui/icons/css/coreui-icons.min.css" rel="stylesheet">
    <link href="/vendors/flag-icon-css/css/flag-icon.min.css" rel="stylesheet">
    <link href="/vendors/font-awesome/css/font-awesome.min.css" rel="stylesheet">
    <link href="/vendors/simple-line-icons/css/simple-line-icons.css" rel="stylesheet">
    <!-- Main styles for this application-->
    <link href="/css/style.css" rel="stylesheet">
    <link href="/vendors/pace-progress/css/pace.min.css" rel="stylesheet">
  </head>
  <body class="app header-fixed sidebar-fixed aside-menu-fixed sidebar-lg-show">

   {{ template "header" .}}

    <div class="app-body">
      {{ template "side" .}}
      {{ template "content" .}}
    </div>

    {{ template "footer" }}

    <!-- CoreUI and necessary plugins-->
    <script src="/vendors/jquery/js/jquery.min.js"></script>
    <script src="/js/jquery.rest.js"></script>
    <script src="/vendors/popper.js/js/popper.min.js"></script>
    <script src="/vendors/bootstrap/js/bootstrap.min.js"></script>
    <script src="/vendors/pace-progress/js/pace.min.js"></script> 
    <script src="/vendors/perfect-scrollbar/js/perfect-scrollbar.min.js"></script>
    <script src="/vendors/@coreui/coreui/js/coreui.min.js"></script>
    <!-- Plugins and scripts required by this view-->
    <script src="vendors/chart.js/js/Chart.min.js"></script>
    <script src="/js/echarts.js"></script>
    <script src="/js/mermaid.min.js"></script>
    <!-- <script src="vendors/@coreui/coreui-plugin-chartjs-custom-tooltips/js/custom-tooltips.min.js"></script> -->
    
  </body>
</html>