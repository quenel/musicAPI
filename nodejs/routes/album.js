"use strict";
const express = require("express");
const rpc = require("../lib/rpc");
const log = require("../lib/log");
const entities = require("../entities/");

const app = module.exports = express.Router();
app.get("/", async (req, res, next) => {
	try {
		let resp = await rpc.call("getAlbumsIds");
		let albums = [];
		for (var i = 0 ; i < resp.length ; i ++ ) {
			albums[i] = entities.album.build(resp[i]);
		}
		res.json(albums);
	} catch (err) {
		next(err);
	}
});

app.get("/:id", async (req, res, next) => {
	let expandArtist = (req.query.expand_artist === 'true');
	try {
		let resp = await rpc.call("getAlbum", {"id" : parseInt(req.params.id)});
		let album = entities.album.build(resp.id, resp.name, resp.tracks);

		let artist;
		if (expandArtist) {
			let respArt = await rpc.call("getArtist", {"id" : parseInt(resp.artist.id)});
			artist = entities.artist.build(respArt.id, respArt.name, respArt.albums);
		} else {
			artist = entities.artist.build(resp.artist.id);
		}
		album["artist"] = artist;

		res.json(album);
	} catch (err) {
		next(err)
	}
});

app.get("/:id/wordcloud/", async (req, res, next) => {
	let id = parseInt(req.params.id);
	let size = parseInt(req.query.size);
	try {
		let wordcloudResp = await rpc.call("getAlbumCloudWord", {"id" : id, "size" : size});
		let albumResp = await rpc.call("getAlbum", {"id" : id});
		let artistWordCloud = entities.album.buildCloud(albumResp.id, albumResp.name, wordcloudResp);
		res.json(artistWordCloud);
	} catch (err) {	
		next(err);
	}
});