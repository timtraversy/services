FROM golang:latest
RUN apt-get update && apt-get install -y pandoc
WORKDIR /website
# CMD ls
# CMD find markdown -name '*.md' -type f -exec \
#  pandoc -o '${}%.md'.html {} -s -H css/stylesheet.html \;
# CMD find ./ -iname "*.md" -type f -exec sh -c 'pandoc "${0}" -o "${0%}.html"' {} \;
# pandoc -o home.html menu.md home.md -s -H stylesheet.html