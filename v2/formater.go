package lets_go_module_format

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"os"
)

type patient struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

type records []patient

func Do(sourceFile, resultFile string) {
	f, err := os.Open(sourceFile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var p patient
	r := records{}
	decoder := json.NewDecoder(f)
	for decoder.More() {
		err = decoder.Decode(&p)
		if err != nil {
			log.Fatal(err)
		}
		r = append(r, p)
	}

	// v2.0.0 - запись структуры в xml

	f, err = os.CreateTemp("./", resultFile)
	if err != nil {
		log.Fatalln(err)
	}
	err = xml.NewEncoder(f).Encode(&r)
	if err != nil {
		log.Fatalln(err)
	}
	err = f.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
