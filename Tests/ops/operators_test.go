package ops

import (
	"testing"
)
type args struct{
	x,y int
}
func TestKeyOpGenerate(t *testing.T) {
	tests := []struct{
		name string
		args args
		want string
	}{
		{"success", args{1,2}, "1_2"},
		{"success large integers", args{9999999,33333333},"9999999_33333333"},
	}
	for _, test := range tests{
		t.Run(test.name, func(t *testing.T){
			kp := GetKeyOperator()
			if got := kp.Generate(test.args.x, test.args.y); got != test.want{
				t.Errorf("keyOp.Generate() = %v, want %v", got, test.want)
			}
		})
	}

}

func TestKeyOpDegenerate(t *testing.T) {
	tests := []struct{
		name string
		input string
		wantX, wantY int
		wantErr bool
	}{
		{"success","1_2",1,2,false},
		{"bad format", "1_%2", 0, 0,true},
	}
	for _, test := range tests{
		t.Run(test.name, func(t *testing.T){
			kp := GetKeyOperator()
			  gotX, gotY, gotErr := kp.Degenerate(test.input)
			  //|| gotY != test.wantY ||(gotErr != nil) != test.wantErr
			 if gotX != test.wantX{
			 	t.Errorf("got: x = %v, want: x = %v", gotX, test.wantX)
			 }
			if gotY != test.wantY{
				t.Errorf("got: y = %v, want: y = %v", gotY, test.wantY)
			}
			if (gotErr != nil) != test.wantErr{
				t.Errorf("keyOp.Degenerate() error = %v, wantErr %v", gotErr, test.wantErr)
			}
		})
	}
}
