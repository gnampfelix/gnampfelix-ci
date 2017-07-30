PUML_FILES := $(wildcard doc/*.puml)
IMAGE_FILES := $(addprefix doc/,$(notdir $(PUML_FILES:.puml=.svg)))

doc: $(IMAGE_FILES)

doc/%.svg: doc/%.puml
	java -jar /bin/plantuml.jar -tsvg $^

.PHONY: images
images:
	docker build -t ci-builder images/builder
	#docker push ci-builder

clean:
	rm doc/*.svg
