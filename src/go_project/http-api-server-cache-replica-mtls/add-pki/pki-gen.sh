#!/bin/bash

# 1.1 Generate CA private key
openssl genrsa -out ca-local.key 4096

# 1.2 Generate self-signed CA certificate
openssl req -x509 -new -nodes -key ca-local.key -sha256 -days 3650 -out ca-local.crt \
  -subj "/C=AU/ST=Victoria/L=Melbourne/O=MyLocal/CN=MyLocalRootCA"

# 2.1 Generate server private key
openssl genrsa -out server-local.key 2048

# 2.2 Create server CSR (Certificate Signing Request)
openssl req -new -key server-local.key -out server-local.csr \
  -subj "/C=AU/ST=Victoria/L=Melbourne/O=MyLocal/CN=localhost"

# 2.3 Create a config file for SAN (Subject Alternative Name)
cat > server-local.ext <<EOF
authorityKeyIdentifier=keyid,issuer
basicConstraints=CA:FALSE
keyUsage = digitalSignature, keyEncipherment
subjectAltName = @alt_names

[alt_names]
DNS.1 = localhost
IP.1 = 127.0.0.1
EOF

# 2.4 Sign the server cert with the CA
openssl x509 -req -in server-local.csr -CA ca-local.crt -CAkey ca-local.key -CAcreateserial \
  -out server-local.crt -days 365 -sha256 -extfile server-local.ext

# You now have:
# server.key (private key)
# server.crt (signed server certificate)
# ca.crt (CA to trust the cert)

# 3.1 Generate client private key
openssl genrsa -out client-local.key 2048

# 3.2 Create client CSR
openssl req -new -key client-local.key -out client-local.csr \
  -subj "/C=AU/ST=Victoria/L=Melbourne/O=MyLocal/CN=client-user"

# 3.3 Optional: client cert config
cat > client-local.ext <<EOF
basicConstraints=CA:FALSE
keyUsage = digitalSignature
extendedKeyUsage = clientAuth
EOF

# 3.4 Sign client cert with the CA
openssl x509 -req -in client-local.csr -CA ca-local.crt -CAkey ca-local.key -CAcreateserial \
  -out client-local.crt -days 365 -sha256 -extfile client-local.ext

# You now have:
# client.key (private key)
# client.crt (signed client certificate)
# ca.crt (used by server to validate client cert)


# curl https://localhost:8443/replicas --cert client.crt --key client.key --cacert ca.crt
