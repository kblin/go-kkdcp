Kerberos Key Distribution Center Proxy
======================================

A Go implementation of [MS-KKDCP](http://msdn.microsoft.com/en-us/library/hh553774.aspx),
a protocol to proxy Kerberos ticket requests via HTTP(S).

Deployment
----------

After installing the dependencies and building with `go build`, you can simply run the `go-kkdcp` executable.

Ideally, you use a reverse proxy server to front this and handle SSL, like so:

```
server {
	listen 443;
	listen [::]:443;
	server_name kdcproxy.demo.kblin.org;

	ssl on;
	ssl_certificate /etc/ssl/certs/kdcproxy.pem;
	ssl_certificate_key /etc/ssl/private/kdcproxy.key;

	root /var/www/kdxproxy;
	index index.html;

	location /KdcProxy {
		proxy_pass http://127.0.0.1:8124/;
		include proxy_params;
		add_header Cache-Control "no-cache, no-store, must-revalidate";
		add_header Pragma no-cache;
		add_header Expires 0;
	}
}
```

Microsoft suggests `/KdcProxy` as an endpoint, but at least with MIT Kerberos' client tools other paths work as well.

License
-------

This software is licensed under the GNU GPL v3 (or later), see [`LICENSE`](LICENSE) for details.
