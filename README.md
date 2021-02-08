


生成私钥

```bash
$ openssl genrsa -des3 -out private.pem 2048
```

生成公钥

```bash
$ openssl rsa -in private.pem -outform PEM -pubout -out public.pem
```

