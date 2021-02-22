{{define "content" }}

<main class="main">

    <ol class="breadcrumb">
        <li class="breadcrumb-item">
            <a href="/index">Home</a>
        </li>
        <li class="breadcrumb-item active">Test Mock Stub Collector</li>

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
                                        router>router {{.metrics.HubCollectorStubRouterCounter}}] --> |{{.metrics.HubCollectorStubGatewayLifeSupervisorCounter}}| gateway_life_supervisor
                                        gateway_life_supervisor --> |{{.metrics.HubCollectorStubCOGAdapterCounter}}| gateway_adapter

                                        router --> |{{.metrics.HubCollectorStubCoreSourceLifeSupervisorCounter}}| coresource_life_supervisor
                                        coresource_life_supervisor --> |{{.metrics.HubCollectorStubCORAdapterCounter}}| cor_adapter

                                        router --> |{{.metrics.HubCollectorStubCOVAdapterCounter}}| cov_adapter
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

<script src="/js/view/stub.js"></script>
{{end}}