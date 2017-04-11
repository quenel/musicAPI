"use strict";
const http = require("http");
const express = require("express");
const log = require("./lib/log");

const app = express();

app.use("/artists", require("./routes/artist"));
app.use("/albums", require("./routes/album"));
app.use("/tracks", require("./routes/track"));

app.use((req, res) => {
	log.warn({"path": req.path}, "404 not found");
	res.sendStatus(404);
});

app.use((err, req, res, next) => {
    if (err instanceof Error) {
        if (err.code !== 0) {
            log.error({err});
        } else {
            log.error({err}, "MISSING ERROR CODE");
        }
        res.status(500).json({"error": err.message, "code": err.code});
    } else {
        log.error({err}, "MISSING ERROR CODE");
        res.status(500).json({"error": err.message, "code": 0});
    }
});

http.createServer(app).listen(8081);	
log.info({"port": 80}, "http server started");