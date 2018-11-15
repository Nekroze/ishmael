# vim: ts=4 sw=4 sts=4 noet
Feature: Find subcommand gets the ID for a running instance of a docker compose project's service

	@smoke
	Scenario Outline: Can get help info by running ishmael with inquisitive inputs
		When I run `ishmael find <INPUT>`

		Then it should pass with "Usage:"

		Examples:
			| INPUT  |
			| -h     |
			| --help |

	@smoke
	Scenario: The find subcommand requires a project and service name
		When I run `ishmael find`

		Then it should fail with "Usage:"

	Scenario: Can get an existing container ID
		When I run `ishmael find ishmael tests`

		Then the exit status should be 0
		And the output should match /^[0-9a-z]+$/
