package module3rd

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/AikoCute-Offical/AikoR-Nginx/command"
	"github.com/AikoCute-Offical/AikoR-Nginx/util"
)

func DownloadAndExtractParallel(m Module3rd) {
	if util.FileExists(m.Name) {
		log.Printf("%s already exists.", m.Name)
		return
	}

	if m.Form != "local" {
		if len(m.Rev) > 0 {
			log.Printf("Download %s-%s.....", m.Name, m.Rev)
		} else {
			log.Printf("Download %s.....", m.Name)
		}

		logName := fmt.Sprintf("%s.log", m.Name)

		err := download(m, logName)
		if err != nil {
			util.PrintFatalMsg(err, logName)
		}
	} else if !util.FileExists(m.Url) {
		log.Fatalf("no such directory:%s", m.Url)
	}
}

func download(m Module3rd, logName string) error {
	form := m.Form
	url := m.Url

	switch form {
	case "git":
		fallthrough
	case "hg":
		args := []string{form, "clone", url}
		if command.VerboseEnabled {
			return command.Run(args)
		}

		f, err := os.Create(logName)
		if err != nil {
			return command.Run(args)
		}
		defer f.Close()

		cmd, err := command.Make(args)
		if err != nil {
			return err
		}

		writer := bufio.NewWriter(f)
		defer writer.Flush()

		cmd.Stderr = writer

		return cmd.Run()
	case "local": // not implemented yet
		return nil
	}

	return fmt.Errorf("form=%s is not supported", form)
}
