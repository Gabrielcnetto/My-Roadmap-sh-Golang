# 2. Generate web client's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-cert.pem -subj "/C=BR/ST=Sao Paulo/L=Sao Paulo/O=Office/OU=Computer/CN=*.tec1.com/emailAddress=silas1@tec1.com"

# 3. Use CA's private key to sign web client's CSR and get back the signed certificate
openssl x509 -req -in client-cert.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out client-cert.pem -extfile client-ext.cnf

echo "client's signed certificate"
openssl x509 -in ca-cert.pem -noout -text