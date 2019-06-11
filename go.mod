module github.com/xbonlinenet/alter

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190605123033-f99c8df09eb5
	golang.org/x/net => github.com/golang/net v0.0.0-20190607181551-461777fb6f67
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190610200419-93c9922d18ae
	golang.org/x/text => github.com/golang/text v0.3.2
	google.golang.org/appengine => github.com/golang/appengine v1.6.1
	gopkg.in/fsnotify.v1 => github.com/fsnotify/fsnotify v1.4.7
)

go 1.12
