package frp

type Common struct {
	server_addr        string
	server_port        int
	http_proxy         string
	log_file           string
	log_level          string
	log_max_days       int
	disable_log_color  bool
	token              string
	admin_addr         string
	admin_port         int
	admin_user         string
	admin_pwd          string
	assets_dir         string
	pool_count         int
	tcp_mux            bool
	user               bool
	login_fail_exit    bool
	protocol           string
	tls_enable         bool
	dns_server         string
	start              string
	heartbeat_interval int
	heartbeat_timeout  int
	meta_var1          int
	meta_var2          int
}
