package rpi

var Commands = map[string]string{
	"core_temperature": "vcgencmd measure_temp",
	"camera_status": "vcgencmd get_camera",
	"throttle_status": "vcgencmd get_throttled",
	"voltage_core": "vcgencmd measure_volts core",
	"voltage_sdram_c": "vcgencmd measure_volts sdram_c",
	"voltage_sdram_i": "vcgencmd measure_volts sdram_i",
	"voltage_sdram_p": "vcgencmd measure_volts sdram_p",
	"clockrate_arm": "vcgencmd measure_clock arm",
	"clockrate_core": "vcgencmd measure_clock core",
	"clockrate_h264": "vcgencmd measure_clock h264",
	"clockrate_isp": "vcgencmd measure_clock isp",
	"clockrate_v3d": "vcgencmd measure_clock v3d",
	"clockrate_uart": "vcgencmd measure_clock uart",
	"clockrate_pwm": "vcgencmd measure_clock pwm",
	"clockrate_emmc": "vcgencmd measure_clock emmc",
	"clockrate_pixel": "vcgencmd measure_clock pixel",
	"clockrate_vec": "vcgencmd measure_clock vec",
	"clockrate_hdmi": "vcgencmd measure_clock hdmi",
	"clockrate_dpi": "vcgencmd measure_clock dpi",
}


/*
for single-shot info:
vcgencmd get_mem arm && vcgencmd get_mem gpu
vcgencmd version
vcgencmd get_config int
 */