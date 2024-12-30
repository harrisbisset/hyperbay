# Webring Server

## Routes

Basic:

- GET /list
  - Returns JSON of all available sites
- GET /refresh
  - Requires the request to have a ```hash``` header matching the server
  - Requests the server to update the relay list

Optional:

- GET /random
  - NO RECOMMENDED, if possible implement from the client-side
  - Gets a random site from the site list

## Docker
