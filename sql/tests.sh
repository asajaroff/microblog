#!/usr/bin/bash

set -x

curl http://localhost:1323/login -d email=user@example.com&password=yourpassword

curl -v POST http://localhost:1323/login -d "email=jdoe@example.com&password=d21fdbcf8b5b4c1e324a7b2f1a1a"
