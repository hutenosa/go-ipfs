package config

// CurrentCommit is the current git commit, this is set as a ldflag in the Makefile
var CurrentCommit string

// CurrentVersionNumber is the current application's version literal
var CurrentVersionNumber = "0.4.2"

var ApiVersion = "/go-ipfs/" + CurrentVersionNumber + "/"
