var bunyan = require("bunyan");

var options = {
    "name": "music_api",
    "streams": [
        {
            "level": "info",
            "stream": process.stdout
        }
    ],
    "serializers": bunyan.stdSerializers
};

var logger = bunyan.createLogger(options);

module.exports = logger;
