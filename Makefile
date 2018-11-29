install:
	brew install flatbuffers
build-schema:
	cat src/lib/schema/ntriples.fbs src/lib/schema/curve.fbs src/lib/schema/layout.fbs src/lib/schema/album.fbs > src/lib/schema/dalmatian.fbs
	cd src/lib/schema; flatc --ts --no-fb-import dalmatian.fbs
	cat src/lib/schema/fbsheader.txt src/lib/schema/dalmatian_generated.ts > tmp && mv tmp src/lib/schema/dalmatian_generated.ts
