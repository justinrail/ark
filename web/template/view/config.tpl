{{define "content"}}

<main class="main">

    <ol class="breadcrumb">
        <li class="breadcrumb-item">
            <a href="/index">Home</a>
        </li>
        <li class="breadcrumb-item active">Config</li>

    </ol>
    <div class="container-fluid">
        <div id="ui-view">
            <div>
                <div class="row">
                    <div class="col-sm-4 col-md-4">
                        <div class="card">
                            <div class="card-header">App</div>
                            <div class="card-body">
                                <div class="card">
                                    <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                                        <thead>
                                            <tr>
                                                <th>Item</th>
                                                <th>Description</th>
                                                <th>Value</th>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            {{range $i, $x := $.cfgs}} 
                                            {{if eq $x.Category "app"}}
                                            <tr>
                                                <td>{{$x.ItemName}}</td>
                                                <td>{{$x.Remark}}</td>
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
                    <div class="col-sm-4 col-md-4">
                        <div class="card">
                            <div class="card-header">Hub</div>
                            <div class="card-body">
                                <div class="card">
                                    <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                                        <thead>
                                            <tr>
                                                <th>Item</th>
                                                <th>Description</th>
                                                <th>Value</th>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            {{range $i, $x := $.cfgs}} 
                                            {{if eq $x.Category "hub"}}
                                            <tr>
                                                <td>{{$x.ItemName}}</td>
                                                <td>{{$x.Remark}}</td>
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
                    <div class="col-sm-4 col-md-4">
                        <div class="card">
                            <div class="card-header">Phoenix</div>
                            <div class="card-body">
                                <div class="card">
                                    <table class="table table-responsive-sm table-sm table-borderless table-hover table-striped">
                                        <thead>
                                            <tr>
                                                <th>Item</th>
                                                <th>Description</th>
                                                <th>Value</th>
                                            </tr>
                                        </thead>
                                        <tbody>
                                            {{range $i, $x := $.cfgs}} 
                                            {{if eq $x.Category "phoenix"}}
                                            <tr>
                                                <td>{{$x.ItemName}}</td>
                                                <td>{{$x.Remark}}</td>
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
</main>

{{end}}