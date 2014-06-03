ZoneInfo
========

Iterate through tzinfo database from iana.org/time-zones and look for bogus entries. *Found Some!!*

```
 $ go run ./seeplaces.go | grep ERR
ERR at Location: 'unknown time zone America/Godthab'
ERR at Location: 'unknown time zone America/Santiago'
ERR at Location: 'unknown time zone Antarctica/Palmer'
ERR at Location: 'unknown time zone Asia/Tel_Aviv'
ERR at Location: 'unknown time zone Asia/Jerusalem'
ERR at Location: 'unknown time zone Asia/Gaza'
ERR at Location: 'unknown time zone Asia/Hebron'
ERR at Location: 'unknown time zone Chile/Continental'
ERR at Location: 'unknown time zone Chile/EasterIsland'
ERR at Location: 'unknown time zone Israel'
ERR at Location: 'unknown time zone Pacific/Fiji'
ERR at Location: 'unknown time zone Pacific/Easter'
```
The small script makenamelist.sh does most of the work of making the list of places.

At this point I can't be sure whether the error is in the Go time library or the zoneinfo.zip from iana.org.  They both look right under light perusal.
