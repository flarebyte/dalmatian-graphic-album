SCHEMA = model/schema
NAME = dalmatian
TYPESCRIPT = typescript/src/lib

install-ts:
	cd typescript; yarn install

install:
	brew install flatbuffers

build-schema:
	cat $(SCHEMA)/parts/*.fbs > $(SCHEMA)/$(NAME).fbs

generate-ts:
	cd $(SCHEMA); flatc --ts --no-fb-import -o ../../$(TYPESCRIPT) $(NAME).fbs
	cat $(SCHEMA)/parts/fbsheader.txt $(TYPESCRIPT)/$(NAME)_generated.ts > tmp && mv tmp $(TYPESCRIPT)/$(NAME)_generated.ts

generate-go:
	cd $(SCHEMA); flatc --go --no-fb-import -o ../../go $(NAME).fbs
