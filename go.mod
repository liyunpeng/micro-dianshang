module micro

go 1.13

require (
	github.com/armon/go-metrics v0.3.0 // indirect
	github.com/cenkalti/backoff/v3 v3.1.1 // indirect
	github.com/cloudflare/cloudflare-go v0.10.9 // indirect
	github.com/coreos/etcd v3.3.18+incompatible // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/hashicorp/go-immutable-radix v1.1.0 // indirect
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/serf v0.8.5 // indirect
	github.com/jinzhu/gorm v1.9.11
	github.com/klauspost/cpuid v1.2.2 // indirect
	github.com/lucas-clemente/quic-go v0.14.1 // indirect
	github.com/micro/go-grpc v1.0.1 // indirect
	github.com/micro/go-log v0.1.0
	github.com/micro/go-micro v1.17.1
	github.com/micro/go-plugins v1.5.1
	github.com/micro/micro v1.17.1 // indirect
	github.com/miekg/dns v1.1.24 // indirect
	github.com/nats-io/jwt v0.3.2 // indirect
	github.com/olekukonko/tablewriter v0.0.4 // indirect
	github.com/onsi/ginkgo v1.10.3 // indirect
	go.uber.org/atomic v1.5.1 // indirect
	go.uber.org/multierr v1.4.0 // indirect
	go.uber.org/zap v1.13.0 // indirect
	golang.org/x/lint v0.0.0-20191125180803-fdd1cda4f05f // indirect
	golang.org/x/net v0.0.0-20191207000613-e7e4b65ae663 // indirect
	golang.org/x/sys v0.0.0-20191206220618-eeba5f6aabab // indirect
	golang.org/x/tools v0.0.0-20191206204035-259af5ff87bd // indirect
	golang.org/x/vgo v0.0.0-20180912184537-9d567625acf4 // indirect
	google.golang.org/genproto v0.0.0-20191206224255-0243a4be9c8f // indirect
	gopkg.in/square/go-jose.v2 v2.4.0 // indirect
	shopping/user/handler v0.0.0-00010101000000-000000000000
	shopping/user/model v0.0.0-00010101000000-000000000000
	shopping/user/proto/user v0.0.0-00010101000000-000000000000
	shopping/user/repository v0.0.0-00010101000000-000000000000
)

replace (
	github.com/golang/lint => F:\GoWorkSpace\micro-dianshang\vendor\golang.org\x\lint
	github.com/testcontainers/testcontainer-go => F:\GoWorkSpace\testcontainer\testcontainers-go
	shopping/user/handler => F:\GoWorkSpace\micro-dianshang\shopping\handler
	shopping/user/model => F:\GoWorkSpace\micro-dianshang\shopping\model
	shopping/user/proto/user => F:\GoWorkSpace\micro-dianshang\shopping\proto\user
	shopping/user/repository => F:\GoWorkSpace\micro-dianshang\shopping\repository
)
