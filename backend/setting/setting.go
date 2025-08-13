package setting

// ساختار تنظیمات
type Config struct {
	VPNApiBaseURL string
	VPNApiKey     string
}

// متغیر سراسری برای استفاده در کل پروژه
var AppConfig = Config{
	VPNApiBaseURL: "https://vpn-server.com/api", // آدرس پایه API
	VPNApiKey:     "YOUR_API_KEY",               // اگه نیاز به کلید API داری
}
