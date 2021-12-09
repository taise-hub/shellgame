url = 'ws://'+document.location.host + "/wstimeattack";
conn = new WebSocket(url);

conn.onmessage = function(event){
    console.log(event.data);
    obj = JSON.parse(event.data);
    switch (obj.DataType) {
        case 'cmd':
            var si = document.getElementById("standard-input");
            var messages = b64DecodeUnicode(obj.StdOut).split('\n');
            for (var i=messages.length-1; i> -1; i--) {
                var item = document.createElement("span");
                item.setAttribute("class", "result");
                item.innerText = messages[i];
                si.after(item);
            }
            changeTarget();//4
            changeRead();
            newLine();//6
            break;
        case 'score':
            console.log("DEBUG: Score: " + obj.Score);
            console.log("DEBUG: judgement: " + obj.Judgement)
            var score = document.getElementById("myScore");
            console.log(score);
            score.value = obj.Score;

            // 終了判定のところ。
            if (obj.Judgement) {
                if (obj.Owner) {
                    alert("win!!")
                    return
                }else {
                    alert("loose...")
                    return
                }
            }
    }
}

window.addEventListener("load", function(){
    si = document.getElementById("standard-input");
    si.focus();
});

window.addEventListener("keyup", function(e){
    let key = e.code;
    if(key=="Enter") {
        cmd = document.getElementById("standard-input").value; //1
        if(cmd === "clear") {
            clear();
            return;
        }
        if(!conn) {
            return false;
        }
        if(!cmd) {
            return false;
        }
        
        var msg = {
            DataType: "cmd",
            Cmd: cmd,
        };

        conn.send(JSON.stringify(msg));
		console.log("[DEBUG] " + cmd);
		return;
    }
})

function newLine() {
    var list = document.getElementsByClassName("terminal")[0];
    var li = document.createElement("li");
    var span = document.createElement("span");
    span.setAttribute("id", "console");
    var console = document.createTextNode("$");
    span.appendChild(console);
    var input = document.createElement("input");
    input.setAttribute("type", "text");
    input.setAttribute("id", "standard-input");
    input.setAttribute("spellcheck", "false");
    li.appendChild(span);
    li.appendChild(input);
    list.appendChild(li);
    input.focus();
}

function changeTarget() {
    var console = document.getElementById("console");
    console.removeAttribute("id");
    console.setAttribute("class", "console");

    var si = document.getElementById("standard-input");
    si.removeAttribute("id");
    si.setAttribute("class","standard-input");
}

function changeRead() {
    var standardInputs = document.getElementsByClassName("standard-input");
    for (var i = 0; i < standardInputs.length; i++) {
        var si = standardInputs[i];
        si.readOnly = true;
    }
}

function clear() {
    var terminal = document.getElementsByClassName("terminal")[0];
    while( terminal.firstChild ){
        terminal.removeChild( terminal.firstChild );
      }
    newLine();
}

function requestAnswer() {
    var qname = document.getElementById("qname").value;
    var answer = document.getElementById("answer").value;
    var msg = {
        DataType: "score",
        AnswerReq: {
            Name: qname,
            Answer: answer,
        },
        Cmd: null,
    };

    conn.send(JSON.stringify(msg));
    console.log("[DEBUG] " + JSON.stringify(msg));
    return;

}

function b64DecodeUnicode(str) {
    var messages = atob(str);
    const decoded_array = new Uint8Array(Array.prototype.map.call(messages, c => c.charCodeAt()));
    const decode = new TextDecoder().decode(decoded_array);
    return decode
}