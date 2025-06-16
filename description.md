# Go Challenge - Flight Price

## Description

We need to monitor flight prices between two cities: Miami, USA, and Santiago, Chile. Develop a Go microservice that retrieves flight prices from airline APIs and provides a structured, sorted response for a given origin and destination. The solution should emphasize concurrency, security, efficient data structures for processing and storing data, and comprehensive testing.

Users will enter a destination and receive a response containing:
* The cheapest available flight.
* The fastest available flight.
* A comparison of prices from multiple providers.

## Assignment
* Fetch flight prices from all five airline APIs provided: Americandes, Avionca, Copana, Latangos and Skylux.
* Compare the prices and time.
* Sort and structure the response.
* Implement JWT authentication.
* Ensure parallel API requests.

## API documentation
All five airline APIs—Americandes, Avionca, Copana, Latangos, and Skylux—are fully documented. Please review their documentation and perform a health check on each API before starting your implementation.

| API | Endpoint |
|-----|-----|
| Americandes | https://americandes.flylang.jobsityapps.com |
| Avionca | https://avionca.flylang.jobsityapps.com/graphql |
| Copana | https://copana.flylang.jobsityapps.com |
| Latangos | https://latangos.flylang.jobsityapps.com |
| Skylux | https://skylux.flylang.jobsityapps.com |


## Mandatory features
### RESTful Endpoints & Authentication
* GET /flights/search?origin=XXX&destination=YYY&date=YYYY-MM-DD → Returns flight price comparisons.
* Implement JWT authentication to restrict API access.

### Parallel API Fetching & Data Processing
* Make concurrent requests to at least three APIs.
* Aggregate and compare flight data from multiple providers.
* Identify and sort the cheapest and fastest flights.

### Security
* Keep credentials and tokens secure.
* Ensure only authenticated users can access the endpoints.
* Document how HTTPS/TLS should be configured for production.

### Testing
* Unit and integration tests for core aspect.
* Validate the accuracy of price comparisons.

## Bonus (optional)

* Docker: Provide a Dockerfile to run the application easily.
* Caching: Store the last fetched price for a short time (e.g., 30 seconds) to minimize API calls.
* WebSocket / SSE: A /subscribe/{symbol} endpoint allowing clients to receive periodic updates (e.g., every 30 seconds) via a time.Ticker.

## Considerations

* You are required to **create a video in which you explain the code you wrote**, the choices you made, how it works, and demonstrate the challenge's complete functionality. Remember that the execution must be shown from the beginning, it should not be running beforehand.
* Keep credentials and tokens secure.
* Show us in the **Readme** all relevant information about your project.
* Keep your code versioned with Git locally.
* Feel free to use small helper libraries.
* The project is focused on the backend, please have the frontend as **simple as you can**.
* This challenge was designed to be completed, including all the bonus tasks, **within 5 hours**.
