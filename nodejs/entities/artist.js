"use strict"
const href = require("../lib/href");

exports.build = (id, name, albumsIds) => {
	let artist = { "href" : href.uri("artists/" + id)};
	
	if (name !== undefined) {
		artist["name"] = name;
	}

 	if (albumsIds !== undefined) {
 		let albums = [];
 		for ( var i = 0 ; i < albumsIds.length ; i ++ ) {
 			albums[i] = {"href" : href.uri("albums/" + albumsIds[i])};
 		}
 		artist["albums"] = albums;
 	}
 	return artist;
}

exports.buildCloud = (id, name, values) => {
	let wordcloud = {
		"href" : href.uri("artist/" + id + "/wordcloud?size=" + values.length),
		"values" : values
	};
	let artist = {"href" : href.uri("artists/" + id), "name" : name};

	return {"artist" : artist, "cloud" : wordcloud};
}