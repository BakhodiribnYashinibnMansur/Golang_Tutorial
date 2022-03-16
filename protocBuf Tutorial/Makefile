CURRENT_DIR=$(shell pwd)

proto-gen:
	./scripts/gen-proto.sh	${CURRENT_DIR}

clone-protos:
	rm -rf protos/* && cp -R ur_protos/* protos

pull-proto-module:
	git submodule update --init --recursive

update-proto-module:
	git submodule update --remote --merge


.PHONY: proto