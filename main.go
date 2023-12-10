package main

func main() {
	bc := NewBlockchain()
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}

//bc.AddBlock("王显完成了区块链实验的第一步！")
//bc.AddBlock("王显实现了pow")
