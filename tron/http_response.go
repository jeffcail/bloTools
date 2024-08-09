package tron

type TransactionRes struct {
	Data    []*Transaction `json:"data"`
	Success bool           `json:"success"`
	Meta    struct {
		At          int64  `json:"at"`
		Fingerprint string `json:"fingerprint"`
		Links       struct {
			Next string `json:"next"`
		} `json:"links"`
		PageSize int64 `json:"page_size"`
	} `json:"meta"`
}

type Transaction struct {
	TransactionId string `json:"transaction_id"`
	TokenInfo     struct {
		Symbol   string `json:"symbol"`
		Address  string `json:"address"`
		Decimals int    `json:"decimals"`
		Name     string `json:"name"`
	} `json:"token_info"`
	BlockTimestamp int64  `json:"block_timestamp"`
	From           string `json:"from"`
	To             string `json:"to"`
	Type           string `json:"type"`
	Value          string `json:"value"`
}

type GetTrc10TokenPrecision struct {
	AssetIssue []*AssetIssue `json:"assetIssue"`
}
type AssetIssue struct {
	OwnerAddress string `json:"owner_address"`
	Name         string `json:"name"`
	Abbr         string `json:"abbr"`
	TotalSupply  int    `json:"total_supply"`
	TrxNum       int    `json:"trx_num"`
	Precision    int    `json:"precision"`
	Num          int    `json:"num"`
	StartTime    int64  `json:"start_time"`
	EndTime      int64  `json:"end_time"`
	Description  string `json:"description"`
	Url          string `json:"url"`
	Id           string `json:"id"`
}

type GetTransactionByID struct {
	Ret        interface{} `json:"ret"`
	Signature  interface{} `json:"signature"`
	TxID       string      `json:"txID"`
	RawData    *RawData2   `json:"raw_data"`
	RawDataHex string      `json:"raw_data_hex"`
}

type RawData2 struct {
	Data          string          `json:"data,omitempty"`
	Contract      []*ContractData `json:"contract"`
	RefBlockBytes string          `json:"ref_block_bytes"`
	RefBlockHash  string          `json:"ref_block_hash"`
	Expiration    int             `json:"expiration"`
	FeeLimit      int             `json:"fee_limit"`
	Timestamp     int             `json:"timestamp"`
}

type ContractData struct {
	Parameter *ParameterData `json:"parameter"`
	Type      string         `json:"type"`
}

type ParameterData struct {
	Value   *ValueData `json:"value"`
	TypeUrl string     `json:"type_url"`
}

type ValueData struct {
	Data            string `json:"data,omitempty"`
	Amount          int64  `json:"amount,omitempty"`
	AssetName       string `json:"asset_name,omitempty"`
	OwnerAddress    string `json:"owner_address"`
	ContractAddress string `json:"contract_address"`
}

type GetAssetIssueByID struct {
	OwnerAddress string `json:"owner_address"`
	Name         string `json:"name"`
	Abbr         string `json:"abbr"`
	TotalSupply  int64  `json:"total_supply"`
	TxrNum       int64  `json:"txr_num"`
	Precision    int64  `json:"precision"`
	Num          int64  `json:"num"`
	StartTime    int    `json:"start_time"`
	EndTime      int64  `json:"end_time"`
	Description  string `json:"description"`
	Url          string `json:"url"`
	Id           string `json:"id"`
}
