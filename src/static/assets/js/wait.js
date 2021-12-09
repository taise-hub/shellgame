url = 'ws://'+document.location.host + "/wswait";
conn = new WebSocket(url);

conn.onmessage = function(event){
    redirect();
    return;
}


function redirect(){
    location.href='/join';
}