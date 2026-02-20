package config

import "os"

type Config struct {
	Port              string
	EthereumRPC       string
	PrivateKey        string
	IPFSEndpoint      string
	EncryptionEnabled bool
	ContractAddress   string
	ChainID           int64
}

func Load() *Config {
	return &Config{
		Port:              getEnv("PORT", "8080"),
		EthereumRPC:       os.Getenv("ETH_RPC_URL"),
		PrivateKey:        os.Getenv("PRIVATE_KEY"),
		IPFSEndpoint:      os.Getenv("IPFS_ENDPOINT"),
		EncryptionEnabled: getEnv("ENCRYPTED_ENABLED", "false") == "true",
		ContractAddress:   getEnv("CONTRACT_ADDRESS", ""),
		ChainID:           11155111,
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
