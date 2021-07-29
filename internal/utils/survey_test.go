package utils

import (
	"reflect"
	"testing"

	"github.com/ozoncp/ocp-survey-api/internal/models"
)

func TestSplitToChunks(t *testing.T) {
	type args struct {
		surveys   []models.Survey
		chunkSize int
	}
	tests := []struct {
		name    string
		args    args
		want    [][]models.Survey
		wantErr bool
	}{
		{
			name: "Empty slice",
			args: args{
				surveys:   []models.Survey{},
				chunkSize: 10,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Invalid chunk size",
			args: args{
				surveys: []models.Survey{
					{Id: 0}, {Id: 1}, {Id: 2},
				},
				chunkSize: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Smaller than chunk",
			args: args{
				surveys: []models.Survey{
					{Id: 0}, {Id: 1}, {Id: 2},
				},
				chunkSize: 10,
			},
			want: [][]models.Survey{
				{{Id: 0}, {Id: 1}, {Id: 2}},
			},
			wantErr: false,
		},
		{
			name: "Same size as chunk",
			args: args{
				surveys: []models.Survey{
					{Id: 0}, {Id: 1}, {Id: 2},
				},
				chunkSize: 3,
			},
			want: [][]models.Survey{
				{{Id: 0}, {Id: 1}, {Id: 2}},
			},
			wantErr: false,
		},
		{
			name: "Larger than chunk",
			args: args{
				surveys: []models.Survey{
					{Id: 0}, {Id: 1}, {Id: 2},
					{Id: 3}, {Id: 4},
				},
				chunkSize: 3,
			},
			want: [][]models.Survey{
				{{Id: 0}, {Id: 1}, {Id: 2}},
				{{Id: 3}, {Id: 4}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SplitToChunks(tt.args.surveys, tt.args.chunkSize)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitToChunks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitToChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceToMap(t *testing.T) {
	type args struct {
		surveys []models.Survey
	}
	tests := []struct {
		name    string
		args    args
		want    map[uint64]models.Survey
		wantErr bool
	}{
		{
			name: "Empty slice",
			args: args{
				surveys: []models.Survey{},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Duplicate IDs",
			args: args{
				surveys: []models.Survey{
					{Id: 1, UserId: 10, Link: "aaa"},
					{Id: 1, UserId: 20, Link: "bbb"},
					{Id: 2},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Valid case",
			args: args{
				surveys: []models.Survey{
					{Id: 1}, {Id: 2}, {Id: 3},
				},
			},
			want: map[uint64]models.Survey{
				1: {Id: 1}, 2: {Id: 2}, 3: {Id: 3},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SliceToMap(tt.args.surveys)
			if (err != nil) != tt.wantErr {
				t.Errorf("SliceToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SliceToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
