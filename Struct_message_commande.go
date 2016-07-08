package main


type CNF1 []struct{
	ID string `json:"id"`
	Fact string `json:"fact"`
	Date string `json:"date"`
	Vote string `json:"vote"`
	Points string `json:"points"`
}


type CNF2 struct {
	IconURL string `json:"icon_url"`
	ID string `json:"id"`
	URL string `json:"url"`
	Value string `json:"value"`
}

type JPIG struct {

	prenom string
	nom string

}