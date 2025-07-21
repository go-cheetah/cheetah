package command

import "testing"

func TestCommand_gen(t *testing.T) {
	type fields struct {
		Name           string
		Path           string
		Type           string
		Title          string
		AbsProjectPath string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Command{
				Name:           tt.fields.Name,
				Path:           tt.fields.Path,
				Type:           tt.fields.Type,
				Title:          tt.fields.Title,
				AbsProjectPath: tt.fields.AbsProjectPath,
			}
			if err := s.gen(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("Command.gen() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
