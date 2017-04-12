# musicAPI

## Description:
This API gives access to music data coming from musixmatch and allows you to get word cloud for artists albums and tracks. This data are made available by the microsevice that will first fetch data from musixmatch and then build the top used word for the different resources.

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

Title : Show all tracks.
<br>
URL : /tracks
<br>
Method : GET
<br>
URL Params : none

Title : Show one track by id.
<br>
URL : /tracks/:id
<br>
Method : GET
<br>
URL Params : Optional: expand_artist=[boolean] ; expand_album=[boolean]

Title : Show one track cloud word by album id.
<br>
URL : /tracks/:id/wordcloud
<br>
Method : GET
<br>
URL Params : Required: size=[integer]