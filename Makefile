install:
	brew install flatbuffers
build-schema:
	cd src/lib/schema; flatc --ts dalmatian.fbs