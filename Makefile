.PHONY: gen
gen: zwiftpower.proto
	protoc --go_opt=module=github.com/jonahgeorge/zwiftpower-go --go_out=. --go-json_out=allow_unknown:. zwiftpower.proto

test:
	go test -v
