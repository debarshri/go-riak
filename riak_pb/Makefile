UPSTREAM_TAG := 2.1.0.2

TARGET = riak_pb.go

PROTO = riak_kv.proto \
		riak.proto \
		riak_search.proto \
		riak_dt.proto \
		riak_yokozuna.proto \

GOFILES	= $(PROTO:.proto=.pb.go)

define TEMPLATE
package riak

import (
endef

export TEMPLATE

all: $(GOFILES)
	@echo "$$TEMPLATE" > $(TARGET)
	@for file in $(GOFILES); do \
		sed -n '/import riak/b; s/import //p' $$file >> $(TARGET); \
		sed -i '/^import/d' $$file; \
		sed -i '/^package/d' $$file; \
	done
	@echo ")" >> $(TARGET)
	@cat $(GOFILES) >> $(TARGET)
	@sed -i 's/riak.RpbPair/RpbPair/g' $(TARGET)
	@echo "formatting file..."
	@go fmt $(TARGET)
	@mv $(TARGET) ../

%.pb.go : %.proto
	protoc --plugin=protoc-gen-go=$(GOPATH)/bin/protoc-gen-go --go_out=./ $<

clean:
	rm -f $(GOFILES)
