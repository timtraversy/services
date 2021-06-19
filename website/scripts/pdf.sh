docker run --rm --volume "`pwd`:/data" --user `id -u`:`id -g` pandoc/latex resume.md -o README.pdf
