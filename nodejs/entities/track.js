"use strict"
const href = require("../lib/href");

exports.build = (id, name) => {
	let track = { "href" : href.uri("tracks/" + id)};
	
	if (name !== undefined) {
		track["name"] = name;
	}

 	return track;
}

exports.buildCloud = (id, name, values) => {
	let wordcloud = {
		"href" : href.uri("track/" + id + "/wordcloud?size=" + values.length),
		"values" : values
	};
	let track = {"href" : href.uri("tracks/" + id), "name" : name};

	return {"track" : track, "cloud" : wordcloud};
}