package main

import "fmt"

type Imagen struct {
	titulo  string
	formato string
	canales string
}

func (img *Imagen) mostrar() {
	fmt.Println(img.titulo)
	fmt.Println(img.formato)
	fmt.Println(img.canales)
}

type Audio struct {
	titulo   string
	formato  string
	duracion string
}

func (aud *Audio) mostrar() {
	fmt.Println(aud.titulo)
	fmt.Println(aud.formato)
	fmt.Println(aud.duracion)
}

type Video struct {
	titulo  string
	formato string
	frames  string
}

func (vid *Video) mostrar() {
	fmt.Println(vid.titulo)
	fmt.Println(vid.formato)
	fmt.Println(vid.frames)
}

type Multimedia interface {
	mostrar()
}

type ContenidoWeb struct {
	ArchivosMultimedia []Multimedia
}

func (cw *ContenidoWeb) mostrar() {
	for _, c := range cw.ArchivosMultimedia {
		c.mostrar()
	}
}

func main() {
	var opc int
	var valor1, valor2, valor3 string
	for opc != 4 {
		fmt.Println("Selecciona la opcion que quieras capturar!!!")
		fmt.Println("1.- Imagen")
		fmt.Println("2.- Audio")
		fmt.Println("3.- Video")
		fmt.Println("4.- Salir")
		fmt.Scan(&opc)
		switch opc {
		case 1:
			fmt.Println("Capturar Imagen")
			fmt.Println("1.- Titulo")
			fmt.Scan(&valor1)
			fmt.Println("2.- Formato")
			fmt.Scan(&valor2)
			fmt.Println("3.- Canales")
			fmt.Scan(&valor3)
			img := Imagen{titulo: valor1, formato: valor2, canales: valor3}
			cwd := ContenidoWeb{
				ArchivosMultimedia: []Multimedia{
					&img,
				},
			}
			cwd.mostrar()
		case 2:
			fmt.Println("Capturar Audio")
			fmt.Println("1.- Titulo")
			fmt.Scan(&valor1)
			fmt.Println("2.- Formato")
			fmt.Scan(&valor2)
			fmt.Println("3.- Duracion")
			fmt.Scan(&valor3)
			aud := Audio{titulo: valor1, formato: valor2, duracion: valor3}
			cwd := ContenidoWeb{
				ArchivosMultimedia: []Multimedia{
					&aud,
				},
			}
			cwd.mostrar()
		case 3:
			fmt.Println("Capturar Video")
			fmt.Println("1.- Titulo")
			fmt.Scan(&valor1)
			fmt.Println("2.- Formato")
			fmt.Scan(&valor2)
			fmt.Println("3.- Frames")
			fmt.Scan(&valor3)
			vid := Video{titulo: valor1, formato: valor2, frames: valor3}
			cwd := ContenidoWeb{
				ArchivosMultimedia: []Multimedia{
					&vid,
				},
			}
			cwd.mostrar()
		default:
			fmt.Println("opcion incorrecta")

		}
	}
}
