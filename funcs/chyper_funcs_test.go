package funcs

import (
	"DEMO-neo4j/core"
	"DEMO-neo4j/utility"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"log"
	"reflect"
	"testing"
)

// creo una var condivisa per il driver, altrimenti se lo creo in ogni test diventa troppo dispendioso
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

func TestRepoDriver_ContainsItem(t *testing.T) {
	type fields struct {
		drv neo4j.Driver
	}
	type args struct {
		nome string
		sku  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "Non presente",
			fields: fields{drv: drv},
			args: args{
				nome: "x",
				sku:  123,
			},
			want: false,
		},

		{
			name: "Presente",
			fields: fields{
				drv: drv,
			},
			args: args{
				nome: "Air Force 1",
				sku:  11,
			},
			want: true,
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			r := &RepoDriver{drv: tt.fields.drv}
			got := r.ContainsItem(tt.args.nome, tt.args.sku)
			if got != tt.want {
				t.Errorf("ContainsItem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepoDriver_InsertItem(t *testing.T) {
	type fields struct {
		drv neo4j.Driver
	}
	type args struct {
		nome string
		sku  int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "Insert 1",
			fields: fields{drv: drv},
			args: args{
				nome: "InsertTest",
				sku:  100,
			},
			want: true,
		},
		{
			name:   "Insert Uguale 1",
			fields: fields{drv: drv},
			args: args{
				nome: "InsertTest",
				sku:  100,
			},
			want: false,
		},
		{
			name:   "Insert Uguale 2",
			fields: fields{drv: drv},
			args: args{
				nome: "InsertTest",
				sku:  100,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &RepoDriver{
				drv: tt.fields.drv,
			}
			session := r.drv.NewSession(neo4j.SessionConfig{AccessMode: neo4j.AccessModeRead})
			res, err := session.ReadTransaction(func(tx neo4j.Transaction) (interface{}, error) {
				query := `MATCH (m:Item) WHERE m.nome = $wantedNome AND m.sku = $wantedSku RETURN m`
				result, err := tx.Run(query, map[string]interface{}{"wantedNome": tt.args.nome, "wantedSku": tt.args.sku})
				if err != nil {
					log.Fatal(err)
					return false, err
				}
				var i core.Item
				for result.Next() {

					props := result.Record().Values[0].(neo4j.Node).Props
					i = core.Item{
						Name: props["nome"].(string),
						Sku:  int(props["sku"].(int64)),
					}
				}
				return i, nil
			})
			if err != nil {
				log.Fatal(err)
			}
			got := r.InsertItem(tt.args.nome, tt.args.sku)

			// creo Item per fare il confronto con quello che ottengo dalla query
			wantedItem := core.Item{
				Name: tt.args.nome,
				Sku:  tt.args.sku,
			}

			if !(cmp.Equal(res.(core.Item), wantedItem)) {
				fmt.Printf("\t %v wanted %v", res.(core.Item), wantedItem)
			}
			if got != tt.want {
				t.Errorf("InsertItem() = %v, want %v", got, tt.want)
			}
		})
	}
}
