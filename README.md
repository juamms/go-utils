# Go Utils package

## Available utilities

### `GetExecutablePath`
Returns the path of the current executable.

### `SafeJoinPaths`
Joins two given paths and applies the correct path separator.

### `StringToInt`
Converts a string to an int, and if the conversion fails uses the provided default value.

### `LoadJSONFromFile`
Loads a JSON file from a given path and populates the given model with the JSON file's contents.

### `WriteJSONToFile`
Writes a given model as JSON string to a given path.

### `ParseConfig`
Reads a `config.json` file from the same directory as the executable and parses its JSON to the given model.