# pwa-blog-server

## Docker
```
docker build -t pwa-server .
docker run --rm --name pwa-server-c -p 8080:8080 pwa-server
```

## nginx
```
server {
  listen 443 ssl;
  server_name ...;
  root /home/ubuntu/blog;

  ssl_certificate            /etc/letsencrypt/live/.../fullchain.pem;
  ssl_certificate_key        /etc/letsencrypt/live/.../privkey.pem;
  ssl_protocols              TLSv1 TLSv1.1 TLSv1.2;
  ssl_prefer_server_ciphers  on;
  ssl_ciphers                AES256+EECDH:AES256+EDH:!aNULL;

  location /api {
    rewrite ^.+api/?(.*)$ /$1 break;
    proxy_pass http://127.0.0.1:8080;
  }
}
```

## routes
```
- http://127.0.0.1:8080/getDetail/123
- http://127.0.0.1:8080/getList
```
