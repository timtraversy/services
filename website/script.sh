# docker run --rm -it \
#     -v $(pwd):/src \
#     --user $(id -u):$(id -g) \
#     -p 1313:1313 \
#     klakegg/hugo:ext-alpine \
#     shell
docker-compose run --user $(id -u):$(id -g) --service-ports server $@