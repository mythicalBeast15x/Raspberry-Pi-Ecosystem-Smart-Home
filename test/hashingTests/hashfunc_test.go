package hashingTests

import (
	"CMPSC488SP24SecThursday/hashing"
	"fmt"
	"os"
	"testing"
)

func readFileToString(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("Error reading file: %v", err))
	}
	return string(content)
}
func Test_calculateMD5(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{input: "Hello, World!"},
			want: "65a8e27d8879283831b664bd8b7f0ad4",
		},
		{
			name: "Test case 2",
			args: args{input: readFileToString("sampledoc.txt")},
			want: "f633d57005758304ecc189ca68f39fcb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := hashing.CalculateMD5(tt.args.input)
			if got != tt.want {
				t.Errorf("calculateMD5() got = %v, want %v", got, tt.want)
			}

		})
	}
}

func Test_calculateSHA1(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{input: "Hello, World!"},
			want: "0a0a9f2a6772942557ab5355d76af442f8f65e01",
		},
		{
			name: "Test case 2",
			args: args{input: readFileToString("sampledoc.txt")},
			want: "88911d1d7c3b774a61c6306a73d6a8c32355b769",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := hashing.CalculateSHA1(tt.args.input)
			if got != tt.want {
				t.Errorf("calculateSHA1() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateSHA256(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{input: "Hello, World!"},
			want: "dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f",
		},
		{
			name: "Test case 2",
			args: args{input: readFileToString("sampledoc.txt")},
			want: "1d37f8a282ab24c295425654b82067190434cebd2153b179cf3a231dd0abbd73",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := hashing.CalculateSHA256(tt.args.input)
			if got != tt.want {
				t.Errorf("calculateSHA256() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calculateSHA512(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{input: "Hello, World!"},
			want: "374d794a95cdcfd8b35993185fef9ba368f160d8daf432d08ba9f1ed1e5abe6cc69291e0fa2fe0006a52570ef18c19def4e617c33ce52ef0a6e5fbe318cb0387",
		},
		{
			name: "Test case 2",
			args: args{input: readFileToString("sampledoc.txt")},
			want: "312d42e27669b7b2456158745dd4a0eebd13c8b3f7c68be67d4317790bb2a24f818a0b62216b9acf75c43e7d0f1acd34c19eaa6842263d81a4348e92ea769914",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := hashing.CalculateSHA512(tt.args.input)
			if got != tt.want {
				t.Errorf("calculateSHA512() got = %v, want %v", got, tt.want)
			}
		})
	}
}
