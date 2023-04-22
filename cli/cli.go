package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"

	"ca_parks/config"
	"ca_parks/core"
	"ca_parks/core/entity/user"
	"ca_parks/core/module/user/types"

	"github.com/manifoldco/promptui"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "GPG CLI"

	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		{
			Name:  "create-user",
			Usage: "Creates a user",
			Action: func(cliCtx *cli.Context) error {
				ctx := context.Background()
				config := config.NewWithPath("../config/base.yaml")

				fields := map[string]promptContent{

					"name": {
						label:     "Name",
						fieldType: "string",
					},

					"email": {
						label:     "Email",
						fieldType: "string",
					},

					"password": {
						label:     "Password",
						fieldType: "string",
					},
				}

				for k, f := range fields {
					res, err := promptGetInput(f)
					if err != nil {
						return err
					}
					f.value = res
					fields[k] = f
				}

				c, err := core.New(ctx, config)
				if err != nil {
					return err
				}

				res, err := c.User().Upsert(ctx, types.UpsertRequest{

					Name: fields["name"].value,

					Email: fields["email"].value,

					Password: fields["password"].value,

					Status: user.STATUS_ENABLED,
				}, false)

				if err != nil {
					return err
				}

				fmt.Printf("%v", res)

				return nil
			},
		},
		{
			Name:  "create-token",
			Usage: "Creates a token for a given user",
			Action: func(c *cli.Context) error {
				url := "http://localhost:8080/signin"
				fmt.Println("URL: ", url)

				fields := map[string]promptContent{
					"email": {
						label:     "Email",
						fieldType: "string",
					},

					"password": {
						label:     "Password",
						fieldType: "string",
					},
				}

				for k, f := range fields {
					res, err := promptGetInput(f)
					if err != nil {
						return err
					}
					f.value = res
					fields[k] = f
				}

				var jsonStr = []byte(fmt.Sprintf(`{"email":"%s", "password":"%s"}`,
					fields["email"].value,
					fields["password"].value,
				),
				)
				req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
				req.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					return err
				}
				defer resp.Body.Close()

				fmt.Println("Status:", resp.Status)
				body, _ := io.ReadAll(resp.Body)
				resData := map[string]string{}
				err = json.Unmarshal(body, &resData)
				if err != nil {
					return err
				}

				fmt.Println("Token:", resData["Token"])
				fmt.Println("Expires:", resData["Expires"])
				return nil

			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

type promptContent struct {
	label     string
	fieldType string
	value     string
}

func promptGetInput(pc promptContent) (string, error) {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New("required field")
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | bold }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", err
	}

	return result, nil
}
