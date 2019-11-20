# Build Env
FROM bamdockerhub/go-alpine-build AS build-env

ENV GO111MODULE=on

ADD . /go/src/github.com/blockchain-abstraction-middleware/game-jam-abstraction

WORKDIR /go/src/github.com/blockchain-abstraction-middleware/game-jam-abstraction

RUN go build -i -o app ./cmd/serve/main.go

# Application Image
FROM bamdockerhub/alpine-gpg

COPY --from=build-env /go/src/github.com/blockchain-abstraction-middleware/game-jam-abstraction/app /usr/local/bin/app

COPY --from=build-env /go/src/github.com/blockchain-abstraction-middleware/game-jam-abstraction/cmd/serve/config ./config

CMD ["sh", "-c", "gpg --allow-secret-key-import --import /.secret-key/secring.pgp; /usr/local/bin/app"]