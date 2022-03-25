package main

import (
	"context"
	pb "github.com/enginewang/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/enginewang/shippy/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
	"log"
)

type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type ConsignmentRepository struct {
	consignments []*pb.Consignment
}

// Create a new consignment
func (repo *ConsignmentRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *ConsignmentRepository) GetAll() []*pb.Consignment {
	return repo.consignments
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type consignmentService struct {
	repo         repository
	vesselClient vesselProto.VesselService
}

// CreateConsignment - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *consignmentService) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	vesselResponse, err := s.vesselClient.FindAvailable(context.Background(), &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	req.VesselId = vesselResponse.Vessel.Id
	log.Printf("Found vessel: %s \n", vesselResponse.Vessel.Name)
	// Save our consignment
	consignment, err := s.repo.Create(req)
	if err != nil {
		return err
	}
	res.Created = true
	res.Consignment = consignment
	return nil
}

func (s *consignmentService) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consignments = consignments
	return nil
}

func main() {
	repo := &ConsignmentRepository{}
	service := micro.NewService(
		micro.Name("shippy.service.consignment"),
		micro.Version("lastest"),
	)
	service.Init()
	vesselClient := vesselProto.NewVesselService("shippy.service.vessel", service.Client())
	if err := pb.RegisterShippingServiceHandler(service.Server(),
		&consignmentService{repo, vesselClient}); err != nil {
		log.Panic(err)
	}
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
