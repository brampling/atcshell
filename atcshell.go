package main

import "os"
import "gopkg.in/abiosoft/ishell.v2"
import "github.com/brampling/atccustomlist"

func main() {
	debugflag := false

	shell := ishell.New()
	shell.Print("\n")
	shell.Println("ActiveTrust Cloud Custom List Tool")

	shell.AddCmd(&ishell.Cmd{
        Name: "debug",
        Help: "Toggle debug mode",
        Func: func(c *ishell.Context) {
		debugflag = !debugflag
		c.Printf("Debug is %t\n", debugflag)
	},
	})
	shell.AddCmd(&ishell.Cmd{
        Name: "list",
        Help: "List all Custom Lists",
        Func: func(c *ishell.Context) {
		atccustomlist.CustomListList(debugflag)
		c.Print("\n")
	},
	})
	shell.AddCmd(&ishell.Cmd{
        Name: "addlist",
        Help: "Add a Custom List",
        Func: func(c *ishell.Context) {
		c.ShowPrompt(false)
	        defer c.ShowPrompt(true)
		c.Print("Custom List name: ")
		name := c.ReadLine()
		c.Print("Custom List description (optional): ")
		desc := c.ReadLine()
		atccustomlist.CustomListCreate(name, desc, debugflag)
		c.Print("\n")
	},
	})
	shell.AddCmd(&ishell.Cmd{
        Name: "dellist",
        Help: "Delete a Custom List",
        Func: func(c *ishell.Context) {
		c.ShowPrompt(false)
	        defer c.ShowPrompt(true)
		c.Print("Custom List id: ")
		id := c.ReadLine()
		atccustomlist.CustomListDelete(id, debugflag)
		c.Print("\n")
	},
	})
	shell.AddCmd(&ishell.Cmd{
        Name: "listitems",
        Help: "List the items in a Custom List",
        Func: func(c *ishell.Context) {
		c.ShowPrompt(false)
	        defer c.ShowPrompt(true)
		c.Print("Custom List id: ")
		id := c.ReadLine()
		atccustomlist.CustomListItemList(id, debugflag)
		c.Print("\n")
	},
	})
	shell.AddCmd(&ishell.Cmd{
        Name: "additem",
        Help: "Add an item to a Custom List",
        Func: func(c *ishell.Context) {
		c.ShowPrompt(false)
	        defer c.ShowPrompt(true)
		c.Print("Custom List id: ")
		id := c.ReadLine()
		c.Print("FQDN/IP to add: ")
		fqdn := c.ReadLine()
		atccustomlist.CustomListItemAdd(id, fqdn, debugflag)
		c.Print("\n")
	},
	})
	shell.AddCmd(&ishell.Cmd{
        Name: "delitem",
        Help: "Delete an item from a Custom List",
        Func: func(c *ishell.Context) {
		c.ShowPrompt(false)
	        defer c.ShowPrompt(true)
		c.Print("Custom List id: ")
		id := c.ReadLine()
		c.Print("FQDN/IP to delete: ")
		fqdn := c.ReadLine()
		atccustomlist.CustomListItemDelete(id, fqdn, debugflag)
		c.Print("\n")
	},
	})
	shell.Println(shell.HelpText())
	shell.Run()
	os.Exit(0)

}
