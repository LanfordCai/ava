package rpc

import (
	"context"

	"github.com/LanfordCai/ava/api/utils"
	"github.com/LanfordCai/ava/proto/addressvalidator"
	pb "github.com/LanfordCai/ava/proto/addressvalidator"
	"github.com/LanfordCai/ava/validator"
	"google.golang.org/grpc"
)

// AddressValidator ...
type AddressValidator struct {
	addressvalidator.UnimplementedAddressValidatorServer
}

// NewAddressValidatorServer ...
func NewAddressValidatorServer() *grpc.Server {
	s := grpc.NewServer()

	pb.RegisterAddressValidatorServer(s, &AddressValidator{})
	return s
}

// ValidateAddress ...
func (s *AddressValidator) ValidateAddress(ctx context.Context, in *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	v := utils.GetValidatorByChain(in.GetChain())
	if v == nil {
		return nil, ErrUnsupportedChain
	}

	network := validator.Mainnet
	if in.GetIsTestnet() {
		network = validator.Testnet
	}

	r := v.ValidateAddress(in.GetAddress(), network)

	isValid := r.IsValid
	unsupportedTypes := utils.GetUnsupportedAddressTypes(in.GetChain())
	if utils.Contains(unsupportedTypes, string(r.Type)) {
		isValid = false
	}

	resp := &pb.ValidateResponse{
		IsValid: isValid,
		Type:    string(r.Type),
		Msg:     r.Msg,
		Status:  string(r.Status),
	}

	return resp, nil
}

// IsContractAddress ...
func (s *AddressValidator) IsContractAddress(ctx context.Context, in *pb.ValidateRequest) (*pb.ContractCheckResponse, error) {
	return nil, nil
}
