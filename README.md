# gohuedata

Command-line tool written in Go to get data from your Philips Hue system. Just a little work-in-progress that I'm fiddling with.

## How to develop

1. `cp config.yml.example config.yml`
1. Update config.yml with details about your Philips Hue bridges. See [Get started](https://developers.meethue.com/develop/get-started-2/) docs from Philips about how to get your bridge IP address and a username.
1. `go run cmd/gohuedata/main.go`
