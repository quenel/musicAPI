# musicAPI
API root : http://localhost:8081/

API Endpoints:

Title : Show all artists.
<br>
URL : /artists
Method : GET
URL Params : none

Title : Show one artist by id.
URL : /artists/:id
Method : GET
URL Params : none

Title : Show one artist cloud word by artist id.
URL : /artists/:id/wordcloud
Method : GET
URL Params : Required: size=[integer]

Title : Show all albums.
URL : /albums
Method : GET
URL Params : none

Title : Show one album by id.
URL : /albums/:id
Method : GET
URL Params : Optional: expand_artist=[boolean]

Title : Show one album cloud word by album id.
URL : /albums/:id/wordcloud
Method : GET
URL Params : Required: size=[integer]

Title : Show all albums.
URL : /tracks
Method : GET
URL Params : none

Title : Show one album by id.
URL : /tracks/:id
Method : GET
URL Params : Optional: expand_artist=[boolean] ; expand_album=[boolean]

Title : Show one album cloud word by album id.
URL : /tracks/:id/wordcloud
Method : GET
URL Params : Required: size=[integer]