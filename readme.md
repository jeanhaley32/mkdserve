🚧 **Work in Progress** 🚧

# MkdServe (MarkDown Serve)
Very Simple server. Not yet feature complete. 

## Features
- Handles a single page `main.html`
- css is stored in `/assets/`, and can be linked from there.
- images can be stores and served from `/image/`, going to <URL>/image/ will lead to a generated index page that shows links to all images served.
  - The server will scrap all gif/jpeg/png files in `/image/` and create a simple index page of links to those images.
  - You can also link to them directly if you know the image name you're looking for

## How to Use this
- `go run . -ip=$IP -port=$PORT -page=$MKPAGETARGET`
- if left blank `ip` defaults to localhost, and `port` defaults to 8080. `page` defaults to "main.md"
- This should just work, as long as nothing else is serving on that port.

## Future Additions
- Enable Javascript
- rewrite to be simpler. There are a few preconstructed standard Library items that do what I did, but alot better.
- Properly implement traffic limiting, right now it's not really working as it should. Lazily implemented.
- Support multiple subdomains, not just one main.html.
That's it so far, until I can think of anything else. 
