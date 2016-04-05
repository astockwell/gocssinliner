gocssinliner
============
Quick 'n dirty CLI to Campaign Monitor's [CSS inliner service](https://inliner.cm/) for HTML emails.


Usage
-----
`gocssinliner -i infile.html -o outfilename.html`

Requires internet access. Use responsibly, as I'm sure the CM CSS inliner service will rate-limit you if you abuse.


Install
-------
Download the latest compiled binary from [releases](https://github.com/astockwell/gocssinliner/releases) and put it in your PATH (somewhere like /usr/local/bin will usually work on OSX) or project folder to run locally.


Build
-----
1. Clone & cd into the repo
2. cd to `cmd/gocssinliner` and run `go install` (requires go compiler installed)


Credit
------
Big thanks to [Campaign Monitor](https://www.campaignmonitor.com/) for maintaining their awesome [CSS inliner](https://inliner.cm/) service!
