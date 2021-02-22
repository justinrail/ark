{{define "content"}}

<main class="main">

  <ol class="breadcrumb">
    <li class="breadcrumb-item">
      <a href="/index">Home</a>
    </li>
    <li class="breadcrumb-item active">Domain</li>

  </ol>
  <div class="container-fluid">
    <div id="ui-view">
      <div>
        <div class="row">
            <div class="col-sm-8 col-md-8">
                <div class="card">
                  <div class="card-header">Loaded Gateways</div>
                  <div class="card-body">
                    <div class="card">
                      <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                        <thead>
                          <tr>
                            <th>ID</th>
                            <th>UUID</th>
                            <th>Name</th>
                            <th>Collector</th>
                            <th>IP</th>
                            <th>ConState</th>
                            <th>SynState</th>
                            <th>Actions</th>
                          </tr>
                        </thead>
                        <tbody>
                          {{range $i, $x := $.gateways}}
                          <tr>
                            <td>{{$x.ID}}</td>
                            <td>{{$x.UUID}}</td>
                            <td>{{$x.Name}}</td>
                            <td>{{$x.Collector}}</td>
                            <td>{{$x.IP}}</td>
                            <td>{{$x.ConState}}</td>
                            <td>{{$x.SynState}}</td>
                            <td>
                              <a href="javascript:void(0);" onclick="onLogClick({{$x.ID}}); return false;">Log</a> 
                              <a href="/gateway/{{$x.ID}}">Detail</a>
                              <a href="/removegateway/?gatewayId={{$x.ID}}">Remove</a>
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

<script src="js/view/domain.js" ></script>
{{end}}