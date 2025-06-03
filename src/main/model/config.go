package model

type Config struct {
	HidePublicIP   bool   `json:"hide_public_ip"`
	HidePrivateIP  bool   `json:"hide_private_ip"`
	HideGPU0       bool   `json:"hide_gpu0"`
	HideGPU1       bool   `json:"hide_gpu1"`
	HideUsername   bool   `json:"hide_username"`
	HideHostname   bool   `json:"hide_hostname"`
	HideKernel     bool   `json:"hide_kernel"`
	HideUptime     bool   `json:"hide_uptime"`
	HideResolution bool   `json:"hide_resolution"`
	HideShell      bool   `json:"hide_shell"`
	HideDE         bool   `json:"hide_de"`
	HideWM         bool   `json:"hide_wm"`
	ColorOutput    bool   `json:"color_output"`
	UseEmoji       bool   `json:"use_emoji"`
	CompactMode    bool   `json:"compact_mode"`
	OutputFormat   string `json:"output_format"`
	StylePreset    string `json:"style_preset"`
	FontStyle      string `json:"font_style"`
}
