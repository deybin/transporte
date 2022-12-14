package tables

import "transporte/src/models"

func Clientes_GetSchema() ([]models.Base, string) {
	var Clientes []models.Base
	tableName := "Clientes"
	Clientes = append(Clientes, models.Base{
		Name:        "c_docu",
		Description: "c_docu",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Min:  2,
			Max:  2,
		},
	})
	Clientes = append(Clientes, models.Base{
		Name:        "n_docu",
		Description: "n_docu",
		Required:    true,
		Important:   true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Number_DB(),
			Min:  1,
			Max:  11,
		},
	})
	Clientes = append(Clientes, models.Base{
		Name:        "l_clie",
		Description: "l_clie",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       3,
			Max:       100,
			UpperCase: true,
		},
	})
	Clientes = append(Clientes, models.Base{
		Name:        "k_gene",
		Description: "k_gene",
		Type:        "uint64",
		Uint: models.Uints{
			Max: 10,
		},
	})
	Clientes = append(Clientes, models.Base{
		Name:        "f_naci",
		Description: "f_naci",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Date: true,
		},
	})
	Clientes = append(Clientes, models.Base{
		Name:        "l_dire",
		Description: "l_dire",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       20,
			Max:       400,
			UpperCase: true,
		},
	})
	Clientes = append(Clientes, models.Base{
		Name:        "l_refe",
		Description: "l_refe",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       10,
			Max:       400,
			UpperCase: true,
		},
	})
	Clientes = append(Clientes, models.Base{
		Name:        "c_ubig",
		Description: "c_ubig",
		Required:    true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Min:  6,
			Max:  6,
		},
	})
	Clientes = append(Clientes, models.Base{
		Name:        "n_tele",
		Description: "n_tele",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr: *models.Null(),
			Min:  3,
			Max:  30,
		},
	})
	Clientes = append(Clientes, models.Base{
		Name:        "n_celu",
		Description: "n_celu",
		Required:    true,
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       3,
			Max:       30,
			UpperCase: true,
		},
	})
	Clientes = append(Clientes, models.Base{
		Name:        "l_obse",
		Description: "l_obse",
		Update:      true,
		Type:        "string",
		Strings: models.Strings{
			Expr:      *models.Null(),
			Min:       10,
			Max:       100,
			UpperCase: true,
		},
	})
	return Clientes, tableName
}
