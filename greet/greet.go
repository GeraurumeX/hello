package greet

var emoji = "Gerardo"

// Si el nombre de la función comienza con mayuscula se puede exportar
// con minuscula no se puede exportar

// GreetEnglish retorna saludo en ingles
func English() string {
	return "Hi" + emoji
}

//  Italian retorna saludo en ingles
func Italian() string {
	return "Ciao" + emoji
}
