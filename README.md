# union-tokyo.com

This is all data and codes for [union-tokyo.com](http://union-tokyo.com/)

## Dependencies

- Docker

## How deploy

```bash
$ git clone https://github.com/lowply/union-tokyo.com.git
$ cd union-tokyo.com
$ make deploy
```

## Environments

##### Backend

- [Go](https://golang.org/) for application server
- [Martini](https://github.com/go-martini/martini) for URL routing
- [Gin](https://github.com/codegangsta/gin) for realtime compile

##### Frontend

- HTML5 + CSS3
- [jQuery](https://jquery.com/)
- [Bootstrap](http://getbootstrap.com/) for CSS framework
- [SASS](http://sass-lang.com/) for CSS coding
- [Compass](http://compass-style.org/) for realtime compile
- SVGZ for vector graphics

##### Datastore

- JSON file, no database

##### Infrastructure

- [Sakura VPS](http://vps.sakura.ad.jp/) for hosting
- [Amazon Route53](http://aws.amazon.com/route53/) for DNS
- [Nginx](http://nginx.org/) for reverse proxy

##### Deployment

- [GitHub](https://github.com/lowply/union-tokyo.com/) for version control
- [Docker](https://www.docker.com/) for application portability
