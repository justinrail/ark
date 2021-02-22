{{define "content" }}

<main class="main">

    <ol class="breadcrumb">
        <li class="breadcrumb-item">
            <a href="/index">Home</a>
        </li>
        <li class="breadcrumb-item active">Topology</li>

    </ol>
    <div class="container-fluid">
        <div id="ui-view">
            <div>
                <div class="row">
                    <div class="col-sm-4 col-md-4">
                        <div class="card">
                            <div class="card-header">Data Queues</div>
                            <div class="card-body">
                                <div class="progress-group">
                                    <div class="progress-group-header align-items-end">
                                        <i class="icon-reload progress-group-icon"></i>
                                        <div>BUS/COG</div>
                                        <div class="ml-auto font-weight-bold mr-2">Total:{{.metrics.HubBusCOGCounter}}
                                        </div>
                                        <div class="text-muted small">(Wait:{{.metrics.HubBusCOGPercent}}%)</div>
                                    </div>
                                    <div class="progress-group-bars">
                                        <div class="progress progress-xs">
                                            <div class="progress-bar bg-success" role="progressbar"
                                                style="width: {{.metrics.HubBusCOGPercent}}%"
                                                aria-valuenow="{{.metrics.HubBusCOGPercent}}" aria-valuemin="0"
                                                aria-valuemax="100"></div>
                                        </div>
                                    </div>
                                </div>
                                <div class="progress-group">
                                    <div class="progress-group-header align-items-end">
                                        <i class="icon-reload progress-group-icon"></i>
                                        <div>BUS/COR</div>
                                        <div class="ml-auto font-weight-bold mr-2">Total:{{.metrics.HubBusCORCounter}}
                                        </div>
                                        <div class="text-muted small">(Wait:{{.metrics.HubBusCORPercent}}%)</div>
                                    </div>
                                    <div class="progress-group-bars">
                                        <div class="progress progress-xs">
                                            <div class="progress-bar bg-success" role="progressbar"
                                                style="width: {{.metrics.HubBusCOGPercent}}%"
                                                aria-valuenow="{{.metrics.HubBusCOGPercent}}" aria-valuemin="0"
                                                aria-valuemax="100"></div>
                                        </div>
                                    </div>
                                </div>
                                <div class="progress-group">
                                    <div class="progress-group-header align-items-end">
                                        <i class="icon-reload progress-group-icon"></i>
                                        <div>BUS/COV</div>
                                        <div class="ml-auto font-weight-bold mr-2">Total:{{.metrics.HubBusCOVCounter}}
                                        </div>
                                        <div class="text-muted small">(Wait:{{.metrics.HubBusCOVPercent}}%)</div>
                                    </div>
                                    <div class="progress-group-bars">
                                        <div class="progress progress-xs">
                                            <div class="progress-bar bg-success" role="progressbar"
                                                style="width: {{.metrics.HubBusCOGPercent}}%"
                                                aria-valuenow="{{.metrics.HubBusCOGPercent}}" aria-valuemin="0"
                                                aria-valuemax="100"></div>
                                        </div>
                                    </div>
                                </div>
                                <div class="progress-group">
                                    <div class="progress-group-header align-items-end">
                                        <i class="icon-reload progress-group-icon"></i>
                                        <div>Sink/HisCorePoint</div>
                                        <div class="ml-auto font-weight-bold mr-2">
                                            Total:{{.metrics.HubSinkHisCorePointCounter}}
                                            Lost:{{.metrics.HubSinkHisCorePointLostCounter}}</div>
                                        <div class="text-muted small">(Wait:{{.metrics.HubSinkHisCorePointPercent}}%)
                                        </div>
                                    </div>
                                    <div class="progress-group-bars">
                                        <div class="progress progress-xs">
                                            <div class="progress-bar bg-primary" role="progressbar"
                                                style="width: {{.metrics.HubSinkHisCorePointPercent}}%"
                                                aria-valuenow="{{.metrics.HubSinkHisCorePointPercent}}"
                                                aria-valuemin="0" aria-valuemax="100"></div>
                                        </div>
                                    </div>
                                </div>
                                <div class="progress-group">
                                    <div class="progress-group-header align-items-end">
                                        <i class="icon-reload progress-group-icon"></i>
                                        <div>Sink/HisComplexIndex</div>
                                        <div class="ml-auto font-weight-bold mr-2">
                                            Total:{{.metrics.HubSinkHisComplexIndexCounter}}
                                            Lost:{{.metrics.HubSinkHisComplexIndexLostCounter}}</div>
                                        <div class="text-muted small">(Wait:{{.metrics.HubSinkHisComplexIndexPercent}}%)
                                        </div>
                                    </div>
                                    <div class="progress-group-bars">
                                        <div class="progress progress-xs">
                                            <div class="progress-bar bg-primary" role="progressbar"
                                                style="width: {{.metrics.HubSinkHisComplexIndexPercent}}%"
                                                aria-valuenow="{{.metrics.HubSinkHisComplexIndexPercent}}"
                                                aria-valuemin="0" aria-valuemax="100"></div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="col-sm-6 col-md-6">
                        <div class="card">
                            <div class="card-header">DAG Graph</div>
                            <div class="card-body">
                                <div class="card">
                                    <!-- 
                        <div id="graph" style="width: auto;height:500px;"></div> -->
                                    <div class="mermaid">
                                        graph TD
                                        cog_spout --> |{{.metrics.HubTopoCOGStateUpdaterCounter}}| cog_state_updater
                                        cor_spout --> |{{.metrics.HubTopoCORStateUpdaterCounter}}| cor_state_updater
                                        cov_spout --> |{{.metrics.HubTopoCOVDataUpdaterCounter}}| cov_data_updater
                                        cov_data_updater --> |{{.metrics.HubTopoCOVStateNormalizerCounter}}| cov_state_normalizer
                                        cov_state_normalizer --> |{{.metrics.HubTopoCorePointStateUpdaterCounter}}| corepoint_state_updater
                                        corepoint_state_updater --> |{{.metrics.HubTopoCoreActiveEventUpdaterCounter}}| coreactiveevent_updater
                                        corepoint_state_updater --> |{{.metrics.HubTopoPhoenixCOVCounter}}| phoenix_cov_hooker
                                        coreactiveevent_updater --> |{{.metrics.HubTopoHisCoreActiveEventAppenderCounter}}| hiscoreactiveevent_appender
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div class="col-sm-2 col-md-2">
                        <div class="card">
                            <div class="card-header">Fuse</div>
                            <div class="card-body">
                                <button class="btn btn-block btn-ghost-success active" type="button"
                                    aria-pressed="true">Primary</button>
                                <button class="btn btn-block btn-ghost-success" type="button"
                                    aria-pressed="true">Primary</button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</main>

<script src="/js/view/topology.js"></script>
{{end}}