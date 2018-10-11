package main

func initFile(_ []string) error {
	return eventStore().Init()
}
