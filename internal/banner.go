package internal

import "github.com/tomiok/fuego-cache/logs"

func PrintBanner() {
	logs.StdInfo(`

8888888888 888     888 8888888888  .d8888b.   .d88888b.       .d8888b.         d8888  .d8888b.  888    888 8888888888 
888        888     888 888        d88P  Y88b d88P" "Y88b     d88P  Y88b       d88888 d88P  Y88b 888    888 888        
888        888     888 888        888    888 888     888     888    888      d88P888 888    888 888    888 888        
8888888    888     888 8888888    888        888     888     888            d88P 888 888        8888888888 8888888    
888        888     888 888        888  88888 888     888     888           d88P  888 888        888    888 888        
888        888     888 888        888    888 888     888     888    888   d88P   888 888    888 888    888 888        
888        Y88b. .d88P 888        Y88b  d88P Y88b. .d88P     Y88b  d88P  d8888888888 Y88b  d88P 888    888 888        
888         "Y88888P"  8888888888  "Y8888P88  "Y88888P"       "Y8888P"  d88P     888  "Y8888P"  888    888 8888888888   ` +
		fire())
}

func fire() string {
	return `                        .                      
                          /                       
                          (                       
                    .     ((                      
                      (/   ((  .                  
                      ((* ./(  (                  
                     (*((((*( (                   
                    (**((/*(((  (                 
                 . (**/(/*((/***(                 
                ( */(*****(*/(/*(,                
                 ( **(...****(***(((,             
                 ((*.*(...****..*((/(             
                 (**.**,...*..../**(*             
                 (**..*........**,*(              
                  (*,............/                
                    .*.........     `
}

func symbol() string {
	s := `

				    xxxxxxx
                               x xxxxxxxxxxxxx x
                            x     xxxxxxxxxxx     x
                                   xxxxxxxxx
                         x          xxxxxxx          x
                                     xxxxx
                        x             xxx             x
                                       x
                       xxxxxxxxxxxxxxx   xxxxxxxxxxxxxxx
                        xxxxxxxxxxxxx     xxxxxxxxxxxxx
                         xxxxxxxxxxx       xxxxxxxxxxx
                          xxxxxxxxx         xxxxxxxxx
                            xxxxxx           xxxxxx
                              xxx             xxx
                                  x         x
                                       x`
	return s
}
