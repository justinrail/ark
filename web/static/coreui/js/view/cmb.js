
function run (){
    mermaid.initialize({ theme: 'forest', startOnLoad:true }) 
}


//因为domain.js在jquery前加载，所以juery的加载判断是无法执行的
if (document.readyState !== 'loading') {
    run();
  } else {
    document.addEventListener('DOMContentLoaded', run);
}