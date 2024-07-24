package services

import (
	"bytes"
	"encoding/json"
	"errors"
	models "fast/model"
	"fmt"
	"net/http"
)

type CreaResponse struct {
	PrfCadCodRnp  string      `json:"prfCadCodRnp"`
	AnoPag        int         `json:"anoPag"`
	TpoSit        string      `json:"tpoSit"`
	DtaAlt        string      `json:"dtaAlt"`
	CreCadCod     string      `json:"creCadCod"`
	SisUsuLgn     string      `json:"sisUsuLgn"`
	QtdParcelas   int         `json:"qtdParcelas"`
	Valor         float64     `json:"valor"`
	ValorRecebido float64     `json:"valorRecebido"`
	DtaPag        string      `json:"dtaPag"`
	NroBoleto     string      `json:"nroBoleto"`
	DscObservacao string      `json:"dscObservacao"`
	Parcelamentos interface{} `json:"parcelamentos"`
}

func SendCreaRequest(prfCadCodRnp string) (*CreaResponse, error) {
	apiURL := fmt.Sprintf("https://hackathon.teste.confea.org.br/Anuidades/BuscaPorRNP?prfCadCodRnp=%s", prfCadCodRnp)
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get a valid response from CREA API")
	}

	var creaResponse CreaResponse
	if err := json.NewDecoder(resp.Body).Decode(&creaResponse); err != nil {
		return nil, err
	}

	return &creaResponse, nil
}

func SendCreaRequestPost(creaData models.CreaData) (*models.CreaData, error) {
	apiURL := "https://hackathon.teste.confea.org.br/Carteiras/Listar"
	jsonData, _ := json.Marshal(creaData)
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get a valid response from CREA API")
	}

	var creaResponse models.CreaData
	if err := json.NewDecoder(resp.Body).Decode(&creaResponse); err != nil {
		return nil, err
	}

	return &creaResponse, nil
}
