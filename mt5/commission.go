package mt5

type Commission struct {
	Name             string
	Description      string
	Path             string
	Mode             string
	RangeMode        string
	ChargeMode       string
	TurnoverCurrency string
	EntryMode        string
	ActionMode       string
	ProfitMode       string
	ReasonMode       string
	Tiers            []*CommissionTier
}

type CommissionTier struct {
	Mode      string
	Type      string
	Value     string
	Minimal   string
	Maximal   string
	RangeFrom string
	RangeTo   string
	Currency  string
}
