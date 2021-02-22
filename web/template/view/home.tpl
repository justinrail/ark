{{define "content"}}

<main class="main">

  <ol class="breadcrumb">
    <li class="breadcrumb-item">
      <a href="/index">Home</a>
    </li>
    <li class="breadcrumb-item active">Home</li>

  </ol>
  <div class="container-fluid">
    <div id="ui-view">
      <div>
        <div class="row">
          <div class="col-sm-10 col-md-10">
            <div class="row">
              <div class="col-sm-6 col-md-6">
                <div class="card text-white bg-primary">
                  <div class="card-body pb-0">
                    <div class="text-value">9.82%</div>
                    <div>CPU Usage</div>
                  </div>
                  <div class="chart-wrapper mt-3 mx-3" style="height:70px;">
                    <canvas class="chart chartjs-render-monitor" id="cpuline" height="70" width="327"
                      style="display: block; width: 327px; height: 70px;"></canvas>
                  </div>
                </div>
              </div>

              <div class="col-sm-6 col-md-6">
                <div class="card text-white bg-info">
                  <div class="card-body pb-0">
                    <div class="text-value">43.23%</div>
                    <div>Memory Usage</div>
                  </div>
                  <div class="chart-wrapper mt-3 mx-3" style="height:70px;">
                    <canvas class="chart chartjs-render-monitor" id="membar" height="70" width="327"
                      style="display: block; width: 327px; height: 70px;"></canvas>
                  </div>
                </div>

              </div>
            </div>
            <div class="card">
              <div class="card-body">
                <div class="row">
                  <div class="col-sm-5">
                    <h4 class="card-title mb-0">Key Data I/O Rate</h4>
                    <div class="small text-muted">pcs/second</div>
                  </div>
                </div>

                <div class="chart-wrapper" style="height:300px;margin-top:40px;">
                  <canvas class="chart chartjs-render-monitor" id="main-chart" height="300" width="806"
                    style="display: block; width: 806px; height: 300px;"></canvas>
                </div>
              </div>
              <div class="card-footer">
                <div class="row text-center">
                  <div class="col-sm-12 col-md mb-sm-2 mb-0">
                    <div class="text-muted">Gateways</div>
                    <strong>29 Connected (40%)</strong>
                    <div class="progress progress-xs mt-2">
                      <div class="progress-bar bg-success" role="progressbar" style="width: 40%" aria-valuenow="40"
                        aria-valuemin="0" aria-valuemax="100"></div>
                    </div>
                  </div>
                  <div class="col-sm-12 col-md mb-sm-2 mb-0">
                    <div class="text-muted">CoreSources</div>
                    <strong>124 Online (20%)</strong>
                    <div class="progress progress-xs mt-2">
                      <div class="progress-bar bg-info" role="progressbar" style="width: 20%" aria-valuenow="20"
                        aria-valuemin="0" aria-valuemax="100"></div>
                    </div>
                  </div>
                  <div class="col-sm-12 col-md mb-sm-2 mb-0">
                    <div class="text-muted">CorePoints</div>
                    <strong>78.706 Valid (60%)</strong>
                    <div class="progress progress-xs mt-2">
                      <div class="progress-bar bg-warning" role="progressbar" style="width: 60%" aria-valuenow="60"
                        aria-valuemin="0" aria-valuemax="100"></div>
                    </div>
                  </div>
                  <div class="col-sm-12 col-md mb-sm-2 mb-0">
                    <div class="text-muted">Active Events</div>
                    <strong>2512 Ended (80%)</strong>
                    <div class="progress progress-xs mt-2">
                      <div class="progress-bar bg-danger" role="progressbar" style="width: 80%" aria-valuenow="80"
                        aria-valuemin="0" aria-valuemax="100"></div>
                    </div>
                  </div>
                  <div class="col-sm-12 col-md mb-sm-2 mb-0">
                    <div class="text-muted">History Points</div>
                    <strong>5670303 Saved</strong>
                    <div class="progress progress-xs mt-2">
                      <div class="progress-bar" role="progressbar" style="width: 40%" aria-valuenow="40"
                        aria-valuemin="0" aria-valuemax="100"></div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="row">
              <div class="col-sm-6 col-lg-3">
                <div class="brand-card">
                  <div class="brand-card-header bg-facebook">
                    <i class="display-4">CPU</i>
                  </div>
                  <div class="brand-card-body">
                    <div>
                      <div class="text-value">4</div>
                      <div class="text-uppercase text-muted small">cores</div>
                    </div>
                    <div>
                      <div class="text-value">3.1</div>
                      <div class="text-uppercase text-muted small">Ghz</div>
                    </div>
                  </div>
                </div>
              </div>

              <div class="col-sm-6 col-lg-3">
                <div class="brand-card">
                  <div class="brand-card-header bg-twitter">
                    <i class="display-4">OS</i>
                  </div>
                  <div class="brand-card-body">
                    <div>
                      <div class="text-value">Debian</div>
                      <div class="text-uppercase text-muted small">Plantform</div>
                    </div>
                    <div>
                      <div class="text-value">9.0</div>
                      <div class="text-uppercase text-muted small">version</div>
                    </div>
                  </div>
                </div>
              </div>

              <div class="col-sm-6 col-lg-3">
                <div class="brand-card">
                  <div class="brand-card-header bg-linkedin">
                    <i class="display-4">Threads</i>
                  </div>
                  <div class="brand-card-body">
                    <div>
                      <div class="text-value">11</div>
                      <div class="text-uppercase text-muted small">Threads</div>
                    </div>
                    <div>
                      <div class="text-value">29</div>
                      <div class="text-uppercase text-muted small">Routines</div>
                    </div>
                  </div>
                </div>
              </div>

              <div class="col-sm-6 col-lg-3">
                <div class="brand-card">
                  <div class="brand-card-header bg-google-plus">
                    <i class="display-4">Disk</i>
                  </div>
                  <div class="brand-card-body">
                    <div>
                      <div class="text-value">40Gb</div>
                      <div class="text-uppercase text-muted small">total</div>
                    </div>
                    <div>
                      <div class="text-value">21.2Gb</div>
                      <div class="text-uppercase text-muted small">free</div>
                    </div>
                  </div>
                </div>
              </div>

            </div>
            <div class="row">
              <div class="col-sm-6 col-lg-3">
                <div class="card">
                  <div class="card-body">
                    <div class="text-value">892</div>
                    <div>Total of Complex Index</div>
                    <div class="progress progress-xs my-2">
                      <div class="progress-bar bg-success" role="progressbar" style="width: 25%" aria-valuenow="25"
                        aria-valuemin="0" aria-valuemax="100"></div>
                    </div>
                    <small class="text-muted">Index caculateion using IOT data.</small>
                  </div>
                </div>
              </div>

              <div class="col-sm-6 col-lg-3">
                <div class="card">
                  <div class="card-body">
                    <div class="text-value">12</div>
                    <div>Total of Notify Rules</div>
                    <div class="progress progress-xs my-2">
                      <div class="progress-bar bg-info" role="progressbar" style="width: 25%" aria-valuenow="25"
                        aria-valuemin="0" aria-valuemax="100"></div>
                    </div>
                    <small class="text-muted">Notification rule engine service.</small>
                  </div>
                </div>
              </div>

              <div class="col-sm-6 col-lg-3">
                <div class="card">
                  <div class="card-body">
                    <div class="text-value">981</div>
                    <div>Cameras of Stream Server</div>
                    <div class="progress progress-xs my-2">
                      <div class="progress-bar bg-warning" role="progressbar" style="width: 25%" aria-valuenow="25"
                        aria-valuemin="0" aria-valuemax="100"></div>
                    </div>
                    <small class="text-muted">Video Server stream channels.</small>
                  </div>
                </div>
              </div>

              <div class="col-sm-6 col-lg-3">
                <div class="card">
                  <div class="card-body">
                    <div class="text-value">234691</div>
                    <div>Graph Nodes Count</div>
                    <div class="progress progress-xs my-2">
                      <div class="progress-bar bg-danger" role="progressbar" style="width: 25%" aria-valuenow="25"
                        aria-valuemin="0" aria-valuemax="100"></div>
                    </div>
                    <small class="text-muted">Graph database nodes count.</small>
                  </div>
                </div>
              </div>

            </div>
          </div>
          <div class="col-sm-2 col-md-2">
            <div class="tab-pane p-3 active show" id="settings" role="tabpanel">
              <h6>Modules</h6>
              <div class="aside-options">
                <div class="clearfix mt-4">
                  <small>
                    <b>Stub Collector</b>
                  </small>
                  <label class="switch switch-label switch-pill switch-success switch-sm float-right">
                    <input class="switch-input" type="checkbox" checked="">
                    <span class="switch-slider" data-checked="On" data-unchecked="Off"></span>
                  </label>
                </div>
                <div>
                  <small class="text-muted">Mock collector for test and demo.</small>
                </div>
              </div>
              <div class="aside-options">
                <div class="clearfix mt-3">
                  <small>
                    <b>CMB Collector</b>
                  </small>
                  <label class="switch switch-label switch-pill switch-success switch-sm float-right">
                    <input class="switch-input" type="checkbox">
                    <span class="switch-slider" data-checked="On" data-unchecked="Off"></span>
                  </label>
                </div>
                <div>
                  <small class="text-muted">China Mobile B-Interface collector for FSU connected in.</small>
                </div>
              </div>
              <div class="aside-options">
                <div class="clearfix mt-3">
                  <small>
                    <b>C2C Collector</b>
                  </small>
                  <label class="switch switch-label switch-pill switch-success switch-sm float-right">
                    <input class="switch-input" type="checkbox">
                    <span class="switch-slider" data-checked="On" data-unchecked="Off"></span>
                  </label>
                </div>
                <div>
                  <small class="text-muted">ark to ark RPC Aggregator.</small>
                </div>
              </div>
              <div class="aside-options">
                <div class="clearfix mt-3">
                  <small>
                    <b>Stream Server</b>
                  </small>
                  <label class="switch switch-label switch-pill switch-success switch-sm float-right">
                    <input class="switch-input" type="checkbox" checked=“”>
                    <span class="switch-slider" data-checked="On" data-unchecked="Off"></span>
                  </label>
                </div>
                <div>
                  <small class="text-muted">Video Stream RSTP to HLS Server.</small>
                </div>
              </div>
              <hr>
              <h6>System Utilization</h6>
              <div class="text-uppercase mb-1 mt-4">
                <small>
                  <b>CPU Usage</b>
                </small>
              </div>
              <div class="progress progress-xs">
                <div class="progress-bar bg-info" role="progressbar" style="width: 25%" aria-valuenow="25"
                  aria-valuemin="0" aria-valuemax="100"></div>
              </div>
              <small class="text-muted">348 Processes. 1/4 Cores.</small>
              <div class="text-uppercase mb-1 mt-2">
                <small>
                  <b>Memory Usage</b>
                </small>
              </div>
              <div class="progress progress-xs">
                <div class="progress-bar bg-warning" role="progressbar" style="width: 70%" aria-valuenow="70"
                  aria-valuemin="0" aria-valuemax="100"></div>
              </div>
              <small class="text-muted">11444GB/16384MB</small>
              <div class="text-uppercase mb-1 mt-2">
                <small>
                  <b>SSD 1 Usage</b>
                </small>
              </div>
              <div class="progress progress-xs">
                <div class="progress-bar bg-danger" role="progressbar" style="width: 95%" aria-valuenow="95"
                  aria-valuemin="0" aria-valuemax="100"></div>
              </div>
              <small class="text-muted">243GB/256GB</small>
              <div class="text-uppercase mb-1 mt-2">
                <small>
                  <b>SSD 2 Usage</b>
                </small>
              </div>
              <div class="progress progress-xs">
                <div class="progress-bar bg-success" role="progressbar" style="width: 10%" aria-valuenow="10"
                  aria-valuemin="0" aria-valuemax="100"></div>
              </div>
              <small class="text-muted">25GB/256GB</small>
            </div>
          </div>
        </div>
      </div>
    </div>
</main>
<script src="/js/view/home.js"></script>
{{end}}