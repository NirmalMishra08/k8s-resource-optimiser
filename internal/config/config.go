package config

type Config struct {
	CPUBuffer        float64
	MemoryBuffer     float64
	MinCPU           int
	MinMemory        int
}

func DefaultConfig() Config {
	return Config{
		CPUBuffer:    1.5,
		MemoryBuffer: 1.3,
		MinCPU:       100,
		MinMemory:    256,
	}
}