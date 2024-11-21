// TODO: This package is copied from the ror monorepo/internnal have to check is it is shared ands hould be put back in pkg
package vaultmodels

type VaultClusterModel struct {
	Data Data `json:"data"`
}

type Data struct {
	RorClientSecret string `json:"rorClientSecret"`
}

type VaultDexModel struct {
	Data DexData `json:"data"`
}

type DexData struct {
	DexSecret string `json:"dexSecret"`
}
