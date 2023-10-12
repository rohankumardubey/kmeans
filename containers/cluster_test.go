package containers

import (
	"testing"
)

func TestCluster_Recenter(t *testing.T) {
	type fields struct {
		Center  Vector
		Members []Vector
	}
	tests := []struct {
		name       string
		fields     fields
		wantCenter Vector
	}{
		{
			name: "test1",
			fields: fields{
				Center:  Vector{1, 1},
				Members: []Vector{{1, 1}, {2, 2}},
			},
			wantCenter: Vector{1.5, 1.5},
		},
		{
			name: "test2",
			fields: fields{
				Center:  Vector{1, 1},
				Members: []Vector{{1, 1}, {2, 2}, {3, 3}},
			},
			wantCenter: Vector{2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cluster{
				center:  tt.fields.Center,
				members: tt.fields.Members,
			}
			c.Recenter()
			if c.center.Compare(tt.wantCenter) != 0 {
				t.Errorf("Recenter() gotCenter = %v, want %v", c.center, tt.wantCenter)
			}
		})
	}
}

func TestCluster_RecenterReturningMovedDistance(t *testing.T) {
	type fields struct {
		Center  Vector
		Members []Vector
	}
	type args struct {
		distFn DistanceFunction
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		wantMoveDistances float64
		wantCenter        Vector
		wantErr           bool
	}{
		{
			name: "empty cluster",
			fields: fields{
				Center:  Vector{1, 1},
				Members: []Vector{},
			},
			args:       args{distFn: EuclideanDistance},
			wantCenter: Vector{1, 1}, // unchanged
		},
		{
			name: "non-empty cluster",
			fields: fields{
				Center:  Vector{1, 1},
				Members: []Vector{{1, 1}, {2, 2}},
			},
			args:              args{distFn: EuclideanDistance},
			wantMoveDistances: 0.7071067811865476,
			wantCenter:        Vector{1.5, 1.5}, // changed
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cluster{
				center:  tt.fields.Center,
				members: tt.fields.Members,
			}
			gotMoveDistance, err := c.RecenterWithMovedDistance(tt.args.distFn)
			if (err != nil) != tt.wantErr {
				t.Errorf("RecenterReturningMovedDistance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMoveDistance != tt.wantMoveDistances {
				t.Errorf("RecenterReturningMovedDistance() gotMoveDistance = %v, want %v", gotMoveDistance, tt.wantMoveDistances)
			}
			if c.center.Compare(tt.wantCenter) != 0 {
				t.Errorf("Recenter() gotCenter = %v, want %v", c.center, tt.wantCenter)
			}
		})
	}
}

func TestCluster_Reset(t *testing.T) {
	type fields struct {
		Center  Vector
		Members []Vector
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Test1",
			fields: fields{
				Center:  Vector{1, 1},
				Members: []Vector{{1, 1}, {2, 2}, {3, 3}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cluster{
				center:  tt.fields.Center,
				members: tt.fields.Members,
			}
			c.Reset()
			if len(c.members) != 0 {
				t.Errorf("Reset() = %v, want %v", c.members, []Vector{})
			}
			if c.center.Compare(tt.fields.Center) != 0 {
				t.Errorf("Reset() = %v, want %v", c.center, tt.fields.Center)
			}
		})
	}
}
