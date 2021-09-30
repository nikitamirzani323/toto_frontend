FROM golang:alpine AS totobuild
WORKDIR /go/src/bitbucket.org/isbtotogroup/frontend_svelte
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

# ---- Svelte Base ----
FROM node:lts-alpine AS totosveltebase
WORKDIR /svelteapp
COPY [ "svelte/package.json" , "svelte/yarn.lock" , "svelte/rollup.config.js" , "./"]

# ---- Svelte Dependencies ----
FROM totosveltebase AS totosveltedep
RUN yarn
RUN cp -R node_modules prod_node_modules

#
# ---- Svelte Builder ----
FROM totosveltebase AS totosveltebuilder
COPY --from=totosveltedep /svelteapp/prod_node_modules ./node_modules
COPY ./svelte .
RUN yarn build

# Moving the binary to the 'final Image' to make it smaller
FROM alpine:latest as totosvelterelease
WORKDIR /app
RUN mkdir -p ./svelte/public
COPY --from=totosveltebuilder /svelteapp/public ./svelte/public
COPY --from=totobuild /go/src/bitbucket.org/isbtotogroup/frontend_svelte/app .
EXPOSE 7071
CMD ["./app"]