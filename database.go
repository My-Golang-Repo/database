package main

import (
	"database/sql"
	"fmt"

	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:SADAasga1234@/test?parseTime=true")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	{
		sorgu := `

    CREATE TABLE IF NOT EXISTS kullanicilar (
        id INT AUTO_INCREMENT,
        kullanici TEXT NOT NULL,
        sifre TEXT NOT NULL,
        tarih DATETIME,
        PRIMARY KEY (id)
    );`
		//Sorguyu çalıştırma
		_, err = db.Exec(sorgu)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

/* 	{
		kullaniciDegeri := "johndoe"
		sifreDegeri := "secret"
		tarihDegeri := time.Now()
		sonuc, err := db.Exec(`INSERT INTO kullanicilar (kullanici, sifre, tarih) VALUES (?, ?, ?)`, kullaniciDegeri, sifreDegeri, tarihDegeri)
		if err != nil {
			log.Fatal(err.Error())
		}
		kullaniciID, err := sonuc.LastInsertId()
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println("Eklenen kullanıcının id'si:", kullaniciID)

	} */
/* 
	{
		var (
			id int
			kullanici string
			sifre string
			tarih time.Time
			sorguid int
		)

		sorguid=99
		//Tabloyu sorgulayıp sonuçları değişkenlere yazdıralım
		sorgu := `SELECT id, kullanici, sifre, tarih FROM kullanicilar WHERE id = ?`
		err := db.QueryRow(sorgu, sorguid).Scan(&id, &kullanici, &sifre, &tarih)
		if err != nil {
			log.Fatal(err.Error())
		}
		//Çıkan aldığımız verileri ekrana bastıralım
		log.Println(id, kullanici, sifre, tarih)
	}  */

/* 	{
		var (
			id   int
			user string
			pass string
			date time.Time
		)
		err := db.QueryRow("select * from kullanicilar where id=?", 5).Scan(&id, &user, &pass, &date)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Println(id, user)

	} */
 	{
		type kullanici struct {
			id int
			kullanici string
			sifre string
			tarih time.Time
		}
		
		veri, err := db.Query("select * from kullanicilar")
		if err != nil {
			log.Fatal(err.Error())
		}
		defer veri.Close()
		log.Println(*veri)
		var kullanicilar []kullanici
		for veri.Next(){
			var k kullanici
			err := veri.Scan(&k.id, &k.kullanici, &k.sifre, &k.tarih)
			if err !=nil {
				log.Fatal(err.Error())
			}
			kullanicilar = append(kullanicilar, k)
			//log.Println(kullanicilar)
			fmt.Printf("%#v", kullanicilar)
			log.Println("Ikinci sorgu", k.id, k.kullanici, k.sifre, k.tarih)
		}

		{
			_, err = db.Exec("delete from kullanicilar where id = 134")	
			if err != nil {
			 log.Fatal(err.Error())
			}

		}

	} 

	{

	}

}
