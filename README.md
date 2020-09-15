# [Blog](https://kuderblog.com)
Created my own blog on golang and Nuxt.js. Got a lot of features to add, but for now i can create, edit, publish and delete posts, all this from frontend (so it ofcourse has users and auth logic, but right now there is only one user: me). For editing posts using rich text editor [tiptap-vuetify](https://github.com/iliyaZelenko/tiptap-vuetify). Wanted to deploy version 1.0 of blog to start posting my ideas, so may have some little bugs, but work continues.

Repo is public for limited time for presentation only purposes.

# NGINX conf

```bash
server {
      listen         80;
      server_name    kuderblog.com www.kuderblog.com;
      return         301 https://$server_name$request_uri;
}

server {
      server_name kuderblog.com www.kuderblog.com;
      listen 443 ssl http2;
      ssl on;
      ssl_certificate     /etc/letsencrypt/live/kuderblog.com/fullchain.pem;
      ssl_certificate_key /etc/letsencrypt/live/kuderblog.com/privkey.pem;
      ssl_protocols TLSv1.2 TLSv1.3;

      error_log /var/log/blog/blog_error.log error;
      access_log /var/log/blog/blog_access.log;

      location / {
        proxy_pass http://localhost:3000;
      }

      location /api {
        proxy_pass http://localhost:5000;
      }
}
```
