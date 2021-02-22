{{define "content"}}

<main class="main">

  <ol class="breadcrumb">
    <li class="breadcrumb-item">
      <a href="/index">Home</a>
    </li>
    <li class="breadcrumb-item active">Runtime</li>

  </ol>
  <div class="container-fluid">
    <div id="ui-view">
      <div>
          <div class="col-sm-6 col-md-6">
              <div class="card">
                <div class="card-header">NET IO Counters</div>
                <div class="card-body">
                  <div class="card">
                    <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                      <thead>
                        <tr>
                          <th> Item</th>
                          <th> Value</th>
                        </tr>
                      </thead>
                      <tbody>
                        {{range $i, $x := $.runtimes}}
                        {{if eq $x.Category "netio"}}
                        <tr>
                          <td>{{$x.ItemName}}</td>
                          <td>{{$x.ItemValue}}</td>
                        </tr>
                        {{end}}
                        {{end}}
  
                      </tbody>
                    </table>
                  </div>
                </div>
              </div>
            </div>

          <div class="col-sm-6 col-md-6">
              <div class="card">
                <div class="card-header">Connections</div>
                <div class="card-body">
                  <div class="card">
                    <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                      <thead>
                        <tr>
                          <th>Local -- Remote </th>
                          <th>Status</th>
                        </tr>
                      </thead>
                      <tbody>
                        {{range $i, $x := $.runtimes}}
                        {{if eq $x.Category "connection"}}
                        <tr>
                          <td>{{$x.ItemName}}</td>
                          <td>{{$x.ItemValue}}</td>
                        </tr>
                        {{end}}
                        {{end}}
  
                      </tbody>
                    </table>
                  </div>
                </div>
              </div>
            </div>

          <div class="col-sm-12 col-md-12">
            <div class="card">
              <div class="card-header">Processes</div>
              <div class="card-body">
                <div class="card">
                  <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                    <thead>
                      <tr>
                        <th> Item</th>
                        <th> Value</th>
                      </tr>
                    </thead>
                    <tbody>
                      {{range $i, $x := $.runtimes}}
                      {{if eq $x.Category "process"}}
                      <tr>
                        <td>{{$x.ItemName}}</td>
                        <td>{{$x.ItemValue}}</td>
                      </tr>
                      {{end}}
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