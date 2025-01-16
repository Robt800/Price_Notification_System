package output

import (
	"context"
	"testing"
)

func TestOutputs(t *testing.T) {

}

func TestOutputs1(t *testing.T) {
	type args struct {
		producedData chan []byte
		ctx          context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Outputs(tt.args.producedData, tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Outputs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
