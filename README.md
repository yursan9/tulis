Tulis
=====

Tulis is a live markdown blog from idiot for everyone

> P.S Building blogging platform is hard.
> In the end I only use this for few weeks.
> I guess this will be testimony how I learn Golang.

## Installation ##

How can I install this?

```shell
go get -u github.com/yursan9/tulis
```

## Run the Program ##

If you already configured your GOPATH and PATH to GOPATH/bin, you can run the
program like this:

```shell
tulis
```

And you can open `localhost:port` in your browser (Port is default to 8080).

## Run the Example ##

If you want to see how the application work, you can `cd` to the directory
and execute `go run main.go`

```shell
cd $GOPATH/src/github.com/yursan9/tulis
TULIS_BASE="example" TULIS_RELATIVE="true" go run main.go 
```

## Configuration ##

Because it's originally made to power up my simple blog, there isn't many thing
that can be configured. You can configure this using environment variables. You
need to add prefix `TULIS` before the name of options.

Example:
```shell
export TULIS_PORT=":1234"
tulis
```

| Option        | Description                                   | Type     |
|---------------|-----------------------------------------------|----------|
| PORT          | Port number                                   | string   |
| BASEDIR       | Base directory for your site                  | string   |
| POSTDIR       | Path to directory of your blog's posts        | string   |
| TEMPLDIR      | Path to directory of your blog's templates    | string   |
| STATICDIR     | Path to directory of your blog's static files | string   |
| RELATIVE      | If true make *DIR relative to BASEDIR         | bool     |
| MAXPOSTS      | Number of max posts in one page               | uint     |

## Thanks to ##
+ julienschmidt's [httprouter](github.com/julienschmidt/httprouter)
+ russross's [blackfriday](github.com/russross/blackfriday)
+ BurntSushi's [toml](https://github.com/BurntSushi/toml)
+ gosimple's [slug](https://github.com/gosimple/slug)
