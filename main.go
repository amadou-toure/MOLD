package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	//var Framework string
	var username string
	var ProjectName string
	banner := `
		 _____                    _______                   _____            _____          
         /\    \                 /::\    \                 /\    \          /\    \         
        /::\____\               /::::\    \               /::\____\        /::\    \        
       /::::|   |              /::::::\    \             /:::/    /       /::::\    \       
      /:::::|   |             /::::::::\    \           /:::/    /       /::::::\    \      
     /::::::|   |            /:::/~~\:::\    \         /:::/    /       /:::/\:::\    \     
    /:::/|::|   |           /:::/    \:::\    \       /:::/    /       /:::/  \:::\    \    
   /:::/ |::|   |          /:::/    / \:::\    \     /:::/    /       /:::/    \:::\    \   
  /:::/  |::|___|______   /:::/____/   \:::\____\   /:::/    /       /:::/    / \:::\    \  
 /:::/   |::::::::\    \ |:::|    |     |:::|    | /:::/    /       /:::/    /   \:::\ ___\ 
/:::/    |:::::::::\____\|:::|____|     |:::|    |/:::/____/       /:::/____/     \:::|    |
\::/    / ~~~~~/:::/    / \:::\    \   /:::/    / \:::\    \       \:::\    \     /:::|____|
 \/____/      /:::/    /   \:::\    \ /:::/    /   \:::\    \       \:::\    \   /:::/    / 
             /:::/    /     \:::\    /:::/    /     \:::\    \       \:::\    \ /:::/    /  
            /:::/    /       \:::\__/:::/    /       \:::\    \       \:::\    /:::/    /   
           /:::/    /         \::::::::/    /         \:::\    \       \:::\  /:::/    /    
          /:::/    /           \::::::/    /           \:::\    \       \:::\/:::/    /     
         /:::/    /             \::::/    /             \:::\    \       \::::::/    /      
        /:::/    /               \::/____/               \:::\____\       \::::/    /       
        \::/    /                 ~~                      \::/    /        \::/____/     
         \/____/                                           \/____/   
`

	fmt.Println(banner)
	//to init a fiber project: "go mod init github.com/username/projectname","go get github.com/gofiber/fiber"
	fmt.Print("Enter your username: ")
	fmt.Scanln(&username)
	// fmt.Println("choose your project type:")
	// fmt.Println("1. fiber")
	// fmt.Println("2. gin")
	// fmt.Println("3. Fast HTTP")
	// fmt.Println("4. Gorilla")
	// fmt.Print("Enter your choice: ")
	// choice := cmd.GetUserInput()
	// switch choice {
	// case 1:
	// 	Framework:="Fiber"
	// case 2:
	// 	Framework:="Gin"
	// case 3:
	// 	Framework:="Fast HTTP"
	// case 4:
	// 	Framework:="Gorilla"
	// default:
	// 	fmt.Println("Invalid choice. Please try again.")
	fmt.Print("Enter your project name: ")
	fmt.Scanln(&ProjectName)
	// fmt.Print("choose your dbms: ")
	// fmt.Println("1. MySQL")
	// fmt.Println("2. PostgreSQL")
	// fmt.Println("3. MongoDB")
	// fmt.Println("4. SQLite")
	// fmt.Print("Enter your choice: ")
	// choice := cmd.GetUserInput()
	// switch choice {
	// case 1:
	// 	bd:="MySQL"
	// case 2:
	// 	bd:="PostgreSQL"
	// case 3:
	// 	bd:="MongoDB"
	// case 4:
	// 	bd:="SQLite"
	// default:
	// 	fmt.Println("Invalid choice. Please try again.")
	// }
	fmt.Println("Creating project...")
	os.MkdirAll("ProjectName/models", 0755)
	os.Mkdir("ProjectName/routes", 0755)
	os.Mkdir("ProjectName/config", 0755)
	os.Mkdir("ProjectName/models", 0755)
	os.Mkdir("ProjectName/public", 0755)
	os.Mkdir("ProjectName/tests", 0755)
	os.Mkdir("ProjectName/docs", 0755)
	os.Create("ProjectName/main.go")
	os.Create("ProjectName/config/config.go")
	os.Create("ProjectName/models/models.go")
	os.Create("ProjectName/routes/routes.go")
	os.Create("ProjectName/tests/tests.go")
	os.Create("ProjectName/docs/README.md")
	exec.Command("cd", ProjectName).Run()
	exec.Command("go", "mod", "init", "github.com/"+username+"/"+ProjectName).Run()
	exec.Command("go", "get", "github.com/gofiber/fiber").Run()

	fmt.Println("Project created successfully!")

}

// func initFiber() {}
// func initGin() {}
// func initFastHTTP() {}
// func initGorilla() {}
