user www-data;
worker_processes auto;
pid /run/nginx.pid;

events {
}

http {
  server {
    listen 80;

    location / {
      proxy_pass http://localhost:8080;
    }
  }
}