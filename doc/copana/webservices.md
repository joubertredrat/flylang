# Copana Webservice API Documentation

## Overview
The Copana Webservice provides endpoints for system health checks, token authentication, and flight search operations. This API uses XML for request and response formats, adhering to webservice standards.

### Base URL
```
http://127.0.0.1:19003
```

---

## Endpoints

### 1. **System Ping**
**Endpoint**: `/webservices/copana/version27_8_5/system/ping.wsd`

**Method**: `POST`

**Description**: Checks the health of the system.

#### Response Example:
```xml
<response>
  <response>pong</response>
  <time>2025-04-12 14:30:00</time>
</response>
```

---

### 2. **Token Authentication**
**Endpoint**: `/webservices/copana/version27_8_5/authentication/token.wsd`

**Method**: `POST`

**Description**: Generates a token for authentication. Tokens are valid for 2 minutes.

#### Response Example:
```xml
<Token>
  <value>abc123xyz456</value>
  <expires_in>120 seconds</expires_in>
</Token>
```

---

### 3. **Flight Search**

**Endpoint**: `/webservices/copana/version27_8_5/flight/search.wsd`

**Method**: `POST`

**Description**: Searches for flights available.

#### Request Parameters:
| Parameter             | Type   | Required | Description                              | Example       |
|-----------------------|--------|----------|------------------------------------------|---------------|
| `iata:departure_code` | String | Yes      | IATA code of the departure airport       | `MIA`         |
| `iata:arrival_code`   | String | Yes      | IATA code of the arrival airport         | `SCL`         |
| `date:departure`      | String | Yes      | Departure date in `DD/MM/YY` format      | `12/04/25`    |
| `token`               | String | Yes      | Authentication token                     | `abc123xyz456`|

#### Response Example (Success):
```xml
<Flights>
  <Flight>
    <flight:code>CP001</flight:code>
    <date:departure>12/04/25 14:30</date:departure>
    <date:arrival>12/04/25 18:45</date:arrival>
    <amount:pricing>200.00</amount:pricing>
    <flight:route>MIA-SCL</flight:route>
  </Flight>
</Flights>
```

#### Response Example (Invalid Token):
```xml
<Error>
  <code>7</code>
  <message>Invalid or expired token</message>
</Error>
```

#### Response Example (Missing Fields):
```xml
<Error>
  <code>9</code>
  <message>Missing required fields: iata:from, iata:to, or date:departure</message>
</Error>
```

#### Response Example (Invalid Date Format):
```xml
<Error>
  <code>15</code>
  <message>Invalid date format. Use DD/MM/YY (e.g., 12/04/91).</message>
</Error>
```

---

## Error Codes
| Code | Message                          | Description                              |
|------|----------------------------------|------------------------------------------|
| 7    | Invalid or expired token         | The provided token is invalid or expired.|
| 9    | Missing required fields          | One or more required fields are missing. |
| 15   | Invalid date format              | The provided date format is incorrect.   |

---

## Token Management
- Tokens are valid for **2 minutes**.
- A new token can be generated using the `/authentication/token.wsd` endpoint.

---

## Notes
- All responses are in XML format.
- Ensure that the token is valid before making requests to the `/flight/search.wsd` endpoint.
- Use the correct date format (`DD/MM/YY`) for the `date:departure` parameter.

---

## Contact
For support or inquiries, contact the Webservices core team.
