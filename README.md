# musicAPI

## Description:
This api gives access to music data coming from musixmatch. The API allows you to get word cloud for artists albums and tracks. The microsevice will give you the top used word by artist, album or tracks and how often they occur.

## Installation:
Clone the project, git clone git@github.com:quenel/musicAPI.git
<br>
Go into the musicAPI folder, cd musicAPI.
<br>
Start microserive, ./start_microservice.sh
<br>
Start api, ./start_microservice.sh

## API Endpoints:

API root : http://localhost:8081/

Title : Show all artists.
<br>
URL : /artists
<br>
Method : GET
<br>
URL Params : none

Title : Show one artist by id.
<br>
URL : /artists/:id
<br>
Method : GET
<br>
URL Params : none

Title : Show one artist cloud word by artist id.
<br>
URL : /artists/:id/wordcloud
<br>
Method : GET
<br>
URL Params : Required: size=[integer]

Title : Show all albums.
<br>
URL : /albums
<br>
Method : GET
<br>
URL Params : none

Title : Show one album by id.
<br>
URL : /albums/:id
<br>
Method : GET
<br>
URL Params : Optional: expand_artist=[boolean]

Title : Show one album cloud word by album id.
<br>
URL : /albums/:id/wordcloud
<br>
Method : GET
<br>
URL Params : Required: size=[integer]

Title : Show all albums.
<br>
URL : /tracks
<br>
Method : GET
<br>
URL Params : none

Title : Show one album by id.
<br>
URL : /tracks/:id
<br>
Method : GET
<br>
URL Params : Optional: expand_artist=[boolean] ; expand_album=[boolean]

Title : Show one album cloud word by album id.
<br>
URL : /tracks/:id/wordcloud
<br>
Method : GET
<br>
URL Params : Required: size=[integer]