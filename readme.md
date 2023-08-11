🚧 **Work in Progress** 🚧

# MkdServe (MarkDown Serve)
very simple single page HTTP server. Converts MarkDown (.md) files to HTML. Serves jpeg, gifs, and pngs from `/image/` subdirectory, addressable via `/image/` subdomain.

## Features
- supports CSS stylsheets within `/assets/` subdirectory
- serves images out of `/image/` subdirectory
- Will serve main.md within root path of program. Converts it to HTML. This seems to make it mostly compatible with HTML tags. 

## Something Cool 
 Because this service reads the local md page each time it responds to a request, you can actively modify the page and that change will automatically
 go live without the need to restart the service. As far as I know this is pretty normal for web development, but I thought it was a cool semi unintentional feature 
 based on how this is constructed.
 I can live update my site, and have those updates immediately take effect. 

## How to Use this
- `go run . -ip=$IP -port=$PORT -page=$MKPAGETARGET`
- if left blank `ip` defaults to localhost, and `port` defaults to 8080. `page` defaults to "main.md"
- This should just work, as long as nothing else is serving on that port. 
