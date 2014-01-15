aboutme
=======

A page explaining what I've been doing on stack, github, and twitter.

Currently running at [bclymer.com/](http://bclymer.com/).

You can make this all about you by switching the ids in

- twitter.go
- stack.go
- github.go

You'll also need to configure twitter by including `twitterAuth.json` in the root directory.

It should look like

`{"consumerKey": "xxxxxxxxxxxxxxxxxxxx", "consumerSecret": "xxxxxxxxxxxxxxxxxxxx", "accessToken": "111111111-xxxxxxxxxxxxxxxxxxxx", "accessTokenSecret": "xxxxxxxxxxxxxxxxxxxx"}`

Although Redis isn't required, I highly recommend using it, because you'll quickly eat through your API usage limits if you don't cache responses.
