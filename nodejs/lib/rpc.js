"use strict"
const request = require("request-promise");

exports.call = (rpcName, input) => {
    const opt = {
        "method": "POST",
        "uri": "http://localhost:8080/rpc/" + rpcName,
        "body": input || {},
        "json": true,
    };

    return request(opt);
};