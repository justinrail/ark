{{define "content" }}

<main class="main">

    <ol class="breadcrumb">
        <li class="breadcrumb-item">
            <a href="/index">Home</a>
        </li>
        <li class="breadcrumb-item active">China Mobile B Collector</li>

    </ol>
    <div class="container-fluid">
        <div id="ui-view">
            <div>
                <div class="row">
                    <div class="col-sm-12 col-md-12">
                        <div class="card">
                            <div class="card-header">DAG Graph</div>
                            <div class="card-body">
                                <div class="card">
                                    <!-- 
                        <div id="graph" style="width: auto;height:500px;"></div> -->
                                    <div class="mermaid">
                                        graph TD
                                        router>router {{.metrics.HubCollectorCMBRouterCounter}}] --> |{{.metrics.HubCollectorCMBLoginProcessorCounter}}| login_processor
                                        
                                        login_processor --> cog_bus>cog_bus {{.metrics.HubCollectorCMBCOGCounter}}]

                                        router --> |{{.metrics.HubCollectorCMBSendConfigProcessorCounter}}| sendconfig_processor
                                        sendconfig_processor --> cog_bus

                                        router --> |{{.metrics.HubCollectorCMBSendAlarmProcessorCounter}}| sendalarm_processor
                                        sendalarm_processor --> cov_bus>cov_bus {{.metrics.HubCollectorCMBCOVCounter}}]

                                        cov_task>cov_task {{.metrics.HubCollectorCMBGetDataTaskCounter}}] --> cov_bus
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</main>

<script src="/js/view/cmb.js"></script>
{{end}}