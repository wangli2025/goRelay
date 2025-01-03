package relayclient_main

var (
	configFile string
)

type Config struct {
	ListenRelayClientAddr string       `json:"listen_relay_client_addr"`
	WhiteIpList           []string     `json:"white_ip_list"`
	DebugLog              bool         `json:"debug_log"`
	RealServerInfo        []RealServer `json:"realServerInfo"`
}

type RealServer struct {
	ID             string `json:"id"`
	RealServerAddr string `json:"real_Server_Addr"`
}
