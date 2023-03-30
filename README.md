# Uniform

## Dependencies

* [go-nats](https://github.com/nats-io/nats.go)
* [go-diary](https://github.com/go-diary/diary)

```
sudo openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /etc/ssl/private/uniform-nats.key -out /etc/ssl/certs/uniform-nats.crt
sudo openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /etc/ssl/private/uniform-amqps.key -out /etc/ssl/certs/uniform-amqps.crt
sudo chmod +r /etc/ssl/private/uniform-nats.key
sudo chmod +r /etc/ssl/private/uniform-amqps.key
```