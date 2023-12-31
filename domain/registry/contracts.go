package registry

import "github.com/ethereum/go-ethereum/common"

type RegistryContracts struct {
	Dispatch            common.Address
	AgentRegistry       common.Address
	ScannerRegistry     common.Address
	ScannerPoolRegistry common.Address
	ScannerNodeVersion  common.Address
	FortaStaking        common.Address
	Forta               common.Address
	Migration           common.Address
	Rewards             common.Address
	StakeAllocator      common.Address
}
