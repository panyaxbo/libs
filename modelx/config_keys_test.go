package modelx

import "testing"

func TestBuildConfigKey(t *testing.T) {
	type args struct {
		configKeys []string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty string",
			args: struct {
				configKeys []string
			}{
				configKeys: []string{},
			},
			want: "",
		},
		{
			name: "normal string",
			args: struct {
				configKeys []string
			}{
				configKeys: []string{"a", "B"},
			},
			want: "A.B",
		},
		{
			name: "1 string",
			args: struct {
				configKeys []string
			}{
				configKeys: []string{"a"},
			},
			want: "A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildConfigKey(tt.args.configKeys...); got != tt.want {
				t.Errorf("BuildConfigKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
