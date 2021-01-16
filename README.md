# TLS Testing

A quick and dirty implementation of a client and server to test the performance differences between TLS and Mutual TLS (mTLS).

## How to build

1. Compile the server ``go build server.go``
2. Compile the client ``go build client.go``

## How to generate cert and key

You'll need a certificate and a key to run this called ``cert.pem`` and ``key.pem``

```bash
openssl req -newkey rsa:2048 \
            -new -nodes -x509 \
            -days 3650 \
            -out cert.pem \
            -keyout key.pem \
            -subj "/C=AU/ST=Melbourne/L=rjk/O=rjk/OU=rjk/CN=localhost"
```

## How to run

This can run in two modes, with mTLS and without.

### With mTLS

1. Run the server ``./server -mtls``
2. Run the client ``./client -mtls -loops=10`` (replace the loops number with the number of times you'd like the client to request from the server)

### Without mTLS

Same as above, but omit the ``-mtls`` flag.

## Acknowledgement

The code is based on an example here https://venilnoronha.io/a-step-by-step-guide-to-mtls-in-go