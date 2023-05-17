# XunYouRB
Basic functionality on CiYuanJi
Based on the Android client API, the body decryption is simply implemented, and all request parameters pass the verification normally, which can be easily downloaded

------------


### Instructions for use
Configure your own token in the variable Token of the request.go file (obtained by capturing packets)
```go
var (
	headers = map[string]string{
		"Channel":     "35",
		"User-Agent":  "Mozilla/5.0 (Linux; Android 11; Pixel 4 XL Build/RP1A.200720.009; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/92.0.4515.115 Mobile Safari/537.36",
		"Targetmodel": "SM-N9700",
		"Platform":    "1",
		"Deviceno":    "d0b7cef20c3c6b5f",
		"Version":     "3.3.2",
		"Token":       "",
	}
)
```
