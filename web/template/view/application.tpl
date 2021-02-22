{{define "content"}}

<main class="main">

  <ol class="breadcrumb">
    <li class="breadcrumb-item">
      <a href="/index">Home</a>
    </li>
    <li class="breadcrumb-item active">Application</li>

  </ol>
  <div class="container-fluid">
    <div id="ui-view">
      <div>
        <div class="row">
          <div class="col-sm-6 col-md-6">
            <div class="card">
              <div class="card-header">Basic</div>
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
                      {{range $i, $x := $.apps}}
                      {{if eq $x.Category "basic"}}
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
                  <div class="card-header">Probe</div>
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
                          {{range $i, $x := $.apps}}
                          {{if eq $x.Category "probe"}}
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