# mac-wifi-fixer
tests for internet connection and resets mac wifi when needed for mac bug

It should be able to automatically detect your wifi adapter 

Usage: mac-wifi-fixer [-i] [url]

-i interval in minutes to run check ( default: 5 min )

url is a list of urls separated by spaces like https://google.com http://yahoo.com
that will be used to test connectivity. It also has a nice side effect of testing if those sites are active and serving Status OK messages. (default: https://google.com)

When a reset occurs it uses the "say" command to let you know it reset the interface.
