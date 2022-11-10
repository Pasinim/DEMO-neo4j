package funcs

import (
	"DEMO-neo4j/core"
	"DEMO-neo4j/utility"
	"github.com/google/go-cmp/cmp"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"reflect"
	"testing"
)

var drv = utility.InitDriver()

func TestRepoDriver_GetItemFromSku(t *testing.T) {
	type fields struct {
		drv neo4j.Driver
	}
	type args struct {
		sku int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    core.Item
		wantErr bool
	}{
		{
			name:   "Art 1",
			fields: fields{drv},
			args:   args{sku: 11},
			want: core.Item{
				Name: "Air Force 1",
				Sku:  11,
			},
			wantErr: false,
		},
		{
			name:   "Art 2",
			fields: fields{drv},
			args:   args{sku: 22},
			want: core.Item{
				Name: "Free Run",
				Sku:  22,
			},
			wantErr: false,
		},
		{
			name:   "Art 3",
			fields: fields{drv},
			args:   args{sku: 33},
			want: core.Item{
				Name: "AirMax 98",
				Sku:  33,
			},
			wantErr: false,
		},
		{
			name:   "Art 4",
			fields: fields{drv},
			args:   args{sku: 44},
			want: core.Item{
				Name: "Felpa con Cappuccio",
				Sku:  44,
			},
			wantErr: false,
		},
		{
			name:   "Art 5",
			fields: fields{drv},
			args:   args{sku: 55},
			want: core.Item{
				Name: "Felpa senza Cappuccio",
				Sku:  55,
			},
			wantErr: false,
		},
		{
			name:    "Art null",
			fields:  fields{drv},
			args:    args{sku: 100},
			want:    *new(core.Item),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RepoDriver{
				drv: tt.fields.drv,
			}
			got, err := r.GetItemFromSku(tt.args.sku)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetItemFromSku() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItemFromSku() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepoDriver_GetItems(t *testing.T) {
	type fields struct {
		drv neo4j.Driver
	}
	tests := []struct {
		name   string
		fields fields
		want   []core.Item
	}{
		{
			name:   "All Articoli",
			fields: fields{drv},
			want: []core.Item{
				{
					Name: "Air Force 1",
					Sku:  11,
				},
				{
					Name: "Free Run",
					Sku:  22,
				},
				{
					Name: "AirMax 98",
					Sku:  33,
				},
				{
					Name: "Felpa con Cappuccio",
					Sku:  44,
				},
				{
					Name: "Felpa senza Cappuccio",
					Sku:  55,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RepoDriver{
				drv: tt.fields.drv,
			}
			if got := r.GetItems(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItems() = %v, want %v\n", got, tt.want)
				t.Errorf("Gli elementi non corrispondono:\n%s", cmp.Diff(got, tt.want))

			}
		})
	}
}
