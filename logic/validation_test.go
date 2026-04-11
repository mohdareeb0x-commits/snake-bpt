package logic

import "testing"

func TestCheckDependency_Exists(t *testing.T) {
	if err := CheckDependency("git"); err != nil {
		t.Errorf("expected git to be found, got: %v", err)
	}
}

func TestCheckDependency_Missing(t *testing.T) {
	err := CheckDependency("not-a-real-binary")
	if err == nil {
		t.Error("expected error for missing binary, got nil")
	}
}

func TestCheckInternet(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping internet check test")
	}
	if err := CheckInternet(); err != nil {
		t.Errorf("expected internet to be available, got: %v", err)
	}
}

func TestValidateIP_Valid(t *testing.T) {
	validation := ValidateIP("127.0.0.1");
	if validation == false {
		t.Errorf("expected valid host, got invalid")
	}
}

func TestValidateIP_Invalid(t *testing.T) {
	validation := ValidateIP("invalid-host");
	if validation == true {
		t.Errorf("expected invalid host, got valid")
	}
}

func TestValidatePort_Valid(t *testing.T) {
	validation := ValidatePort("8080")
	if validation == false {
		t.Errorf("expected valid port, got invalid")
	}

	validation = ValidatePort("56700")
	if validation == false {
		t.Errorf("expected valid port, got invalid")
	}
}

func TestValidatePort_Invalid(t *testing.T) {
	validation := ValidatePort("99999") 
	if validation == true{
		t.Errorf("expected invalid port, got valid")
	}

	validation = ValidatePort("Invalid-Port") 
	if validation == true{
		t.Errorf("expected invalid port, got valid")
	}
}
