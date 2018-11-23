install:
	brew install flatbuffers
build-schema:
	cd src/lib/schema; flatc --ts *.fbs
	cat src/lib/schema/fbsheader.txt src/lib/schema/ntriples_generated.ts > tmp && mv tmp src/lib/schema/ntriples_generated.ts
	cat src/lib/schema/fbsheader.txt src/lib/schema/curve_generated.ts > tmp && mv tmp src/lib/schema/curve_generated.ts
	cat src/lib/schema/fbsheader.txt src/lib/schema/dalmatian_generated.ts > tmp && mv tmp src/lib/schema/dalmatian_generated.ts
