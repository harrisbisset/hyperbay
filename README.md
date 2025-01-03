# Webring Server

## relay.toml

```toml
hash = ""
hostUser = ""
hostEmail = ""

[[sites]]
slug = "exmpl" # text without spaces, used when name may not be appropriate (e.g. routes)
name = "example" # name of site
src = "" # link to site
url = "example.com" # display text for site link
dead = true # whether link is dead, service will modify this field, if the link is found to be dead
```
