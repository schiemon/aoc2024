package main

type Config struct {
	day             int
	puzzle          int
	pathToInputFile string
}

func NewConfig(day int, puzzle int, pathToInputFile string) *Config {
	return &Config{
		day:             day,
		puzzle:          puzzle,
		pathToInputFile: pathToInputFile,
	}
}

func NewConfigFromInputArgs(args InputArgs) (*Config, error) {
	pathToInputFile, err := GetPathToInputFile(args.pathToInputFolder, args.day, args.puzzle)

	if err != nil {
		return nil, err
	}

	return NewConfig(args.day, args.puzzle, pathToInputFile), nil
}

func (c *Config) Day() int {
	return c.day
}

func (c *Config) Puzzle() int {
	return c.puzzle
}

func (c *Config) PathToInputFile() string {
	return c.pathToInputFile
}
