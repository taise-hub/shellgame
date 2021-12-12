var url = 'ws://'+document.location.host + "/battle/ws";
var conn = new WebSocket(url);

// 自分のディレクトリ情報
var leftCurrentDir = "/home/sechacker";
var leftPreviousDir = "";
var leftDefaultDir = "/home/sechacker";
//　対戦相手のディレクトリ情報
var rightCurrentDir = "/home/sechacker";
var rightPreviousDir = "";
var rightDefaultDir = "/home/sechacker";
var commandHistory = [];
var currentCommand = 0;

conn.onmessage = function(event){
    var response = JSON.parse(event.data);
    switch (response.Type) {
        case 'question':
            // 問題の表示
            setQuestion(response.QuestionNumebers)
            break;
        case 'tick':
            if (response.Tick === 300 ) {
                alert("")
                // TODO: 勝者の判定
            }
            elem = document.getElementById("elapsed-seconds");
            elem.innerText = response.Tick;
            return
        case 'command':
            // 最初に自分のコマンドか相手のコマンドか判定して処理を変える。
            if (response.Personally) {
                switchConsoleLeft(response);
            }
            else {
                switchConsoleRight(response);
            }
            break;
        case 'answer':
            if (response.Personally) {
                var score = document.getElementById("myScore"); 
                if (response.Correct) {
                    score.value += 100/3
                }
            }
            else {
                var score = document.getElementById("opScore");
                if (response.Correct) {
                    score.value += 100/3
                }
            }            
            if (response.Complete) {
                (response.Owner? alert("勝ちっちの〜ち！"): alert("ん〜〜〜負け！！！"));
            }
    }
}

window.addEventListener("load", function(){
    newLine("left");
    newLine("right");
    var si = document.getElementById("standard-input-left");
    si.focus();
});

// コマンドの送信
window.addEventListener("keyup", function(e){
    if(e.code=="Enter") {
        rawCommand = document.getElementById("standard-input-left").value;
        var command;
        if (rawCommand.endsWith(";")) {
            command = "cd " + leftCurrentDir + "; " + rawCommand + " pwd";
        }else {
            command = "cd " + leftCurrentDir + "; " + rawCommand + "; pwd";
        }
        if(rawCommand === "clear") {
            clearLeft();
            return;
        }
        if(!conn) {
            return false;
        }
        var msg = {
            Type: "command",
            Command: command,
        };
        conn.send(JSON.stringify(msg));
        commandHistory.push(rawCommand);
        currentCommand = commandHistory.length;
		return;
    }
})

// ひとつ前のコマンドを取得
function previousCommand() {
    if (currentCommand != 0) {
        switchCommand(currentCommand - 1)
    }
}

// ひとつ先のコマンドを取得
function nextCommand() {
    if (currentCommand != commandHistory.length) {
        switchCommand(currentCommand + 1)
    }
}

// 入力するコマンドを変更
function switchCommand(newCommand) {
    currentCommand = newCommand;
    if (currentCommand == commandHistory.length) {
        document.getElementById("standard-input-left").value = "";
    }else {
        document.getElementById("standard-input-left").value = commandHistory[currentCommand];
        setTimeout(function(){ document.getElementById("standard-input-left").selectionStart = document.getElementById("standard-input-left").selectionEnd = 10000; }, 0);
    }
}

//自分のコマンド送信時の処理
function switchConsoleLeft(commandResponse) {
    var si = document.getElementById("standard-input-left");
    var parsedResponse = b64DecodeUnicode(commandResponse.CommandResult.StdOut).split('\n');
    if (parsedResponse.length > 1) {
        parsedResponse.pop();
        leftCurrentDir = parsedResponse.pop();
    }
    for (var i=parsedResponse.length-1; i> -1; i--) {
        var item = document.createElement("span");
        item.setAttribute("class", "result");
        item.innerText = parsedResponse[i];
        si.after(item);
    }
    changeTarget("left");
    changeRead("left");
    newLine("left");
}

// 右画面の制御
function switchConsoleRight(commandResponse) {
    var si = document.getElementById("standard-input-right");
    var parsedResponse = b64DecodeUnicode(commandResponse.CommandResult.StdOut).split('\n');
    if (parsedResponse.length > 1) {
        parsedResponse.pop();
        rightCurrentDir = parsedResponse.pop();
    }
    for (var i=parsedResponse.length-1; i> -1; i--) {
        var item = document.createElement("span");
        item.setAttribute("class", "result");
        item.innerText = parsedResponse[i];
        si.after(item);
    }
    changeTarget("right");
    newLine("right");
}

// 新しいコンソールを追加
function newLine(mode) {
    switch(mode) {
        case 'left':
            var list = document.getElementsByClassName("terminal")[0];
            currentDir = leftCurrentDir;
            defaultDir = leftDefaultDir;
            break
        case 'right':
            var list = document.getElementsByClassName("terminal")[1];
            currentDir = rightCurrentDir;
            defaultDir = rightDefaultDir;
            break
    }
    var li = document.createElement("li");
    li.setAttribute("id", "console-line");
    var markField = document.createElement("span");
    markField.setAttribute("id", "console-mark-" + mode);
    markField.appendChild(document.createTextNode("$"));
    var currentDirField = document.createElement("span");
    currentDirField.setAttribute("id", "console-current-dir-" + mode);
    if (currentDir.indexOf(defaultDir) === 0) {
        currentDirField.appendChild(document.createTextNode("[" + currentDir.replace(defaultDir, "~") + "]"));
    }else {
        currentDirField.appendChild(document.createTextNode("[" + currentDir + "]"));
    }
    var input = document.createElement("input");
    input.setAttribute("type", "text");
    input.setAttribute("id", "standard-input-" + mode);
    input.setAttribute("spellcheck", "false");
    li.appendChild(currentDirField);
    li.appendChild(markField);
    li.appendChild(input);
    list.appendChild(li);
    if (mode=='right') {
        input.readOnly = true;
        return
    }
    updateCurrentConsoleWidth(mode)
    input.focus();
}

// カーソル位置の変更
function changeTarget(mode) {
    var consoleCurrentDir = document.getElementById("console-current-dir-" + mode);
    consoleCurrentDir.removeAttribute("id");
    consoleCurrentDir.setAttribute("class", "console-current-dir-" + mode);

    var si = document.getElementById("standard-input-" + mode);
    si.removeAttribute("id");
    si.setAttribute("class","standard-input-" +  mode);
}

function changeRead(mode) {
    var standardInputs = document.getElementsByClassName("standard-input-" + mode);
    for (var i = 0; i < standardInputs.length; i++) {
        var si = standardInputs[i];
        si.readOnly = true;
    }
}

// clearコマンド実行時の処理
function clearLeft() {
    var terminal = document.getElementsByClassName("terminal")[0];
    while( terminal.firstChild ){
        terminal.removeChild( terminal.firstChild );
      }
    newLine('left');
}

//　回答の送信
function requestAnswer() {
    var qname = document.getElementById("qname").value;
    var answer = document.getElementById("answer").value;
    var msg = {
        Type: "answer",
        AnswerName: qname,
        Answer: answer,
    };
    conn.send(JSON.stringify(msg));
    return;

}

function setQuestion(questionNumebers) {
    var questionlist = document.getElementById("question-list");
    for (const elem of questionNumebers) {
        q = document.createElement("li");
        q.innerText = "/home/sechacker/" + elem;
        questionlist.appendChild(q);
    }

}

function b64DecodeUnicode(str) {
    var messages = atob(str);
    const decoded_array = new Uint8Array(Array.prototype.map.call(messages, c => c.charCodeAt()));
    const decode = new TextDecoder().decode(decoded_array);
    return decode
}

function updateCurrentConsoleWidth(mode) {
    var consoleCurrentDirWidth = document.getElementById("console-current-dir-" + mode).scrollWidth + 8;
    var consoleWidth = document.getElementById("console-mark-" + mode).scrollWidth + 8;
    var container = document.getElementById('standard-input-' + mode);
    var size = consoleWidth + consoleCurrentDirWidth;
    container.setAttribute("style", "width: calc(100% - " + size + "px); display: block; border: none; resize: none; background-color: #000000; color: white; font-size: 12pt; line-height: 1.4em; display: inline-block;")
}