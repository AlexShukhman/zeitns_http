# ZeitToken Name Server HTTP + Golang + Redis

A simple stasher and grabber for ZeitTokens over HTTP

## Quick Start

```sh
redis-server # or whatever you use to start redis
cd /path/to/zeitns_http
go build .
./zeitns_http
```

## Notes

* As of right now, this is a POC project and really just set up to run locally
* This app expects to hit a non-sharded, local Redis DB. No testing has been done with any other implementations or instantiations of Redis

## Contributers

* [Alex Shukhman](https://github.com/alexshukhman) -- Code Owner

## Questions

* See email details for Alex Shukhman [here.](https://email-alex.com)