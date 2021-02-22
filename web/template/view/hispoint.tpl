{{define "content" }}

<main class="main">

  <ol class="breadcrumb">
    <li class="breadcrumb-item">
      <a href="/index">Home</a>
    </li>
    <li class="breadcrumb-item active">History Point</li>

  </ol>
  <div class="container-fluid">
    <div id="ui-view">
      <div>
        <div class="row">
            <div class="col-sm-12 col-md-12">
                <div class="card">
                  <div class="card-header">History of Points</div>
                  <div class="card-body">
                    <div class="card">
                      <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                        <thead>
                          <tr>
                            <th>CoreSourceID</th>                            
                            <th>CorePointID</th>
                            <th>CurrentNumericValueString</th>
                            <th>CurrentStringValue</th>
                            <th>LimitState</th>
                            <th>TimestampString</th>
                            <th>StandardID</th>
                          </tr>
                        </thead>
                        <tbody>
                          {{range $i, $x := $.hispointviews}}
                          <tr>
                            <td>{{$x.CoreSourceID}}</td>                            
                            <td>{{$x.CorePointID}}</td>
                            <td>{{$x.CurrentNumericValueString}}</td>
                            <td>{{$x.CurrentStringValue}}</td>
                            <td>{{$x.LimitState}}</td>
                            <td>{{$x.TimestampString}}</td>
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

{{end}}