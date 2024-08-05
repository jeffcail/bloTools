package tron

type AccountResourceMessage struct {
	FreeNetUsed          int64
	FreeNetLimit         int64
	NetUsed              int64
	NetLimit             int64
	AssetNetUsed         map[string]int64
	AssetNetLimit        map[string]int64
	TotalNetLimit        int64
	TotalNetWeight       int64
	TotalTronPowerWeight int64
	TronPowerUsed        int64
	TronPowerLimit       int64
	EnergyUsed           int64
	EnergyLimit          int64
	TotalEnergyLimit     int64
	TotalEnergyWeight    int64
	StorageUsed          int64
	StorageLimit         int64
}

type DelegatedResourceList struct {
	From                   string `json:"from"`
	To                     string `json:"to"`
	FrozenBalanceForEnergy int64  `json:"frozen_balance_for_energy"`
	ExpireTimeForEnergy    int64  `json:"expire_time_for_energy"`
}
