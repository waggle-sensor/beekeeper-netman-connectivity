# Beekeeper NetworkManager Connectivity

This is a simple http response service that replies to node issued NetworkManager connectivity checks to verify Internet access. This service is being created within Beekeeper so that the nodes can do a connectivity check over the "wan tunnel" established with Beekeeper.

## NetworkManager `curl` results

The NetworkManager connectivity test does a interface bound curl command:

```
curl --interface wan0 -v https://connectivity-check.sagecontinuum.org

* Rebuilt URL to: https://connectivity-check.sagecontinuum.org/
*   Trying 192.5.86.2...
* TCP_NODELAY set
* Connected to connectivity-check.sagecontinuum.org (192.5.86.2) port 443 (#0)
* ALPN, offering h2
* ALPN, offering http/1.1
* successfully set certificate verify locations:
*   CAfile: /etc/ssl/certs/ca-certificates.crt
  CApath: /etc/ssl/certs
* TLSv1.3 (OUT), TLS handshake, Client hello (1):
* TLSv1.3 (IN), TLS handshake, Server hello (2):
* TLSv1.3 (IN), TLS Unknown, Certificate Status (22):
* TLSv1.3 (IN), TLS handshake, Unknown (8):
* TLSv1.3 (IN), TLS Unknown, Certificate Status (22):
* TLSv1.3 (IN), TLS handshake, Certificate (11):
* TLSv1.3 (IN), TLS Unknown, Certificate Status (22):
* TLSv1.3 (IN), TLS handshake, CERT verify (15):
* TLSv1.3 (IN), TLS Unknown, Certificate Status (22):
* TLSv1.3 (IN), TLS handshake, Finished (20):
* TLSv1.3 (OUT), TLS change cipher, Client hello (1):
* TLSv1.3 (OUT), TLS Unknown, Certificate Status (22):
* TLSv1.3 (OUT), TLS handshake, Finished (20):
* SSL connection using TLSv1.3 / TLS_AES_256_GCM_SHA384
* ALPN, server accepted to use h2
* Server certificate:
*  subject: CN=*.sagecontinuum.org
*  start date: Apr 13 00:00:00 2022 GMT
*  expire date: May 14 23:59:59 2023 GMT
*  subjectAltName: host "connectivity-check.sagecontinuum.org" matched cert's "*.sagecontinuum.org"
*  issuer: C=FR; ST=Paris; L=Paris; O=Gandi; CN=Gandi Standard SSL CA 2
*  SSL certificate verify ok.
* Using HTTP2, server supports multi-use
* Connection state changed (HTTP/2 confirmed)
* Copying HTTP/2 data in stream buffer to connection buffer after upgrade: len=0
* TLSv1.3 (OUT), TLS Unknown, Unknown (23):
* TLSv1.3 (OUT), TLS Unknown, Unknown (23):
* TLSv1.3 (OUT), TLS Unknown, Unknown (23):
* Using Stream ID: 1 (easy handle 0x55928ce9a0)
* TLSv1.3 (OUT), TLS Unknown, Unknown (23):
> GET / HTTP/2
> Host: connectivity-check.sagecontinuum.org
> User-Agent: curl/7.58.0
> Accept: */*
>
* TLSv1.3 (IN), TLS Unknown, Certificate Status (22):
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* TLSv1.3 (IN), TLS Unknown, Certificate Status (22):
* TLSv1.3 (IN), TLS handshake, Newsession Ticket (4):
* TLSv1.3 (IN), TLS Unknown, Unknown (23):
* Connection state changed (MAX_CONCURRENT_STREAMS updated)!
* TLSv1.3 (OUT), TLS Unknown, Unknown (23):
* TLSv1.3 (IN), TLS Unknown, Unknown (23):
* TLSv1.3 (IN), TLS Unknown, Unknown (23):
< HTTP/2 204
< date: Fri, 20 May 2022 22:49:19 GMT
< x-networkmanager-status: online
< strict-transport-security: max-age=15724800; includeSubDomains
<
* Connection #0 to host connectivity-check.sagecontinuum.org left intact
```

## References

https://wiki.archlinux.org/title/NetworkManager#Checking_connectivity
https://man.archlinux.org/man/NetworkManager.conf.5#CONNECTIVITY_SECTION
