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

	const blueColor = "\033[34m"
	const resetColor = "\033[0m"
	const banner = blueColor + `

                                                                                                                                               
                                                                                                                                               
MMMMMMMM               MMMMMMMM     OOOOOOOOO     LLLLLLLLLLL             DDDDDDDDDDDDD                        GGGGGGGGGGGGG     OOOOOOOOO     
M:::::::M             M:::::::M   OO:::::::::OO   L:::::::::L             D::::::::::::DDD                  GGG::::::::::::G   OO:::::::::OO   
M::::::::M           M::::::::M OO:::::::::::::OO L:::::::::L             D:::::::::::::::DD              GG:::::::::::::::G OO:::::::::::::OO 
M:::::::::M         M:::::::::MO:::::::OOO:::::::OLL:::::::LL             DDD:::::DDDDD:::::D            G:::::GGGGGGGG::::GO:::::::OOO:::::::O
M::::::::::M       M::::::::::MO::::::O   O::::::O  L:::::L                 D:::::D    D:::::D          G:::::G       GGGGGGO::::::O   O::::::O
M:::::::::::M     M:::::::::::MO:::::O     O:::::O  L:::::L                 D:::::D     D:::::D        G:::::G              O:::::O     O:::::O
M:::::::M::::M   M::::M:::::::MO:::::O     O:::::O  L:::::L                 D:::::D     D:::::D        G:::::G              O:::::O     O:::::O
M::::::M M::::M M::::M M::::::MO:::::O     O:::::O  L:::::L                 D:::::D     D:::::D        G:::::G    GGGGGGGGGGO:::::O     O:::::O
M::::::M  M::::M::::M  M::::::MO:::::O     O:::::O  L:::::L                 D:::::D     D:::::D        G:::::G    G::::::::GO:::::O     O:::::O
M::::::M   M:::::::M   M::::::MO:::::O     O:::::O  L:::::L                 D:::::D     D:::::D        G:::::G    GGGGG::::GO:::::O     O:::::O
M::::::M    M:::::M    M::::::MO:::::O     O:::::O  L:::::L                 D:::::D     D:::::D        G:::::G        G::::GO:::::O     O:::::O
M::::::M     MMMMM     M::::::MO::::::O   O::::::O  L:::::L         LLLLLL  D:::::D    D:::::D          G:::::G       G::::GO::::::O   O::::::O
M::::::M               M::::::MO:::::::OOO:::::::OLL:::::::LLLLLLLLL:::::LDDD:::::DDDDD:::::D            G:::::GGGGGGGG::::GO:::::::OOO:::::::O
M::::::M               M::::::M OO:::::::::::::OO L::::::::::::::::::::::LD:::::::::::::::DD    ......    GG:::::::::::::::G OO:::::::::::::OO 
M::::::M               M::::::M   OO:::::::::OO   L::::::::::::::::::::::LD::::::::::::DDD      .::::.      GGG::::::GGG:::G   OO:::::::::OO   
MMMMMMMM               MMMMMMMM     OOOOOOOOO     LLLLLLLLLLLLLLLLLLLLLLLLDDDDDDDDDDDDD         ......         GGGGGG   GGGG     OOOOOOOOO     
                                                                                                                                               
                                                                                                                                               
                                                                                                                                               
                                                                                                                                               
                                                                                                                                               
                                                                                                                                               
                                                                                                                                            
` + resetColor

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
	os.MkdirAll(ProjectName+"/models", 0755)
	os.Mkdir(ProjectName+"/routes", 0755)
	os.Mkdir(ProjectName+"/config", 0755)
	os.Mkdir(ProjectName+"/models", 0755)
	os.Mkdir(ProjectName+"/public", 0755)
	os.Mkdir(ProjectName+"/tests", 0755)
	os.Mkdir(ProjectName+"/docs", 0755)
	os.Create(ProjectName + "/main.go")
	os.Create(ProjectName + "/config/config.go")
	os.Create(ProjectName + "/models/models.go")
	os.Create(ProjectName + "/routes/routes.go")
	os.Create(ProjectName + "/tests/tests.go")
	os.Create(ProjectName + "/docs/README.md")

	cmd1 := exec.Command("go", "mod", "init", "github.com/"+username+"/"+ProjectName)
	cmd1.Dir = ProjectName
	cmd1.Run()
	cmd2 := exec.Command("go", "get", "github.com/gofiber/fiber")
	cmd2.Dir = ProjectName
	cmd2.Run()
	fmt.Println("Project created successfully!")

}
func CMD() {

}

// func initFiber() {}
// func initGin() {}
// func initFastHTTP() {}
// func initGorilla() {}
