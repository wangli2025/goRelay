package pipeserver_main

type Config struct {
	ListenPipeServerAddr string   `json:"listen_pipe_server_addr"`
	WhiteIpList          []string `json:"white_ip_list"`
	BlackIpList          []string `json:"black_ip_list"`
	DebugLog             bool     `json:"debug_log"`
}
