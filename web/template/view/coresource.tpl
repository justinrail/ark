{{define "content" }}

<main class="main">

  <ol class="breadcrumb">
    <li class="breadcrumb-item">
      <a href="/index">Home</a>
    </li>
    <li class="breadcrumb-item"><a href="/domain">Domain</a></li>
    <li class="breadcrumb-item"><a href="/gateway/{{$.GatewayID}}">Gateway {{$.GatewayID}}</a></li>
    <li class="breadcrumb-item active">CoreSource</li>

  </ol>
  <div class="container-fluid">
    <div id="ui-view">
      <div>
        <div class="row">
            <div class="col-sm-8 col-md-8">
                <div class="card">
                  <div class="card-header">CorePoints</div>
                  <div class="card-body">
                    <div class="card">
                      <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                        <thead>
                          <tr>
                            <th>ID</th>
                            <th>Name</th>
                            <th>Cron</th>
                            <th>Expression</th>
                            <th>NumericValue</th>
                            <th>StringValue</th>
                            <th>EventState</th>
                            <th>IsAvailabe</th>
                            <th>LimitState</th>
                            <th>UpdateTime</th>
                          </tr>
                        </thead>
                        <tbody>
                          {{range $i, $x := $.corepointviews}}
                          <tr>
                            <td>{{$x.CorePointID}}</td>
                            <td>{{$x.PointName}}</td>
                            <td>{{$x.Cron}}</td>
                            <td>{{$x.Expression}}</td>
                            <td>{{$x.CurrentNumericValue}}</td>
                            <td>{{$x.CurrentStringValue}}</td>
                            <td>{{$x.CurrentEventState}}</td>
                            <td>{{$x.IsAvailabe}}</td>
                            <td>{{$x.LimitState}}</td>
                            <td>{{$x.UpdateTimeString}}</td>
                            
                            <td><a href="javascript:void(0);" onclick="onLogClick({{$x.GatewayID}},{{$x.CoreSourceID}},{{$x.CorePointID}}); return false;">Log</a>
                            </td>
                          </tr>
                          {{end}}    
                        </tbody>
                      </table>
                    </div>
                  </div>
                </div>
              </div>
              <div class="col-sm-4 col-md-4">
                <div class="card">
                  <div class="card-header">Recent Packets</div>
                  <div class="card-body">
                    <div class="card">
                      <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                        <thead>
                          <tr>
                            <th>Value</th>
                            <th>Timestamp</th>
                          </tr>
                        </thead>
                        <tbody id="packetList">

                        </tbody>
                      </table>
                    </div>
                  </div>
                </div>
                <div class="card">
                  <div class="card-header">Recent States</div>
                  <div class="card-body">
                    <div class="card">
                      <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                        <thead>
                          <tr>
                            <th>State</th>
                            <th>Timestamp</th>
                          </tr>
                        </thead>
                        <tbody id="stateList">
                        </tbody>
                      </table>
                    </div>
                  </div>
                </div>
              </div>
        </div>
      </div>
    </div>
  </div>
</main>

<script src="/js/view/coresource.js" ></script>
{{end}}