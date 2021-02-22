{{define "content" }}

<main class="main">

  <ol class="breadcrumb">
    <li class="breadcrumb-item">
      <a href="/index">Home</a>
    </li>
    <li class="breadcrumb-item active">LiveEvent</li>

  </ol>
  <div class="container-fluid">
    <div id="ui-view">
      <div>
        <div class="row">
            <div class="col-sm-12 col-md-12">
                <div class="card">
                  <div class="card-header">LiveEvents</div>
                  <div class="card-body">
                    <div class="card">
                      <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                        <thead>
                          <tr>
                            <th>GatewayID</th>
                            <th>CoreSourceID</th>                            
                            <th>CorePointID</th>
                            <th>CorePointName</th>
                            <th>TriggerValue</th>
                            <th>EventState</th>
                            <th>StartTime</th>
                            <th>EndTime</th>
                            <th>StandardID</th>
                          </tr>
                        </thead>
                        <tbody>
                          {{range $i, $x := $.liveeventviews}}
                          <tr>
                            <td>{{$x.GatewayID}}</td>
                            <td>{{$x.CoreSourceID}}</td>                            
                            <td>{{$x.CorePointID}}</td>
                            <td>{{$x.CorePointName}}</td>
                            <td>{{$x.CurrentNumericValueString}}</td>
                            <td>{{$x.CurrentEventState}}</td>
                            <td>{{$x.StartTimeString}}</td>
                            <td>{{$x.EndTimeString}}</td>
                            <td>{{$x.StandardID}}</td>
                          </tr>
                          {{end}}    
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

<!-- <script src="/js/view/gateway.js" ></script> -->
{{end}}