FROM alpine AS builder

RUN mkdir -p /nuxt-blog

WORKDIR /nuxt-blog

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add --no-cache --update nodejs nodejs-npm

COPY package.json  package-lock.json ./

RUN npm config set registry=https://registry.npm.taobao.org

RUN npm install --production


FROM alpine

# MAINTAINER
LABEL name="nuxt-blog"
LABEL version="1.0.1"
LABEL author="bigfool <1063944784@qq.com>"
LABEL maintainer="bigfool <1063944784@qq.com>"
LABEL description="nuxt-blog application"

RUN mkdir -p /nuxt-blog \
    && mkdir -p /nuxt-blog/logs

# 新建一个用户www 并设置项目目录用户组
RUN adduser -D -H www \
    && chown -R www /nuxt-blog

WORKDIR /nuxt-blog

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

RUN apk add --no-cache --update nodejs nodejs-npm

COPY --from=builder /nuxt-blog/node_modules ./node_modules
COPY . .

RUN npm config set registry=https://registry.npm.taobao.org \
    && npm install node-sass --sass_binary_site=https://npm.taobao.org/mirrors/node-sass/ \
    && npm install -g pm2

RUN npm run build

RUN rm -rf `ls | egrep -v '(.nuxt|node_modules|static|pm2.json|nuxt.config.js|package.json|package-lock.json)'`

ENV NODE_ENV=production

EXPOSE 8081

CMD ["pm2-docker", "pm2.json"]
