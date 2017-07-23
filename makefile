PUML_FILES := $(wildcard doc/*.puml)
IMAGE_FILES := $(addprefix doc/,$(notdir $(PUML_FILES:.puml=.png)))

doc: $(IMAGE_FILES)

doc/%.png: doc/%.puml
	java -jar /bin/plantuml.jar $^

clean:
	rm doc/*.png
