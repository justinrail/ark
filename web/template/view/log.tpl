{{define "content"}}

<main class="main">

  <ol class="breadcrumb">
    <li class="breadcrumb-item">
      <a href="/index">Home</a>
    </li>
    <li class="breadcrumb-item active">Log</li>

  </ol>
  <div class="container-fluid">
    <div id="ui-view">
      <div>
        <div class="row">
          <div class="col-sm-12 col-md-12">
            <div class="card">
              <div class="card-header">Log History</div>
              <div class="card-body">
                  <div class="card">
                    <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                      <thead>
                        <tr>
                          <th>Content</th>
                        </tr>
                      </thead>
                      <tbody>
                        {{range $i, $x := $.LogContents}}
                        <tr>
                          <td>{{$x}}</td>
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
</main>

{{end}}