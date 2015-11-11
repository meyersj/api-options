var express = require('express');
var url = require('url'); 
var app = express();

function build_response(success, msg) {
    return {Success:success, Msg:msg};
}

function authenticate(user_id, password) {
	var response;
    if (user_id == null || password == null) {
		response = build_response(false, "Missing id and/or pass parameters");
	}
    else if (user_id == "admin" && password == "admin") {
		response = build_response(true, "User authenticated");
	} else {
		response = build_response(false, "Try user:admin password:admin");
	}
    return response;
}

app.get('/', function(req, res) {
    res.type('text/plain');
    res.send('Index');
})

app.get('/user', function(req, res) {
    // grab query parameters
    var url_parts = url.parse(req.url, true);
    var user_id = url_parts.query.id;
    var password = url_parts.query.pass;
    
    // authenticate user and return json response
    res.type('application/json');
    res.send(authenticate(user_id, password));
})

// start server
app.listen(5002);
console.log("Starting server at 127.0.0.1:5002")
