{{define "content"}}

<main class="main">

  <ol class="breadcrumb">
    <li class="breadcrumb-item">
      <a href="/index">Home</a>
    </li>
    <li class="breadcrumb-item active">Machine</li>

  </ol>
  <div class="container-fluid">
    <div id="ui-view">
      <div>
        <div class="row">
                <div class="col-sm-6 col-md-6">
                        <div class="card">
                          <div class="card-header">Host</div>
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
                                  {{range $i, $x := $.machines}}
                                  {{if eq $x.Category "host"}}
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
          <div class="col-sm-3 col-md-3">
            <div class="card">
              <div class="card-header">Disk</div>
              <div class="card-body">
                <div class="card">
                  <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                    <thead>
                      <tr>
                        <th>Item</th>
                        <th>Value</th>
                      </tr>
                    </thead>
                    <tbody>
                      {{range $i, $d := $.machines}}
                      {{if eq $d.Category "disk"}}
                      <tr>
                        <td>{{$d.ItemName}}</td>
                        <td>{{$d.ItemValue}}</td>
                      </tr>
                      {{end}}
                      {{end}}

                    </tbody>
                  </table>
                </div>
              </div>
            </div>
          </div>
          <div class="col-sm-3 col-md-3">
                <div class="card">
                  <div class="card-header">Memory</div>
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
                          {{range $i, $x := $.machines}}
                          {{if eq $x.Category "mem"}}
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
              <div class="card-header">CPU</div>
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
                      {{range $i, $c := $.machines}}
                      {{if eq $c.Category "cpu"}}
                      <tr>
                        <td>{{$c.ItemName}}</td>
                        <td>{{$c.ItemValue}}</td>
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
              <div class="card-header">Net Interface</div>
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
                      {{range $i, $x := $.machines}}
                      {{if eq $x.Category "interface"}}
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