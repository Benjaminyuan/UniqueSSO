FROM node:latest As builder

WORKDIR /app

COPY . /app/

RUN yarn config set registry https://mirrors.sjtug.sjtu.edu.cn/npm-registry && \
    yarn build

FROM xylonx/cn-ubuntu:latest As prod

LABEL maintainer="xylon"

ARG PROJECT_NAME="uniquesso-dashboard"

WORKDIR /opt/${PROJECT_NAME}
COPY --from=0 /app/dist .