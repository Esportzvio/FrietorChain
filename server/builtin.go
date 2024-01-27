package server

import (
	"github.com/esportzvio/frietorchain/chain"
	"github.com/esportzvio/frietorchain/consensus"
	consensusDev "github.com/esportzvio/frietorchain/consensus/dev"
	consensusDummy "github.com/esportzvio/frietorchain/consensus/dummy"
	consensusIBFT "github.com/esportzvio/frietorchain/consensus/ibft"
	consensusPolyBFT "github.com/esportzvio/frietorchain/consensus/polybft"
	"github.com/esportzvio/frietorchain/forkmanager"
	"github.com/esportzvio/frietorchain/secrets"
	"github.com/esportzvio/frietorchain/secrets/awsssm"
	"github.com/esportzvio/frietorchain/secrets/gcpssm"
	"github.com/esportzvio/frietorchain/secrets/hashicorpvault"
	"github.com/esportzvio/frietorchain/secrets/local"
	"github.com/esportzvio/frietorchain/state"
)

type GenesisFactoryHook func(config *chain.Chain, engineName string) func(*state.Transition) error

type ConsensusType string

type ForkManagerFactory func(forks *chain.Forks) error

type ForkManagerInitialParamsFactory func(config *chain.Chain) (*forkmanager.ForkParams, error)

const (
	DevConsensus     ConsensusType = "dev"
	IBFTConsensus    ConsensusType = "ibft"
	PolyBFTConsensus ConsensusType = consensusPolyBFT.ConsensusName
	DummyConsensus   ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:     consensusDev.Factory,
	IBFTConsensus:    consensusIBFT.Factory,
	PolyBFTConsensus: consensusPolyBFT.Factory,
	DummyConsensus:   consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

var genesisCreationFactory = map[ConsensusType]GenesisFactoryHook{
	PolyBFTConsensus: consensusPolyBFT.GenesisPostHookFactory,
}

var forkManagerFactory = map[ConsensusType]ForkManagerFactory{
	PolyBFTConsensus: consensusPolyBFT.ForkManagerFactory,
}

var forkManagerInitialParamsFactory = map[ConsensusType]ForkManagerInitialParamsFactory{
	PolyBFTConsensus: consensusPolyBFT.ForkManagerInitialParamsFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
