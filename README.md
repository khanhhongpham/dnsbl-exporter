# Dnsbl Exporter (Experimental)
Prometheus exporter for checking domains listed on blacklists
## Build and Run
### With go
build executable file
```
cd dnsbl-exporter
go build
```
running with [default](default.yml) config 
```
./dnsbl-exporter 
```
running with custom config
``` 
./dnsbl-exporter --config config.yml
```
### With docker
build docker image
```
docker build -t dnsbl-exporter .
```
running with [default](default.yml) config 
```
docker run -p 8881:8881 dnsbl-exporter 
```
running with custom config 
```
docker run -p 8881:8881 dnsbl-exporter --config config.yml -v config.yml:config.yml
```
### Default blacklists
```
  - 0spam.fusionzero.com
  - all.rbl.webiron.net
  - all.s5h.net
  - babl.rbl.webiron.net
  - badnets.spameatingmonkey.net
  - b.barracudacentral.org
  - blacklist.woody.ch
  - bl.blocklist.de
  - bl.nordspam.com
  - bl.mailspike.net
  - bl.spamcop.net
  - bl.spameatingmonkey.net
  - bl.tiopan.com
  - bogons.cymru.com
  - cabl.rbl.webiron.net
  - cbl.abuseat.org
  - cbl.anti-spam.org.cn
  - cdl.anti-spam.org.cn
  - cml.anti-spam.org.cn
  - db.wpbl.info
  - dnsbl-1.uceprotect.net
  - dnsbl-2.uceprotect.net
  - dnsbl-3.uceprotect.net
  - dnsbl.anticaptcha.net
  - dnsbl.cobion.com
  - dnsbl.inps.de
  - dnsbl.spfbl.net
  - dnsbl.zapbl.net
  - dnsrbl.swinog.ch
  - drone.abuse.ch
  - dul.dnsbl.sorbs.net
  - dyna.spamrats.com
  - httpbl.abuse.ch
  - http.dnsbl.sorbs.net
  - images.rbl.msrbl.net
  - ips.backscatterer.org
  - ix.dnsbl.manitu.net
  - korea.services.net
  - misc.dnsbl.sorbs.net
  - netbl.spameatingmonkey.net
  - nomail.rhsbl.sorbs.ne
  - noptr.spamrats.com
  - pbl.spamhaus.org
  - phishing.rbl.msrbl.net
  - rbl2.triumf.ca
  - rbl.megarbl.net
  - rbl.realtimeblacklist.com
  - rbl.schulte.org
  - relays.nether.net
  - sbl.spamhaus.org
  - smtp.dnsbl.sorbs.net
  - socks.dnsbl.sorbs.net
  - spam.abuse.ch
  - spam.dnsbl.sorbs.net
  - spamguard.leadmon.net
  - spamrbl.imp.ch
  - spam.rbl.msrbl.net
  - spamsources.fabel.dk
  - spam.spamrats.com
  - srnblack.surgate.net
  - stabl.rbl.webiron.net
  - st.technovision.dk
  - tor.dan.me.uk
  - truncate.gbudb.net
  - ubl.unsubscore.com
  - virus.rbl.msrbl.net
  - web.dnsbl.sorbs.net
  - web.rbl.msrbl.net
  - wormrbl.imp.ch
  - xbl.spamhaus.org
  - zombie.dnsbl.sorbs.net
  - bl.ipv6.spameatingmonkey.net
  - ipv6.blacklist.woody.ch
  - v6.fullbogons.cymru.com
  - black.uribl.com
  - dbl.nordspam.com
  - dbl.spamhaus.org
  - dbl.tiopan.com
  - dyndns.rbl.jp
  - grey.uribl.com
  - list.anonwhois.net
  - multi.surbl.org
  - red.uribl.com
  - rhsbl.sorbs.net
  - rhsbl.zapbl.net
  - uri.blacklist.woody.ch
  - uribl.spameatingmonkey.net
  - uribl.swinog.ch
  - url.rbl.jp
```