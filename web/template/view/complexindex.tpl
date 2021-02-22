{{define "content" }}

<main class="main">

  <ol class="breadcrumb">
    <li class="breadcrumb-item">
      <a href="/index">Home</a>
    </li>
    <li class="breadcrumb-item active">ComplexIndex</li>

  </ol>
  <div class="container-fluid">
    <div id="ui-view">
      <div>
        <div class="row">
            <div class="col-sm-12 col-md-12">
                <div class="card">
                  <div class="card-header">Complex Indexes</div>
                  <div class="card-body">
                    <div class="card">
                      <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                        <thead>
                          <tr>
                            <th>ID</th>
                            <th>Name</th>                            
                            <th>ObjectTypeID</th>
                            <th>BusinessID</th>
                            <th>Label</th>
                            <th>GRID</th>
                            <th>CalcCron</th>
                            <th>CalcType</th>
                            <th>AfterCalc</th>
                            <th>SaveCron</th>
                            <th>Expression</th>
                            <th>LastValue</th>
                            <th>LastTimestamp</th>
                            <th>CurrentValue</th>
                            <th>Timestamp</th>                            
                            <th>IsValid</th>
                          </tr>
                        </thead>
                        <tbody>
                          {{range $i, $x := $.complexindexs}}
                          <tr>
                            <td>{{$x.ComplexIndexID}}</td>
                            <td><a href="/hiscomplexindex/?complexindexId={{$x.ComplexIndexID}}">{{$x.ComplexIndexName}}</a></td>                            
                            <td>{{$x.ObjectTypeID}}</td>
                            <td>{{$x.BusinessID}}</td>
                            <td>{{$x.Label}}</td>
                            <td>{{$x.GlobalResourceID}}</td>
                            <td>{{$x.CalcCron}}</td>
                            <td>{{$x.CalcType}}</td>
                            <td>{{$x.AfterCalc}}</td>
                            <td>{{$x.SaveCron}}</td>
                            <td>{{$x.Expression}}</td>
                            <td>{{$x.LastValueString}}</td>
                            <td>{{$x.LastTimestampString}}</td>
                            <td>{{$x.CurrentValueString}}</td>
                            <td>{{$x.TimestampString}}</td>
                            <td>{{$x.IsValid}}</td>
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