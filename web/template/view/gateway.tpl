{{define "content" }}

<main class="main">

  <ol class="breadcrumb">
    <li class="breadcrumb-item">
      <a href="/index">Home</a>
    </li>
    <li class="breadcrumb-item"><a href="/domain">Domain</a></li>
    <li class="breadcrumb-item active">Gateway</li>

  </ol>
  <div class="container-fluid">
    <div id="ui-view">
      <div>
        <div class="row">
            <div class="col-sm-8 col-md-8">
                <div class="card">
                  <div class="card-header">CoreSources</div>
                  <div class="card-body">
                    <div class="card">
                      <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                        <thead>
                          <tr>
                            <th>CoreSourceID</th>
                            <th>GatewayID</th>
                            <th>UniqueID</th>
                            <th>SourceName</th>
                            <th>State</th>
                          </tr>
                        </thead>
                        <tbody>
                          {{range $i, $x := $.coresources}}
                          <tr>
                            <td>{{$x.CoreSourceID}}</td>
                            <td>{{$x.GatewayID}}</td>
                            <td>{{$x.UniqueID}}</td>
                            <td>{{$x.SourceName}}</td>
                            <td>{{$x.State}}</td>
                            <td><a href="javascript:void(0);" onclick="onLogClick({{$x.GatewayID}},{{$x.CoreSourceID}}); return false;">Log</a>
                                <a href="/coresource/?gatewayId={{$x.GatewayID}}&coresourceId={{$x.CoreSourceID}}">Detail</a>
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
                            <th>Flag</th>
                            <th>Timestamp</th>
                          </tr>
                        </thead>
                        <tbody id="packetList">

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

<script src="/js/view/gateway.js" ></script>
{{end}}