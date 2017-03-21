// this file was generated by gomacro command: import "log/syslog"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	"log/syslog"
)

func init() {
	Packages["log/syslog"] = Package{
	Binds: map[string]Value{
		"Dial":	ValueOf(syslog.Dial),
		"LOG_ALERT":	ValueOf(syslog.LOG_ALERT),
		"LOG_AUTH":	ValueOf(syslog.LOG_AUTH),
		"LOG_AUTHPRIV":	ValueOf(syslog.LOG_AUTHPRIV),
		"LOG_CRIT":	ValueOf(syslog.LOG_CRIT),
		"LOG_CRON":	ValueOf(syslog.LOG_CRON),
		"LOG_DAEMON":	ValueOf(syslog.LOG_DAEMON),
		"LOG_DEBUG":	ValueOf(syslog.LOG_DEBUG),
		"LOG_EMERG":	ValueOf(syslog.LOG_EMERG),
		"LOG_ERR":	ValueOf(syslog.LOG_ERR),
		"LOG_FTP":	ValueOf(syslog.LOG_FTP),
		"LOG_INFO":	ValueOf(syslog.LOG_INFO),
		"LOG_KERN":	ValueOf(syslog.LOG_KERN),
		"LOG_LOCAL0":	ValueOf(syslog.LOG_LOCAL0),
		"LOG_LOCAL1":	ValueOf(syslog.LOG_LOCAL1),
		"LOG_LOCAL2":	ValueOf(syslog.LOG_LOCAL2),
		"LOG_LOCAL3":	ValueOf(syslog.LOG_LOCAL3),
		"LOG_LOCAL4":	ValueOf(syslog.LOG_LOCAL4),
		"LOG_LOCAL5":	ValueOf(syslog.LOG_LOCAL5),
		"LOG_LOCAL6":	ValueOf(syslog.LOG_LOCAL6),
		"LOG_LOCAL7":	ValueOf(syslog.LOG_LOCAL7),
		"LOG_LPR":	ValueOf(syslog.LOG_LPR),
		"LOG_MAIL":	ValueOf(syslog.LOG_MAIL),
		"LOG_NEWS":	ValueOf(syslog.LOG_NEWS),
		"LOG_NOTICE":	ValueOf(syslog.LOG_NOTICE),
		"LOG_SYSLOG":	ValueOf(syslog.LOG_SYSLOG),
		"LOG_USER":	ValueOf(syslog.LOG_USER),
		"LOG_UUCP":	ValueOf(syslog.LOG_UUCP),
		"LOG_WARNING":	ValueOf(syslog.LOG_WARNING),
		"New":	ValueOf(syslog.New),
		"NewLogger":	ValueOf(syslog.NewLogger),
	},
	Types: map[string]Type{
		"Priority":	TypeOf((*syslog.Priority)(nil)).Elem(),
		"Writer":	TypeOf((*syslog.Writer)(nil)).Elem(),
	},
	Proxies: map[string]Type{
	} }
}
