package main

type SrConfig struct {
	OutputDir   string
	ContractID  string
	AccessToken string
	EndPoint    string
	TableNames  []string
	Conditions  map[string]*string
	FileFormat  string
}
