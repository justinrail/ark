{{define "content" }}

<main class="main">

  <ol class="breadcrumb">
    <li class="breadcrumb-item">
      <a href="/index">Home</a>
    </li>
    <li class="breadcrumb-item"><a href="/complexindex">ComplexIndex {{$.ComplexIndexID}}</a></li>
    <li class="breadcrumb-item active">History Complex Index</li>

  </ol>
  <div class="container-fluid">
    <div id="ui-view">
      <div>
        <div class="row">
            <div class="col-sm-12 col-md-12">
                <div class="card">
                  <div class="card-header">History of ComplexIndex  {{$.ComplexIndexID}}</div>
                  <div class="card-body">
                    <div class="card">
                      <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                        <thead>
                            <tr>
                              <th>Name</th>                            
                              <th>Category</th>
                              <th>Title</th>
                              <th>Label</th>
                              <th>GlobalResourceID</th>
                              <th>CalcType</th>
                              <th>LastValue</th>
                              <th>LastTimestamp</th>
                              <th>CurrentValue</th>
                              <th>Timestamp</th>                            
                            </tr>
                          </thead>
                          <tbody>
                            {{range $i, $x := $.hiscomplexindexs}}
                            <tr>
                              <td>{{$x.ComplexIndexName}}</td>                            
                              <td>{{$x.Category}}</td>
                              <td>{{$x.Title}}</td>
                              <td>{{$x.Label}}</td>
                              <td>{{$x.GlobalResourceID}}</td>
                              <td>{{$x.CalcType}}</td>
                              <td>{{$x.LastValueString}}</td>
                              <td>{{$x.LastTimestampString}}</td>
                              <td>{{$x.CurrentValueString}}</td>
                              <td>{{$x.TimestampString}}</td>
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