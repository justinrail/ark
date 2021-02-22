{{define "content" }}

<main class="main">

  <ol class="breadcrumb">
    <li class="breadcrumb-item">
      <a href="/index">Home</a>
    </li>
    <li class="breadcrumb-item active">NotifyRule</li>

  </ol>
  <div class="container-fluid">
    <div id="ui-view">
      <div>
        <div class="row">
            <div class="col-sm-12 col-md-12">
                <div class="card">
                  <div class="card-header">Notify Rules</div>
                  <div class="card-body">
                    <div class="card">
                      <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                        <thead>
                          <tr>
                            <th>UUID</th>            
                            <th>NotifyRuleName</th>
                            <th>ProvideType</th>
                            <th>TriggerClause</th>
                            <th>CleanerClause</th>
                            <th>MessageQueueID</th>
                          </tr>
                        </thead>
                        <tbody>
                          {{range $i, $x := $.notifyrules}}
                          <tr>
                            <td>{{$x.UUID}}</td>
                            <td><a href="/notifymessage/?messagequeueId={{$x.MessageQueueID}}&&notifyruleName={{$x.NotifyRuleName}}">{{$x.NotifyRuleName}}</a></td>                            
                            <td>{{$x.ProvideType}}</td>
                            <td>{{$x.TriggerClause}}</td>
                            <td>{{$x.CleanerClause}}</td>
                            <td>{{$x.MessageQueueID}}</td>                            
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