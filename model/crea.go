package models

import (
	"time"
)

type CreaData struct {
	ID            int64     `xorm:"pk autoincr" json:"id"`
	PrfCadCodRnp  string    `json:"prfCadCodRnp"`
	CrtCpf        string    `json:"crtCpf"`
	PrfCrtCod     int64     `json:"prfCrtCod"`
	TpoSit        string    `json:"tpoSit"`
	DtaEms        string    `json:"dtaEms"`
	DtaVal        string    `json:"dtaVal"`
	DtaTpoSit     string    `json:"dtaTpoSit"`
	DtaAlt        string    `json:"dtaAlt"`
	SisUsuLgn     string    `json:"sisUsuLgn"`
	MalCrtTpoSit  int       `json:"malCrtTpoSit"`
	DtaMalTpoSit  string    `json:"dtaMalTpoSit"`
	MalTpoSit     int       `json:"malTpoSit"`
	MalCrtCodRet  string    `json:"malCrtCodRet"`
	MalCrtIndErro string    `json:"malCrtIndErro"`
	IndEtg        string    `json:"indEtg"`
	DtaEtg        string    `json:"dtaEtg"`
	CanMot        int       `json:"canMot"`
	IndDev        string    `json:"indDev"`
	DtaDev        string    `json:"dtaDev"`
	TpoEntrada    string    `json:"tpoEntrada"`
	CrtNroProt    string    `json:"crtNroProt"`
	CrtNmeL1      string    `json:"crtNmeL1"`
	CrtNmeL2      string    `json:"crtNmeL2"`
	CrtNmePai     string    `json:"crtNmePai"`
	CrtNmeMae     string    `json:"crtNmeMae"`
	CrtRgOrgEmsUf string    `json:"crtRgOrgEmsUf"`
	CrtTpoSng     string    `json:"crtTpoSng"`
	CrtDtaNsc     string    `json:"crtDtaNsc"`
	CrtNat        string    `json:"crtNat"`
	CrtUfnat      string    `json:"crtUfnat"`
	CrtNac        string    `json:"crtNac"`
	CrtCreaReg    string    `json:"crtCreaReg"`
	CrtNroRegCrtM string    `json:"crtNroRegCrtM"`
	CrtTpoVal     string    `json:"crtTpoVal"`
	CrtDtaVal     string    `json:"crtDtaVal"`
	CrtTit1       string    `json:"crtTit1"`
	CrtTit2       string    `json:"crtTit2"`
	CrtTit3       string    `json:"crtTit3"`
	CrtCodReem    string    `json:"crtCodReem"`
	CrtUfcreaSol  string    `json:"crtUfcreaSol"`
	CrtIspNme     string    `json:"crtIspNme"`
	CrtRnp        string    `json:"crtRnp"`
	CrtDtaEnv     time.Time `json:"crtDtaEnv"`
	CanJust       string    `json:"canJust"`
	CrtTit4       string    `json:"crtTit4"`
	NroPisPasep   string    `json:"nroPisPasep"`
	FlgDoador     bool      `json:"flgDoador"`
	TpoCarteira   string    `json:"tpoCarteira"`
	CrtNmeSocial  string    `json:"crtNmeSocial"`
}

type ShareData struct {
	ID            int64  `xorm:"pk autoincr" json:"id"`
	DocumentPhoto string `json:"document_photo"`
	CodeCrea      int    `json:"code_crea"`
	Email         string `json:"email"`
	Senha         string `json:"senha"`
	GovIntegred   bool   `json:"gov_integred"`
	MecIntegred   bool   `json:"mec_integred"`
}
