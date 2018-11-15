# vim: ts=4 sw=4 sts=4 noet
Feature: Alive subcommand checks container running status

	@smoke
	Scenario Outline: Can get help info by running ishmael with inquisitive inputs
		When I run `ishmael alive <INPUT>`

		Then it should pass with "Usage:"

		Examples:
			| INPUT  |
			| -h     |
			| --help |

	@smoke
	Scenario: The alive subcommand requires a container identifier/name
		When I run `ishmael alive`

		Then it should fail with "Usage:"

	Scenario: Can check existing container is running immediately
		Given I successfully run `docker run -dit --rm --name ishmael_tests_alive_1 alpine sh`

		When I run `ishmael alive ishmael_tests_alive_1 `

		Then the exit status should be 0
