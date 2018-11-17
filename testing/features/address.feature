# vim: ts=4 sw=4 sts=4 noet
Feature: Address subcommand gets container addresses

	@smoke
	Scenario Outline: Can get help info by running ishmael with inquisitive inputs
		When I run `ishmael address <INPUT>`

		Then it should pass with "Usage:"

		Examples:
			| INPUT  |
			| -h     |
			| --help |

	@smoke
	Scenario: The alive subcommand requires a container identifier/name
		When I run `ishmael address`

		Then it should fail with "Usage:"

	Scenario: Can get address to running container exposing single port
		Given I successfully run `docker run -dit --rm --expose 8080 --name ishmael_tests_address_1 alpine sh`

		When I run `ishmael address ishmael_tests_address_1`

		Then the exit status should be 0
		# https://regex101.com/r/Roie27/1
		And the output should match /^([0-9]+\.){3}[0-9]+:8080$/

	Scenario: Can get addresses to running container with host netowrking and exposing two ports
		Given I successfully run `docker run -dit --rm --net host --expose 8081 --expose 8082 --name ishmael_tests_address_2 alpine sh`

		When I run `ishmael address ishmael_tests_address_2`

		Then the exit status should be 0
		And the output should contain "127.0.0.1:8081"
		And the output should contain "127.0.0.1:8082"
