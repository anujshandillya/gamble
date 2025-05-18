package types

type VerificationData struct {
	HashInput      string  `json:"hashInput"`
	Hash           string  `json:"hash"`
	UsedServerSeed *string `json:"usedServerSeed"`
}

type BetResultResponseLimbo struct {
	Currency     string           `json:"currency"`
	Result       float64          `json:"result"`
	Outcome      string           `json:"outcome"`
	Payout       float64          `json:"payout"`
	Wager        float64          `json:"wager"`
	Profit       float64          `json:"profit"`
	Nonce        int              `json:"nonce"`
	ClientSeed   string           `json:"clientSeed"`
	Verification VerificationData `json:"verification"`
}
type BetResultResponseDice struct {
	Currency     string           `json:"currency"`
	Result       float64          `json:"result"`
	Outcome      string           `json:"outcome"`
	OverUnder    string           `json:"overUnder"`
	Payout       float64          `json:"payout"`
	Wager        float64          `json:"wager"`
	Profit       float64          `json:"profit"`
	Nonce        int              `json:"nonce"`
	ClientSeed   string           `json:"clientSeed"`
	Verification VerificationData `json:"verification"`
}

type BetResultResponseWheel struct {
	Currency     string           `json:"currency"`
	Result       float64          `json:"result"`
	Index        int              `json:"index"`
	Outcome      string           `json:"outcome"`
	OverUnder    string           `json:"overUnder"`
	Payout       float64          `json:"payout"`
	Wager        float64          `json:"wager"`
	Profit       float64          `json:"profit"`
	Nonce        int              `json:"nonce"`
	ClientSeed   string           `json:"clientSeed"`
	Verification VerificationData `json:"verification"`
}
