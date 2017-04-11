"use strict";
const express = require("express");
const rpc = require("../lib/rpc");
const log = require("../lib/log");
const entities = require("../entities/");

const app = module.exports = express.Router();
app.get("/", async (req, res, next) => {
	try {
		let resp = await rpc.call("getTracksIds");
		let tracks = [];
		for (var i = 0 ; i < resp.length ; i ++ ) {
			tracks[i] = entities.track.build(resp[i]);
		}
		res.json(tracks);
	} catch (err) {
		next(err);
	}
});

app.get("/:id", async (req, res, next) => {
	let expandArtist = (req.query.expand_artist  === 'true');
	let expandAlbum = (req.query.expand_album  === 'true');
	try {
		let resp = await rpc.call("getTrack", {"id" : parseInt(req.params.id)});
		let track = entities.track.build(resp.id, resp.name);
		
		let album;
		if (expandAlbum) {
			let respAlb = await rpc.call("getAlbum", {"id" : parseInt(resp.album.id)});
			album = entities.album.build(respAlb.id, respAlb.name, respAlb.tracks);
		} else {
			album = entities.album.build(resp.album.id);
		}
		track["album"] = album;

		let artist;
		if (expandArtist) {
			let respArt = await rpc.call("getArtist", {"id" : parseInt(resp.artist.id)});
			artist = entities.artist.build(respArt.id, respArt.name, respArt.albums);
		} else {
			artist = entities.artist.build(resp.artist.id);
		}
		track["artist"] = artist;

		res.json(track);
	} catch (err) {
		next(err)
	}
});

app.get("/:id/wordcloud/", async (req, res, next) => {
	let id = parseInt(req.params.id);
	let size = parseInt(req.query.size);
	try {
		let wordcloudResp = await rpc.call("getTrackCloudWord", {"id" : id, "size" : size});
		let trackResp = await rpc.call("getTrack", {"id" : id});
		let trackCloudWord = entities.track.buildCloud(trackResp.id, trackResp.name, wordcloudResp);
		res.json(trackCloudWord);
	} catch (err) {
		next(err);
	}
});