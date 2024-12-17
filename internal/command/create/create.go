package create

import (
	"fmt"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/spf13/cobra"
	"github.com/truxcoder/trux/internal/pkg/helper"
	"github.com/truxcoder/trux/tpl"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

type Create struct {
	ProjectName                string
	CreateType                 string
	FilePath                   string
	FileName                   string
	StructName                 string
	StructNameLowerFirst       string
	StructNameFirstChar        string
	StructNameSnakeCase        string
	StructNamePlural           string
	StructNamePluralLowerFirst string
	IsFull                     bool
}

func NewCreate() *Create {
	return &Create{}
}

var CmdCreate = &cobra.Command{
	Use:     "create [type] [handler-name]",
	Short:   "Create a new handler/service/repository/model",
	Example: "trux create handler user",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
var (
	tplPath   string
	pluralize *helper.Pluralize
)

func init() {
	CmdCreateHandler.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	CmdCreateService.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	CmdCreateRepository.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	CmdCreateModel.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	CmdCreateAll.Flags().StringVarP(&tplPath, "tpl-path", "t", tplPath, "template path")
	pluralize = helper.NewPluralize()
}

var CmdCreateHandler = &cobra.Command{
	Use:     "handler",
	Short:   "Create a new handler",
	Example: "trux create handler user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CmdCreateService = &cobra.Command{
	Use:     "service",
	Short:   "Create a new service",
	Example: "trux create service user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CmdCreateRepository = &cobra.Command{
	Use:     "repository",
	Short:   "Create a new repository",
	Example: "trux create repository user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CmdCreateModel = &cobra.Command{
	Use:     "model",
	Short:   "Create a new model",
	Example: "trux create model user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CmdCreateAPI = &cobra.Command{
	Use:     "api",
	Short:   "Create a new api",
	Example: "trux create api user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}
var CmdCreateAll = &cobra.Command{
	Use:     "all",
	Short:   "Create a new handler & service & repository & model & api",
	Example: "trux create all user",
	Args:    cobra.ExactArgs(1),
	Run:     runCreate,
}

func runCreate(cmd *cobra.Command, args []string) {
	c := NewCreate()
	c.ProjectName = helper.GetProjectName(".")
	c.CreateType = cmd.Use
	c.FilePath, c.StructName = filepath.Split(args[0])
	c.FileName = strings.ReplaceAll(c.StructName, ".go", "")
	c.StructName = strutil.UpperFirst(strutil.CamelCase(c.FileName))
	c.StructNameLowerFirst = strutil.LowerFirst(c.StructName)
	c.StructNameFirstChar = string(c.StructNameLowerFirst[0])
	c.StructNameSnakeCase = strutil.SnakeCase(c.StructName)
	c.StructNamePlural = pluralize.Plural(c.StructName)
	c.StructNamePluralLowerFirst = strutil.LowerFirst(c.StructNamePlural)

	switch c.CreateType {
	case "handler", "service", "repository", "model":
		c.genFile()
	case "api":
		c.FilePath = "api/v1/"
		c.genFile()
	case "all":
		c.CreateType = "handler"
		c.genFile()

		c.CreateType = "service"
		c.genFile()

		c.CreateType = "repository"
		c.genFile()

		c.CreateType = "model"
		c.genFile()

		c.CreateType = "api"
		c.FilePath = "api/v1"
		c.genFile()
	default:
		log.Fatalf("Invalid handler type: %s", c.CreateType)
	}

}
func (c *Create) genFile() {
	filePath := c.FilePath
	if filePath == "" {
		filePath = fmt.Sprintf("internal/%s/", c.CreateType)
	}
	f := createFile(filePath, strings.ToLower(c.FileName)+".go")
	if f == nil {
		log.Printf("warn: file %s%s %s", filePath, strings.ToLower(c.FileName)+".go", "already exists.")
		return
	}
	defer f.Close()
	var t *template.Template
	var err error
	if tplPath == "" {
		t, err = template.ParseFS(tpl.CreateTemplateFS, fmt.Sprintf("create/%s.tpl", c.CreateType))
	} else {
		t, err = template.ParseFiles(path.Join(tplPath, fmt.Sprintf("%s.tpl", c.CreateType)))
	}
	if err != nil {
		log.Fatalf("create %s error: %s", c.CreateType, err.Error())
	}
	err = t.Execute(f, c)
	if err != nil {
		log.Fatalf("create %s error: %s", c.CreateType, err.Error())
	}
	log.Printf("Created new %s: %s", c.CreateType, filePath+strings.ToLower(c.FileName)+".go")

}
func createFile(dirPath string, filename string) *os.File {
	filePath := filepath.Join(dirPath, filename)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create dir %s: %v", dirPath, err)
	}
	stat, _ := os.Stat(filePath)
	if stat != nil {
		return nil
	}
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Failed to create file %s: %v", filePath, err)
	}

	return file
}
