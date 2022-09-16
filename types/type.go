package types

type Status struct {
	Status    uint32 `json:"status"`
	NetworkId uint32 `json:"network_id"`
}

type AccountAsset struct {
	Id      uint32 `json:"id"`
	Name    string `json:"name"`
	Balance string `json:"balance"`
	Price   string `json:"price"`
}

type AccountLp struct {
	Index  uint32 `json:"index"`
	Amount string `json:"amount"`
}

type Account struct {
	Status uint32          `json:"status"`
	Index  int64           `json:"index"`
	Name   string          `json:"name"`
	Pk     string          `json:"pk"`
	Nonce  int64           `json:"nonce"`
	Assets []*AccountAsset `json:"assets"`
	Lps    []*AccountLp    `json:"lps"`
}

type SimpleAccount struct {
	Index int64  `json:"index"`
	Name  string `json:"name"`
	Pk    string `json:"pk"`
}

type Accounts struct {
	Total    uint32           `json:"total"`
	Accounts []*SimpleAccount `json:"accounts"`
}

type Asset struct {
	Id         uint32 `json:"id"`
	Name       string `json:"name"`
	Decimals   uint32 `json:"decimals"`
	Symbol     string `json:"symbol"`
	Address    string `json:"address"`
	Price      string `json:"price"`
	IsGasAsset uint32 `json:"is_gas_asset"`
}

type Assets struct {
	Total  uint32   `json:"total"`
	Assets []*Asset `json:"assets"`
}

type Block struct {
	Commitment                      string `json:"commitment"`
	Height                          int64  `json:"height"`
	StateRoot                       string `json:"state_root"`
	PriorityOperations              int64  `json:"priority_operations"`
	PendingOnChainOperationsHash    string `json:"pending_on_chain_operations_hash"`
	PendingOnChainOperationsPubData string `json:"pending_on_chain_operations_pub_data"`
	CommittedTxHash                 string `json:"committed_tx_hash"`
	CommittedAt                     int64  `json:"committed_at"`
	VerifiedTxHash                  string `json:"verified_tx_hash"`
	VerifiedAt                      int64  `json:"verified_at"`
	Txs                             []*Tx  `json:"txs"`
	Status                          int64  `json:"status"`
}

type Blocks struct {
	Total  uint32   `json:"total"`
	Blocks []*Block `json:"blocks"`
}

type CurrentHeight struct {
	Height int64 `json:"height"`
}

type ContractAddress struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Layer2BasicInfo struct {
	BlockCommitted            int64             `json:"block_committed"`
	BlockVerified             int64             `json:"block_verified"`
	TotalTransactionCount     int64             `json:"total_transaction_count"`
	YesterdayTransactionCount int64             `json:"yesterday_transaction_count"`
	TodayTransactionCount     int64             `json:"today_transaction_count"`
	YesterdayActiveUserCount  int64             `json:"yesterday_active_user_count"`
	TodayActiveUserCount      int64             `json:"today_active_user_count"`
	ContractAddresses         []ContractAddress `json:"contract_addresses"`
}

type GasFee struct {
	GasFee string `json:"gas_fee"`
}

type GasAccount struct {
	Status int64  `json:"status"`
	Index  int64  `json:"index"`
	Name   string `json:"name"`
}

type GasFeeAssets struct {
	Assets []Asset `json:"assets"`
}

type Search struct {
	DataType int32 `json:"data_type"`
}

type SwapAmount struct {
	AssetId     uint32 `json:"asset_id"`
	AssetName   string `json:"asset_name"`
	AssetAmount string `json:"asset_amount"`
}

type Pair struct {
	Index         uint32 `json:"index"`
	AssetAId      uint32 `json:"asset_a_id"`
	AssetAName    string `json:"asset_a_name"`
	AssetAAmount  string `json:"asset_a_amount"`
	AssetAPrice   string `json:"asset_a_price"`
	AssetBId      uint32 `json:"asset_b_id"`
	AssetBName    string `json:"asset_b_name"`
	AssetBPrice   string `json:"asset_b_price"`
	AssetBAmount  string `json:"asset_b_amount"`
	FeeRate       int64  `json:"fee_rate"`
	TreasuryRate  int64  `json:"treasury_rate"`
	TotalLpAmount string `json:"total_lp_amount"`
}

type Pairs struct {
	Pairs []*Pair `json:"pairs"`
}

type LpValue struct {
	AssetAId     uint32 `json:"asset_a_id"`
	AssetAName   string `json:"asset_a_name"`
	AssetAPrice  string `json:"asset_a_price"`
	AssetAAmount string `json:"asset_a_amount"`
	AssetBId     uint32 `json:"asset_b_id"`
	AssetBName   string `json:"asset_b_name"`
	AssetBPrice  string `json:"asset_b_price"`
	AssetBAmount string `json:"asset_b_amount"`
}

type Tx struct {
	Hash          string `json:"hash"`
	Type          int64  `json:"type,range=[1:64]"`
	Amount        string `json:"amount"`
	Info          string `json:"info"`
	Status        int64  `json:"status"`
	Index         int64  `json:"index"`
	GasFeeAssetId int64  `json:"gas_fee_asset_id"`
	GasFee        string `json:"gas_fee"`
	NftIndex      int64  `json:"nft_index"`
	CollectionId  int64  `json:"collection_id"`
	PairIndex     int64  `json:"pair_index"`
	AssetId       int64  `json:"asset_id"`
	AssetName     string `json:"asset_name"`
	NativeAddress string `json:"native_address"`
	ExtraInfo     string `json:"extra_info"`
	Memo          string `json:"memo"`
	AccountIndex  int64  `json:"account_index"`
	AccountName   string `json:"account_name"`
	Nonce         int64  `json:"nonce"`
	ExpiredAt     int64  `json:"expire_at"`
	BlockHeight   int64  `json:"block_height"`
	CreatedAt     int64  `json:"created_at"`
	StateRoot     string `json:"state_root"`
}

type Txs struct {
	Total uint32 `json:"total"`
	Txs   []*Tx  `json:"txs"`
}

type TxHash struct {
	TxHash string `json:"tx_hash"`
}

type NextNonce struct {
	Nonce uint64 `json:"nonce"`
}

type EnrichedTx struct {
	Tx
	CommittedAt int64 `json:"committed_at"`
	VerifiedAt  int64 `json:"verified_at"`
	ExecutedAt  int64 `json:"executed_at"`
	AssetAId    int64 `json:"asset_a_id"`
	AssetBId    int64 `json:"asset_b_id"`
}

type MaxOfferId struct {
	OfferId uint64 `json:"offer_id"`
}

type Nft struct {
	Index               int64  `json:"index"`
	CreatorAccountIndex int64  `json:"creator_account_index"`
	CreatorAccountName  string `json:"creator_account_name"`
	OwnerAccountIndex   int64  `json:"owner_account_index"`
	OwnerAccountName    string `json:"owner_account_name"`
	ContentHash         string `json:"content_hash"`
	L1Address           string `json:"l1_address"`
	L1TokenId           string `json:"l1_token_id"`
	CreatorTreasuryRate int64  `json:"creator_treasury_rate"`
	CollectionId        int64  `json:"collection_id"`
}

type Nfts struct {
	Total int64  `json:"total"`
	Nfts  []*Nft `json:"nfts"`
}
