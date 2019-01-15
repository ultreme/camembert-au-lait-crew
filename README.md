# calcbiz - Camembert au lait crew's website & API
:hamburger: Camembert au lait crew website &amp; API

## Links

* Live: https://www.camembertaulaitcrew.biz
* Install/Dev: `go get ultre.me/calcbiz`

## API Usages (HTTP)

### SoundCloud (`pkg/soundcloud`)

#### Fetch me

```console
❯ http localhost:9000/api/soundcloud/me
HTTP/1.1 200 OK
Content-Length: 690
Content-Type: application/json
Date: Tue, 15 Jan 2019 10:52:22 GMT
Grpc-Metadata-Content-Type: application/grpc
Keep-Alive: timeout=38

{
    "avatar_url": "https://i1.sndcdn.com/avatars-000524030097-1slyvk-large.jpg",
    "city": "Partout",
    "country": "France",
    "description": "Les chips\n\nnous@camembertaulaitcrew.biz",
    "followers_count": "498",
    "followings_count": "187",
    "full_name": "camembert au lait crew",
    "id": "96137699",
    "permalink": "camembert-au-lait-crew",
    "permalink_url": "http://soundcloud.com/camembert-au-lait-crew",
    "playlist_count": "28",
    "public_favorites_count": "259",
    "track_count": "583",
    "uri": "https://api.soundcloud.com/users/96137699",
    "username": "Camembert au lait crew",
    "website": "http://www.camembertaulaitcrew.biz/",
    "website_title": "camembertaulaitcrew.biz"
}
```

#### Fetch all playlists/albums/releases

```console
$ http localhost:9000/api/soundcloud/playlists | head -n 5
{
  "playlists": [
    {
      "id": "683036052",
      "created_at": "2019/01/10 23:05:09 +0000",
```

#### Fetch a specific playlist/album/release

```console
$ http localhost:9000/api/soundcloud/playlists/678762324 | head -n 5
{
  "id": "678762324",
  "created_at": "2019/01/05 00:31:46 +0000",
  "title": "OGR001",
  "sharing": "public",
```

#### Fetch a random playlist/album/release

```console
$ http localhost:9000/api/soundcloud/playlists/0 | head -n 5
{
  "id": "36518144",
  "created_at": "2014/05/25 12:08:15 +0000",
  "title": "Camembert au lait crew - Des Trucs",
  "sharing": "public",
```

#### Fetch all tracks

```console
$ http localhost:9000/api/soundcloud/tracks | head -n 5
{
  "tracks": [
    {
      "id": "559269942",
      "CreatedAt": "2019/01/15 06:02:10 +0000",
```

#### Fetch a specific track

```console
$ http localhost:9000/api/soundcloud/tracks/475706220 | head -n 5
{
  "id": "475706220",
  "CreatedAt": "2018/07/23 19:10:11 +0000",
  "Title": "Bannalec",
  "Sharing": "public",
```

#### Fetch a random track

```console
$ http localhost:9000/api/soundcloud/tracks/0 | head -n 5
{
  "id": "559269885",
  "CreatedAt": "2019/01/15 06:01:58 +0000",
  "Title": "Grand écart (feat. Rouge Gorge)",
  "Sharing": "public",
```

### Call Numberinfo (`pkg/numberinfo`)

```console
$ http localhost:9000/api/numberinfo/1337
HTTP/1.1 200 OK
Content-Length: 100
Content-Type: application/json
Date: Tue, 15 Jan 2019 10:46:55 GMT
Grpc-Metadata-Content-Type: application/grpc
Keep-Alive: timeout=38

{
    "facts": {
        "is-prime": "false",
        "number": "1337",
        "sqrt": "36.565010597564445"
    }
}
```

### Call Ping

```console
$ http :9000/api/ping
HTTP/1.1 200 OK
Content-Length: 20
Content-Type: application/json
Date: Tue, 15 Jan 2019 10:50:10 GMT
Grpc-Metadata-Content-Type: application/grpc
Keep-Alive: timeout=38

{
    "pong": "pong"
}
```

### Call Crew

```console
$ http :9000/api/crew | head -n 5
{
  "name": "Camembert au lait crew",
  "website": "http://www.camembertaulaitcrew.biz",
  "members": [
    {
```

### Call WOTD

```console
$ http localhost:9000/api/wotd
HTTP/1.1 200 OK
Content-Length: 22
Content-Type: application/json
Date: Tue, 15 Jan 2019 10:51:05 GMT
Grpc-Metadata-Content-Type: application/grpc
Keep-Alive: timeout=38

{
    "word": "cancun"
}
```

### Call Alternate Logo

```console
$ http localhost:9000/api/alternate-logo
HTTP/1.1 200 OK
Content-Length: 88
Content-Type: application/json
Date: Tue, 15 Jan 2019 10:51:30 GMT
Grpc-Metadata-Content-Type: application/grpc
Keep-Alive: timeout=38

{
    "path": "https://camembertaulaitcrew.github.io/assets/logo-alternate-300/trim.jpg"
}
```
