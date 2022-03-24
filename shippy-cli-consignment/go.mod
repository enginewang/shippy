module github.com/enginewang/shippy/shippy-cli-consignment

replace github.com/enginewang/shippy/shippy-service-consignment => ../shippy-service-consignment

go 1.17

require (
	github.com/enginewang/shippy/shippy-service-consignment v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.45.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/micro/go-micro/v2 v2.9.1 // indirect
	github.com/miekg/dns v1.1.27 // indirect
	github.com/oxtoacart/bpool v0.0.0-20190530202638-03653db5a59c // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200523222454-059865788121 // indirect
	golang.org/x/text v0.3.2 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
