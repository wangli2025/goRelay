package pipeclient_main

var (
	configFile string
)

type Config struct {
	PipeServerAddr  string `json:"pipe_server_addr"`
	RelayClientAddr string `json:"relay_client_addr"`
	DebugLog        bool   `json:"debug_log"`
}
