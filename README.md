# 🧮 Test Calculator

[![test status](https://img.shields.io/github/actions/workflow/status/mikesmithgh/TestCalculator/test.yml?style=flat-square&logo=github&logoColor=c7c7c7&label=tests&labelColor=282828&event=push)](https://github.com/mikesmithgh/TestCalculator/actions/workflows/test.yml?query=event%3Apush)

A simple CLI calculator.

## Testing

Test Calculator provides two ways of testing the application, unit tests and integration tests.

Run unit tests and integration tests with the command:
```sh
make test
```

### Unit Tests

Unit tests follow the conventional structure of having test files in the same package as implementation files. For example,
the file [calculator.go](./calculator/calculator.go) in the `calculator` package has unit tests in the file [calculator_test.go](./calculator/calculator_test.go). 
The unit tests have access to the package variables and functions. Unit tests should test the independent components of the application.

Run unit tests with the command:
```sh
make test_unit
```

### Integration Tests

Integration tests have a dedicated `integration` package. This prevents exposing package level variables and functions that should
not be used during integration tests. Integration tests should test the entire application verifying that each component integrates 
with other components as expected.

Integration tests are initialized by the [init_test.go](./integration/init_test.go) `TestMain` function such that before each test it:
- Creates a temporary directory
- Builds the binary for `TestCalculator` in the temporary directory 
 
Each test has access to the `builtBinaryPath` variable in the `integration` package. Currently, there is only one test `TestCalculator` 
that follows the same table-driven structure as the unit tests. The test is provided an `input` slice of strings, `expected` string 
result, and `err` error result. The `input` is passed to the binary `builtBinaryPath`. Each `input` is successful if the output of 
`TestCalcuatlor` contains the string `expected` and equals the expected `err`, if `err` is defined.

Run integration tests with the command:
```sh
make test_integration
```

### Usage

The [Makefile](./Makefile) contains commands for building and testing Test Calculator.

To see the available commands run `make help`.

```
help             ==> describe make commands                                                                                                                                                                                                                                                  
test             ==> run unit tests and integration tests                                                                                                                                                                                                                                    
testjsonfmt      ==> run unit tests and integration tests with json output                                                                                                                                                                                                                   
test_unit        ==> run unit tests                                                                                                                                                                                                                                                          
test_integration ==> run integration tests                                                                                                                                                                                                                                                   
cover            ==> run test coverage                                                                                                                                                                                                                                                       
```

### Continuous Integration

Every pull request runs both unit and integration tests when it is open or a new commit is pushed. All tests must pass in order to merge
to the `main` branch. The output of the test results can be views in [Github Actions](https://github.com/mikesmithgh/TestCalculator/actions/workflows/test.yml) 
and is formatted with [gotestfmt](https://github.com/GoTestTools/gotestfmt) to improve readability. 

#### Example passings tests
![test-success](https://github.com/mikesmithgh/TestCalculator/assets/10135646/48805941-966f-4d37-8e35-4525b2164201)

#### Example test failures
![test-failure](https://github.com/mikesmithgh/TestCalculator/assets/10135646/05b87d57-3a8a-477e-b458-15fb6bbf1bc8)


The json output of the test run is saved as the artifact `test-log` for each run to allow downloading the results. `test-log` is available 
on the summary page of the [test](https://github.com/mikesmithgh/TestCalculator/actions/workflows/test.yml) workflow.

#### Example test-log artifact
![test-artifacts](https://github.com/mikesmithgh/TestCalculator/assets/10135646/c899c8bd-0d11-489f-bf8e-8e85c481d6f9)

To see the test results for the pipelines, view the [test](https://github.com/mikesmithgh/TestCalculator/actions/workflows/test.yml) workflow in Github Actions. Additionally, a [![test status](https://img.shields.io/github/actions/workflow/status/mikesmithgh/TestCalculator/test.yml?style=flat-square&logo=github&logoColor=c7c7c7&label=tests&labelColor=282828&event=push)](https://github.com/mikesmithgh/TestCalculator/actions/workflows/test.yml?query=event%3Apush)
status badge is provided at the top of this README to indicate if tests have passed for the latest commit on the `main` branch.




