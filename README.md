# union-tokyo.com

This is all data and codes for [union-tokyo.com](http://union-tokyo.com/)

## Prerequirements

- [golang/dep](https://github.com/golang/dep)
- Local port 3000

## How deploy

```bash
$ git clone https://github.com/lowply/union-tokyo.com.git
$ cd union-tokyo.com
$ make deps
$ DEPDIR=/path/to/dir make deploy
```

Then reverse proxy to `localhost:3000` and set static file path to `/public`

## Environments

##### Backend

- [Go](https://golang.org/) for application server
- [Martini](https://github.com/go-martini/martini) for URL routing

##### Frontend

- HTML5 + CSS3
- [jQuery](https://jquery.com/)
- [Bootstrap](http://getbootstrap.com/) for CSS framework
- [SASS](http://sass-lang.com/) for CSS coding
- [Compass](http://compass-style.org/) for realtime compile
- SVGZ for vector graphics

##### Datastore

- JSON files, no database
