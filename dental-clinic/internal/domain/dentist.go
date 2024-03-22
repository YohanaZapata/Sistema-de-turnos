package domain

type Dentist struct {
	ID       string
	Name     string
	Lastname string
	Matricula string
}

func (d Dentist) Validate() error {
	return nil
}

type Paciente struct {
	ID       string
	Name     string
	Lastname string
	Domicilio string
	DNI      string
	FechaDeAlta string
}

func (p Paciente) Validate() error {
	return nil
}

type Turno struct {
	ID string
	PacienteID string
	DentistaID string
	FechaHora string
	Descripcion string
}

func (t Turno) Validate() error {
	return nil
}