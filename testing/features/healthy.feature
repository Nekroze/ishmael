# vim: ts=4 sw=4 sts=4 noet
Feature: Healthy subcommand checks container healthy status

	@smoke
	Scenario Outline: Can get help info by running ishmael with inquisitive inputs
		When I run `ishmael healthy <INPUT>`

		Then it should pass with "Usage:"

		Examples:
			| INPUT  |
			| -h     |
			| --help |

	@smoke
	Scenario: The healthy subcommand requires a container identifier/name
		When I run `ishmael healthy`

		Then it should fail with "Usage:"

	Scenario: Can check existing container is healthy immediately
		Given I successfully run `docker run -dit --rm --health-cmd true --health-interval 500ms --name ishmael_tests_healthy_1 alpine sh`
		And I successfully run `sleep 1`

		When I run `ishmael healthy ishmael_tests_healthy_1 `

		Then the exit status should be 0
