package model

type DisplaySettings struct {
	HidePublicIP   bool `json:"hide_public_ip"`
	HidePrivateIP  bool `json:"hide_private_ip"`
	HideGPU0       bool `json:"hide_gpu0"`
	HideGPU1       bool `json:"hide_gpu1"`
	HideUsername   bool `json:"hide_username"`
	HideHostname   bool `json:"hide_hostname"`
	HideKernel     bool `json:"hide_kernel"`
	HideUptime     bool `json:"hide_uptime"`
	HideResolution bool `json:"hide_resolution"`
	HideShell      bool `json:"hide_shell"`
	HideDE         bool `json:"hide_de"`
	HideWM         bool `json:"hide_wm"`
}

type ThemeSettings struct {
	ColorOutput     string `json:"color_output"`
	CompactMode     bool   `json:"compact_mode"`
	FontStyle       string `json:"font_style"`
	UseDifferentimg bool   `json:"use_differentimg"`
	ImageSource     string `json:"image_source"`
}

type Config struct {
	Display DisplaySettings `json:"display"`
	Theme   ThemeSettings   `json:"theme"`
}
