"use strict";
const express = require("express");
const rpc = require("../lib/rpc");
const log = require("../lib/log");
const entities = require("../entities/");

const app = module.exports = express.Router();
app.get("/", async (req, res, next) => {
	try {
		let resp = await rpc.call("getArtistsIds");
		let artists = [];
		for (var i = 0 ; i < resp.length ; i ++ ) {
			artists[i] = entities.artist.build(resp[i]);
		}
		res.json(artists);
	} catch (err) {
		next(err);
	}
});

app.get("/:id", async (req, res, next) => {
	try {
		let resp = await rpc.call("getArtist", {"id" : parseInt(req.params.id)});
		let artist = entities.artist.build(resp.id, resp.name, resp.albums);
		res.json(artist);
	} catch (err) {
		next(err)
	}
});

app.get("/:id/wordcloud/", async (req, res, next) => {
	let id = parseInt(req.params.id);
	let size = parseInt(req.query.size);
	try {
		let wordcloudResp = await rpc.call("getArtistCloudWord", {"id" : id, "size" : size});
		let artistResp = await rpc.call("getArtist", {"id" : id});
		let artistWordCloud = entities.artist.buildCloud(artistResp.id, artistResp.name, wordcloudResp);
		res.json(artistWordCloud);
	} catch (err) {	
		next(err);
	}
});