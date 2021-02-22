{{define "side"}}
<div class="sidebar">
  <nav class="sidebar-nav">
    <ul class="nav">
      <li class="nav-item">
        <a class="nav-link" href="index.html">
          <i class="nav-icon cui-dashboard"></i> Home
        </a>
      </li>
      <li class="nav-title">Hub</li>
      <li class="nav-item  nav-dropdown">
        <a class="nav-link nav-dropdown-toggle" href="#">
          <i class="nav-icon icon-grid"></i> Collector</a>
        <ul class="nav-dropdown-items">
          <li class="nav-item">
            <a class="nav-link" href="\stub">
              <i class="nav-icon icon-cursor"></i> Stub</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="\stub">
              <i class="nav-icon icon-cursor"></i> S2/3</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="\cmb">
              <i class="nav-icon icon-cursor"></i> CMB</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="\stub">
              <i class="nav-icon icon-cursor"></i> SNMP</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="\stub">
              <i class="nav-icon icon-cursor"></i> C2C</a>
          </li>
        </ul>
      </li>
      <li class="nav-item nav-dropdown">
        <a class="nav-link nav-dropdown-toggle" href="#">
          <i class="nav-icon icon-grid"></i> Domain</a>
        <ul class="nav-dropdown-items">
          <li class="nav-item">
            <a class="nav-link" href="\domain">
              <i class="nav-icon icon-cursor"></i> Gateways</a>
          </li>
          <li class="nav-item ">
            <a class="nav-link" href="\liveevent">
              <i class="nav-icon icon-cursor"></i> LiveEvents</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="\hisliveevent">
              <i class="nav-icon icon-cursor"></i> HisLiveEvents</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="\hispoint">
              <i class="nav-icon icon-cursor"></i> HisPoints</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="\complexindex">
              <i class="nav-icon icon-cursor"></i> ComplexIndexs</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="\notifyrule">
              <i class="nav-icon icon-cursor"></i> NotifyRule</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" href="buttons/brand-buttons.html">
              <i class="nav-icon icon-cursor"></i> HisCommands</a>
          </li>
        </ul>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="\topology">
          <i class="nav-icon icon-graph"></i> Topology</a>
      </li>
      <li class="nav-title">Self</li>
      <li class="nav-item">
        <a class="nav-link" href="\application">
          <i class="nav-icon cui-laptop"></i>Application</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="\config">
          <i class="nav-icon icon-wrench"></i>Config</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="\log">
          <i class="nav-icon icon-speech"></i>Log</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="\runtime">
          <i class="nav-icon icon-screen-desktop"></i>Runtime</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="\machine">
          <i class="nav-icon cui-monitor"></i>Machine</a>
      </li>
      <li class="divider"></li>
      <li class="nav-title">Extras</li>
      <li class="nav-item">
        <a class="nav-link" href="colors.html">
          <i class="nav-icon icon-info"></i>License Server</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="colors.html">
          <i class="nav-icon icon-info"></i>Time Server</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="colors.html">
          <i class="nav-icon icon-info"></i>Push Server</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="colors.html">
          <i class="nav-icon icon-info"></i>Stream Server</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="colors.html">
          <i class="nav-icon icon-info"></i>Email Server</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="colors.html">
          <i class="nav-icon icon-info"></i>Access Server</a>
      </li>
      <li class="nav-item">
        <a class="nav-link" href="colors.html">
          <i class="nav-icon icon-info"></i>OAuth Server</a>
      </li>
      <li class="nav-item">
          <a class="nav-link" href="colors.html">
            <i class="nav-icon icon-info"></i>Graph Server</a>
        </li>      
      <li class="nav-item">
          <a class="nav-link" href="colors.html">
            <i class="nav-icon icon-info"></i>About</a>
        </li>      
    </ul>
  </nav>
  <button class="sidebar-minimizer brand-minimizer" type="button"></button>
</div>

{{end}}