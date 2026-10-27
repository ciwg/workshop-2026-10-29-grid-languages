package main

import (
	"bytes"

	"github.com/spf13/viper"
	. "github.com/stevegt/goadapt"
)

// a quick viper demo

var configYamlDemo = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

/*
func demoExample() {
	// read the config using viper
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(configYaml))
	// show the "age" value
	age := viper.GetInt("age")
	Pl("age", age)
	// show the "hobbies" value
	hobbies := viper.GetStringSlice("hobbies")
	Pl("hobbies", hobbies)
	// show the "eyes" value
	eyes := viper.GetString("eyes")
	Pl("eyes", eyes)
}
*/

// compiler is a dummy function to simulate a compiler execution.
func compiler(inputFile, outputFile string) {
	// simulate compiler execution
	Pl("Compiling", inputFile, "to", outputFile)
	// here you would add the actual compilation logic
}

// processConfig demonstrates how to read a config file and process
// the results
func processConfig(configYaml []byte) {
	// set defaults
	viper.SetDefault("compiler_input_file", "README.md")
	viper.SetDefault("compiler_output_file", "README.md")
	viper.SetDefault("fetch_file", "README.md")

	// read the config using viper
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(configYaml))
	inputFile := viper.GetString("compiler_input_file")
	outputFile := viper.GetString("compiler_output_file")
	// fetchFile := viper.GetString("fetch_file")

	// if the input file is the same as the output file, skip compilation
	if inputFile == outputFile {
		Pl("Skipping compilation as input file is the same as output file.")
	} else {
		outputFile := viper.GetString("compiler_output_file")
		// call the compiler function
		compiler(inputFile, outputFile)
	}

	Pl("compiler_input_file:", viper.GetString("compiler_input_file"))
	Pl("compiler_output_file:", viper.GetString("compiler_output_file"))
	Pl("fetch_file:", viper.GetString("fetch_file"))
}

func main() {
	// demoExample()

	/*
	   * no compiler execution:
	       * empty or missing config file
	           * if input_filename == fetch_filename, then compiler is skipped
	       * user edits [README.md](README.md)
	       * github displays [README.md](README.md)
	       * js calls fetch([README.md](README.md))
	       * i.e. no change from current behavior
	*/
	Pl("No compiler execution example:")
	var configNoCompiler = []byte(``)
	processConfig(configNoCompiler)

	/*
	* compiler execution:
	* compiler_input_filename: [README-in.md](README-in.md)
	* no other entries
	* user edits [README-in.md](README-in.md)
	* compiler creates [README.md](README.md)
	* github displays [README.md](README.md)
	* js fetch([README.md](README.md))
	 */
	Pl("Compiler execution example:")
	var configWithCompiler = []byte(`
	compiler_input_file: "README-in.md"
	compiler_output_file: "README.md"
	`)
	// XXX bug
	processConfig(configWithCompiler)

	/*
	   * compiler execution with post-compile customization:
	       * compiler_input_filename: [README-in.md](http://README-in.md)
	       * compiler_output_filename: [README-out.md](http://README-in.md)
	       * fetch_filename: [README.md](README.md)
	       * user edits [README-in.md](http://README-in.md)
	       * compiler creates [README-out.md](http://README.md)
	       * Makefile creates [README.md](README.md)
	       * github displays [README.md](http://README.md)
	       * js fetch([README.md](http://README.md))

	*/

}
