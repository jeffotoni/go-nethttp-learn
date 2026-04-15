Read this in: [🇧🇷 Português](README_pt.md) | [🇪🇸 Español](README_es.md) | **🇺🇸 English**
---

# Backend, HTTP and API Architecture
### From the fundamentals of web communication to implementation with `net/http` in Go

<br>

![Go Version](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=flat-square&logo=go&logoColor=white)
![Status](https://img.shields.io/badge/status-in%20development-F59E0B?style=flat-square)
![License](https://img.shields.io/badge/license-MIT-10B981?style=flat-square)
![Author](https://img.shields.io/badge/author-jeffotoni-0EA5E9?style=flat-square&logo=github)

<br>

> **A consistent backend is not born from endpoints or frameworks.**
> It is born from a deep understanding of protocols, well-defined contracts, isolated business rules, state management, layered security, and operations thought out from the first commit.
> Code is a consequence. Decision is the foundation.

This material does not start at the route, the handler, or the framework. It starts before, at the root: at the origin of backend as a discipline, in the server's role within distributed systems, and in the semantics that make an API predictable, understandable, and operable in production. HTTP, REST, contracts, serialization, security, and observability appear here as parts of the same cohesive system, not as isolated topics. The implementation in Go comes later, when the conceptual base already sustains technical choices with precision and intention.

---

## ✦ About the author

Developed by **Jefferson Otoni Lima (Jeffotoni)**, Senior Software Engineer and Solutions Architect with over **22 years of experience** building high-performance distributed systems. Specialist in API design, cloud-native architecture, Go, and backend ecosystems at scale. Creator of the **Quick Framework**, author of **Go Bootcamp**, and active contributor to the Go community in Brazil and worldwide.

[![LinkedIn](https://img.shields.io/badge/LinkedIn-jeffotoni-0A66C2?style=flat-square&logo=linkedin)](https://www.linkedin.com/in/jeffotoni)
[![GitHub](https://img.shields.io/badge/GitHub-jeffotoni-181717?style=flat-square&logo=github)](https://github.com/jeffotoni)
[![Site](https://img.shields.io/badge/Site-jeffotoni.com-10B981?style=flat-square)](http://jeffotoni.com)

---

## 🚀 Quick Start

Get the project and host this documentation locally in seconds.

### 1. Clone the repository
```bash
git clone https://github.com/jeffotoni/go-nethttp-learn
cd go-nethttp-learn
```

### 2. Host the documentation locally
Choose your favorite tool to serve the `index.html`, CSS, and JS on port 3000:

**Using Go (The Gopher way)**
```bash
go run static/main.go
```

**Using Python**
```bash
python3 -m http.server 3000
```

**Using Node.js (npx)**
```bash
npx serve .
```

---

## Material Structure

| | Part | What you will find |
|:---:|---|---|
| **I** | Backend Fundamentals | Origin of the backend, role of the server, responsibilities, and pillars of modern backend |
| **II** | HTTP and API Semantics | Methods, status codes, body, cache, REST constraints, and correct API behavior |
| **III** | Servers and Infrastructure | Types of servers, serialization, operational context, and frequent decisions in real backend |
| **IV** | Security, Observability, and Contracts | Authentication, authorization, HTTPS, structured logs, metrics, tracing, and API versioning |
| **V** | Implementation with Go | `net/http` package in depth, without unnecessary abstraction and with a consolidated conceptual base |

---

## What is covered

| Block | What you will find |
|---|---|
| Backend Fundamentals | Origin of the backend, role of the server, responsibilities, and pillars of modern backend |
| HTTP and protocols | Evolution of HTTP, keep-alive, network layers, TCP, UDP, and communication context |
| REST and semantics | Constraints, resources, methods, status codes, payloads, and API maturity |
| Servers and infrastructure | Web servers, reverse proxy, serialization, and tools surrounding backend in production |
| Security | Authentication vs authorization, JWT, API keys, CORS, HTTPS, and rate limiting as a design decision |
| Observability | Structured logs, metrics, distributed tracing, and what an operable backend means |
| Contract Design | API versioning, documentation as part of the contract, and evolution without breaking clients |
| Go and `net/http` | Handlers, `Request`, `ResponseWriter`, `ServeMux`, and `Server` |
| API server in practice | Response standardization, validation, health endpoints, middleware, and local execution |

---

## Course Objectives

- Understand backend as a **system**, not as a collection of routes and handlers
- Master the fundamentals of **HTTP, REST, and API semantics** with conceptual precision
- Understand **security and observability** as design pillars, not as additional layers
- Learn to **define and evolve API contracts** without breaking clients or accumulating technical debt
- Connect the entire conceptual base with **practical and grounded implementation in Go**
- Develop the reasoning to **make technical decisions with clarity**, not just reproduce patterns
- Leave the course capable of building, operating, and evolving a **quality production backend**

---

## Content Track

| Stage | Theme |
|:---:|---|
| `1` | Backend fundamentals, web services, and client-server architecture |
| `2` | HTTP, connections, keep-alive, and protocol stack |
| `3` | REST, resources, API semantics, and maturity levels |
| `4` | Servers, reverse proxy, serialization, and infrastructure context |
| `5` | Security: authentication, authorization, HTTPS, JWT, and rate limiting |
| `6` | Observability: structured logs, metrics, and distributed tracing |
| `7` | Contract design, versioning, and API documentation |
| `8` | Overview of Go for building APIs |
| `9` | Fundamentals of `net/http` |
| `10` | Handlers, `ServeMux`, `Server`, and request/response flow |
| `11` | API server, security, and local execution with Docker |

---

## Official Manual Resources

| Channel | Link | Objective |
|---|---|---|
| Podcast | [Diving into backend](https://youtu.be/uk1hwBAKGLc) | Reinforce the conceptual context of the material |
| NotebookLM | [NotebookLM, Manual Chat](https://notebooklm.google.com/notebook/0421b1d1-9c27-415a-a3d2-bc83ce397b1f) | Ask, explore, and learn through chat, presentations, podcasts, and more |
| LinkedIn | [linkedin.com/in/jeffotoni](https://www.linkedin.com/in/jeffotoni) | Author's professional profile |
| GitHub | [github.com/jeffotoni](https://github.com/jeffotoni) | Author's repositories and projects |
| Go Roadmap | [github.com/jeffotoni/groadmap](https://github.com/jeffotoni/groadmap) | Macro view of study and evolution in Go |
| Site | [go-nethttp-learn](https://jeffotoni.github.io/go-nethttp-learn/) | Version of the repository published as a website |

---

## Go References

| Reference | Type | Link | Focus |
|---|:---:|---|---|
| Official Go Site | Official | [go.dev](https://go.dev/) | Main language portal |
| ChatBoot with Go | Official | [ChatBoot Google Go](https://codewiki.google/github.com/golang/go#community-guidelines-and-support) | Official language assistant |
| Official Tutorial | Official | [go.dev/doc/tutorial](https://go.dev/doc/tutorial/) | Step-by-step to get started |
| Tour of Go | Official | [go.dev/tour/welcome/1](https://go.dev/tour/welcome/1) | Interactive learning |
| Language Specification | Official | [go.dev/ref/spec](https://go.dev/ref/spec) | Formal language rules |
| Effective Go | Official | [go.dev/doc/effective_go](https://go.dev/doc/effective_go) | Style and best practices |
| Release Notes | Official | [go.dev/doc/devel/release](https://go.dev/doc/devel/release) | Version history |
| Go 1.26 Version Notes | Official | [go.dev/doc/go1.26](https://go.dev/doc/go1.26) | What's new in version 1.26 |
| Go 1.26 Blog | Official | [go.dev/blog/go1.26](https://go.dev/blog/go1.26?ref=dailydev) | Practical release explanations |
| Go by Example | Community | [gobyexample.com](https://gobyexample.com) | Direct and short examples |
| Quick Framework | Community | [github.com/jeffotoni/quick](https://github.com/jeffotoni/quick) | Lightweight framework for APIs in Go |

---

## Jeffotoni References: Go and Architectureeferences: Go and Architecture

<details>
<summary><strong>View all projects and repositories</strong></summary>

<br>

| Project | Link | Focus |
|---|---|---|
| Go Bootcamp | [gobootcamp.jeffotoni.com](https://gobootcamp.jeffotoni.com/br/index.html) | Complete Go learning track |
| Personal Site | [jeffotoni.com](http://jeffotoni.com) | Content, articles, and author materials |
| Go Manual | [gomanual.jeffotoni.com](https://gomanual.jeffotoni.com/) | Go reference manual |
| Go Roadmap | [github.com/jeffotoni/groadmap](https://github.com/jeffotoni/groadmap) | Macro view of the Go journey |
| Quick Framework | [github.com/jeffotoni/quick](https://github.com/jeffotoni/quick) | Lightweight and performant framework for APIs in Go |
| Quick Benchmarks | [github.com/goquick-run/benchmarks](https://github.com/goquick-run/benchmarks) | Performance comparisons between frameworks |
| Go Example | [github.com/jeffotoni/goexample](https://github.com/jeffotoni/goexample) | Collection of practical examples in Go |
| Go Cache | [github.com/jeffotoni/gocache](https://github.com/jeffotoni/gocache) | Cache strategies in Go |
| Go Hexagonal DDD | [github.com/jeffotoni/go-hexagonal-ddd](https://github.com/jeffotoni/go-hexagonal-ddd) | Hexagonal Architecture and DDD in Go |
| Go gRPC Lecture | [github.com/jeffotoni/gogrpc.palestra](https://github.com/jeffotoni/gogrpc.palestra) | gRPC materials and examples in Go |
| Go Workshop DevOps | [github.com/jeffotoni/goworkshopdevops](https://github.com/jeffotoni/goworkshopdevops) | Go practices applied to DevOps |
| Benchmark | [github.com/jeffotoni/benchmark](https://github.com/jeffotoni/benchmark) | Benchmark studies and analysis |

</details>

---

## Table of Contents

- [Material Structure](#material-structure)
- [What is covered](#what-is-covered)
- [Course Objectives](#course-objectives)
- [Content Track](#content-track)
- [Official Manual Resources](#official-manual-resources)
- [Go References](#go-references)
- [Jeffotoni References](#jeffotoni-references-go-and-architecture)
- [1. Context: Web Services, REST, and Protocols](#1-context-web-services-rest-and-protocols)
  - [Web Services Overview](#web-services-overview)
  - [Communication Diagrams](#communication-diagrams)
  - [Rapid Evolution of HTTP](#rapid-evolution-of-http)
  - [Keep-Alive Diagram](#keep-alive-diagram)
  - [Keep-Alive: HTTP/1.0 -> HTTP/1.1 -> HTTP/2](#keep-alive-http10---http11---http2)
  - [HTTP, TCP, and UDP (Quick Difference)](#http-tcp-and-udp-quick-difference)
  - [OSI Model and TCP/IP (Diagram)](#osi-model-7-layers)
  - [TCP/IP Model (4 Layers)](#tcpip-model-4-layers)
  - [REST vs RESTful](#rest-vs-restful)
  - [Acronym Meanings](#acronym-meanings)
  - [REST Constraints](#rest-constraints)
  - [REST Constraints Diagram](#rest-constraints-diagram)
  - [Uniform Interface (detailed in 4 parts)](#uniform-interface-detailed-in-4-parts)
  - [Maturity Levels (Richardson)](#maturity-levels-richardson)
  - [Richardson Diagram](#richardson-diagram)
  - [HTTP Methods (HTTP Verbs)](#http-methods-http-verbs)
  - [Body in REST (request/response) with status in practice](#body-in-rest-requestresponse-with-status-in-practice)
  - [Essential Status Codes for APIs](#essential-status-codes-for-apis)
  - [Serialization Formats](#serialization-formats)
  - [Web and Application Servers](#web-and-application-servers)
  - [Web Servers/Reverse Proxy made in Go](#web-serversreverse-proxy-made-in-go)
  - [Go Ecosystem in DevOps](#go-ecosystem-in-devops)
- [2. Overview of Go for APIs](#2-overview-of-go-for-apis)
  - [What is Go](#what-is-go)
  - [Go Differentials for Building APIs](#go-differentials-for-building-apis)
  - [Concurrency in Go (Simple to Understand)](#concurrency-in-go-simple-to-understand)
  - [Compiled, Static, and Dynamic (In Practice)](#compiled-static-and-dynamic-in-practice)
  - [Built-in HTTP Server](#http-server-built-in)
  - [Official Language Keywords (25)](#official-language-keywords-25)
- [3. `net/http` Fundamentals](#3-nethttp-fundamentals)
  - [The `net/http` Package](#the-nethttp-package)
  - [Mini Reference of Components](#mini-reference-of-components)
  - [Minimal Anatomy of a Handler (`w` and `r`)](#minimal-anatomy-of-a-handler-w-and-r)
- [4. Practical Manual: ListenAndServe (Phase Zero)](#4-practical-manual-listenandserve-phase-zero)
  - [4.1 Essential Difference: `HandleFunc` vs `HandlerFunc`](#41-essential-difference-handlefunc-vs-handlerfunc)
  - [4.2 What `ListenAndServe` Accepts](#42-what-listenandserve-accepts)
  - [4.3 Base Variations (without custom `ServeMux`)](#43-base-variations-without-custom-servemux)
  - [4.4 Some Possibilities](#44-some-possibilities)
  - [4.5 `ServeMux` with method pattern + `http.Server`](#45-servemux-with-method-pattern--httpserver)
  - [4.6 When to use `http.Handler`?](#46-when-to-use-httphandler)
  - [4.7 Graceful Shutdown](#47-graceful-shutdown)
  - [4.8 Testing handlers with `httptest`](#48-testing-handlers-with-httptest)
  - [4.9 Two Servers on Distinct Ports with Goroutines](#49-two-servers-on-distinct-ports-with-goroutines)
- [5. Server API](#5-server-api)
  - [5.0 helpers.go: shared functions](#50-helpersgo-shared-functions)
  - [5.1 Response Standardization](#51-response-standardization)
  - [5.2 Error and Status Map by Scenario](#52-error-and-status-map-by-scenario)
  - [5.3 Route Organization](#53-route-organization)
  - [5.4 Server Input Validation](#54-server-input-validation)
  - [5.5 Health Endpoints](#55-health-endpoints)
  - [5.6 Basic Auth Middleware](#56-basic-auth-middleware)
  - [5.7 Environment Variables](#57-environment-variables)
  - [5.8 CORS Middleware](#58-cors-middleware)
  - [5.9 Documentation with OpenAPI/Swagger](#59-documentation-with-openapiswagger)
- [6. Docker: Build and Run Local](#6-docker-build-and-run-local)
  - [6.1 Multi-stage Dockerfile (Alpine + Brazil timezone)](#61-multi-stage-dockerfile-alpine--brazil-timezone)
  - [6.2 Basic Docker Commands](#62-basic-docker-commands)

---

## 1. Context: Web Services, REST, and Protocols

### Web Services Overview

| Style/Technology | Year (origin/appearance) | Main Characteristic | When it appears most |
|---|---:|---|---|
| SOAP | 1998 | Rigid contract, XML, WSDL | Corporate legacy and formal integrations |
| REST | 2000 | Architectural style over HTTP | Web APIs in general |
| WebHooks | 2007 | Event-oriented HTTP callback (server push to client) | Payments, integrations, CI/CD pipelines |
| WebSocket | 2011 | Bidirectional and persistent connection over HTTP | Chats, real-time dashboards, games |
| SSE (Server-Sent Events) | 2006 | Server pushes events over standard HTTP (unidirectional) | Notifications, response streaming, LLMs |
| GraphQL | 2015 | Client defines response fields | Scenarios with multiple data views |
| gRPC | 2015 | RPC with Protobuf over HTTP/2 | Internal microservices communication |


### Communication Diagrams

#### SOAP
![SOAP](docs/diagrams/soap.svg)

#### REST
![REST](docs/diagrams/rest.svg)

#### WebHooks
![WebHooks](docs/diagrams/webhooks.svg)

#### WebSocket
![WebSocket](docs/diagrams/websocket.svg)

#### SSE (Server-Sent Events)
![SSE](docs/diagrams/sse.svg)

#### GraphQL
![GraphQL](docs/diagrams/graphql.svg)

#### gRPC
![gRPC](docs/diagrams/grpc.svg)

### Rapid Evolution of HTTP

| Protocol | Year | Highlights |
|---|---:|---|
| HTTP/0.9 | 1991 | Original version; GET only, no headers, no status code |
| HTTP/1.0 | 1996 | Headers, status codes, and multiple content types; connection closed by default |
| HTTP/1.1 | 1997 | Keep-alive standard, pipelining, mandatory host; web base for decades |
| HTTP/2 | 2015 | Binary, multiplexing, header compression (HPACK), server push |
| HTTP/3 | 2022 | QUIC over UDP, lower latency on unstable networks, more resilient connection |

### Keep-Alive: HTTP/1.0 -> HTTP/1.1 -> HTTP/2

![Keep-Alive](docs/diagrams/keepalive.svg)


```text
HTTP/1.0 (1996)
├─ Connection closed after EACH request/response
├─ Keep-Alive was OPTIONAL (via header)
└─ Header: Connection: keep-alive (explicit)

HTTP/1.1 (1997)
├─ Keep-Alive is the DEFAULT
├─ Persistent connections by default
├─ To close: Connection: close
└─ Better performance out-of-the-box

HTTP/2 (2015)
├─ Multiplexing over a single connection
├─ Implicit Keep-Alive
└─ Multiple simultaneous requests
```

#### HTTP/1.0 - Activating Keep-Alive

```http
GET /index.html HTTP/1.0
Host: example.com
Connection: keep-alive
```

```http
HTTP/1.0 200 OK
Connection: keep-alive
Keep-Alive: timeout=5, max=100
Content-Type: text/html
Content-Length: 1234

<html>...</html>
```

#### HTTP/1.1 - Keep-Alive by Default

```http
GET /index.html HTTP/1.1
Host: example.com
```

```http
HTTP/1.1 200 OK
Content-Type: text/html
Content-Length: 1234

<html>...</html>
```

#### HTTP/1.1 - Explicitly Closing Connection

```http
GET /logout HTTP/1.1
Host: example.com
Connection: close
```

```http
HTTP/1.1 200 OK
Connection: close
Content-Type: application/json

{"status":"logged out"}
```

### HTTP, TCP, and UDP (Quick Difference)

Educational reference:
- OSI Model (7 layers)
- TCP/IP Model (4 layers, most used in practice)

#### OSI Model (7 Layers)

![OSI Model and TCP/IP](docs/diagrams/osi-tcpip.svg)


| Layer | Name | Main Function | Protocol Examples / Technologies |
|---:|---|---|---|
| 7 | Application | Interface with user and applications | HTTP, HTTPS, FTP, SMTP, DNS |
| 6 | Presentation | Formatting, encryption, compression | SSL/TLS, JPEG, MP3, JSON |
| 5 | Session | Session/connection control | NetBIOS, RPC |
| 4 | Transport | End-to-end communication, error control | TCP, UDP |
| 3 | Network | Logical addressing and routing | IP, ICMP, IPSec |
| 2 | Data Link | Communication within local network | Ethernet, ARP, PPP |
| 1 | Physical | Electrical/optical transmission of bits | Cables, Fiber, Wi-Fi (physical part) |

#### TCP/IP Model (4 Layers)

| TCP/IP Layer | OSI Equivalent | Examples |
|---|---|---|
| Application | 7, 6, and 5 | HTTP, FTP, SMTP, DNS |
| Transport | 4 | TCP, UDP |
| Internet | 3 | IP, ICMP |
| Network Access | 2 and 1 | Ethernet, Wi-Fi |

Quick summary of the stack:
- HTTP/1.1 and HTTP/2: `HTTP -> TCP -> IP`
- HTTP/3: `HTTP -> QUIC(UDP) -> IP`

Educational analogy (message and mail):
1. Application: writing the message
2. Transport: putting it in the envelope (TCP checks if it arrived)
3. Network: choosing the route to the destination
4. Data Link: taking it to the local post office
5. Physical: road and truck

Basic example in Go:

```go
http.ListenAndServe(":8080", nil)
```

That is, `net/http` is at the top of the stack but depends on all layers below.

### REST vs RESTful

- `REST` is an architectural style (set of constraints)
- `RESTful` is the API that applies REST consistently in practice

### Acronym Meanings

| Term | Meaning | Type | Where it fits |
|---|---|---|---|
| HTTP | HyperText Transfer Protocol | Protocol | Application Layer |
| REST | Representational State Transfer | Architectural Style | Uses HTTP |
| SOAP | Simple Object Access Protocol | Protocol | Uses HTTP (usually) |
| gRPC | Google Remote Procedure Call | Framework / RPC | Uses HTTP/2 |

### REST Constraints

![REST Constraints](docs/diagrams/rest-constraints.svg)


```text
┌───────────────────────────────────────────────────────┐
│ 1. Client-Server                                      │
│    Separation of concerns                             │
│                                                       │
│ 2. Stateless                                          │
│    Each request is independent                        │
│                                                       │
│ 3. Cacheable                                          │
│    Responses must indicate caching                    │
│                                                       │
│ 4. Uniform Interface                                  │
│    ├─ Resource identification                         │
│    ├─ Manipulation via representations                │
│    ├─ Self-descriptive messages                       │
│    └─ HATEOAS                                         │
│                                                       │
│ 5. Layered System                                     │
│    Client doesn't know if it connects directly to     │
│    the end server or intermediate layers               │
│                                                       │
│ 6. Code on Demand (optional)                          │
│    Server can send executable code                    │
└───────────────────────────────────────────────────────┘
```

#### Educational preview of each constraint

| Constraint | Practical Preview |
|---|---|
| Client-Server | Frontend and backend evolve independently |
| Stateless | Token/authentication and context go in each request |
| Cacheable | Use of `Cache-Control`, `ETag`, `Last-Modified` |
| Uniform Interface | URI + method + status + consistent representation |
| Layered System | CDN, load balancer, and API gateway between client and app |
| Code on Demand | Ex.: JavaScript delivered to the client (optional) |

#### Uniform Interface (detailed in 4 parts)

**1. Resource Identification**

Each resource has a unique identifier (URI).

Examples:
- `/users/123`
- `/posts/456`

**2. Manipulation via Representations**

Client manipulates resources through representations (`JSON`, `XML`, etc).
The server sends the representation of the resource, not the resource in memory.

Practical example:
- the client receives user JSON
- when doing `PUT /users/123`, it sends a new representation of that user

**3. Self-descriptive Messages**

Each message must contain enough information for processing.

```http
Content-Type: application/json
Accept: application/json
```

With this, client and server understand input/output format without a "hidden agreement".

**4. HATEOAS (Hypermedia As The Engine Of Application State)**

The API returns links for next valid actions, and the client navigates through these links.

```json
{
  "id": 123,
  "name": "John",
  "links": [
    {"rel": "self", "href": "/users/123"},
    {"rel": "posts", "href": "/users/123/posts"},
    {"rel": "delete", "href": "/users/123", "method": "DELETE"}
  ]
}
```

In practice, HATEOAS is the least implemented item in most RESTful APIs.

Common reasons:
- mobile/web clients prefer a fixed contract documented in OpenAPI/Swagger
- teams prioritize simplicity of implementation and maintenance
- gateways, versioning, and SDKs tend to centralize flow outside hypermedia
- extra modeling cost doesn't always generate clear benefit in the product

### Maturity Levels (Richardson)

![Richardson Maturity Levels](docs/diagrams/richardson.svg)


The model was proposed by **Leonard Richardson**, a software architect who wrote about REST APIs and helped popularize best practices in building HTTP services.

Model objective:
- evaluate how RESTful an API is
- classify HTTP APIs into maturity levels
- help evolve APIs from disguised RPC to more well-structured REST

It has 4 levels (0 to 3).

| Level | Name | Description |
|---:|---|---|
| 0 | POX / RPC over HTTP | HTTP only as transport |
| 1 | Resources | Resources identified by URI |
| 2 | Verbs + status | Correct use of HTTP verbs and status codes |
| 3 | HATEOAS | Hypermedia guiding the client |

#### Level 0 - The Swamp of POX

- uses HTTP only as transport
- usually a single endpoint
- common to see `POST` for everything

Example:

```http
POST /api
Content-Type: application/json

{
  "action": "getUser",
  "id": 10
}
```

Here HTTP becomes just a "tunnel" for RPC commands.

#### Level 1 - Resources

- separates by resources (different URLs)
- may still use `POST` for almost everything

Resource examples:
- `/users`
- `/orders`

Main gain: beginning of organization by domain.

#### Level 2 - Correct HTTP Verbs

- uses `GET`, `POST`, `PUT`, `DELETE` correctly
- uses appropriate status codes

Examples:

```http
GET /users/10
DELETE /users/10
```

Here are most of the APIs that the market calls REST in practice.

#### Level 3 - HATEOAS

`HATEOAS` = *Hypermedia As The Engine Of Application State*.

The response includes links to possible next steps.

```json
{
  "id": 10,
  "name": "Jefferson",
  "links": [
    {"rel": "orders", "href": "/users/10/orders"},
    {"rel": "delete", "href": "/users/10"}
  ]
}
```

Here the API guides the client dynamically.

In practice:
- most modern APIs stay at level 2
- few implement HATEOAS fully
- many APIs call themselves REST but are still at level 1

### HTTP Methods (HTTP Verbs)

| Verb | Correct Use | Example |
|---|---|---|
| `GET` | Fetch data | `GET /users/123` |
| `POST` | Create | `POST /users` |
| `PUT` | Replace | `PUT /users/123` |
| `PATCH` | Partially update | `PATCH /users/123` |
| `DELETE` | Remove | `DELETE /users/123` |

### Body in REST (request/response) with status in practice

Simple rules:
- `GET` and `DELETE`: normally without body
- `POST`, `PUT`, `PATCH`: normally with body
- Always define `Content-Type` and validate input

#### GET (fetch resource)

```http
GET /users/123
Accept: application/json
```

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "id": 123,
  "name": "Jeff Otoni"
}
```

```http
HTTP/1.1 404 Not Found
Content-Type: application/json

{
  "error": "user_not_found"
}
```

#### POST (create resource)

```http
POST /users
Content-Type: application/json

{
  "name": "Jeff Otoni",
  "email": "jeff@email.com"
}
```

```http
HTTP/1.1 201 Created
Location: /users/123
Content-Type: application/json

{
  "id": 123,
  "name": "Jeff Otoni",
  "email": "jeff@email.com"
}
```

```http
HTTP/1.1 422 Unprocessable Entity
Content-Type: application/json

{
  "error": "validation_failed",
  "details": {
    "email": "invalid format"
  }
}
```

#### PUT (replace resource)

```http
PUT /users/123
Content-Type: application/json

{
  "name": "Jeff Otoni",
  "email": "new@email.com"
}
```

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "id": 123,
  "name": "Jeff Otoni",
  "email": "new@email.com"
}
```

```http
HTTP/1.1 404 Not Found
Content-Type: application/json

{
  "error": "user_not_found"
}
```

#### PATCH (partial update)

```http
PATCH /users/123
Content-Type: application/json

{
  "email": "patch@email.com"
}
```

```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "id": 123,
  "name": "Jeff Otoni",
  "email": "patch@email.com"
}
```

#### DELETE (remove resource)

```http
DELETE /users/123
```

```http
HTTP/1.1 204 No Content
```

```http
HTTP/1.1 404 Not Found
Content-Type: application/json

{
  "error": "user_not_found"
}
```

### Essential Status Codes for APIs

| Scenario | Status |
|---|---|
| Success with return | `200 OK` |
| Resource creation | `201 Created` |
| Success without body | `204 No Content` |
| Input error | `400 Bad Request` |
| Not authenticated | `401 Unauthorized` |
| No permission | `403 Forbidden` |
| Not found | `404 Not Found` |
| State conflict | `409 Conflict` |
| Semantic validation error | `422 Unprocessable Entity` |
| Internal error | `500 Internal Server Error` |
| Too many requests | `429 Too Many Requests` |

### Serialization Formats

For this course, the main focus will be `JSON` in REST APIs with Go.

| Format | Type | When to use |
|---|---|---|
| JSON | Text | Public REST APIs and simplicity |
| Protobuf | Binary | gRPC and high-performance internal communication |
| Avro | Binary | Streaming/Kafka with strong schema evolution |
| MessagePack | Binary | More compact payload without too much complexity |
| CBOR | Binary | IoT and scenarios with IETF standard |

Minimal example in Go:

```go
package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Serialize (struct -> JSON)
	u := User{Name: "Jeff", Email: "jeff@email.com"}
	b, _ := json.Marshal(u)
	fmt.Println(string(b)) // {"name":"Jeff","email":"jeff@email.com"}

	// Deserialize (JSON -> struct)
	var u2 User
	_ = json.Unmarshal(b, &u2)
	fmt.Println(u2.Name) // Jeff
}
```

### Web and Application Servers

| Server | Year | Category | Note |
|---|---:|---|---|
| Apache HTTP Server | 1995 | Web Server | Historical base of open source web |
| IIS | 1995 | Web Server | Microsoft's web server |
| nginx | 2004 | Web Server/Reverse Proxy | Widely used in high concurrency |
| Caddy | 2015 | Web Server | Automatic HTTPS by default |
| Tomcat | 1999 | Application Server (Java) | Very common in Java applications |
| JBoss / WildFly | 2006 (WildFly 2014) | Application Server (Java) | Enterprise line of Java ecosystem |

### Web Servers/Reverse Proxy made in Go

| Project | Category | Where it appears a lot | Why Go helps here |
|---|---|---|---|
| Caddy | Web server / reverse proxy | APIs, automatic TLS, simple edge | Single binary, native concurrency, and easy deployment |
| Traefik | Reverse proxy / ingress | Docker, Kubernetes, service discovery | Cloud-native integration and high network performance |
| Fabio | Load balancer / reverse proxy | Environments with Consul | Operational simplicity and good concurrent model |
| `httputil.ReverseProxy` | Reverse proxy (stdlib) | Internal APIs, simple proxies without extra dependency | Native in stdlib `net/http/httputil`, zero dependencies |

### Market Share (macro view)

![Web Servers Market Share](docs/diagrams/market-share.svg)

### Go Ecosystem in DevOps

Go has become one of the central languages of the **CNCF/DevOps** ecosystem by delivering:
- Portable and simple to distribute binaries
- Good network performance and concurrency
- Stable toolchain for infrastructure projects

| Tool | Category | Relation with Go |
|---|---|---|
| Docker (Moby/Engine) | Containerization | Central implementation in Go (with parts in other languages) |
| Kubernetes | Orchestration | Core project in Go |
| Consul | Service discovery/config | Core in Go |
| etcd | Distributed KV | Core in Go |
| Terraform | Infrastructure as Code | Core in Go |
| Vault | Secrets management | Core in Go |
| CockroachDB | Distributed SQL Database | Core mostly in Go |
| InfluxDB | Time-series database | Strong use of Go in core |
| Prometheus | Monitoring | Core in Go |
| Grafana | Observability | Backend in Go (frontend in TypeScript) |
| Gitea | Git forge/self-hosted | Core in Go |
| Helm | Kubernetes package manager | Core in Go |
| ArgoCD | GitOps / CD for Kubernetes | Core in Go |
| Cilium | Networking / eBPF for Kubernetes | Core in Go |

---

## 2. Overview of Go for APIs

### What is Go

Go is a compiled, statically typed language with a simple syntax, focused on productivity, performance, and readability.

### Release Year and Key Names

| Item | Information |
|---|---|
| Project start | 2007 (Google) |
| Public release | 2009 |
| Version 1.0 | 2012 |
| Creators | Robert Griesemer, Rob Pike, Ken Thompson |

### Go Differentials for Building APIs

- Strong standard library (`net/http`, `encoding/json`, `context`, `database/sql`)
- Fast compilation and simple deployment (single binary)
- Native concurrency with goroutines and channels
- More predictable code with less accidental complexity
- Excellent robustness for high-load, low-latency APIs
- Integrated tests in the toolchain (`go test`) with practical support for unit and table-driven tests
- Native coverage, benchmark, and fuzz testing (`-cover`, `-bench`, `-fuzz`) to increase API reliability
- Native `context.Context` for cancellation, timeout, and value propagation between handlers and goroutines
- Native cross-compilation: compile to any OS/architecture with `GOOS` and `GOARCH` without additional toolchain

### Concurrency in Go (Simple to Understand)

- `goroutine`: function executing concurrently at low cost
- `channel`: secure channel for communication between goroutines
- `select`: coordinates multiple channels and timeouts

Mental model:
1. Use goroutines to structure concurrent work (not to be confused with parallelism)
2. Exchange data via channels (instead of always sharing memory)
3. Control cancellation and deadline with `context.Context`
4. The runtime/scheduler decides when there is real parallelism (e.g., multiple cores)

Minimal example with goroutine + channel:

```go
package main

import "fmt"

func sum(a, b int, ch chan int) {
	ch <- a + b
}

func main() {
	ch := make(chan int)
	go sum(3, 7, ch) // executes concurrently
	result := <-ch  // waits for the value to arrive
	fmt.Println(result) // 10
}
```

Example with `context.Context`, timeout in handler:

```go
func handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		http.Error(w, "timeout", http.StatusGatewayTimeout)
	case result := <-fetchData(ctx):
		w.Write(result)
	}
}
```

### Compiled, Static, and Dynamic (In Practice)

| Aspect | How it works in Go |
|---|---|
| Compilation | AOT (ahead-of-time), generates native binary |
| Typing | Static and strong at compile time |
| Linking | Usually static; can use dynamic in scenarios with `cgo` |
| Runtime | Dynamic for GC, scheduler, and reflection when necessary |
| Cross-compilation | `GOOS=linux GOARCH=amd64 go build` generates binary for any platform |

Example of cross-compilation:

```bash
# Compile for Linux AMD64 (from any OS)
GOOS=linux GOARCH=amd64 go build -o server-linux ./cmd/api

# Compile for Windows
GOOS=windows GOARCH=amd64 go build -o server.exe ./cmd/api

# Compile for ARM (ex: Raspberry Pi)
GOOS=linux GOARCH=arm64 go build -o server-arm ./cmd/api
```

### Go Today is Written in Go

Since Go 1.5, the main compiler is self-hosted (written in Go).
There are still low-level parts in assembly.

### Built-in HTTP Server

Go already brings a built-in HTTP server in the stdlib via `net/http`.

```go
http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong"))
})
http.ListenAndServe(":8080", nil)
```

This doesn't replace all the roles of an nginx/reverse proxy, but it greatly accelerates API development.

### Official Language Keywords (25)

| 1 | 2 | 3 | 4 | 5 |
|---|---|---|---|---|
| `break` | `default` | `func` | `interface` | `select` |
| `case` | `defer` | `go` | `map` | `struct` |
| `chan` | `else` | `goto` | `package` | `switch` |
| `const` | `fallthrough` | `if` | `range` | `type` |
| `continue` | `for` | `import` | `return` | `var` |

---

## 3. `net/http` Fundamentals

### The `net/http` Package

The package offers:
- HTTP Client
- HTTP Server
- `Request` and `ResponseWriter`
- `Handler`, `HandlerFunc`, and `ServeMux`
- Cookie utilities, headers, and more

Components:
- `http.ListenAndServe`
- `http.Request`
- `http.ResponseWriter`
- `http.HandleFunc`
- `http.HandlerFunc`
- `http.Handle`
- `http.Handler`
- `http.ServeMux`
- `http.Server`

### Mini Reference of Components

**`http.ListenAndServe`**

```go
log.Fatal(http.ListenAndServe(":8080", nil))
```

**`http.Request` and `http.ResponseWriter`:**

```go
func echo(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("method=" + r.Method + " path=" + r.URL.Path + " id=" + id))
}
```

**`http.HandleFunc`:**
```go
http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("pong"))
})

log.Fatal(http.ListenAndServe(":8080", nil))
```

**`http.HandlerFunc`**

```go
handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("pong"))
})

log.Fatal(http.ListenAndServe(":8080", handler))
```

**`http.Handle`**

```go
http.Handle("/ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("pong"))
}))

log.Fatal(http.ListenAndServe(":8080", nil))
```

**`http.Handler`**

```go
type PingHandler struct{}

func (PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("pong"))
}

log.Fatal(http.ListenAndServe(":8080", PingHandler{}))
```

**`http.ServeMux`**

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("pong"))
})

log.Fatal(http.ListenAndServe(":8080", mux))
```

**`http.Server`**

```go
srv := &http.Server{
	Addr:    ":8080",
	Handler: mux,
}

log.Fatal(srv.ListenAndServe())
```

Quick mental rule:
- `HandleFunc`: function
- `HandlerFunc`: function adapted to `Handler`
- `Handle`: registers a `Handler`
- `Handler`: complete behavior (`ServeHTTP`)

### Minimal Anatomy of a Handler (`w` and `r`)

![Handler Anatomy](docs/diagrams/handler-anatomy.svg)


**Standard Signature**

```go
func(w http.ResponseWriter, r *http.Request)
```

**w `http.ResponseWriter`:**
- it is the output of your API (response to the client)
- think in the order: **Headers -> Status -> Body**

**Main methods of `ResponseWriter`**

| Method | What it does | Important notes |
|---|---|---|
| `Header() http.Header` | Manipulates response headers | Define before `WriteHeader` |
| `Write([]byte)` | Writes the body | If `WriteHeader` is not called, it sends `200` automatically |
| `WriteHeader(statusCode int)` | Defines HTTP status | Should be called once |

Short example:

```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
_, _ = w.Write([]byte(`{"ok":true}`))
```

**Important rules**
- after `WriteHeader`, headers are frozen
- `Write()` implicitly calls `WriteHeader(200)` if no status was sent
- correct order:
1. `Header().Set(...)`
2. `WriteHeader(...)`
3. `Write(...)`

**r `*http.Request`:**
- represents everything the client sent in the request

**Most used fields of `Request`**

| Field | Type | Purpose |
|---|---|---|
| `r.Method` | `string` | HTTP Verb (`GET`, `POST`, etc.) |
| `r.URL` | `*url.URL` | Path and query string (`r.URL.Path`, `r.URL.Query().Get("id")`) |
| `r.Header` | `http.Header` | Request headers |
| `r.Body` | `io.ReadCloser` | Request body |
| `r.Host` | `string` | Called host |
| `r.RemoteAddr` | `string` | Client source IP/port |
| `r.Proto` | `string` | Protocol (`HTTP/1.1`, `HTTP/2.0`) |
| `r.ContentLength` | `int64` | Body size |

**URL Anatomy (each part)**

Example:

```text
https://domain.com/api/v1/user?id=123&debug=true#section
```

| URL Part | Example | Where to use in Go server |
|---|---|---|
| Protocol (scheme) | `https` | infer via `r.TLS` (`nil` = http, non-nil = https) |
| Host | `domain.com` | `r.Host` |
| Path | `/api/v1/user` | `r.URL.Path` |
| Raw query string | `id=123&debug=true` | `r.URL.RawQuery` |
| Query params | `id=123`, `debug=true` | `r.URL.Query().Get("id")`, `r.URL.Query().Get("debug")` |
| Fragment | `#section` | does not reach the server (browser does not send in HTTP request) |

Practical example in the handler:

```go
scheme := "http"
if r.TLS != nil {
	scheme = "https"
}

fullURL := scheme + "://" + r.Host + r.URL.RequestURI()
// fullURL => https://domain.com/api/v1/user?id=123&debug=true
```

**Useful fields and methods of `r.URL` (`*url.URL`)**

| Expression | Type | Purpose |
|---|---|---|
| `r.URL.Path` | `string` | Route path without query (`/api/v1/user`) |
| `r.URL.RawQuery` | `string` | Raw query string (`id=10&debug=true`) |
| `r.URL.Query()` | `url.Values` | Map of query parameters |
| `r.URL.Query().Get("id")` | `string` | Gets the first value for the key |
| `r.URL.Query()["tag"]` | `[]string` | Gets all values for the repeated key |
| `r.URL.EscapedPath()` | `string` | URL escaped path |
| `r.URL.String()` | `string` | URL in text format (good for log/debug) |

Working with headers:

```go
r.Header.Get("Authorization")
r.Header.Get("Content-Type")
```

Working with JSON body:

```go
defer r.Body.Close()
_ = json.NewDecoder(r.Body).Decode(&payload)
```

Best practice: limit body size:

```go
r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1MB
```

**`r.PathValue` (Go 1.22+): extracting route segments:**

When the route pattern contains `{name}`, the value is extracted with `r.PathValue("name")`:

```go
mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") // extracts "42" from /users/42
	_, _ = fmt.Fprintf(w, `{"id":"%s"}`, id)
})
```

| Expression | Type | Purpose |
|---|---|---|
| `r.PathValue("id")` | `string` | Named route segment (`/users/{id}`) |
| `r.URL.Query().Get("q")` | `string` | Query parameter (`?q=value`) |
| `r.URL.Query()["tag"]` | `[]string` | Multiple values for the same key (`?tag=a&tag=b`) |

---

## 4. Practical Manual: ListenAndServe (Phase Zero)

In this phase, the `README.md` is the main manual for copying, pasting, and running.

Practical note:
- each example uses port `:8080`

![ListenAndServe Flow](docs/diagrams/listenandserve.svg)

- run one example at a time (stop the previous one before running the next)

### Class reasoning line

| Order | Focus | Result for the student |
|---|---|---|
| 1 | `HandleFunc` vs `HandlerFunc` | Avoids the most common registration errors |
| 2 | What `ListenAndServe` accepts | Knows how to pass `nil`, `HandlerFunc`, `ServeMux`, or custom type |
| 3 | Base variations without custom `ServeMux` | Masters the basic HTTP flow |
| 4 | Reading the request: method, path, query, headers, body, `PathValue` | Extracts any data from the request |
| 5 | Writing the response: headers, status, body | Knows the correct order and avoids bugs |
| 6 | Full CRUD: GET, POST, PUT, PATCH, DELETE | Covers verbs with real examples |
| 7 | `http.Server` with timeouts | Configures server for production |
| 8 | Middleware: Logger, Auth, chain | Composes reusable behavior |
| 9 | Response standardization: motivation for `writeJSON` | Understands why section 5 exists |
| 10 | Graceful shutdown | Shuts down the server without losing in-progress requests |
| 11 | Testing handlers with `httptest` | Tests handlers without starting a real server |

### 4.1 Essential Difference: `HandleFunc` vs `HandlerFunc`

`HandleFunc` is a registration function.
`HandlerFunc` is an adapter type (becomes `http.Handler`).

```go
// WRONG - HandleFunc does not return anything
http.Handle("/route", http.HandleFunc(...))

// RIGHT - HandlerFunc is a type
http.Handle("/route", http.HandlerFunc(...))

// RIGHT - HandleFunc registers directly
http.HandleFunc("/route", ...)
```

### 4.2 What `ListenAndServe` Accepts

Signature:

```go
func ListenAndServe(addr string, handler Handler) error
```

The second argument accepts anything that implements the `http.Handler` interface, i.e., any type that has the `ServeHTTP(w, r)` method.

![What ListenAndServe accepts](docs/diagrams/listenandserve-handler.svg)

| Option | When to use |
|---|---|
| `nil` | Uses the global `DefaultServeMux`. Simple for examples, but avoid in production |
| `http.HandlerFunc(fn)` | Adapts a function directly as a handler. Useful for single-route server |
| `http.NewServeMux()` | Dedicated and isolated router. Recommended for any real API |
| Custom type with `ServeHTTP` | When you need internal state, composition, or middleware chain |

**About `DefaultServeMux` and why to avoid it in production:**

`DefaultServeMux` is a global `*ServeMux` automatically created by the `net/http` package. By passing `nil`, the server uses this mux implicitly. The problem: any imported package can register routes on it via `init()`, creating invisible and potentially unintended exposed routes.

```go
// CAUTION: any import might have done this in init()
import _ "some/package" // might have registered /debug/pprof, /metrics, etc.

http.ListenAndServe(":8080", nil) // exposes these routes without your knowledge
```

Always prefer your own `ServeMux`:

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /ping", pingHandler)
http.ListenAndServe(":8080", mux) // only your routes
```

**Custom type with `ServeHTTP` in practice:**

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

type App struct {
	version string
}

func (a App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/ping":
		_, _ = w.Write([]byte("pong"))
	case "/version":
		_, _ = fmt.Fprintf(w, `{"version":"%s"}`, a.version)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	app := App{version: "1.0.0"}
	log.Fatal(http.ListenAndServe(":8080", app))
}
```

**`http.Server` as the recommended alternative for production:**

`http.ListenAndServe` is convenient but doesn't allow configuring timeouts. For production always use `http.Server`:

```go
srv := &http.Server{
	Addr:              ":8080",
	Handler:           mux,           // your dedicated ServeMux
	ReadHeaderTimeout: 5 * time.Second,
	ReadTimeout:       15 * time.Second,
	WriteTimeout:      15 * time.Second,
	IdleTimeout:       60 * time.Second,
}
log.Fatal(srv.ListenAndServe())
```

### 4.3 Base Variations (without custom `ServeMux`)

#### Example 4.3.1 - `DefaultServeMux` with `HandleFunc`

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Home"))
	})

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("API"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/
curl -i localhost:8080/api
```

#### Example 4.3.2 - `DefaultServeMux` with `Handle`

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Home"))
	}))

	http.Handle("/api", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("API"))
	}))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/
curl -i localhost:8080/api
```

#### Example 4.3.3 - Direct single handler (manual routing)

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			switch r.URL.Path {
			case "/":
				_, _ = w.Write([]byte("Home"))
			case "/api":
				_, _ = w.Write([]byte("API"))
			default:
				w.WriteHeader(http.StatusNotFound)
				_, _ = w.Write([]byte("Not Found"))
			}
		}),
	))
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/
curl -i localhost:8080/api
curl -i localhost:8080/x
```

#### Example 4.3.4 - Extract `HandlerFunc` to variable

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	myHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Extracted handler"))
	})

	log.Fatal(http.ListenAndServe(":8080", myHandler))
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/
```

#### Example 4.3.5 - Convert to `HandlerFunc`

```go
package main

import (
	"log"
	"net/http"
)

func myFunc(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Normal function converted to Handler"))
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(myFunc)))
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/
```

### 4.4 Some Possibilities

#### Example 4.4.1 - URL Parameters (`r.URL.Query`)

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		_, _ = fmt.Fprintf(w, "Hello, %s!", name)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Run:

```bash
go run main.go
curl -i "localhost:8080/hello?name=jeff"
curl -i localhost:8080/hello
```

#### Example 4.4.2 - JSON Response

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]string{
			"name": "John",
			"role": "Developer",
		})
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/api/user
```

#### Example 4.4.3 - Custom headers + status code

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Custom-Header", "DevopsBH")
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Response with custom headers"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/
```


#### Example 4.4.4 - Query params with multiple values

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		// ?tag=go&tag=api&tag=http -> ["go", "api", "http"]
		tags := r.URL.Query()["tag"]
		q := r.URL.Query().Get("q")
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w,
			`{"q":"%s","tags":["%s"]}`,
			q, strings.Join(tags, `","`))
	})
	 log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Run:

```bash
go run main.go
curl -i "localhost:8080/search?q=golang&tag=go&tag=api&tag=http"
```

#### Example 4.4.5 - Reading custom headers from request

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		traceID := r.Header.Get("X-Trace-Id")
		clientVersion := r.Header.Get("X-Client-Version")
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w,
			`{"auth_present":%v,"trace_id":"%s","client_version":"%s"}`,
			auth != "", traceID, clientVersion)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/info \
	-H "Authorization: Bearer token123" \
	-H "X-Trace-Id: abc-456" \
	-H "X-Client-Version: 2.1.0"
```

#### Example 4.4.6 - `r.PathValue` (Go 1.22+)

`r.PathValue` extracts named segments directly from the route pattern.

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{"id":"%s"}`, id)
	})
	mux.HandleFunc("GET /posts/{postID}/comments/{commentID}", func(w http.ResponseWriter, r *http.Request) {
		postID := r.PathValue("postID")
		commentID := r.PathValue("commentID")
		w.Header().Set("Content-Type", "application/json")
		_, _ = fmt.Fprintf(w, `{"post_id":"%s","comment_id":"%s"}`, postID, commentID)
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/users/42
curl -i localhost:8080/posts/10/comments/99
```

#### Example 4.4.7 - PATCH `/api/v1/users/{id}` (partial update)

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type PatchUserInput struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}

func patchUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnsupportedMediaType)
		_, _ = w.Write([]byte(`{"error":"content_type_must_be_application_json"}`))
		return
	}
	var in PatchUserInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"invalid_json"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf(
		`{"message":"user partially updated","id":"%s","name":"%s","email":"%s"}`,
		id, in.Name, in.Email,
	)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("PATCH /api/v1/users/{id}", patchUser)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Run:

```bash
go run main.go
curl -i -X PATCH localhost:8080/api/v1/users/42 \
	-H "Content-Type: application/json" \
	-d '{"email":"new@email.com"}'
```

#### Example 4.4.8 - DELETE `/api/v1/users/{id}`

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"missing_id"}`))
		return
	}
	// simulation: id "0" does not exist
	if id == "0" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"error":"user_not_found","id":"%s"}`, id)))
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204: no body
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /api/v1/users/{id}", deleteUser)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Run:

```bash
go run main.go
curl -i -X DELETE localhost:8080/api/v1/users/42
curl -i -X DELETE localhost:8080/api/v1/users/0
```

#### Example 4.4.9 - `http.Redirect`

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// Permanent redirect (301): URL moved definitely
	mux.HandleFunc("GET /old-path", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/new-path", http.StatusMovedPermanently)
	})
	// Temporary redirect (302): URL might change back
	mux.HandleFunc("GET /temp", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/new-path", http.StatusFound)
	})
	mux.HandleFunc("GET /new-path", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("you have arrived at the new destination"))
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/old-path
curl -iL localhost:8080/old-path
curl -i localhost:8080/temp
```

### 4.5 `ServeMux` with method pattern + `http.Server`

Yes, this syntax exists and is official:

```go
mux.HandleFunc("POST /api/v1/user", handler)
```

Note:
- method pattern (`"GET /x"`, `"POST /x"`) requires Go 1.22+

**Keep-Alive** (important for API):
- in **HTTP/1.1**, **keep-alive** is default; the client tends to reuse the connection
- the server doesn't "enable keep-alive manually", but controls idle time
- in Go, `IdleTimeout` is a key configuration for persistent connections
- proxies/load balancers on the way might also close connections

#### Main properties of `http.Server`

| Field | What it does | Example |
|---|---|---|
| `Addr` | Address/port the server will listen on | `":8080"` |
| `Handler` | Who will process the routes (`mux`, custom handler, etc.) | `mux` |
| `IdleTimeout` | Idle connection time waiting for next request (keep-alive) | `60 * time.Second` |
| `ReadTimeout` | Maximum time to read the complete request (headers + body) | `15 * time.Second` |
| `ReadHeaderTimeout` | Maximum time to read only headers (anti-slowloris) | `5 * time.Second` |
| `WriteTimeout` | Maximum time to write the response | `15 * time.Second` |
| `MaxHeaderBytes` | Maximum size of received headers | `1 << 20` (1MB) |

Complete and educational example:

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Request %s processed\n", r.URL.Path)
	fmt.Printf("Connection from: %s\n", r.RemoteAddr)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		IdleTimeout:       60 * time.Second,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      15 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1MB
	}

	fmt.Println("Server listening on :8080")
	fmt.Println("IdleTimeout:", srv.IdleTimeout)
	log.Fatal(srv.ListenAndServe())
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/
```

Quick connection reuse test (HTTP/1.1):

```bash
curl -v --http1.1 http://localhost:8080/ http://localhost:8080/
```

Critical points to explain in class:
- `ReadTimeout` covers complete reading (headers + body)
- `ReadHeaderTimeout` protects against slow header sending (slowloris)
- `IdleTimeout` controls for how long keep-alive connection stays open without new request
- `MaxHeaderBytes: 1 << 20` uses bit shift to define 1MB limit
- keep-alive depends on the client/proxy reusing the connection; the server defines limits and policies

*Note: **Slowloris** is a type of DoS (Denial of Service) attack that keeps several HTTP connections open by sending data very slowly, without finishing the request.*

#### POST `/api/v1/user` with `json.NewDecoder`

```go
package main

import (
	"errors"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

const maxBodyBytes = 1 << 20 // 1MB -> Bit Shift Operator
// bit shift left) which means -> "Shift the number 1 to the left 20 positions"
// Math behind it:
// 1 << n  =  1 × 2^n  =  2^n
// Examples:
// 1 << 0  = 2^0  = 1
// 1 << 1  = 2^1  = 2
// 1 << 2  = 2^2  = 4
// 1 << 3  = 2^3  = 8
// 1 << 4  = 2^4  = 16
// 1 << 5  = 2^5  = 32
// 1 << 10 = 2^10 = 1024      (1 KB)
// 1 << 20 = 2^20 = 1048576   (1 MB)
// 1 << 30 = 2^30 = 1073741824 (1 GB)

func postUserWithDecoder(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnsupportedMediaType)
		_, _ = w.Write([]byte(`{"error":"content_type_must_be_application_json"}`))
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	var in User
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		w.Header().Set("Content-Type", "application/json")
		var maxErr *http.MaxBytesError
		if errors.As(err, &maxErr) {
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			_, _ = w.Write([]byte(`{"error":"body_too_large","max_bytes":1048576}`))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"invalid_json"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(fmt.Sprintf(
		`{"message":"user created (decoder)","name":"%s","email":"%s"}`,
		in.Name, in.Email,
	)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/user", postUserWithDecoder)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Run:

```bash
go run main.go
curl -i -X POST localhost:8080/api/v1/user \
  -H "Content-Type: application/json" \
  -d '{"name":"Jeff","email":"jeff@email.com"}'
```

#### POST `/api/v1/user` reading `Body` + `json.Unmarshal`

```go
package main

import (
	"errors"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

const maxBodyBytes = 1 << 20 // 1MB

func postUserWithUnmarshal(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnsupportedMediaType)
		_, _ = w.Write([]byte(`{"error":"content_type_must_be_application_json"}`))
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)

	raw, err := io.ReadAll(r.Body)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		var maxErr *http.MaxBytesError
		if errors.As(err, &maxErr) {
			w.WriteHeader(http.StatusRequestEntityTooLarge)
			_, _ = w.Write([]byte(`{"error":"body_too_large","max_bytes":1048576}`))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"cannot_read_body"}`))
		return
	}

	var in User
	if err := json.Unmarshal(raw, &in); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"invalid_json"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(fmt.Sprintf(
		`{"message":"user created (unmarshal)","name":"%s","email":"%s"}`,
		in.Name, in.Email,
	)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/user", postUserWithUnmarshal)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Run:

```bash
go run main.go
curl -i -X POST localhost:8080/api/v1/user \
  -H "Content-Type: application/json" \
  -d '{"name":"Jeff","email":"jeff@email.com"}'
```

#### When to use `Decoder` vs `Unmarshal`?

| Option | When to use | Advantage | Attention |
|---|---|---|---|
| `json.NewDecoder(r.Body).Decode(&v)` | Standard HTTP flow reading directly from body | Simple and direct in the handler | Less control over the raw `[]byte` |
| `io.ReadAll(r.Body)` + `json.Unmarshal(raw, &v)` | When you need the raw body before converting | Allows log, audit, signature, prior validation | More verbose and uses memory to store the entire body |

Minimum checklist for API endpoint:
- limit body size (`http.MaxBytesReader`)
- validate `Content-Type: application/json` for JSON endpoints
- validate input JSON and return clear error
- respond with consistent `Content-Type`
- keep handlers named outside of `main` as the flow grows

#### GET `/api/v1/user`

```go
package main

import (
	"log"
	"net/http"
)

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`[{"name":"Jeff","email":"jeff@email.com"}]`))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/user", getUsers)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/api/v1/user
```

#### PUT `/api/v1/user`

```go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func putUser(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnsupportedMediaType)
		_, _ = w.Write([]byte(`{"error":"content_type_must_be_application_json"}`))
		return
	}

	var in User
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"invalid_json"}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(fmt.Sprintf(
		`{"message":"user updated","name":"%s","email":"%s"}`,
		in.Name, in.Email,
	)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("PUT /api/v1/user", putUser)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Run:

```bash
go run main.go
curl -i -X PUT localhost:8080/api/v1/user \
  -H "Content-Type: application/json" \
  -d '{"name":"Jeff Otoni","email":"new@email.com"}'
```

### 4.6 When to use `http.Handler`?

![Middleware Chain](docs/diagrams/middleware-chain.svg)


Golden rule:
- use `http.Handler` when you want to **compose behavior** (middleware, chain, reuse)
- use `http.HandlerFunc` when you want to **respond to a route directly** (simple and fast)

`http.Handler` is the base of everything:

```go
type Handler interface {
    ServeHTTP(http.ResponseWriter, *http.Request)
}
```

It's not a function.
It's behavior: any type that implements `ServeHTTP` becomes an HTTP handler.

#### When `HandlerFunc` is sufficient

Use direct function when:
- simple code
- no internal state
- no middleware composition

```go
package main

import (
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("pong"))
}

func main() {
	http.HandleFunc("/ping", ping)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/ping
```

#### Example with `ServeHTTP` (customized type)

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

type helloHandler struct {
	msg string
}

func (h helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte(fmt.Sprintf("%s | path=%s", h.msg, r.URL.Path)))
}

func main() {
	h := helloHandler{msg: "custom handler implementing ServeHTTP"}
	http.Handle("/hello", h)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/hello
```

#### Example with `http.Handler` for composition (struct + `finalHandler`)

```go
package main

import (
	"log"
	"net/http"
)

type LoggerMiddleware struct {
	next http.Handler
}

func (l LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)
	l.next.ServeHTTP(w, r)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"name":"Jeff","email":"jeff@email.com"}`))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/user", getUser)

	finalHandler := LoggerMiddleware{next: mux}
	srv := &http.Server{
		Addr:    ":8080",
		Handler: finalHandler,
	}

	log.Fatal(srv.ListenAndServe())
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/api/v1/user
```

#### Same composition, another way (middleware func + `finalHandler`)

```go
package main

import (
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"name":"Jeff","email":"jeff@email.com"}`))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/user", getUser)

	finalHandler := loggingMiddleware(mux)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: finalHandler,
	}

	log.Fatal(srv.ListenAndServe())
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/api/v1/user
```


### 4.7 Graceful Shutdown

When calling `log.Fatal(srv.ListenAndServe())`, the server exits immediately upon receiving an OS signal (SIGINT, SIGTERM). This cuts open connections abruptly. `Shutdown` allows the server to finish in-progress requests before closing.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	// Channel that receives OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Starts server in separate goroutine
	go func() {
		fmt.Println("Server listening on :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	// Blocks until signal received
	<-quit
	fmt.Println("Shutting down server...")

	// Context with timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server exited cleanly")
}
```

Run:

```bash
go run main.go
# In another terminal:
curl -i localhost:8080/ping
# To quit with graceful shutdown:
Ctrl+C
```

Important points:
- `signal.Notify` captures `SIGINT` (Ctrl+C) and `SIGTERM` (used by Docker/Kubernetes)
- `srv.ListenAndServe()` returns `http.ErrServerClosed` when `Shutdown` is called; this is expected, not an error
- `context.WithTimeout(10s)` ensures the server doesn't wait indefinitely
- In Kubernetes, `SIGTERM` is sent before removing the pod from the load balancer; graceful shutdown gives time for in-progress requests to finish

### 4.8 Testing handlers with `httptest`

The `net/http/httptest` package allows testing handlers without starting a real server on a port. It is the Go community standard for API testing.

```go
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Handler we want to test
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		_, _ = w.Write([]byte(`{"error":"unsupported_media_type"}`))
		return
	}
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error":"invalid_json"}`))
		return
	}
	if input.Name == "" || input.Email == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte(`{"error":"validation_error"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(`{"message":"user created"}`))
}

func TestCreateUserHandler(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		contentType string
		wantStatus int
	}{
		{
			name:        "success",
			body:        `{"name":"Jeff","email":"jeff@email.com"}`,
			contentType: "application/json",
			wantStatus:  http.StatusCreated,
		},
		{
			name:        "wrong content-type",
			body:        `{"name":"Jeff","email":"jeff@email.com"}`,
			contentType: "text/plain",
			wantStatus:  http.StatusUnsupportedMediaType,
		},
		{
			name:        "missing fields",
			body:        `{"name":"Jeff"}`,
			contentType: "application/json",
			wantStatus:  http.StatusUnprocessableEntity,
		},
		{
			name:        "invalid json",
			body:        `{invalid}`,
			contentType: "application/json",
			wantStatus:  http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/v1/users",
				strings.NewReader(tt.body))
			req.Header.Set("Content-Type", tt.contentType)

			rec := httptest.NewRecorder()
			createUserHandler(rec, req)

			if rec.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantStatus)
			}
		})
	}
}
```

Run:

```bash
go test -v ./...
```

`httptest` components:
- `httptest.NewRequest(method, target, body)`: creates an `*http.Request` for testing without real connection
- `httptest.NewRecorder()`: implements `http.ResponseWriter` and captures status, headers, and body
- `rec.Code`: status code written by the handler
- `rec.Body.String()`: response body as string
- `rec.Header()`: response headers

### 4.9 Two Servers on Distinct Ports with Goroutines

Use one `ServeMux` per port and start each server in its own goroutine.
For all stop/blocking variations, see:
- [`examples/10-dual-listenandserve-goroutines/README.md`](examples/10-dual-listenandserve-goroutines/README.md)
- [`examples/10-dual-listenandserve-goroutines/README_pt.md`](examples/10-dual-listenandserve-goroutines/README_pt.md)

Endpoints:
- `:8080` -> `GET /api/v1/user`, `POST /api/v1/user`
- `:3000` -> `GET /api/v1/mock/user`, `POST /api/v1/mock/user`

Practical example in repository:
- [`examples/10-dual-listenandserve-goroutines/main.go`](examples/10-dual-listenandserve-goroutines/main.go)
- [`examples/10-dual-listenandserve-goroutines/main_test.go`](examples/10-dual-listenandserve-goroutines/main_test.go)
- [`examples/10-dual-listenandserve-goroutines/README.md`](examples/10-dual-listenandserve-goroutines/README.md) (all scenario links)
- [`examples/10-dual-listenandserve-goroutines/README_pt.md`](examples/10-dual-listenandserve-goroutines/README_pt.md) (versão em português)
- [`examples/10-1-dual-listenandserve-basic/main.go`](examples/10-1-dual-listenandserve-basic/main.go) (10.1 compact version)
- [`examples/10-1-dual-listenandserve-basic/main_test.go`](examples/10-1-dual-listenandserve-basic/main_test.go)

```go
package main

import (
	"log"
	"net/http"
)

// newMainMux registers routes for port 8080.
func newMainMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/user", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"user read from :8080"}`))
	})
	mux.HandleFunc("POST /api/v1/user", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"message":"user created on :8080"}`))
	})
	return mux
}

// newMockMux registers routes for port 3000.
func newMockMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/mock/user", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"message":"mock user read from :3000"}`))
	})
	mux.HandleFunc("POST /api/v1/mock/user", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"message":"mock user created on :3000"}`))
	})
	return mux
}

// startServer starts one HTTP server in a goroutine.
func startServer(srv *http.Server) {
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server %s error: %v", srv.Addr, err)
		}
	}()
}

// main starts both servers and blocks the process.
func main() {
	mainServer := &http.Server{Addr: ":8080", Handler: newMainMux()}
	mockServer := &http.Server{Addr: ":3000", Handler: newMockMux()}

	startServer(mainServer)
	startServer(mockServer)
	select {} // 01) Simple, direct, less controllable.
	// done := make(chan struct{}); <-done // 02) Explicit blocking channel.
	// var wg sync.WaitGroup; wg.Add(2); ...; wg.Wait() // 03) WaitGroup.
	// stop := make(chan os.Signal, 1); signal.Notify(stop, os.Interrupt, syscall.SIGTERM); <-stop // 04) Signal channel.
	// time.Sleep(10 * time.Minute) // 05) Fixed sleep window.
	// ctx, cancel := context.WithCancel(context.Background()); <-ctx.Done(); cancel() // 06) Context cancel.
	// for { time.Sleep(time.Hour) } // 07) Infinite loop.
	// errCh := make(chan error, 2); if err := <-errCh; err != nil { ... } // 08) Error channel group.
	// ctx, cancel := context.WithCancel(context.Background()); select { case err := <-errCh: cancel() } // 09) Error channel + context.
	// var mu sync.Mutex; mu.Lock(); mu.Lock() // 10) Mutex deadlock (didactic).
	// runtime.Goexit() // 11) Ends main goroutine only.
	// var mu sync.Mutex; cond := sync.NewCond(&mu); mu.Lock(); cond.Wait() // 12) sync.Cond wait.
	// ch := make(chan struct{}); for range ch {} // 13) Channel range block.
	// fmt.Scanln() // 14) Block on stdin.
	// go runMock(); runMain() // 15) One blocking server + one goroutine.
	// ticker := time.NewTicker(time.Hour); defer ticker.Stop(); for range ticker.C {} // 16) Infinite ticker.
	// stop := make(chan struct{}); select { case <-stop: case err := <-errCh: _ = err } // 17) Select with multiple channels.
}
```

Run:

```bash
go run ./examples/10-dual-listenandserve-goroutines
curl -i localhost:8080/api/v1/user
curl -i -X POST localhost:8080/api/v1/user
curl -i localhost:3000/api/v1/mock/user
curl -i -X POST localhost:3000/api/v1/mock/user
```


---

## 5. Server API

In the section 4 examples, each handler wrote its response directly with `w.Header().Set(...)`, `w.WriteHeader(...)`, and `w.Write(...)`. It works, but as the API grows this becomes repetitive and inconsistent: one handler returns `{"error":"..."}`, another returns `{"message":"..."}`, a third forgets the `Content-Type`.

Section 5 solves this with a centralized response pattern (`writeJSON` and `writeError`), route organization, input validation, health endpoints, and basic authentication. It is the natural evolution of what was built in section 4.

**Focus of this block:**
- server side only
- small examples to copy, paste, and evolve
- standardize API before moving to a larger structure

**Selected items:**

| Item |
|---|
| 5.0 helpers.go: shared functions for the section |
| 5.1 Response Standardization |
| 5.2 Error and status map by scenario |
| 5.3 Route Organization |
| 5.4 Server input validation |
| 5.5 Health Endpoints |
| 5.6 Basic Auth Middleware |
| 5.7 Environment Variables |
| 5.8 CORS Middleware |
| 5.9 Documentation with OpenAPI/Swagger |

### 5.0 helpers.go: shared functions

The following examples reuse `writeJSON` and `writeError`. Instead of redeclaring in each file, extract to a `helpers.go` in the same package:

```go
// helpers.go
package main

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	b, err := json.Marshal(payload)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":{"code":"internal_error","message":"json_encode_failed"}}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(b)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, map[string]any{
		"error": APIError{Code: code, Message: message},
	})
}
```

From here, all examples assume that `writeJSON` and `writeError` are available via `helpers.go` in the same package.

### 5.1 Response Standardization

**Objective:**
- centralize response writing in a single point
- avoid repetition of `Header`, `WriteHeader`, `Write`
- maintain consistent success and error format

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	b, err := json.Marshal(payload)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":{"code":"internal_error","message":"json_encode_failed"}}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(b)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, map[string]any{
		"error": APIError{
			Code:    code,
			Message: message,
		},
	})
}

func okHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{
		"message": "ok",
		"data":    map[string]string{"course": "net/http"},
	})
}

func badHandler(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusBadRequest, "invalid_input", "missing required field")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ok", okHandler)
	mux.HandleFunc("GET /bad", badHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/ok
curl -i localhost:8080/bad
```

### 5.2 Error and status map by scenario

| Scenario | Status | When to use |
|---|---:|---|
| Invalid JSON | `400` | Malformed body |
| Invalid / missing field | `422` | Valid body, but invalid business rule |
| Incorrect Content-Type | `415` | Expected `application/json` |
| Resource not found | `404` | ID/path does not exist |
| Method not allowed | `405` | Endpoint exists, method does not |
| State conflict | `409` | Ex.: email already registered |
| Internal error | `500` | Unexpected server failure |

Example of standardized error return:

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	b, err := json.Marshal(payload)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":{"code":"internal_error","message":"json_encode_failed"}}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(b)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, map[string]any{
		"error": APIError{
			Code:    code,
			Message: message,
		},
	})
}

func errorByScenario(w http.ResponseWriter, r *http.Request) {
	switch r.PathValue("scenario") {
	case "bad-json":
		writeError(w, http.StatusBadRequest, "invalid_json", "malformed request body")
	case "validation":
		writeError(w, http.StatusUnprocessableEntity, "validation_error", "name and email are required")
	case "content-type":
		writeError(w, http.StatusUnsupportedMediaType, "unsupported_media_type", "use application/json")
	case "not-found":
		writeError(w, http.StatusNotFound, "not_found", "resource not found")
	case "method":
		writeError(w, http.StatusMethodNotAllowed, "method_not_allowed", "method not allowed for this endpoint")
	case "conflict":
		writeError(w, http.StatusConflict, "conflict", "email already exists")
	default:
		writeError(w, http.StatusInternalServerError, "internal_error", "unexpected server error")
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /error/{scenario}", errorByScenario)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/error/bad-json
curl -i localhost:8080/error/validation
curl -i localhost:8080/error/not-found
```

### 5.3 Route Organization

Objective:
- keep routes predictable
- separate by context (`health`, `users`, etc.)
- version API (`/api/v1`)

```go
package main

import (
	"log"
	"net/http"
)

func registerRoutes(mux *http.ServeMux) {
	// Health
	mux.HandleFunc("GET /healthz", healthz)
	mux.HandleFunc("GET /readyz", readyz)
	mux.HandleFunc("GET /livez", livez)

	// API v1 - users
	mux.HandleFunc("POST /api/v1/users", createUser)
	mux.HandleFunc("GET /api/v1/users/{id}", getUserByID)
	mux.HandleFunc("PUT /api/v1/users/{id}", updateUser)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("ok"))
}

func readyz(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("ready"))
}

func livez(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("alive"))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(`{"message":"user created"}`))
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	_, _ = w.Write([]byte(`{"id":"` + id + `"}`))
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	_, _ = w.Write([]byte(`{"message":"user updated","id":"` + id + `"}`))
}

func main() {
	mux := http.NewServeMux()
	registerRoutes(mux)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Tip:
- if using `{id}` in the pattern (Go 1.22+), read with `r.PathValue("id")`

Run:

```bash
go run main.go
curl -i localhost:8080/healthz
curl -i localhost:8080/api/v1/users/123
```

### 5.4 Server Input Validation

Short checklist for `POST/PUT`:
1. validate `Content-Type`
2. limit body size
3. decode JSON with `DisallowUnknownFields`
4. validate mandatory fields
5. return correct status (`400`, `415`, `422`)

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type CreateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	b, err := json.Marshal(payload)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":{"code":"internal_error","message":"json_encode_failed"}}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(b)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, map[string]any{
		"error": APIError{
			Code:    code,
			Message: message,
		},
	})
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		writeError(w, http.StatusUnsupportedMediaType, "unsupported_media_type", "use application/json")
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1MB

	var in CreateUserInput
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&in); err != nil {
		writeError(w, http.StatusBadRequest, "invalid_json", "malformed request body")
		return
	}

	if strings.TrimSpace(in.Name) == "" || strings.TrimSpace(in.Email) == "" {
		writeError(w, http.StatusUnprocessableEntity, "validation_error", "name and email are required")
		return
	}

	writeJSON(w, http.StatusCreated, map[string]any{
		"message": "user created",
		"user":    in,
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/users", createUser)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Run:

```bash
go run main.go
curl -i -X POST localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Jeff","email":"jeff@email.com"}'
curl -i -X POST localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Jeff","email":"jeff@email.com","extra":"x"}'
```

### 5.5 Health Endpoints

![Health Endpoints](docs/diagrams/health-endpoints.svg)


Simple pattern:
- `GET /healthz`: server responded (up)
- `GET /livez`: process alive
- `GET /readyz`: ready to receive traffic (dependencies OK)

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var ready = true

func writeJSON(w http.ResponseWriter, status int, payload any) {
	b, err := json.Marshal(payload)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":{"code":"internal_error","message":"json_encode_failed"}}`))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(b)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, map[string]any{
		"error": APIError{
			Code:    code,
			Message: message,
		},
	})
}

func healthz(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func livez(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "alive"})
}

func readyz(w http.ResponseWriter, r *http.Request) {
	if !ready {
		writeError(w, http.StatusServiceUnavailable, "not_ready", "dependencies are not ready")
		return
	}

	writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", healthz)
	mux.HandleFunc("GET /readyz", readyz)
	mux.HandleFunc("GET /livez", livez)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
```

Run:

```bash
go run main.go
curl -i localhost:8080/healthz
curl -i localhost:8080/readyz
curl -i localhost:8080/livez
```

### 5.6 Basic Auth Middleware

When to use (simple and educational):
- protect laboratory/homologation internal endpoint
- test basic HTTP authentication before adopting JWT/OAuth2

```go
package main

import (
	"crypto/subtle"
	"log"
	"net/http"
	"os"
)

func unauthorized(w http.ResponseWriter) {
	w.Header().Set("WWW-Authenticate", `Basic realm="api", charset="UTF-8"`)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write([]byte(`{"error":"unauthorized"}`))
}

func basicAuthMiddleware(next http.Handler) http.Handler {
	// Reads credentials from environment variables
	expectedUser := os.Getenv("API_USER")
	expectedPass := os.Getenv("API_PASS")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok {
			unauthorized(w)
			return
		}
		userOK := subtle.ConstantTimeCompare([]byte(user), []byte(expectedUser)) == 1
		passOK := subtle.ConstantTimeCompare([]byte(pass), []byte(expectedPass)) == 1
		if !userOK || !passOK {
			unauthorized(w)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`{"name":"Jeff","email":"jeff@email.com"}`))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/user", getUser)
	finalHandler := basicAuthMiddleware(mux)
	log.Fatal(http.ListenAndServe(":8080", finalHandler))
}
```

Run:

```bash
API_USER=admin API_PASS=s3cr3t go run main.go
curl -i localhost:8080/api/v1/user
curl -i -u admin:s3cr3t localhost:8080/api/v1/user
```

### 5.7 Environment Variables

In production, sensitive settings (credentials, ports, database URLs) should never be hardcoded. Go offers `os.Getenv` and `os.LookupEnv` for this.

| Function | Behavior |
|---|---|
| `os.Getenv("KEY")` | Returns the value or `""` if it doesn't exist |
| `os.LookupEnv("KEY")` | Returns value + bool `ok`, distinguishes "not defined" from "defined empty" |

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Port with fallback
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Mandatory variable: stops if not defined
	dbURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Fatal("DATABASE_URL not defined")
	}

	// Optional variable with fallback
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	fmt.Printf("Starting at :%s (env=%s db=%s)\n", port, env, dbURL)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
```

Run:

```bash
PORT=9090 APP_ENV=production DATABASE_URL=postgres://localhost/mydb go run main.go
curl -i localhost:9090/ping
```

In Docker:

```dockerfile
ENV PORT=8080
ENV APP_ENV=production
```

Or via `docker run`:

```bash
docker run -e PORT=9090 -e DATABASE_URL=postgres://... -p 9090:9090 nethttp-server:local
```

Best practices:
- use `LookupEnv` for mandatory variables: the server fails immediately if misconfigured
- use `Getenv` with fallback for optional variables with a sensible default value
- never log the value of sensitive variables (passwords, tokens)
- in development, use a `.env` file with a library like `godotenv`; never commit this file


### 5.8 CORS Middleware

CORS (Cross-Origin Resource Sharing) is a browser security mechanism that blocks requests made from an origin different from the API. Every API consumed by a web frontend needs to handle this.

**How it works:**

1. The browser sends an `OPTIONS` request (preflight) before `POST/PUT/DELETE` with custom headers
2. The server responds with `Access-Control-*` headers indicating what is allowed
3. The browser releases (or blocks) the actual request based on this response

```go
package main

import (
	"log"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allows any origin (restrict in production)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Trace-Id")

		// Responds to preflight and exits
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"name":"Jeff","email":"jeff@email.com"}]`))
	})

	// CORS applied to the whole API
	finalHandler := corsMiddleware(mux)
	log.Fatal(http.ListenAndServe(":8080", finalHandler))
}
```

Run:

```bash
go run main.go

# Simulates browser preflight
curl -i -X OPTIONS localhost:8080/api/v1/users \
  -H "Origin: http://localhost:3000" \
  -H "Access-Control-Request-Method: POST" \
  -H "Access-Control-Request-Headers: Content-Type"

# Normal request
curl -i localhost:8080/api/v1/users -H "Origin: http://localhost:3000"
```

| Header | What it controls |
|---|---|
| `Access-Control-Allow-Origin` | Which origins can access (`*` = any, or `https://mysite.com`) |
| `Access-Control-Allow-Methods` | Which HTTP verbs are allowed |
| `Access-Control-Allow-Headers` | Which headers the client can send |
| `Access-Control-Max-Age` | How long the browser can cache the preflight (seconds) |

**In production:** replace `"*"` with your frontend's real origin. Using `"*"` with credentials (`Authorization`) does not work, as the browser requires an explicit origin in this case.

### 5.9 Documentation with OpenAPI/Swagger

#### A little bit of history

In 2010, Tony Tam created Swagger while working at Wordnik to solve a simple problem: how to describe a REST API so that humans and machines could understand. The solution was a specification file in JSON/YAML describing endpoints, parameters, schemas, and responses.

In 2015, the specification was donated to the OpenAPI Initiative (part of the Linux Foundation) and renamed to OpenAPI Specification. The name Swagger remained popular, but technically refers to the tools (Swagger UI, Swagger Editor), not the specification.

| Version | Year | Highlights |
|---|---:|---|
| Swagger 1.x | 2010 | Origin at Wordnik, proprietary format |
| Swagger 2.0 | 2014 | Broad standardization, massive adoption |
| OpenAPI 3.0 | 2017 | Restructured, support for webhooks and links |
| OpenAPI 3.1 | 2021 | Full alignment with JSON Schema |

Today OpenAPI 3.1 is the standard. When someone says "generate the Swagger", they usually mean "generate the OpenAPI file".

**Why it matters:**
- other teams can consume your API without asking for explanations
- SDK clients can be automatically generated
- serves as a living contract between frontend and backend
- facilitates manual testing via web interface

**With `swaggo/swag` in Go:**

```bash
# Installs the generator
go install github.com/swaggo/swag/cmd/swag@latest

# Adds dependencies to the project
go get github.com/swaggo/http-swagger
go get github.com/swaggo/swag
```

Annotate handlers with special comments:

```go
package main

import (
	"log"
	"net/http"

	_ "myproject/docs" // generated by swag
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           User API
// @version         1.0
// @description     Example of API documented with Swagger
// @host            localhost:8080
// @BasePath        /api/v1

// getUsers godoc
// @Summary      List users
// @Description  Returns all registered users
// @Tags         users
// @Produce      json
// @Success      200  {array}   User
// @Failure      500  {object}  APIError
// @Router       /users [get]
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(`[{"name":"Jeff","email":"jeff@email.com"}]`))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/v1/users", getUsers)

	// Serves Swagger UI interface at /swagger/
	mux.Handle("GET /swagger/", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Generate documentation:

```bash
# Generates docs/ folder with OpenAPI JSON/YAML
swag init

# Starts the server
go run main.go

# Accesses visual interface
open http://localhost:8080/swagger/index.html
```

**Recommended workflow:**

| Stage | Action |
|---|---|
| 1 | Writes handler with `// @...` annotations |
| 2 | Runs `swag init` to regenerate `docs/` |
| 3 | Commits `docs/` along with code |
| 4 | CI/CD can validate that `docs/` is updated |

**Alternatives to swaggo:**

| Tool | Approach | Pollutes code? |
|---|---|:---:|
| `swaggo/swag` | Annotations in Go comments, generates OpenAPI 2.0/3.0 | Yes |
| `deepmap/oapi-codegen` | Contract-first: generates Go code from OpenAPI YAML | No |
| `huma` | Framework that automatically generates OpenAPI via Go types | No |
| Manual YAML + static handler | You write the contract, Go serves the file | No |
| Postman → OpenAPI | Exports Postman collection and converts to YAML | No |

---

#### Alternative 1: Manual YAML + Swagger UI without dependency

The cleanest approach. You write the `openapi.yaml` by hand (or via online editor), commit to the repository and serve with a single Go handler. Zero code annotations, zero extra dependencies.

**Step 1: Write the contract in the online Swagger Editor:**

Access [editor.swagger.io](https://editor.swagger.io), write or paste your YAML and validate in real-time. When ready, download the file.

```yaml
# openapi.yaml: save at root or in docs/openapi.yaml
openapi: "3.0.3"
info:
  title: User API
  version: "1.0.0"
paths:
  /api/v1/users:
    get:
      summary: List users
      responses:
        "200":
          description: List of users
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/User"
  /api/v1/users/{id}:
    get:
      summary: Find user by ID
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
      responses:
        "200":
          description: User found
        "404":
          description: Not found
components:
  schemas:
    User:
      type: object
      properties:
        name:
          type: string
        email:
          type: string
```

**Step 2: Serve YAML and UI via Go (zero dependencies):**

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// API Routes
	mux.HandleFunc("GET /api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"name":"Jeff","email":"jeff@email.com"}]`))
	})

	// Serves openapi.yaml file directly
	mux.HandleFunc("GET /openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		http.ServeFile(w, r, "docs/openapi.yaml")
	})

	// Serves Swagger UI via CDN (pure HTML, no Go dependency)
	mux.HandleFunc("GET /docs", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, _ = w.Write([]byte(`<!DOCTYPE html>
<html>
<head>
  <title>API Docs</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist/swagger-ui.css"/>
</head>
<body>
  <div id="swagger-ui"></div>
  <script src="https://unpkg.com/swagger-ui-dist/swagger-ui-bundle.js"></script>
  <script>
    SwaggerUIBundle({ url: "/openapi.yaml", dom_id: "#swagger-ui" })
  </script>
</body>
</html>`))
	})

	log.Println("Docs at http://localhost:8080/docs")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Run:

```bash
go run main.go
open http://localhost:8080/docs
```

Or serve via Redoc (cleaner visual for public documentation):

```go
mux.HandleFunc("GET /redoc", func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	_, _ = w.Write([]byte(`<!DOCTYPE html>
<html>
<head><title>API Docs</title></head>
<body>
  <redoc spec-url="/openapi.yaml"></redoc>
  <script src="https://cdn.jsdelivr.net/npm/redoc/bundles/redoc.standalone.js"></script>
</body>
</html>`))
})
```

---

#### Alternative 2: Postman Collection to OpenAPI YAML

If you already have a Postman collection, you can convert it directly to OpenAPI without writing the YAML from scratch.

**Option A: Through Postman itself (interface):**

1. Open collection in Postman
2. Click on collection's three dots
3. Select `Export` → choose format `Collection v2.1`
4. Save as `collection.json`

**Option B: Convert with Node script (postman-to-openapi):**

```bash
# Installs tool
npm install -g postman-to-openapi

# Converts collection.json to openapi.yaml
p2o collection.json -f docs/openapi.yaml
```

**Option C: Convert with Python script:**

```python
#!/usr/bin/env python3
# convert_postman.py
# pip install pyyaml

import json
import yaml
import sys

def postman_to_openapi(collection_path: str, output_path: str):
    with open(collection_path) as f:
        collection = json.load(f)

    info = collection.get("info", {})
    paths = {}

    def extract_items(items, prefix=""):
        for item in items:
            if "item" in item:
                extract_items(item["item"], prefix)
            elif "request" in item:
                req = item["request"]
                url = req.get("url", {})
                raw = url.get("raw", "") if isinstance(url, dict) else url
                # normalizes path params {param} -> {param}
                path = "/" + "/".join(
                    seg.get("value", seg) if isinstance(seg, dict) else seg
                    for seg in (url.get("path", []) if isinstance(url, dict) else [])
                )
                method = req.get("method", "GET").lower()
                paths.setdefault(path, {})[method] = {
                    "summary": item.get("name", ""),
                    "responses": {"200": {"description": "OK"}},
                }

    extract_items(collection.get("item", []))

    spec = {
        "openapi": "3.0.3",
        "info": {"title": info.get("name", "API"), "version": "1.0.0"},
        "paths": paths,
    }

    with open(output_path, "w") as f:
        yaml.dump(spec, f, allow_unicode=True, sort_keys=False)

    print(f"OpenAPI generated in {output_path}")

if __name__ == "__main__":
    postman_to_openapi(sys.argv[1], sys.argv[2])
```

Run:

```bash
python3 convert_postman.py collection.json docs/openapi.yaml
```

Then serve `openapi.yaml` with the Alternative 1 handler. The result is complete visual documentation without any annotations in your Go code.

---

#### Alternative 3: Contract-first with oapi-codegen

Writes the YAML first, automatically generates Go interfaces and types. The code stays clean because the specification lives outside the code.

```bash
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
oapi-codegen -package api docs/openapi.yaml > api/api.gen.go
```

Go generates the interfaces you implement, without any annotation and without any special comment.

---

**Which approach to choose:**

| Situation | Recommendation |
|---|---|
| Team already uses Postman | Export collection and convert with `postman-to-openapi` or Python script |
| Wants zero Go dependency | Manual YAML + `http.ServeFile` + Swagger UI via CDN |
| Wants rigorous contract and code generation | `oapi-codegen` (contract-first) |
| Wants to iterate fast with annotations | `swaggo/swag` (accepts pollution as tradeoff) |
| Nice public documentation | Redoc via CDN over the same `openapi.yaml` |

---

## 6. Docker: Build and Run Local

Objective:
- compile Go server in lean image (multi-stage)
- run local on port `8080`
- have basic commands for operation and debug

Prerequisite:
- have a functional `main.go` at root (you can use any example from this README)
- use `.dockerignore` to not send unnecessary files in build context

### 6.1 Multi-stage Dockerfile (Alpine + Brazil timezone)

`Dockerfile` file (at project root):

```dockerfile
FROM golang:1.25.6-alpine AS builder
WORKDIR /src
RUN apk add --no-cache ca-certificates
COPY . .
ARG APP_FILE=main.go
ENV CGO_ENABLED=0
RUN go build -trimpath -ldflags="-s -w" -o /out/server "./${APP_FILE}"

FROM alpine:3.20 AS runtime

ENV TZ=America/Sao_Paulo
RUN apk add --no-cache ca-certificates tzdata \
  && ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
  && echo $TZ > /etc/timezone

WORKDIR /app
COPY --from=builder /out/server /app/server

EXPOSE 8080
USER nobody:nobody
ENTRYPOINT ["/app/server"]
```

Notes:
- final image is lean (no Go toolchain)
- timezone set for Brazil (`America/Sao_Paulo`)
- if your entrypoint is another file, use `--build-arg APP_FILE=your_file.go`
- builder image must be compatible with `go.mod` version (ex.: `go 1.25.6`)

### 6.2 Basic Docker Commands

Image build:

```bash
docker build -t nethttp-server:local .
```

Image build without using cache (forces rebuild):

```bash
docker build --no-cache -t nethttp-server:local .
```

Build choosing another Go file:

```bash
docker build -t nethttp-server:local --build-arg APP_FILE=cmd/api/main.go .
```

Clear build cache (builder cache):

```bash
docker builder prune -f
```

Run local container:

```bash
docker run -d --name nethttp-server -p 8080:8080 nethttp-server:local
```

List containers:

```bash
docker ps -a
```

View logs:

```bash
docker logs -f nethttp-server
```

Validate routes:

```bash
curl -i localhost:8080/api/v1/user
```

```bash
curl -i -X POST localhost:8080/api/v1/user \
  -H "Content-Type: application/json" \
  -d '{"name":"Jeff","email":"jeff@email.com"}'
```

```bash
curl -i -X PUT localhost:8080/api/v1/user \
  -H "Content-Type: application/json" \
  -d '{"name":"Jeff Updated","email":"jeff.updated@email.com"}'
```

Check timezone in container:

```bash
docker exec -it nethttp-server date
```

Stop and remove container:

```bash
docker stop nethttp-server
docker rm nethttp-server
```

---

<br>

<sub>Made with dedication by <a href="https://github.com/jeffotoni">@jeffotoni</a> · Go is the language, teaching is the mission.</sub>
