package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strings"
)

// ServerVersion : the version from the upstream repository
var ServerVersion = getDefaultVersion()

// GrafeasRepository : the repository of the Grafeas server to use for downloading Swagger files
var GrafeasRepository = "https://github.com/grafeas/grafeas"

// Reference : the commit, branch, or tag to use for downloading Swagger files from the defined Grafeas repo
var Reference = "master"

// APIVersion : the version of the Grafeas API per the Google Cloud scheme - https://github.com/grafeas/grafeas/blob/master/docs/versioning.md#api
var APIVersion = "v1beta1"

// SwaggerCodegenVersion : the version of the Swagger CodeGen CLI to download
var SwaggerCodegenVersion = "2.4.5"

// MergedClient : whether to keep the generated Swagger clients separate or to merge the paths
// TODO: implement function that merges the paths of each Swagger spec into one file before running codegen
var MergedClient = false

// trimVersionTag : remove the "v" prefix from the version tag string
func trimVersionTag(tag string) string {
	if strings.HasPrefix(tag, "v") {
		return strings.TrimPrefix(tag, "v")
	}
	return tag
}

func getDefaultVersion() string {
	version := os.Getenv("VERSION")
	if version == "" {
		tag := getMostRecentTag()
		trimmedVersion := trimVersionTag(tag)
		version = trimmedVersion
	}
	return version
}

func getMostRecentTag() string {
	tagsBytes, _ := get("https://api.github.com/repos/grafeas/grafeas/tags")
	var tags []map[string]interface{}
	json.Unmarshal(tagsBytes, &tags)
	return tags[0]["name"].(string)
}

func swaggerCompatibility(jsonInput string) string {
	re := regexp.MustCompile(`\/{.*(=.*)}`)
	jsonOutput := re.ReplaceAllStringFunc(jsonInput, func(substringMatch string) string {
		return strings.Split(substringMatch, "=")[0] + "}"
	})
	return jsonOutput
}

func get(url string) ([]byte, error) {
	var responseData []byte
	resp, err := http.Get(url)
	if err != nil {
		return responseData, err
	}
	defer resp.Body.Close()

	responseData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return responseData, err
	}
	return responseData, err
}

func downloadCompatibleSwaggerSpec(url string, filename string) error {
	if filename == "" {
		filename = path.Base(url)
	}
	responseBytes, err := get(url)
	if err != nil {
		return err
	}

	swaggerString := swaggerCompatibility(string(responseBytes))

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	responseReader := bytes.NewReader([]byte(swaggerString))

	_, err = io.Copy(f, responseReader)
	return err
}

func downloadFile(url string, filename string) error {
	if filename == "" {
		filename = path.Base(url)
	}
	responseBytes, err := get(url)
	if err != nil {
		return err
	}

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	responseReader := bytes.NewReader(responseBytes)

	_, err = io.Copy(f, responseReader)
	return err
}

type swaggerCodegenConfig struct {
	language        string
	inputSpec       string
	outputDirectory string
	configFile      string
	jar             string
}

func stringOrDefault(currentValue string, defaultValue string) string {
	if currentValue == "" {
		return defaultValue
	}
	return currentValue
}

func (c *swaggerCodegenConfig) generate() {
	c.language = stringOrDefault(c.language, "go")
	c.jar = stringOrDefault(c.jar, "./swagger-codegen-cli.jar")
	c.configFile = stringOrDefault(c.configFile, "./config.go.json")
	c.outputDirectory = stringOrDefault(c.outputDirectory, ServerVersion)
	args := []string{"-jar", c.jar, "generate", "-i", c.inputSpec, "-l", c.language, "-o", c.outputDirectory, "-c", c.configFile}
	log.Println("[CMD] java " + strings.Join(args, " "))
	cmd := exec.Command("java", args...)

	stderr, _ := cmd.StderrPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stderr)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	cmd.Wait()
}

func downloadSwaggerCodegenCli(version string) {
	downloadURL := fmt.Sprintf("https://repo1.maven.org/maven2/io/swagger/swagger-codegen-cli/%s/swagger-codegen-cli-%s.jar", version, version)
	downloadFile(downloadURL, "swagger-codegen-cli.jar")
}

func getSwaggerNames() []string {
	return []string{"grafeas.swagger.json", "project.swagger.json"}
}

func swaggerGenerate(mergedClient bool) {
	swaggerSpecNames := getSwaggerNames()
	for _, swaggerSpecName := range swaggerSpecNames {
		swaggerSpecURL := strings.Join([]string{GrafeasRepository, "/raw/", Reference, "/proto/", APIVersion, "/swagger/", swaggerSpecName}, "")
		log.Printf("[DOWNLOAD] %s\n", swaggerSpecURL)
		downloadCompatibleSwaggerSpec(swaggerSpecURL, "")
		fileBasename := strings.Split(swaggerSpecName, ".")[0]
		outputDirectory := strings.Join([]string{ServerVersion, fileBasename}, "/")
		if fileBasename == "grafeas" && !mergedClient {
			outputDirectory = ServerVersion
		}
		swaggerCodeGen := swaggerCodegenConfig{
			inputSpec:       swaggerSpecName,
			outputDirectory: outputDirectory,
		}
		swaggerCodeGen.generate()
	}
}

func main() {
	log.Printf(
		`

Grafeas Repository: 	 %s
Reference:		 %s
Server Version:     	 %s
API Version:    	 %s
Swagger Codegen Version: %s

`, GrafeasRepository, Reference, ServerVersion, APIVersion, SwaggerCodegenVersion)
	downloadSwaggerCodegenCli(SwaggerCodegenVersion)
	swaggerGenerate(MergedClient)
}
