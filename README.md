Tulis
=====

Tulis is a live markdown blog from idiot for everyone

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

## Configuration ##

Because it's originally made to power up my simple blog, there isn't many thing
that can be configured. You can configure this using environment variables. You
need to add prefix `TULIS` before the name of options.

Example:
```shell
export TULIS_PORT=":1234"
tulis
```

|Option         | Description   					|
|---------------|-------------------------------------------------------|
| PORT          | Port number 						|
| SITENAME      | Website name in browser          			|
| BASEDIR       | Base directory for your site 				|
| POSTDIR       | Path to directory of your blog's posts 		|
| TEMPLDIR      | Path to directory of your blog's templates 		|
| STATICDIR     | Path to directory of your blog's static files		|
| RELATIVE      | If true make *DIR relative to BASEDIR 		|
| MAXPOSTS	| Number of max posts in one page 			|

## Thanks to ##
+ julienschmidt's [httprouter](github.com/julienschmidt/httprouter)
+ russross's [blackfriday](github.com/russross/blackfriday)
