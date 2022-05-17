# dyn-dns
The goal of this application is to connect a domain name to a machine where the
IP changes. Eg. one running behind a router where the ISP will occasionally
change the IP og the router.

This setup will probably have some lag when the IP changes because:
1. The service will need to detect the change.
    * Depending on how the IP is detected, we probably don't want to poll to
    often.
2. The DNS record will have to be updated.
    * Want to set a low TTL.

Will not be good enough for a production environment, but then one should
probably not use a setup like this anyways.

Currently using http://ip-api.com/json/ to poll public IP.
There are, of course, other options. Eg. making a simple lambda in AWS, or a
cloud function in Google Cloud. This would make the setup more (and
unnecessarily) complicated for this simple use case. Will however try to keep
the setup extendable to be able to implement other ways of retrieving the public
IP.