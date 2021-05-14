package usu_historial_puntos

import "testing"

func TestVerificarFactura(t *testing.T) {
	t.Run("caso 1, debería pasar", func(t *testing.T) {
		d := DtoSalida{
			Tipo:       "carga",
			Factura:    "808-PRUEBA",
			IdComercio: 3,
			IdSucursal: 808,
		}

		got := d.VerificarFactura()
		want := true

		if got != want {
			t.Errorf("got: %t, want %t", got, want)
		}
	})

	t.Run("caso 2, falla por la factura id_comercio-id_sucursal", func(t *testing.T) {
		d := DtoSalida{
			Tipo:       "carga",
			Factura:    "3-808-PRUEBA",
			IdComercio: 3,
			IdSucursal: 808,
		}

		got := d.VerificarFactura()
		want := false

		if got != want {
			t.Errorf("got: %t, want %t", got, want)
		}
	})

	t.Run("caso 3, falla por el Regex", func(t *testing.T) {
		d := DtoSalida{
			Tipo:       "carga",
			Factura:    "U111PRUEBA",
			IdComercio: 3,
			IdSucursal: 808,
		}

		got := d.VerificarFactura()
		want := false

		if got != want {
			t.Errorf("got: %t, want %t", got, want)
		}
	})

	t.Run("caso 4, falla por el tipo de tx", func(t *testing.T) {
		d := DtoSalida{
			Tipo:       "acumulacion",
			Factura:    "prueba",
			IdComercio: 3,
			IdSucursal: 808,
		}

		got := d.VerificarFactura()
		want := false

		if got != want {
			t.Errorf("got: %t, want %t", got, want)
		}
	})
}
