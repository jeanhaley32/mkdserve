🚧 **Work in Progress** 🚧
# Objective
- Create my own webserver to host my own [personal website](http://haley.nyc)

# Features
- Image and asset sub-domain `/assets/` (for JS, and CSS assets)
- Serves a single main page. `/main.html`
- Very simple. uses pre-built golang functionality
- Serves anything located in `/pages/` subfolder as a subdomain
- TLS enabled with flags `-key` and `-crt`

# TODO
- Implement light traffic management (limit amount of requests per X time frame)
