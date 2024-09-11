FROM golang:1.17-alpine AS sdsb4dbuildm
WORKDIR /go/src/bitbucket.org/isbtotogroup/sdsb4d-backend
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

# ---- Svelte Base ----
FROM node:lts-alpine AS sdsb4dbesvelte
WORKDIR /svelteapp
COPY [ "sveltemdb/package.json" , "sveltemdb/yarn.lock" , "sveltemdb/rollup.config.js" , "./"]

# ---- Svelte Dependencies ----
FROM sdsb4dbesvelte AS sdsb4ddepsvelte
RUN yarn
RUN cp -R node_modules prod_node_modules

#
# ---- Svelte Builder ----
FROM sdsb4dbesvelte AS sdsb4dbuilder
COPY --from=sdsb4ddepsvelte /svelteapp/prod_node_modules ./node_modules
COPY ./sveltemdb .
RUN yarn build

# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest as totosvelterelease
WORKDIR /app
RUN apk add tzdata
RUN mkdir -p ./sveltemdb/public
COPY --from=sdsb4dbuilder /svelteapp/public ./sveltemdb/public
COPY --from=sdsb4dbuildm /go/src/bitbucket.org/isbtotogroup/sdsb4d-backend/app .
COPY --from=sdsb4dbuildm /go/src/bitbucket.org/isbtotogroup/sdsb4d-backend/env-sample /app/.env

ENV TZ=Asia/Jakarta
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

EXPOSE 7071
CMD ["./app"]