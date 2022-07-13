
Tool to change IP address of cloudflare DNS records.

Build:
```
go build
```

Config file:
```
cat app.env
ZONE_NAME=mydomain.net
```

Examples:

Get my public IP address:
```sh
./homeDDNS publicIp
```

List DNS entries:
```sh
API_TOKEN=<Cloudflare API token> ./homeDDNS --config=app.env dns list
```

Check if IP address of records is different from current public IP
```sh
API_TOKEN=<Cloudflare API token> ./homeDDNS --config=app.env dns update --dry-run
```

Update record if IP address is different from current public IP
```sh
API_TOKEN=<Cloudflare API token> ./homeDDNS --config=app.env dns update
```



