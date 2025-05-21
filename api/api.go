package api

import (
	"net/http"
	"fmt"
	"ci-environment/utils"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		province := r.URL.Query().Get("provincia")

		capital, err := utils.GetCapital(province)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		
		if province == "" {
			fmt.Fprint(w, "Por favor ingrese una provincia")
		} else if capital == "" {
			fmt.Fprint(w, "No se encontró la provincia ", province)
		} else {	
			fmt.Fprint(w, "La capital es ", capital)
		}
		
	} else {
		fmt.Fprintf(w, "Método http no válido")
	}
}