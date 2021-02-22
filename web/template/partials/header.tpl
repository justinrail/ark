{{define "header"}}
<header class="app-header navbar">
        <button class="navbar-toggler sidebar-toggler d-lg-none mr-auto" type="button" data-toggle="sidebar-show">
          <span class="navbar-toggler-icon"></span>
        </button>
        <a class="navbar-brand" href="#">
            <h2><span>ark</span></h2>
          <!-- <img class="navbar-brand-full" src="img/brand/logo.svg" width="89" height="25" alt="CoreUI Logo"> -->
          <!-- <img class="navbar-brand-minimized" src="img/brand/sygnet.svg" width="30" height="30" alt="CoreUI Logo">  -->
        </a>
        <button class="navbar-toggler sidebar-toggler d-md-down-none" type="button" data-toggle="sidebar-lg-show">
          <span class="navbar-toggler-icon"></span>
        </button>

        <ul class="nav navbar-nav ml-auto">
          <li class="nav-item d-md-down-none">
            <a class="nav-link" href="#">
              <!-- <i class="icon-location-pin"><span ></span></i> -->
              <i class="icon-clock icons d-block "> {{ .systemTime }} </i>
            </a>
          </li>
          <li class="nav-item d-md-down-none">
                <a class="nav-link" href="#">
                </a>
              </li>          
        </ul>
      </header>
{{end}}
