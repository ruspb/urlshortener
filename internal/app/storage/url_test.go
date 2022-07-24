package storage

import "testing"

func TestRetrieve(t *testing.T) {
	// подготовка
	longValue := "longValue"
	shortValue, _ := Store(longValue)

	urls["someKey"] = "someValue"

	type args struct {
		shortURL string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "Сохраненное через апи",
			args:    args{shortURL: shortValue},
			want:    longValue,
			wantErr: false,
		},
		{
			name:    "Сохраненное нативно",
			args:    args{shortURL: "someKey"},
			want:    "someValue",
			wantErr: false,
		},
		{
			name:    "Получение несуществующего",
			args:    args{shortURL: "undefined"},
			want:    "tyty",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Retrieve(tt.args.shortURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Retrieve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want && !tt.wantErr {
				t.Errorf("Retrieve() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore(t *testing.T) {
	longValue := "longValue"
	shortValue, _ := Store(longValue)

	urls["someKey"] = "someValue"

	type args struct {
		longURL string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "dsfdsf",
			args:    args{longURL: longValue},
			want:    shortValue,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Store(tt.args.longURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Store() got = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
func Test_find(t *testing.T) {
	type args struct {
		longURL string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := find(tt.args.longURL)
			if got != tt.want {
				t.Errorf("find() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("find() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_randomString(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomString(tt.args.n); got != tt.want {
				t.Errorf("randomString() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
