package rpc

import (
	"context"
	"strings"

	"github.com/LanfordCai/ava/api/rpc/addressvalidator"
	pb "github.com/LanfordCai/ava/api/rpc/addressvalidator"
	"github.com/LanfordCai/ava/api/utils"
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
	chain := strings.ToLower(in.GetChain())
	addr := in.GetAddress()
	v := utils.GetValidatorByChain(chain)
	if v == nil {
		return nil, ErrUnsupportedChain
	}

	network := validator.Mainnet
	if in.GetIsTestnet() {
		network = validator.Testnet
	}

	r := v.ValidateAddress(addr, network)

	isValid := r.IsValid
	validateStatus := string(r.Status)
	if isValid {
		isValid = !utils.IsUnsupportedAddressType(chain, r.Type)
	}

	if isValid && utils.NeedCheckContractAddress(chain) && !utils.AddressInContractWhiteList(chain, addr) {
		if c, ok := v.(validator.ContractChecker); ok {
			contractDeployed, err := c.IsContractDeployed(addr)
			if err != nil {
				validateStatus = string(validator.Failure)
				isValid = false
			}

			if contractDeployed {
				isValid = false
			}
		}
	}

	resp := &pb.ValidateResponse{
		IsValid: isValid,
		Type:    string(r.Type),
		Msg:     r.Msg,
		Status:  validateStatus,
	}

	return resp, nil
}
