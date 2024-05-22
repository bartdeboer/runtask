module github.com/bartdeboer/runtask

go 1.22.1

// replace github.com/bitfield/script => ../script

require (
	github.com/bartdeboer/script/v2 v2.0.0
	github.com/bitfield/script v0.22.0
	github.com/traefik/yaegi v0.16.1
)

require (
	github.com/bartdeboer/pipeline v0.0.3 // indirect
	github.com/itchyny/gojq v0.12.13 // indirect
	github.com/itchyny/timefmt-go v0.1.5 // indirect
	mvdan.cc/sh/v3 v3.7.0 // indirect
)
