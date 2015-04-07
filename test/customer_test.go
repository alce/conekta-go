package test

import "testing"

var customer *Customer

func TestSetUp(t *testing.T) {
	customer = NewCustomer("cus_zZ42sCRYK1br5zajw", "")
	customer.ApiKey = "<api_key>"
}

func TestPaused(t *testing.T) {
	if err := customer.Pause(); err != nil {
		t.Logf("No se pudo actualizar la subscripción:  %v", err)
		return
	}
	t.Logf("Subscripción actualizada!")
}

func TestResume(t *testing.T) {
	if err := customer.Resume(); err != nil {
		t.Logf("No se pudo actualizar la subscripción:  %v", err)
		return
	}
	t.Logf("Subscripción actualizada!")
}
