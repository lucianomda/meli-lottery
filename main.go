package main

import (
	"fmt"
	"math/rand"
	"time"

	_ "github.com/mercadolibre/fury_credits-charges/src/api/configs"
)

func main() {
	names := []string{
		"estefania.ellena@mercadolibre.com",
		"mauricio.frias@mercadolibre.com",
		"lucas.pecchio@mercadolibre.com",
		"pamela.ferreyra@mercadolibre.com",
		"gimena.martinez@mercadolibre.com",
		"luciano.ferrari@mercadolibre.com",
		"matias.alaniz@mercadolibre.com",
		"ignacio.dalvarez@mercadolibre.com",
		"maria.giglio@mercadolibre.com",
		"leandro.bagur@mercadolibre.com",
		"juan.niza@mercadolibre.com",
		"luciana.guevara@mercadolibre.com",
		"jorge.jimenez@mercadolibre.com",
		"maximiliano.bustos@mercadolibre.com",
		"emiliano.giraudo@mercadolibre.com",
		"gpablo.diaz@mercadolibre.com",
		"nicolas.rivas@mercadolibre.com",
		"rodrigo.torres@mercadolibre.com",
		"alanjulian.deangelis@mercadolibre.com",
		"rodrigo.lopez@mercadolibre.com",
		"matias.fariasfalkiewicz@mercadolibre.com",
		"hernan.malatini@mercadolibre.com",
		"lucas.hernandez@mercadolibre.com",
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })
	for i, names := range names {
		println(fmt.Sprintf("%03d - %s", i + 1, names))
	}
}