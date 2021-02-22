{{define "content" }}

<main class="main">

  <ol class="breadcrumb">
    <li class="breadcrumb-item">
      <a href="/index">Home</a>
    </li>
    <li class="breadcrumb-item"><a href="/notifyrule">NotifyRule: {{$.NotifyRuleName}}</a></li>
    <li class="breadcrumb-item active">NotifyRule Message</li>

  </ol>
  <div class="container-fluid">
    <div id="ui-view">
      <div>
        <div class="row">
            <div class="col-sm-12 col-md-12">
                <div class="card">
                  <div class="card-header">Notify Rule's Message</div>
                  <div class="card-body">
                    <div class="card">
                      <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                        <thead>
                            <tr>
                              <th>Type</th>                            
                              <th>Content</th>
                            </tr>
                          </thead>
                          <tbody>
                            {{range $i, $x := $.messages}}
                            <tr>
                              <td>{{$x.ItemName}}</td>                            
                              <td>{{$x.ItemValue}}</td>
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