Feature: Create datastore
    In order to analyze a datastore
    As a registered User
    I want to register the datastore into the platform

    Scenario: Create a new datastore
        Given I send a POST request to "http://localhost:1323/datastore/4bb2e094-2774-4a48-857a-b7faeebfdcfa" with a body:
        """
        {
            "name": "My greate datastore",
            "ip": "192.168.0.14",
            "path": "/home/dev"
        }
        """
        Then the response status code should be 201
            And the response should be empty