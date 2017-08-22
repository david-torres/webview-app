var ws = new WebSocket("ws://" + window.location.host + "/ws");

// init websocket
ws.onmessage = function (e) {
    var message = JSON.parse(e.data);
    console.log("Websocket Message", message.msg);
};

ws.onopen = function (e) {
    console.log("Websocket Connected");
};

ws.onclose = function (e) {
    console.log("Websocket Disconnected");
};

ws.onerror = function (e) {
    console.log("Websocket Error");
};

// dom-ready?
(function(){
    document.getElementsByTagName('h1')[0].innerHTML = "Hello World from JS!";
    setTimeout(function(){
        ws.send(JSON.stringify({'msg': 'ping'}));
    }, 100);
})();