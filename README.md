aboutme
=======

A page explaining what I've been doing on stack, github, facebook, and twitter.

Currently running at [me.bclymer.com/](http://me.bclymer.com/).

You can make this all about you by switching the ids in

- facebook.go
- twitter.go
- stack.go
- github.go

You'll also need to set up twitter and facebook apps and include `twitterAuth.json` and `facebookAuth.json` in the root directory.

They should look like

`{"consumerKey": "xxxxxxxxxxxxxxxxxxxx", "consumerSecret": "xxxxxxxxxxxxxxxxxxxx", "accessToken": "111111111-xxxxxxxxxxxxxxxxxxxx", "accessTokenSecret": "xxxxxxxxxxxxxxxxxxxx"}`

and

`{"client_id": "111111111111111111", "client_secret": "xxxxxxxxxxxxxxxxxxxx"}`

Although Redis isn't required, I highly recommend using it, because you'll quickly eat through your API usage limits if you don't cache responses.
