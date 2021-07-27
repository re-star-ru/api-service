job "api-service" {
	datacenters = ["restar"] // change for data center
	type = "service"

	group "default" {
		network {
//			port "http" { static = 80 }
//			port "https" { static = 443 }
		}

		task "api-service" {
			driver = "docker"

			volume_mount {
				volume = "certs"
				destination = "/etc/letsencrypt"
				read_only   = true
			}

			config {
				image = "ghcr.io/umputun/reproxy"
				network_mode = "host"
			}

			env {
				PORT = "8080"
			}

			resources {
				cpu = 200
				memory = 128
			}

//			service {
//				name = "global-redis-check"
//				tags = ["global", "cache"]
//				port = "db"
//				check {
//					name = "alive"
//					type = "tcp"
//					interval = "10s"
//					timeout = "2s"
//				}
//			}
		}
	}
}