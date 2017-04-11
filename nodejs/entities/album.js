"use strict"
const href = require("../lib/href");

exports.build = (id, name, tracksIds) => {
	let album = { "href" : href.uri("albums/" + id)};
	
	if (name !== undefined) {
		album["name"] = name;
	}

 	if (tracksIds !== undefined) {
 		let tracks = [];
 		for ( var i = 0 ; i < tracksIds.length ; i ++ ) {
 			tracks[i] = {"href" : href.uri("tracks/" + tracksIds[i])};
 		}
 		album["tracks"] = tracks;
 	}
 	return album;
}

exports.buildCloud = (id, name, values) => {
	let wordcloud = {
		"href" : href.uri("album/" + id + "/wordcloud?size=" + values.length),
		"values" : values
	};
	let album = {"href" : href.uri("albums/" + id), "name" : name};

	return {"album" : album, "cloud" : wordcloud};
}