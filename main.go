package main
import (
			"fmt"
			// "encoding/json"
			"github.com/bwmarrin/discordgo"
		)
		

		// var cmd map[string]func(s *R, ws *websocket.Conn)
		var dg *discordgo.Session
		
func main(){
	var err error
	
	dg, err = discordgo.New("MTkwNDI1ODc3MTA3MDQ4NDQ4.CjrxHQ.pNOxc5T-Z0ltFltNcZ7dpqcDC_M")
	
	if err  != nil {
		fmt.Println("error creating Discord session,", err)
		panic(err)
	}
	
	
		dg.AddHandler(MessageHandler)
		
		// Open the websocket and begin listening.
		dg.Open()
		
		<-make(chan struct{})
		
		
	

				 // else if strings.ToUpper(data.Text) == "!PIG"{
					
					// data.Response("désolée, ce jeu est encore en cours de dévellopement, merci de bien vouloir attendre encore quelque temps.", ws)
					// data.Response("lancement de PIG", ws)
					// PIG(data, ws)
					
					
					
				// } else if strings.ToUpper(data.Text) == "!POKEMON" || strings.ToUpper(data.Text) == "!POKE" {
					
					// data.Response("désolée, ce jeu est encore en cours de dévellopement, merci de bien vouloir attendre encore quelque temps.", ws)
					
									
				// } else if (strings.ToUpper(data.Text) == "!DEGAGE" || strings.ToUpper(data.Text) == "!FAIL") && GetNameById(data.User, &S)== "kiza"{
					
						// data.Response("*en train d'agoniser*", ws)
						// data.Response("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAARRRRRRGGGGG", ws)
						// data.Response(".............................................................................................................................................................................", ws)
						// data.Response("..............................................................................................................................................................................", ws)
						// data.Response("'-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' '-' ", ws)
						// panic(err)
						
						
				// } 

	
	

}


