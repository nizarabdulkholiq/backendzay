package zay

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)



func TestInsertTagihan(t *testing.T) {
	dbname := "tagihandb"
	tagihan:= Tagihan{
		ID:      primitive.NewObjectID(),
		Isitagihan :   "spp : 7,700,000,00",
		
		
		
	}
	insertedID := InsertTagihan(dbname, tagihan)
	if insertedID == nil {
		t.Error("Failed to insert tagihan")
	}
}

func TestInsertUserTagihan(t *testing.T) {
	dbname := "tagihandb"
	usertagihan:= UserTagihan{
		ID:      primitive.NewObjectID(),
		Nama:    "Nizar Abdul Kholiq",
		Email :   "holiq@gmail.com",
		Telepon : "081234567891",
		
		
	}
	insertedID := InsertUserTagihan(dbname, usertagihan)
	if insertedID == nil {
		t.Error("Failed to insert user")
	}
}

func TestInsertDataTagihan(t *testing.T) {
	dbname := "tagihandb"
	datatagihan:= DataTagihan{
		ID:      primitive.NewObjectID(),
		Nama_Mahasiswa:    "Nizar Abdul Kholiq",
		Program_studi :   "Teknik Informatika",
		Jumlah: "200.000",
		Status: "Belum",
		
	}
	insertedID := InsertDataTagihan(dbname, datatagihan)
	if insertedID == nil {
		t.Error("Failed to insert billing")
	}
}

func TestInsertDataSudah(t *testing.T) {
	dbname := "tagihandb"
	datasudah := DataSudah{
		ID:       primitive.NewObjectID(),
		Nama_Mahasiswa: "Zaya wijaya",
		Status: "Sudah",
	}
	insertedID := InsertDataSudah(dbname, datasudah)
	if insertedID == nil {
		t.Error("Failed to insert user")
	}
}


func TestInsertDataBelum(t *testing.T) {
	dbname := "tagihandb"
	databelum := DataBelum{
		ID:       primitive.NewObjectID(),
		Nama_Mahasiswa: "Surya Azi",
		Status: "Belum",
	}
	insertedID := InsertDataBelum(dbname, databelum)
	if insertedID == nil {
		t.Error("Failed to insert user")
	}
}

