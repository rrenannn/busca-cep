package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type ViaCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8080", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	cepParam := r.URL.Query().Get("cep")
	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
	}
	cep, error := BuscaCep(cepParam)
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//result, error := json.Marshal(cep)
	//if error != nil {
	//	w.WriteHeader(http.StatusInternalServerError)
	//	return
	//}
	//w.Write(result) / Sistema utilizando Marshall

	json.NewEncoder(w).Encode(cep) // -> Pega o resultado do CEP e joga no Writer, escrevendo o res

	// Código para buscar o CEP
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func BuscaCep(cep string) (*ViaCEP, error) {
	// Código para buscar o CEP
	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return nil, error
	}
	// Fecha o código no final 
	defer resp.Body.Close()
	
	// Lê o corpo da resposta
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return nil, error
	}

	var c ViaCEP // -> Var que recebe a struct ViaCEP
	error = json.Unmarshal(body, &c) // -> Transforma json em Struct 
	if error != nil {
		return nil, error
	}
	// Retorna o valor correto
	return &c, nil
}