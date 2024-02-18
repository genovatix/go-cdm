package utils

type NetConfig struct {
	Ver            uint32
	Code           []byte
	Global         *GlobalConfig
	BootstrapPeers map[int]string
}

const VERSION uint32 = 1

type GlobalConfig struct {
	Name        string
	RootAddress string
	Mode        string
}

func CreateNetworkConfiguration() *NetConfig {
	nc := &NetConfig{}
	nc.Ver = VERSION
	nc.Code = GetChainCode()
	nc.Global = &GlobalConfig{
		Name:        "TrustMesh",
		RootAddress: "localhost:3001",
		Mode:        "dev_mode",
	}
	nc.BootstrapPeers = make(map[int]string)
	nc.BootstrapPeers[1] = "localhost:3002"
	return nc
}

func WriteConfigToFile(filename string, data []byte) {

}

type ChainCode []byte

func (cc ChainCode) Bytes() []byte {
	return cc
}

func (cc ChainCode) Address() interface{} {
	return nil
}
