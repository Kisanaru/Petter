package main

import (
			"fmt"
			"encoding/json"
			"net/http"
			"strings"
			// "os"
			// "os/exec"
			"github.com/bwmarrin/discordgo"
			"math/rand"
			// "strconv"
			// "time"
			// "golang.org/x/net/html"
			// "bytes"
		)
		var Client http.Client
func MessageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Printf("%20s %20s > %s\n", m.ChannelID, m.Author.Username, m.Content)
	if m.Author.Username != "Petter bots"{
																					
							//------------------------------------------------------message------------------------------------
							
		if (strings.Contains(strings.ToUpper(m.Content), "PETTER!" ) == true){
			var Rep1 = []string{
				"je suis cheper chui pas la laisser moi >3>",
				"Oui?",
				"Je suis partis en vacance lachez moi pendant ces 2 semaine de projet è_é",
			    "Je suis en maintenance et je veux plus de drogue et de jeux video MERCI !!! >_<*",
			}
			rand_rep1 := rand.Intn(len(Rep1))
			dg.ChannelMessageSend(m.ChannelID, Rep1[rand_rep1])
	    } else if (strings.Contains(strings.ToUpper(m.Content), "DISCORD > ALL" ) == true){
			dg.ChannelMessageSend(m.ChannelID, "Discord c'est la base :3, TS.... sa passe encore, mumble.... mouai non on s'en ballec en fait du reste '-'")
		} else if (strings.Contains(strings.ToUpper(m.Content), "VFFT" ) == true){
			dg.ChannelMessageSend(m.ChannelID, "https://www.youtube.com/watch?v=FEbJ2_p9PmM")
		} else if (strings.Contains(strings.ToUpper(m.Content), "SAITAMA!" ) == true){
			dg.ChannelMessageSend(m.ChannelID, "https://www.youtube.com/watch?v=_TUTJ0klnKk")
	    } else if (strings.Contains(strings.ToUpper(m.Content), "DROGUE" ) == true){
			var Rep2 = []string{
				"Chuis cheper la les mec , chui pas la laisser moi >3>",
				"Pas Bien est la drogue, sinon on se fait un bangue??",
				"DEAL WITH IT",
				"quand j'ai pas ma drogue et que je regard dans un mirroire http://www.griffin-tv.com/img_grfn/3.14.jpg",
				"https://66.media.tumblr.com/tumblr_lms4wp3CvG1qf40jso1_500.gif ITS SO GOOD",
				"https://www.youtube.com/watch?v=A2CDdeiJBJk",
				"https://66.media.tumblr.com/c9fcc871f4894aa1df4c1ac1bf6fa5a8/tumblr_inline_o7oomxTbJs1u5hohy_500.gif",
			}
			rand_rep2 := rand.Intn(len(Rep2))
			dg.ChannelMessageSend(m.ChannelID, Rep2[rand_rep2])
	
	    } else if (strings.Contains(strings.ToUpper(m.Content), "KICK" ) == true  && strings.Contains(strings.ToUpper(m.Content), "PETTER" ) == true){
		
			dg.ChannelMessageSend(m.ChannelID, "Nn s'il vous pliz , et je te donne une photo de ma fille a poils '-' alor je peut rester " + m.Author.Username + "?" )
		
		} else if ((strings.Contains(strings.ToUpper(m.Content), "$SEXS" ) == true)){
			
			var perv_girls = []string{
				"YogaPants",
				"boobs",
				"SexyButNotPorn",
				"gonewild",
				"ass",
				"datgap",
				"pawg",
				"girlsinyogapants",
				"rearpussy",
				"assinthong",
			}
			randy := rand.Intn(len(perv_girls))
			
			req, err := http.NewRequest("GET", "https://www.reddit.com/r/"+perv_girls[randy]+".json", nil)
			
			
			if err != nil {
				panic(err)
			}
			
			
			req.Header.Set("User-Agent", "I do things ? ... I'm a stupid bot ! (maxence.dufrenoy@viacesi.fr)")
			resp, err := Client.Do(req)
			
			defer resp.Body.Close()
			
			
			if err != nil {
				panic(err)
			}
			
			
			var P PRV
			dec := json.NewDecoder(resp.Body)
			err = dec.Decode(&P)
			
			
			if err != nil {
				panic(err)
			}
			
			
			if len(P.Data.Children)>0{
				randy2 := rand.Intn(len(P.Data.Children))
				
				dg.ChannelMessageSend(m.ChannelID, P.Data.Children[randy2].Data.Preview.Images[0].Source.URL)
				
				
            }
	    }
	}
}