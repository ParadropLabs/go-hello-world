# Go Hello World Chute

This example demonstrates running a Go application on Paradrop.

## Installing

* Install the Paradrop command line utility. This requires Python.

    sudo pip install pdtools

* Install the chute on a Paradrop device. Replace <address> with the IP address or domain name of your device.

    pdtools device <address> chutes create

* Test the chute. Since it is running a web server, you can use curl.

    curl --location http://<address>/chutes/go-hello-world
