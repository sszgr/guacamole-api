package guacamole

type GuacConnection struct {
	Name              string                   `json:"name"`
	Identifier        string                   `json:"identifier,omitempty"`
	ParentIdentifier  string                   `json:"parentIdentifier"`
	Protocol          string                   `json:"protocol"`
	Attributes        GuacConnectionAttributes `json:"attributes"`
	Properties        GuacConnectionParameters `json:"parameters"`
	ActiveConnections int                      `json:"activeConnections,omitempty"`
	LastActive        int                      `json:"lastActive,omitempty"`
}

type GuacConnectionAttributes struct {
	GuacdEncryption       string `json:"guacd-encryption"`
	FailoverOnly          string `json:"failover-only"`
	Weight                string `json:"weight"`
	MaxConnections        string `json:"max-connections"`
	GuacdHostname         string `json:"guacd-hostname,omitempty"`
	GuacdPort             string `json:"guacd-port"`
	MaxConnectionsPerUser string `json:"max-connections-per-user"`
}

// GuacConnectionParameters See also https://github.com/apache/guacamole-client/tree/master/guacamole-ext/src/main/resources/org/apache/guacamole/protocols
type GuacConnectionParameters struct {
	// rdp & kubernetes & ssh & telnet & vnc
	Port                   string `json:"port"`
	RecordingName          string `json:"recording-name"`
	RecordingExcludeMouse  string `json:"recording-exclude-mouse"`
	Hostname               string `json:"hostname"`
	ReadOnly               string `json:"read-only"`
	RecordingIncludeKeys   string `json:"recording-include-keys"`
	RecordingExcludeOutput string `json:"recording-exclude-output"`
	RecordingPath          string `json:"recording-path"`
	CreateRecordingPath    string `json:"create-recording-path"`

	// kubernetes
	TypescriptName       string `json:"typescript-name"`
	ScrollBack           string `json:"scrollback"`
	Pod                  string `json:"pod"`
	ExecCommand          string `json:"exec-command"`
	Backspace            string `json:"backspace"`
	UseSSL               string `json:"use-ssl"`
	CaCert               string `json:"ca-cert"`
	TypescriptPath       string `json:"typescript-path"`
	FontName             string `json:"font-name"`
	Namespace            string `json:"namespace"`
	CreateTypescriptPath string `json:"create-typescript-path"`
	FontSize             string `json:"font-size"`
	ClientCert           string `json:"client-cert"`
	ColorScheme          string `json:"color-scheme"`
	ClientKey            string `json:"client-key"`
	Container            string `json:"container"`
	IgnoreCert           string `json:"ignore-cert"`

	// rdp
	SftpRootDirectory        string `json:"sftp-root-directory"`
	SftpPassphrase           string `json:"sftp-passphrase"`
	ForceLossless            string `json:"force-lossless"`
	Width                    string `json:"width"`
	SftpPassword             string `json:"sftp-password"`
	Height                   string `json:"height"`
	DisableGlyphCaching      string `json:"disable-glyph-caching"`
	SftpUsername             string `json:"sftp-username"`
	WOLMACAddr               string `json:"wol-mac-addr"`
	SftpDisableUpload        string `json:"sftp-disable-upload"`
	RemoteApp                string `json:"remote-app"`
	DisableBitmapCaching     string `json:"disable-bitmap-caching"`
	DisableDownload          string `json:"disable-download"`
	EnableFontSmoothing      string `json:"enable-font-smoothing"`
	EnableWallpaper          string `json:"enable-wallpaper"`
	DPI                      string `json:"dpi"`
	SftpDirectory            string `json:"sftp-directory"`
	SftpPrivateKey           string `json:"sftp-private-key"`
	ServerLayout             string `json:"server-layout"`
	SftpPort                 string `json:"sftp-port"`
	EnableDrive              string `json:"enable-drive"`
	GatewayPassword          string `json:"gateway-password"`
	CertTOFU                 string `json:"cert-tofu"`
	WolUdpPort               string `json:"wol-udp-port"`
	PrinterName              string `json:"printer-name"`
	DrivePath                string `json:"drive-path"`
	GatewayHostname          string `json:"gateway-hostname"`
	GatewayPort              string `json:"gateway-port"`
	Timezone                 string `json:"timezone"`
	DisableUpload            string `json:"disable-upload"`
	ClientName               string `json:"client-name"`
	Password                 string `json:"password"`
	EnableSFTP               string `json:"enable-sftp"`
	StaticChannels           string `json:"static-channels"`
	SftpServerAliveInterval  string `json:"sftp-server-alive-interval"`
	ResizeMethod             string `json:"resize-method"`
	NormalizeClipboard       string `json:"normalize-clipboard"`
	LoadBalanceInfo          string `json:"load-balance-info"`
	SftpHostKey              string `json:"sftp-host-key"`
	EnableTheming            string `json:"enable-theming"`
	Console                  string `json:"console"`
	DisableCopy              string `json:"disable-copy"`
	DisableAuth              string `json:"disable-auth"`
	DisablePaste             string `json:"disable-paste"`
	EnableFullWindowDrag     string `json:"enable-full-window-drag"`
	DisableGFX               string `json:"disable-gfx"`
	ConsoleAudio             string `json:"console-audio"`
	GatewayUsername          string `json:"gateway-username"`
	CertFingerprints         string `json:"cert-fingerprints"`
	EnableAudioInput         string `json:"enable-audio-input"`
	WOLSendPacket            string `json:"wol-send-packet"`
	DisableOffscreenCaching  string `json:"disable-offscreen-caching"`
	EnablePrinting           string `json:"enable-printing"`
	PreConnectionBlob        string `json:"preconnection-blob"`
	DriveName                string `json:"drive-name"`
	Domain                   string `json:"domain"`
	WOLBroadcastAddr         string `json:"wol-broadcast-addr"`
	PreConnectionID          string `json:"preconnection-id"`
	ColorDepth               string `json:"color-depth"`
	EnableTouch              string `json:"enable-touch"`
	SftpDisableDownload      string `json:"sftp-disable-download"`
	Security                 string `json:"security"`
	InitialProgram           string `json:"initial-program"`
	RemoteAppDir             string `json:"remote-app-dir"`
	Username                 string `json:"username"`
	CreateDrivePath          string `json:"create-drive-path"`
	RemoteAppArgs            string `json:"remote-app-args"`
	GatewayDomain            string `json:"gateway-domain"`
	DisableAudio             string `json:"disable-audio"`
	RecordingExcludeTouch    string `json:"recording-exclude-touch"`
	EnableMenuAnimations     string `json:"enable-menu-animations"`
	SftpHostname             string `json:"sftp-hostname"`
	WOLWaitTime              string `json:"wol-wait-time"`
	EnableDesktopComposition string `json:"enable-desktop-composition"`

	// ssh
	PrivateKey          string `json:"private-key"`
	ServerAliveInterval string `json:"server-alive-interval"`
	Command             string `json:"command"`
	Locale              string `json:"locale"`
	TerminalType        string `json:"terminal-type"`

	// telnet
	LoginSuccessRegex string `json:"login-success-regex"`
	UsernameRegex     string `json:"username-regex"`
	LoginFailureRegex string `json:"login-failure-regex"`
	PasswordRegex     string `json:"password-regex"`

	// vnc
	Cursor          string `json:"cursor"`
	Encodings       string `json:"encodings"`
	AudioServername string `json:"audio-servername"`
	DestPort        string `json:"dest-port"`
	SwapRedBlue     string `json:"swap-red-blue"`
	DestHost        string `json:"dest-host"`
}
