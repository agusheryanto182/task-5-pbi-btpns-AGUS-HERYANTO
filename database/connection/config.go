package connection

import (
	"log"

	"github.com/agusheryanto182/task-5-pbi-btpns-AGUS-HERYANTO/models"
	"github.com/joeshaw/envdecode"
)

func NewConfig() *models.Global {
	var c models.Global
	if err := envdecode.StrictDecode(&c); err != nil {
		log.Fatalf("Failed to decode: %s", err)
	}

	return &c
}
