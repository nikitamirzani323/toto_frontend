server {
   listen 80;
   listen [::]:80;

   server_name tavuli.com;

   location / {
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header Host $http_host;
		proxy_read_timeout 5000;
		proxy_connect_timeout 5000;
		proxy_send_timeout 5000;
		send_timeout 5000;
        proxy_pass http://127.0.0.1:7071;
    }

    location ~ /.well-known {
        allow all;
    }

}
