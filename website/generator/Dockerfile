FROM ubuntu:latest
RUN apt-get update && apt-get install -y pandoc
WORKDIR /website
# CMD ls
CMD find markdown -type f -printf '%p' -exec \
    pandoc -o {} \;
# CMD find ./ -iname "*.md" -type f -exec sh -c 'pandoc "${0}" -o "${0%}.html"' {} \;