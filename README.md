# union-tokyo.com

This is all data and codes for union-tokyo.com

## Dependencies

- Docker

## How to run

```bash
$ git clone https://github.com/lowply/union-tokyo.com.git
$ cd union-tokyo.com
$ docker-compose up
```

## Environments

##### Backend

- Go for application server
- Martini for WAF
- Gin for realtime compile

##### Frontend

- HTML5
- jQuery
- Bootstrap for CSS Framework
- SASS
- Compass for realtime compile
- SVGZ for vector graphics

##### Datastore

- JSON file, no database

##### Infrastructure

- Sakura VPS
- Amazon Route53 for DNS
- Nginx for reverse proxy

##### Deployment

- GitHub for version control
- Docker for application portability
