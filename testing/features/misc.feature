# vim: ts=4 sw=4 sts=4 noet
@ci @smoke
Feature: Root command is informational only

    Scenario: Can get help info by running ishmael with no input
        When I run `ishmael`

        Then it should pass with "Usage:"
