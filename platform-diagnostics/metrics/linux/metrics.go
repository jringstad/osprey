package linux

var Commands = map[string]string{
	"identifier": "hostname",
	"cpu_idle": "vmstat | tail -1 | cut -d ' ' -f42",
	"system_load": "uptime | cut -d ' ' -f 12 | sed 's/,//'",
	"uptime": "cat /proc/uptime | cut -d ' ' -f1",
	"autossh_status": "systemctl status autossh | grep \"Active\" | sed 's/.*Active: //' | sed 's/ since .*;//'",
	"network_eth0": "ifconfig eth0",
	"network_eth1": "ifconfig eth1",
}

/*
for single-shot info:
lsusb
cat /proc/cpuinfo
cat /proc/meminfo
cat /etc/issue.net
cat /etc/timezone
ip addr (full output)
uname -a
*/