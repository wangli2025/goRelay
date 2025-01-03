package relayserver_main

var (
	configFile string
)

type Config struct {
	Id                    string   `json:"id"`
	PipeServerAddr        string   `json:"pipe_server_addr"`
	ListenRelayServerAddr string   `json:"listen_relay_server_addr"`
	WhiteIpList           []string `json:"white_ip_list"`
	DebugLog              bool     `json:"debug_log"`
}
