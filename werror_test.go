package werror

import (
	"errors"
	"os"
	"testing"
)

func TestWrap(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "wrap",
			args: args{
				err: os.ErrClosed,
			},
			want: os.ErrClosed,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.err); !errors.Is(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrapped_Error(t *testing.T) {
	type fields struct {
		err     error
		nextErr error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "error_stringer",
			fields: fields{
				err:     os.ErrInvalid,
				nextErr: nil,
			},
			want: os.ErrInvalid.Error(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Error{
				err:     tt.fields.err,
				nextErr: tt.fields.nextErr,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrapped_Unwrap(t *testing.T) {
	type fields struct {
		err     error
		nextErr error
	}
	tests := []struct {
		name   string
		fields fields
		want   error
	}{
		{
			name: "unwrap",
			fields: fields{
				err: os.ErrNotExist,
				nextErr: &Error{
					err:     os.ErrPermission,
					nextErr: nil,
				},
			},
			want: os.ErrPermission,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Error{
				err:     tt.fields.err,
				nextErr: tt.fields.nextErr,
			}
			if got := e.Unwrap(); !errors.Is(got, tt.want) {
				t.Errorf("Unwrap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWrapped_Wrap(t *testing.T) {
	type fields struct {
		err     error
		nextErr error
	}
	type args struct {
		err error
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		want     error
		unwanted error
	}{
		{
			name: "wrap_want_inner",
			fields: fields{
				err:     os.ErrPermission,
				nextErr: nil,
			},
			args: args{
				err: os.ErrInvalid,
			},
			want:     os.ErrPermission,
			unwanted: os.ErrDeadlineExceeded,
		},
		{
			name: "wrap_want_outer",
			fields: fields{
				err:     os.ErrPermission,
				nextErr: nil,
			},
			args: args{
				err: os.ErrInvalid,
			},
			want:     os.ErrInvalid,
			unwanted: os.ErrNoDeadline,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Error{
				err:     tt.fields.err,
				nextErr: tt.fields.nextErr,
			}
			if got := e.Wrap(tt.args.err); !errors.Is(got, tt.want) {
				t.Errorf("New() + !errors.Is() = %v, want %v", got, tt.want)
			}
			if got := e.Wrap(tt.args.err); errors.Is(got, os.ErrDeadlineExceeded) {
				t.Errorf("New() + errors.Is() = %v, don't want %v", got, tt.unwanted)
			}
		})
	}
}

func TestWrapped_Err(t *testing.T) {
	type fields struct {
		err     error
		nextErr error
	}
	tests := []struct {
		name    string
		fields  fields
		want error
	}{
		{
			name: "err",
			fields: fields{
				err: os.ErrNoDeadline,
				nextErr: nil,
			},
			want: os.ErrNoDeadline,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Error{
				err:     tt.fields.err,
				nextErr: tt.fields.nextErr,
			}
			if err := e.Err(); err != tt.want {
				t.Errorf("Err() error = %v, wantErr %v", err, tt.want)
			}
		})
	}
}
