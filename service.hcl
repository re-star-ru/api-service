job "api-service" {
	datacenters = ["restar"] // change for data center
	type = "service"

	group "default" {
		network {
			port "api-service" { host_network = "private" }
		}

		task "api-service" {
			driver = "docker"

			// client - reproxy
			service {
				port = "api-service"
				tags = [
					"reproxy.enabled=1",
					"reproxy.server=api.re-star.ru",
					"reproxy.route=/v1/stand/"
				]
			}

			resources {
				cpu = 100
				memory = 64
			}

			config {
				image = "ghcr.io/[[.repo]]:[[.tag]]"
				network_mode = "host"
			}

			env {
				PORT = "NOMAD_ADDR_api-service" // or api_service
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