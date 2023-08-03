🚧 **Work in Progress** 🚧

# MkdServe (MarkDown Serve)
very simple single page HTTP server used to serve a markdown file as a webpage. 

## Why?
I am interested in setting up a personal web page, I don't need anything fancy at the moment, and would like for it to just be an easily modifiable
Mark Down page. I am not very familiar with modern web programming, so re-inventing the wheel to a very small degree was actually easier. I will 
eventually adopt a more robust personal webpage, but for now, this seems to do the trick. 

## What does it do?
- Listens for an HTTP request, Handles that request by returning a local MD file converted to HTML using `russcross/blackfriday/v2` library.
- Connection Limit is set using the golang Semaphore library `golang.org/x/sync/semaphore`, and is set using the "ConnectionLimit" variable.
- `page` variable is used to target the md file to be served. This defaults to main.md.

## Something Cool 
 Because this service reads the local md page each time it responds to a request, you can actively modify the page and that change will automatically
 go live without the need to restart the service. As far as I know this is pretty normal for web development, but I thought it was a cool semi unintentional feature 
 based on how this is constructed.
 I can live update my site, and have those updates immediately take effect. 

## How to Use this
- `go run . -ip=$IP -port=$PORT -page=$MKPAGETARGET`
- if left blank `ip` defaults to localhost, and `port` defaults to 8080. `page` defaults to "main.md"
- This should just work, as long as nothing else is serving on that port. 
