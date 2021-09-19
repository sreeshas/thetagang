package config

type Config struct {
	SymbolMap    map[string]float64 `json:"symbolpricemap"`
	TotalCapital float64            `json:"totalcapital"`
	OutputFormat Output   			`json:"outputFormat"`

}

type Output struct {
	GroupByDate bool `json:"groupbydate"`
	GroupBySymbol bool `json:"groupbysymbol"`
}
