Léelo en: [🇺🇸 English](README.md) | [🇧🇷 Português](README_pt.md) | **🇪🇸 Español**
---

# Arquitectura de Backend, HTTP y API
### Desde los fundamentos de la comunicación web hasta la implementación con `net/http` en Go

<br>

![Go Version](https://img.shields.io/badge/Go-1.22%2B-00ADD8?style=flat-square&logo=go&logoColor=white)
![Status](https://img.shields.io/badge/status-en%20desarrollo-F59E0B?style=flat-square)
![License](https://img.shields.io/badge/license-MIT-10B981?style=flat-square)
![Author](https://img.shields.io/badge/author-jeffotoni-0EA5E9?style=flat-square&logo=github)

<br>

> **Un backend consistente no nace de los endpoints o los frameworks.**
> Nace de una comprensión profunda de los protocolos, contratos bien definidos, reglas de negocio aisladas, gestión de estados, seguridad en capas y operaciones pensadas desde el primer commit.
> El código es una consecuencia. La decisión es el fundamento.

Este material no comienza en la ruta, el manejador o el framework. Comienza antes, en la raíz: en el origen del backend como disciplina, en el papel del servidor dentro de los sistemas distribuidos y en la semántica que hace que una API sea predecible, comprensible y operable en producción. HTTP, REST, contratos, serialización, seguridad y observabilidad aparecen aquí como partes de un mismo sistema cohesivo, no como temas aislados. La implementación en Go viene después, cuando la base conceptual ya sustenta las decisiones técnicas con precisión e intención.

---

## ✦ Sobre el autor

Desarrollado por **Jefferson Otoni Lima (Jeffotoni)**, Ingeniero de Software Senior y Arquitecto de Soluciones con más de **22 años de experiencia** construyendo sistemas distribuidos de alto rendimiento. Especialista en diseño de APIs, arquitectura cloud-native, Go y ecosistemas de backend a escala. Creador del **Quick Framework**, autor de **Go Bootcamp** y colaborador activo de la comunidad de Go en Brasil y en el mundo.

[![LinkedIn](https://img.shields.io/badge/LinkedIn-jeffotoni-0A66C2?style=flat-square&logo=linkedin)](https://www.linkedin.com/in/jeffotoni)
[![GitHub](https://img.shields.io/badge/GitHub-jeffotoni-181717?style=flat-square&logo=github)](https://github.com/jeffotoni)
[![Site](https://img.shields.io/badge/Site-jeffotoni.com-10B981?style=flat-square)](http://jeffotoni.com)

---

## 🚀 Inicio Rápido

Obtén el proyecto y aloja esta documentación localmente en segundos.

### 1. Clonar el repositorio
```bash
git clone https://github.com/jeffotoni/go-nethttp-learn
cd go-nethttp-learn
```

### 2. Alojar la documentación localmente
Elige tu herramienta favorita para servir el `index.html`, CSS y JS en el puerto 3000:

**Usando Go (Al estilo Gopher)**
```bash
go run static/main.go
```

**Usando Python**
```bash
python3 -m http.server 3000
```

**Usando Node.js (npx)**
```bash
npx serve .
```

---

## Estructura del Material

| | Parte | Qué encontrarás |
|:---:|---|---|
| **I** | Fundamentos de Backend | Origen del backend, papel del servidor, responsabilidades y pilares del backend moderno |
| **II** | Semántica de HTTP y API | Métodos, códigos de estado, cuerpo, caché, restricciones REST y comportamiento correcto de la API |
| **III** | Servidores e Infraestructura | Tipos de servidores, serialización, contexto operativo y decisiones frecuentes en el backend real |
| **IV** | Seguridad, Observabilidad y Contratos | Autenticación, autorización, HTTPS, logs estructurados, métricas, trazado y versionado de APIs |
| **V** | Implementación con Go | Paquete `net/http` en profundidad, sin abstracciones innecesarias y con una base conceptual consolidada |

---

## Temas cubiertos

| Bloque | Qué encontrarás |
|---|---|
| Fundamentos de Backend | Origen del backend, papel del servidor, responsabilidades y pilares del backend moderno |
| HTTP y protocolos | Evolución de HTTP, keep-alive, capas de red, TCP, UDP y contexto de comunicación |
| REST y semántica | Restricciones, recursos, métodos, códigos de estado, payloads y madurez de la API |
| Servidores e infraestructura | Servidores web, proxy inverso, serialización y herramientas que rodean al backend en producción |
| Seguridad | Autenticación vs autorización, JWT, API keys, CORS, HTTPS y rate limiting como decisión de diseño |
| Observabilidad | Logs estructurados, métricas, trazado distribuido y qué significa un backend operable |
| Diseño de Contratos | Versionado de APIs, documentación como parte del contrato y evolución sin romper a los clientes |
| Go y `net/http` | Manejadores, `Request`, `ResponseWriter`, `ServeMux` y `Server` |
| Servidor API en la práctica | Estandarización de respuestas, validación, endpoints de salud (health), middleware y ejecución local |

---

## Objetivos del Curso

- Entender el backend como un **sistema**, no como una colección de rutas y manejadores
- Dominar los fundamentos de **HTTP, REST y semántica de APIs** con precisión conceptual
- Comprender la **seguridad y observabilidad** como pilares de diseño, no como capas adicionales
- Aprender a **definir y evolucionar contratos de API** sin romper clientes ni acumular deuda técnica
- Conectar toda la base conceptual con una **implementación práctica y fundamentada en Go**
- Desarrollar el razonamiento para **tomar decisiones técnicas con claridad**, no solo reproducir patrones
- Salir del curso capaz de construir, operar y evolucionar un **backend de calidad para producción**

---

## Ruta de Contenido

| Etapa | Tema |
|:---:|---|
| `1` | Fundamentos de backend, servicios web y arquitectura cliente-servidor |
| `2` | HTTP, conexiones, keep-alive y pila de protocolos |
| `3` | REST, recursos, semántica de API y niveles de madurez |
| `4` | Servidores, proxy inverso, serialización y contexto de infraestructura |
| `5` | Seguridad: autenticación, autorización, HTTPS, JWT y rate limiting |
| `6` | Observabilidad: logs estructurados, métricas y trazado distribuido |
| `7` | Diseño de contratos, versionado y documentación de API |
| `8` | Visión general de Go para la construcción de APIs |
| `9` | Fundamentos de `net/http` |
| `10` | Manejadores, `ServeMux`, `Server` y flujo de solicitud/respuesta |
| `11` | Servidor API, seguridad y ejecución local con Docker |

---

## Recursos del Manual Oficial

| Canal | Enlace | Objetivo |
|---|---|---|
| Podcast | [Diving into backend](https://youtu.be/uk1hwBAKGLc) | Reforzar el contexto conceptual del material |
| NotebookLM | [NotebookLM, Manual Chat](https://notebooklm.google.com/notebook/0421b1d1-9c27-415a-a3d2-bc83ce397b1f) | Pregunta, explora y aprende a través de chat, presentaciones, podcasts y más |
| Go Bootcamp | [Trilha completa de aprendizado](https://gobootcamp.jeffotoni.com/br/index.html) | Curso completo y trilha de Go por Jeffotoni |
| LinkedIn | [linkedin.com/in/jeffotoni](https://www.linkedin.com/in/jeffotoni) | Perfil profesional del autor |
| GitHub | [github.com/jeffotoni](https://github.com/jeffotoni) | Repositorios y proyectos del autor |
| Sitio Web | [jeffotoni.com](https://jeffotoni.com) | Sitio web personal y blog del autor |

---

## Referencias de Go

| Referencia | Tipo | Enlace | Enfoque |
|---|:---:|---|---|
| Sitio Oficial de Go | Oficial | [go.dev](https://go.dev/) | Portal principal del lenguaje |
| ChatBoot con Go | Oficial | [ChatBoot Google Go](https://codewiki.google/github.com/golang/go#community-guidelines-and-support) | Asistente oficial del lenguaje |
| Tutorial Oficial | Oficial | [go.dev/doc/tutorial](https://go.dev/doc/tutorial/) | Paso a paso para comenzar |
| Tour of Go | Oficial | [go.dev/tour/welcome/1](https://go.dev/tour/welcome/1) | Aprendizaje interactivo |
| Especificación del Lenguaje | Oficial | [go.dev/ref/spec](https://go.dev/ref/spec) | Reglas formales del lenguaje |
| Effective Go | Oficial | [go.dev/doc/effective_go](https://go.dev/doc/effective_go) | Estilo y mejores prácticas |
| Notas de Versión | Oficial | [go.dev/doc/devel/release](https://go.dev/doc/devel/release) | Historial de versiones |
| Notas de Go 1.26 | Oficial | [go.dev/doc/go1.26](https://go.dev/doc/go1.26) | Novedades de la versión 1.26 |
| Blog de Go 1.26 | Oficial | [go.dev/blog/go1.26](https://go.dev/blog/go1.26?ref=dailydev) | Explicaciones prácticas del lanzamiento |
| Go by Example | Comunidad | [gobyexample.com](https://gobyexample.com) | Ejemplos directos y cortos |
| Quick Framework | Comunidad | [github.com/jeffotoni/quick](https://github.com/jeffotoni/quick) | Framework ligero para APIs en Go |

---

## Referencias de Jeffotoni: Go y Arquitectura

<details>
<summary><strong>Ver todos los proyectos y repositorios</strong></summary>

<br>

| Proyecto | Enlace | Enfoque |
|---|---|---|
| Go Bootcamp | [gobootcamp.jeffotoni.com](https://gobootcamp.jeffotoni.com/br/index.html) | Ruta completa de aprendizaje de Go |
| Sitio Personal | [jeffotoni.com](http://jeffotoni.com) | Contenido, artículos y materiales del autor |
| Go Manual | [gomanual.jeffotoni.com](https://gomanual.jeffotoni.com/) | Manual de referencia de Go |
| Go Roadmap | [github.com/jeffotoni/groadmap](https://github.com/jeffotoni/groadmap) | Visión macro del viaje en Go |
| Quick Framework | [github.com/jeffotoni/quick](https://github.com/jeffotoni/quick) | Framework ligero y performante para APIs en Go |
| Quick Benchmarks | [github.com/goquick-run/benchmarks](https://github.com/goquick-run/benchmarks) | Comparativas de rendimiento entre frameworks |
| Go Example | [github.com/jeffotoni/goexample](https://github.com/jeffotoni/goexample) | Colección de ejemplos prácticos en Go |
| Go Cache | [github.com/jeffotoni/gocache](https://github.com/jeffotoni/gocache) | Estrategias de caché en Go |
| Go Hexagonal DDD | [github.com/jeffotoni/go-hexagonal-ddd](https://github.com/jeffotoni/go-hexagonal-ddd) | Arquitectura Hexagonal y DDD en Go |
| Go gRPC Lecture | [github.com/jeffotoni/gogrpc.palestra](https://github.com/jeffotoni/gogrpc.palestra) | Materiales y ejemplos de gRPC en Go |
| Go Workshop DevOps | [github.com/jeffotoni/goworkshopdevops](https://github.com/jeffotoni/goworkshopdevops) | Prácticas de Go aplicadas a DevOps |
| Benchmark | [github.com/jeffotoni/benchmark](https://github.com/jeffotoni/benchmark) | Estudios y análisis de benchmark |

</details>

---

## Tabla de Contenidos

- [Estructura del Material](#estructura-del-material)
- [Temas cubiertos](#temas-cubiertos)
- [Objetivos del Curso](#objetivos-del-curso)
- [Ruta de Contenido](#ruta-de-contenido)
- [Recursos del Manual Oficial](#recursos-del-manual-oficial)
- [Referencias de Go](#referencias-de-go)
- [Referencias de Jeffotoni](#referencias-de-jeffotoni-go-y-arquitectura)
- [1. Contexto: Servicios Web, REST y Protocolos](#1-contexto-servicios-web-rest-y-protocolos)
  - [Visión General de los Servicios Web](#visión-general-de-los-servicios-web)
  - [Diagramas de Comunicación](#diagramas-de-comunicación)
  - [Evolución Rápida de HTTP](#evolución-rápida-de-http)
  - [Diagrama de Keep-Alive](#diagrama-de-keep-alive)
  - [Keep-Alive: HTTP/1.0 -> HTTP/1.1 -> HTTP/2](#keep-alive-http10---http11---http2)
  - [HTTP, TCP y UDP (Diferencia Rápida)](#http-tcp-y-udp-diferencia-rápida)
  - [Modelo OSI y TCP/IP (Diagrama)](#modelo-osi-7-capas)
  - [Modelo TCP/IP (4 Capas)](#modelo-tcpip-4-capas)
  - [REST vs RESTful](#rest-vs-restful)
  - [Significado de Acrónimos](#significado-de-acrónimos)
  - [Restricciones REST](#restricciones-rest)
  - [Diagrama de Restricciones REST](#diagrama-de-restricciones-rest)
  - [Interfaz Uniforme (detallada en 4 partes)](#interfaz-uniforme-detallada-en-4-partes)
  - [Niveles de Madurez (Richardson)](#niveles-de-madurez-richardson)
  - [Diagrama de Richardson](#diagrama-de-richardson)
  - [Métodos HTTP (Verbos HTTP)](#métodos-http-verbos-http)
  - [Cuerpo en REST (solicitud/respuesta) con estados en la práctica](#cuerpo-en-rest-solicitudrespuesta-con-estados-en-la-práctica)
  - [Códigos de Estado Esenciales para APIs](#códigos-de-estado-esenciales-para-apis)
  - [Formatos de Serialización](#formatos-de-serialización)
  - [Servidores Web y de Aplicaciones](#servidores-web-y-de-aplicaciones)
  - [Servidores Web/Proxy Inverso hechos en Go](#servidores-webproxy-inverso-hechos-en-go)
  - [Ecosistema de Go en DevOps](#ecosistema-de-go-en-devops)
- [2. Visión General de Go para APIs](#2-visión-general-de-go-para-apis)
  - [Qué es Go](#qué-es-go)
  - [Diferenciales de Go para Construir APIs](#diferenciales-de-go-para-construir-apis)
  - [Concurrencia en Go (Simple de Entender)](#concurrencia-en-go-simple-de-entender)
  - [Compilado, Estático y Dinámico (En la Práctica)](#compilado-estático-y-dinámico-en-la-práctica)
  - [Servidor HTTP Incorporado](#servidor-http-incorporado)
  - [Palabras Reservadas Oficiales del Lenguaje (25)](#palabras-reservadas-oficiales-del-lenguaje-25)
- [3. Fundamentos de `net/http`](#3-fundamentos-de-nethttp)
  - [El Paquete `net/http`](#el-paquete-nethttp)
  - [Mini Referencia de Componentes](#mini-referencia-de-componentes)
  - [Anatomía Mínima de un Manejador (`w` y `r`)](#anatomía-mínima-de-un-manejador-w-y-r)
- [4. Manual Práctico: ListenAndServe (Fase Cero)](#4-manual-práctico-listenandserve-fase-cero)
  - [4.1 Diferencia Esencial: `HandleFunc` vs `HandlerFunc`](#41-diferencia-esencial-handlefunc-vs-handlerfunc)
  - [4.2 Qué Acepta `ListenAndServe`](#42-qué-acepta-listenandserve)
  - [4.3 Variaciones Base (sin `ServeMux` personalizado)](#43-variaciones-base-sin-servemux-personalizado)
  - [4.4 Algunas Posibilidades](#44-algunas-posibilidades)
  - [4.5 `ServeMux` con patrón de método + `http.Server`](#45-servemux-con-patrón-de-método--httpserver)
  - [4.6 ¿Cuándo usar `http.Handler`?](#46-cuándo-usar-httphandler)
  - [4.7 Apagado Elegante (Graceful Shutdown)](#47-apagado-elegante-graceful-shutdown)
  - [4.8 Probando manejadores con `httptest`](#48-probando-manejadores-con-httptest)
  - [4.9 Dos Servidores en Puertos Distintos con Goroutines](#49-dos-servidores-en-puertos-distintos-con-goroutines)
- [5. Servidor API](#5-servidor-api)
  - [5.0 helpers.go: funciones compartidas](#50-helpersgo-funciones-compartidas)
  - [5.1 Estandarización de Respuestas](#51-estandarización-de-respuestas)
  - [5.2 Mapa de Errores y Estados por Escenario](#52-mapa-de-errores-y-estados-por-escenario)
  - [5.3 Organización de Rutas](#53-organización-de-rutas)
  - [5.4 Validación de Entrada del Servidor](#54-validación-de-entrada-del-servidor)
  - [5.5 Endpoints de Salud (Health)](#55-endpoints-de-salud-health)
  - [5.6 Middleware de Auth Básica](#56-middleware-de-auth-básica)
  - [5.7 Variables de Entorno](#57-variables-de-entorno)
  - [5.8 Middleware de CORS](#58-middleware-de-cors)
  - [5.9 Documentación con OpenAPI/Swagger](#59-documentación-con-openapiswagger)
- [6. Docker: Construir y Ejecutar Local](#6-docker-construir-y-ejecutar-local)
  - [6.1 Dockerfile multi-etapa (Alpine + zona horaria de Brasil)](#61-dockerfile-multi-etapa-alpine--zona-horaria-de-brasil)
  - [6.2 Comandos Básicos de Docker](#62-comandos-básicos-de-docker)

---

## 1. Contexto: Servicios Web, REST y Protocolos

### Visión General de los Servicios Web

| Estilo/Tecnología | Año (origen/aparición) | Característica Principal | Cuándo aparece más |
|---|---:|---|---|
| SOAP | 1998 | Contrato rígido, XML, WSDL | Legado corporativo e integraciones formales |
| REST | 2000 | Estilo arquitectónico sobre HTTP | APIs web en general |
| WebHooks | 2007 | Callback HTTP orientado a eventos (push del servidor al cliente) | Pagos, integraciones, pipelines de CI/CD |
| WebSocket | 2011 | Conexión bidireccional y persistente sobre HTTP | Chats, dashboards en tiempo real, juegos |
| SSE (Server-Sent Events) | 2006 | El servidor envía eventos sobre HTTP estándar (unidireccional) | Notificaciones, streaming de respuestas, LLMs |
| GraphQL | 2015 | El cliente define los campos de la respuesta | Escenarios con múltiples vistas de datos |
| gRPC | 2015 | RPC con Protobuf sobre HTTP/2 | Comunicación interna entre microservicios |


### Diagramas de Comunicación

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

### Evolución Rápida de HTTP

| Protocolo | Año | Aspectos destacados |
|---|---:|---|
| HTTP/0.9 | 1991 | Versión original; solo GET, sin cabeceras, sin códigos de estado |
| HTTP/1.0 | 1996 | Cabeceras, códigos de estado y múltiples tipos de contenido; conexión cerrada por defecto |
| HTTP/1.1 | 1997 | Keep-alive estándar, pipelining, host obligatorio; base de la web durante décadas |
| HTTP/2 | 2015 | Binario, multiplexación, compresión de cabeceras (HPACK), server push |
| HTTP/3 | 2022 | QUIC sobre UDP, menor latencia en redes inestables, conexión más resiliente |

### Keep-Alive: HTTP/1.0 -> HTTP/1.1 -> HTTP/2

![Keep-Alive](docs/diagrams/keepalive.svg)


```text
HTTP/1.0 (1996)
├─ Conexión cerrada después de CADA solicitud/respuesta
├─ Keep-Alive era OPCIONAL (vía cabecera)
└─ Cabecera: Connection: keep-alive (explícita)

HTTP/1.1 (1997)
├─ Keep-Alive es el valor por DEFECTO
├─ Conexiones persistentes por defecto
├─ Para cerrar: Connection: close
└─ Mejor rendimiento nativo

HTTP/2 (2015)
├─ Multiplexación sobre una sola conexión
├─ Keep-Alive implícito
└─ Múltiples solicitudes simultáneas
```

#### HTTP/1.0 - Activando Keep-Alive

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

#### HTTP/1.1 - Keep-Alive por Defecto

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

#### HTTP/1.1 - Cerrando Conexión Explícitamente

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

### HTTP, TCP y UDP (Diferencia Rápida)

Referencia educativa:
- Modelo OSI (7 capas)
- Modelo TCP/IP (4 capas, el más usado en la práctica)

#### Modelo OSI (7 Capas)

![OSI Model and TCP/IP](docs/diagrams/osi-tcpip.svg)


| Capa | Nombre | Función Principal | Ejemplos de Protocolos / Tecnologías |
|---:|---|---|---|
| 7 | Aplicación | Interfaz con el usuario y las aplicaciones | HTTP, HTTPS, FTP, SMTP, DNS |
| 6 | Presentación | Formateo, cifrado, compresión | SSL/TLS, JPEG, MP3, JSON |
| 5 | Sesión | Control de sesión/conexión | NetBIOS, RPC |
| 4 | Transporte | Comunicación extremo a extremo, control de errores | TCP, UDP |
| 3 | Red | Direccionamiento lógico y enrutamiento | IP, ICMP, IPSec |
| 2 | Enlace de Datos | Comunicación dentro de la red local | Ethernet, ARP, PPP |
| 1 | Física | Transmisión eléctrica/óptica de bits | Cables, Fibra, Wi-Fi (parte física) |

#### Modelo TCP/IP (4 Capas)

| Capa TCP/IP | Equivalente OSI | Ejemplos |
|---|---|---|
| Aplicación | 7, 6 y 5 | HTTP, FTP, SMTP, DNS |
| Transporte | 4 | TCP, UDP |
| Internet | 3 | IP, ICMP |
| Acceso a Red | 2 y 1 | Ethernet, Wi-Fi |

Resumen rápido de la pila:
- HTTP/1.1 y HTTP/2: `HTTP -> TCP -> IP`
- HTTP/3: `HTTP -> QUIC(UDP) -> IP`

Analogía educativa (mensaje y correo):
1. Aplicación: escribir el mensaje
2. Transporte: ponerlo en el sobre (TCP comprueba si llegó)
3. Red: elegir la ruta hacia el destino
4. Enlace de Datos: llevarlo a la oficina de correos local
5. Física: carretera y camión

Ejemplo básico en Go:

```go
http.ListenAndServe(":8080", nil)
```

Es decir, `net/http` está en la cima de la pila pero depende de todas las capas inferiores.

### REST vs RESTful

- `REST` es un estilo arquitectónico (conjunto de restricciones)
- `RESTful` es la API que aplica REST de forma consistente en la práctica

### Significado de Acrónimos

| Término | Significado | Tipo | Dónde encaja |
|---|---|---|---|
| HTTP | HyperText Transfer Protocol | Protocolo | Capa de Aplicación |
| REST | Representational State Transfer | Estilo Arquitectónico | Usa HTTP |
| SOAP | Simple Object Access Protocol | Protocolo | Usa HTTP (normalmente) |
| gRPC | Google Remote Procedure Call | Framework / RPC | Usa HTTP/2 |

### Restricciones REST

![REST Constraints](docs/diagrams/rest-constraints.svg)


```text
┌───────────────────────────────────────────────────────┐
│ 1. Cliente-Servidor                                   │
│    Separación de intereses                            │
│                                                       │
│ 2. Sin estado (Stateless)                             │
│    Cada solicitud es independiente                    │
│                                                       │
│ 3. Cacheable                                          │
│    Las respuestas deben indicar si se pueden guardar  │
│                                                       │
│ 4. Interfaz Uniforme                                  │
│    ├─ Identificación de recursos                      │
│    ├─ Manipulación a través de representaciones       │
│    ├─ Mensajes auto-descriptivos                      │
│    └─ HATEOAS                                         │
│                                                       │
│ 5. Sistema en Capas                                   │
│    El cliente no sabe si se conecta directamente al   │
│    servidor final o a capas intermedias               │
│                                                       │
│ 6. Código bajo demanda (opcional)                     │
│    El servidor puede enviar código ejecutable         │
└───────────────────────────────────────────────────────┘
```

#### Previsualización educativa de cada restricción

| Restricción | Previsualización Práctica |
|---|---|
| Cliente-Servidor | Frontend y backend evolucionan de forma independiente |
| Sin estado | El token/autenticación y el contexto van en cada solicitud |
| Cacheable | Uso de `Cache-Control`, `ETag`, `Last-Modified` |
| Interfaz Uniforme | URI + método + estado + representación consistente |
| Sistema en Capas | CDN, equilibrador de carga y API gateway entre cliente y app |
| Código bajo demanda | Ej.: JavaScript entregado al cliente (opcional) |

#### Interfaz Uniforme (detallada en 4 partes)

**1. Identificación de Recursos**

Cada recurso tiene un identificador único (URI).

Ejemplos:
- `/users/123`
- `/posts/456`

**2. Manipulación a través de representaciones**

El cliente manipula los recursos a través de representaciones (`JSON`, `XML`, etc.).
El servidor envía la representación del recurso, no el recurso en memoria.

Ejemplo práctico:
- el cliente recibe el JSON del usuario
- al hacer `PUT /users/123`, envía una nueva representación de ese usuario

**3. Mensajes auto-descriptivos**

Cada mensaje debe contener información suficiente para su procesamiento.

```http
Content-Type: application/json
Accept: application/json
```

Con esto, cliente y servidor entienden el formato de entrada/salida sin un "acuerdo oculto".

**4. HATEOAS (Hypermedia As The Engine Of Application State)**

La API devuelve enlaces para las siguientes acciones válidas, y el cliente navega a través de estos enlaces.

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

En la práctica, HATEOAS es el elemento menos implementado en la mayoría de las APIs RESTful.

Razones comunes:
- los clientes móviles/web prefieren un contrato fijo documentado en OpenAPI/Swagger
- los equipos priorizan la simplicidad de implementación y mantenimiento
- los gateways, el versionado y los SDKs tienden a centralizar el flujo fuera de la hipermedia
- el coste extra de modelado no siempre genera un beneficio claro en el producto

### Niveles de Madurez (Richardson)

![Richardson Maturity Levels](docs/diagrams/richardson.svg)


El modelo fue propuesto por **Leonard Richardson**, un arquitecto de software que escribió sobre APIs REST y ayudó a popularizar las mejores prácticas en la construcción de servicios HTTP.

Objetivo del modelo:
- evaluar qué tan RESTful es una API
- clasificar las APIs HTTP en niveles de madurez
- ayudar a evolucionar las APIs desde un RPC disfrazado hacia un REST bien estructurado

Tiene 4 niveles (del 0 al 3).

| Nivel | Nombre | Descripción |
|---:|---|---|
| 0 | POX / RPC sobre HTTP | HTTP solo como transporte |
| 1 | Recursos | Recursos identificados por URI |
| 2 | Verbos + estado | Uso correcto de verbos HTTP y códigos de estado |
| 3 | HATEOAS | Hipermedia guiando al cliente |

#### Nivel 0 - El pantano de POX

- usa HTTP solo como transporte
- generalmente un solo endpoint
- es común ver `POST` para todo

Ejemplo:

```http
POST /api
Content-Type: application/json

{
  "action": "getUser",
  "id": 10
}
```

Aquí HTTP se convierte simplemente en un "túnel" para comandos RPC.

#### Nivel 1 - Recursos

- separa por recursos (diferentes URLs)
- puede seguir usando `POST` para casi todo

Ejemplos de recursos:
- `/users`
- `/orders`

Ganancia principal: inicio de la organización por dominio.

#### Nivel 2 - Verbos HTTP Correctos

- usa `GET`, `POST`, `PUT`, `DELETE` correctamente
- usa códigos de estado apropiados

Ejemplos:

```http
GET /users/10
DELETE /users/10
```

Aquí se encuentran la mayoría de las APIs que el mercado llama REST en la práctica.

#### Nivel 3 - HATEOAS

`HATEOAS` = *Hypermedia As The Engine Of Application State*.

La respuesta incluye enlaces a los posibles siguientes pasos.

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

Aquí la API guía al cliente de forma dinámica.

En la práctica:
- la mayoría de las APIs modernas se quedan en el nivel 2
- pocas implementan HATEOAS plenamente
- muchas APIs se autodenominan REST pero aún están en el nivel 1

### Métodos HTTP (Verbos HTTP)

| Verbo | Uso Correcto | Ejemplo |
|---|---|---|
| `GET` | Obtener datos | `GET /users/123` |
| `POST` | Crear | `POST /users` |
| `PUT` | Reemplazar | `PUT /users/123` |
| `PATCH` | Actualizar parcialmente | `PATCH /users/123` |
| `DELETE` | Eliminar | `DELETE /users/123` |

### Cuerpo en REST (solicitud/respuesta) con estados en la práctica

Reglas simples:
- `GET` y `DELETE`: normalmente sin cuerpo
- `POST`, `PUT`, `PATCH`: normalmente con cuerpo
- Definir siempre `Content-Type` y validar la entrada

#### GET (obtener recurso)

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

#### POST (crear recurso)

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

#### PUT (reemplazar recurso)

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

#### PATCH (actualización parcial)

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

#### DELETE (eliminar recurso)

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

### Códigos de Estado Esenciales para APIs

| Escenario | Estado |
|---|---|
| Éxito con retorno | `200 OK` |
| Creación de recurso | `201 Created` |
| Éxito sin cuerpo | `204 No Content` |
| Error de entrada | `400 Bad Request` |
| No autenticado | `401 Unauthorized` |
| Sin permiso | `403 Forbidden` |
| No encontrado | `404 Not Found` |
| Conflicto de estado | `409 Conflict` |
| Error de validación semántica | `422 Unprocessable Entity` |
| Error interno | `500 Internal Server Error` |
| Demasiadas solicitudes | `429 Too Many Requests` |

### Formatos de Serialización

Para este curso, el enfoque principal será `JSON` en APIs REST con Go.

| Formato | Tipo | Cuándo usar |
|---|---|---|
| JSON | Texto | APIs REST públicas y simplicidad |
| Protobuf | Binario | gRPC y comunicación interna de alto rendimiento |
| Avro | Binario | Streaming/Kafka con fuerte evolución de esquema |
| MessagePack | Binario | Payload más compacto sin demasiada complejidad |
| CBOR | Binario | IoT y escenarios con estándar IETF |

Ejemplo mínimo en Go:

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
	// Serializar (struct -> JSON)
	u := User{Name: "Jeff", Email: "jeff@email.com"}
	b, _ := json.Marshal(u)
	fmt.Println(string(b)) // {"name":"Jeff","email":"jeff@email.com"}

	// Deserializar (JSON -> struct)
	var u2 User
	_ = json.Unmarshal(b, &u2)
	fmt.Println(u2.Name) // Jeff
}
```

### Servidores Web y de Aplicaciones

| Servidor | Año | Categoría | Nota |
|---|---:|---|---|
| Apache HTTP Server | 1995 | Servidor Web | Base histórica de la web de código abierto |
| IIS | 1995 | Servidor Web | Servidor web de Microsoft |
| nginx | 2004 | Servidor Web/Proxy Inverso | Muy utilizado en alta concurrencia |
| Caddy | 2015 | Servidor Web | HTTPS automático por defecto |
| Tomcat | 1999 | Servidor de Aplicaciones (Java) | Muy común en aplicaciones Java |
| JBoss / WildFly | 2006 (WildFly 2014) | Servidor de Aplicaciones (Java) | Línea enterprise del ecosistema Java |

### Servidores Web/Proxy Inverso hechos en Go

| Proyecto | Categoría | Dónde aparece mucho | Por qué Go ayuda aquí |
|---|---|---|---|
| Caddy | Servidor web / proxy inverso | APIs, TLS automático, edge simple | Binario único, concurrencia nativa y despliegue fácil |
| Traefik | Proxy inverso / ingress | Docker, Kubernetes, service discovery | Integración nativa con la nube y alto rendimiento de red |
| Fabio | Equilibrador de carga / proxy inverso | Entornos con Consul | Simplicidad operativa y buen modelo concurrente |
| `httputil.ReverseProxy` | Proxy inverso (stdlib) | APIs internas, proxies simples sin dependencias extra | Nativo en la stdlib `net/http/httputil`, cero dependencias |

### Cuota de Mercado (visión macro)

![Web Servers Market Share](docs/diagrams/market-share.svg)

### Ecosistema de Go en DevOps

Go se ha convertido en uno de los lenguajes centrales del ecosistema **CNCF/DevOps** al ofrecer:
- Binarios portátiles y fáciles de distribuir
- Buen rendimiento de red y concurrencia
- Toolchain estable para proyectos de infraestructura

| Herramienta | Categoría | Relación con Go |
|---|---|---|
| Docker (Moby/Engine) | Contenerización | Implementación central en Go (con partes en otros lenguajes) |
| Kubernetes | Orquestación | Proyecto core en Go |
| Consul | Service discovery/config | Core en Go |
| etcd | KV distribuido | Core en Go |
| Terraform | Infraestructura como Código | Core en Go |
| Vault | Gestión de secretos | Core en Go |
| CockroachDB | Base de datos SQL distribuida | Core mayoritariamente en Go |
| InfluxDB | Base de datos de series temporales | Fuerte uso de Go en el core |
| Prometheus | Monitorización | Core en Go |
| Grafana | Observabilidad | Backend en Go (frontend en TypeScript) |
| Gitea | Forja Git/autohospedado | Core en Go |
| Helm | Gestor de paquetes de Kubernetes | Core en Go |
| ArgoCD | GitOps / CD para Kubernetes | Core en Go |
| Cilium | Redes / eBPF para Kubernetes | Core en Go |

---

## 2. Visión General de Go para APIs

### Qué es Go

Go es un lenguaje compilado, de tipado estático y con una sintaxis simple, enfocado en la productividad, el rendimiento y la legibilidad.

### Año de Lanzamiento y Nombres Clave

| Ítem | Información |
|---|---|
| Inicio del proyecto | 2007 (Google) |
| Lanzamiento público | 2009 |
| Versión 1.0 | 2012 |
| Creadores | Robert Griesemer, Rob Pike, Ken Thompson |

### Diferenciales de Go para Construir APIs

- Biblioteca estándar sólida (`net/http`, `encoding/json`, `context`, `database/sql`)
- Compilación rápida y despliegue simple (binario único)
- Concurrencia nativa con goroutines y canales
- Código más predecible con menos complejidad accidental
- Excelente robustez para APIs de alta carga y baja latencia
- Pruebas integradas en el toolchain (`go test`) con soporte práctico para pruebas unitarias y table-driven
- Cobertura nativa, benchmark y fuzz testing (`-cover`, `-bench`, `-fuzz`) para aumentar la fiabilidad de la API
- `context.Context` nativo para cancelación, tiempo de espera y propagación de valores entre manejadores y goroutines
- Compilación cruzada nativa: compila para cualquier SO/arquitectura con `GOOS` y `GOARCH` sin toolchain adicional

### Concurrencia en Go (Simple de Entender)

- `goroutine`: función que se ejecuta de forma concurrente con bajo coste
- `channel`: canal seguro para la comunicación entre goroutines
- `select`: coordina múltiples canales y tiempos de espera

Modelo mental:
1. Usa goroutines para estructurar el trabajo concurrente (no confundir con paralelismo)
2. Intercambia datos vía canales (en lugar de compartir siempre la memoria)
3. Controla la cancelación y los plazos con `context.Context`
4. El runtime/scheduler decide cuándo hay paralelismo real (ej.: múltiples núcleos)

Ejemplo mínimo con goroutine + canal:

```go
package main

import "fmt"

func sum(a, b int, ch chan int) {
	ch <- a + b
}

func main() {
	ch := make(chan int)
	go sum(3, 7, ch) // se ejecuta de forma concurrente
	result := <-ch  // espera a que llegue el valor
	fmt.Println(result) // 10
}
```

Ejemplo con `context.Context`, tiempo de espera en el manejador:

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

### Compilado, Estático y Dinámico (En la Práctica)

| Aspecto | Cómo funciona en Go |
|---|---|
| Compilación | AOT (ahead-of-time), genera binario nativo |
| Tipado | Estático y fuerte en tiempo de compilación |
| Enlace (Linking) | Generalmente estático; puede usar dinámico en escenarios con `cgo` |
| Runtime | Dinámico para GC, scheduler y reflexión cuando es necesario |
| Compilación cruzada | `GOOS=linux GOARCH=amd64 go build` genera el binario para cualquier plataforma |

Ejemplo de compilación cruzada:

```bash
# Compilar para Linux AMD64 (desde cualquier SO)
GOOS=linux GOARCH=amd64 go build -o server-linux ./cmd/api

# Compilar para Windows
GOOS=windows GOARCH=amd64 go build -o server.exe ./cmd/api

# Compilar para ARM (ej: Raspberry Pi)
GOOS=linux GOARCH=arm64 go build -o server-arm ./cmd/api
```

### Go hoy está escrito en Go

Desde Go 1.5, el compilador principal es auto-alojado (escrito en Go).
Todavía existen partes de bajo nivel en ensamblador.

### Servidor HTTP Incorporado

Go ya trae un servidor HTTP incorporado en la stdlib a través de `net/http`.

```go
http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("pong"))
})
http.ListenAndServe(":8080", nil)
```

Esto no reemplaza todos los roles de un nginx/proxy inverso, pero acelera enormemente el desarrollo de APIs.

### Palabras Reservadas Oficiales del Lenguaje (25)

| 1 | 2 | 3 | 4 | 5 |
|---|---|---|---|---|
| `break` | `default` | `func` | `interface` | `select` |
| `case` | `defer` | `go` | `map` | `struct` |
| `chan` | `else` | `goto` | `package` | `switch` |
| `const` | `fallthrough` | `if` | `range` | `type` |
| `continue` | `for` | `import` | `return` | `var` |

---

## 3. Fundamentos de `net/http`

### El Paquete `net/http`

El paquete ofrece:
- Cliente HTTP
- Servidor HTTP
- `Request` y `ResponseWriter`
- `Handler`, `HandlerFunc` y `ServeMux`
- Utilidades para cookies, cabeceras y más

Componentes:
- `http.ListenAndServe`
- `http.Request`
- `http.ResponseWriter`
- `http.HandleFunc`
- `http.HandlerFunc`
- `http.Handle`
- `http.Handler`
- `http.ServeMux`
- `http.Server`

### Mini Referencia de Componentes

**`http.ListenAndServe`**

```go
log.Fatal(http.ListenAndServe(":8080", nil))
```

**`http.Request` y `http.ResponseWriter`:**

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

Regla mental rápida:
- `HandleFunc`: función
- `HandlerFunc`: función adaptada a `Handler`
- `Handle`: registra un `Handler`
- `Handler`: comportamiento completo (`ServeHTTP`)

### Anatomía Mínima de un Manejador (`w` y `r`)

![Handler Anatomy](docs/diagrams/handler-anatomy.svg)


**Firma Estándar**

```go
func(w http.ResponseWriter, r *http.Request)
```

**w `http.ResponseWriter`:**
- es la salida de tu API (respuesta al cliente)
- piensa en el orden: **Cabeceras -> Estado -> Cuerpo**

**Principales métodos de `ResponseWriter`**

| Método | Qué hace | Notas importantes |
|---|---|---|
| `Header() http.Header` | Manipula las cabeceras de respuesta | Definir antes de `WriteHeader` |
| `Write([]byte)` | Escribe el cuerpo | Si no se llama a `WriteHeader`, envía `200` automáticamente |
| `WriteHeader(statusCode int)` | Define el estado HTTP | Debe llamarse una sola vez |

Ejemplo corto:

```go
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(http.StatusCreated)
_, _ = w.Write([]byte(`{"ok":true}`))
```

**Reglas importantes**
- después de `WriteHeader`, las cabeceras se congelan
- `Write()` llama implícitamente a `WriteHeader(200)` si no se envió ningún estado
- orden correcto:
1. `Header().Set(...)`
2. `WriteHeader(...)`
3. `Write(...)`

**r `*http.Request`:**
- representa todo lo que el cliente envió en la solicitud

**Campos más usados de `Request`**

| Campo | Tipo | Propósito |
|---|---|---|
| `r.Method` | `string` | Verbo HTTP (`GET`, `POST`, etc.) |
| `r.URL` | `*url.URL` | Ruta y query string (`r.URL.Path`, `r.URL.Query().Get("id")`) |
| `r.Header` | `http.Header` | Cabeceras de la solicitud |
| `r.Body` | `io.ReadCloser` | Cuerpo de la solicitud |
| `r.Host` | `string` | Host llamado |
| `r.RemoteAddr` | `string` | IP/puerto de origen del cliente |
| `r.Proto` | `string` | Protocolo (`HTTP/1.1`, `HTTP/2.0`) |
| `r.ContentLength` | `int64` | Tamaño del cuerpo |

**Anatomía de la URL (cada parte)**

Ejemplo:

```text
https://domain.com/api/v1/user?id=123&debug=true#section
```

| Parte de la URL | Ejemplo | Dónde usar en el servidor Go |
|---|---|---|
| Protocolo (scheme) | `https` | inferir vía `r.TLS` (`nil` = http, no-nil = https) |
| Host | `domain.com` | `r.Host` |
| Ruta (Path) | `/api/v1/user` | `r.URL.Path` |
| Raw query string | `id=123&debug=true` | `r.URL.RawQuery` |
| Parámetros de consulta | `id=123`, `debug=true` | `r.URL.Query().Get("id")`, `r.URL.Query().Get("debug")` |
| Fragmento | `#section` | no llega al servidor (el navegador no lo envía en la solicitud HTTP) |

Ejemplo práctico en el manejador:

```go
scheme := "http"
if r.TLS != nil {
	scheme = "https"
}

fullURL := scheme + "://" + r.Host + r.URL.RequestURI()
// fullURL => https://domain.com/api/v1/user?id=123&debug=true
```

**Campos y métodos útiles de `r.URL` (`*url.URL`)**

| Expresión | Tipo | Propósito |
|---|---|---|
| `r.URL.Path` | `string` | Ruta sin query (`/api/v1/user`) |
| `r.URL.RawQuery` | `string` | Query string en bruto (`id=10&debug=true`) |
| `r.URL.Query()` | `url.Values` | Mapa de parámetros de consulta |
| `r.URL.Query().Get("id")` | `string` | Obtiene el primer valor para la clave |
| `r.URL.Query()["tag"]` | `[]string` | Obtiene todos los valores para la clave repetida |
| `r.URL.EscapedPath()` | `string` | Ruta con escape URL |
| `r.URL.String()` | `string` | URL en formato de texto (bueno para log/debug) |

Trabajando con cabeceras:

```go
r.Header.Get("Authorization")
r.Header.Get("Content-Type")
```

Trabajando con cuerpo JSON:

```go
defer r.Body.Close()
_ = json.NewDecoder(r.Body).Decode(&payload)
```

Mejor práctica: limitar el tamaño del cuerpo:

```go
r.Body = http.MaxBytesReader(w, r.Body, 1<<20) // 1MB
```

**`r.PathValue` (Go 1.22+): extrayendo segmentos de ruta:**

Cuando el patrón de ruta contiene `{name}`, el valor se extrae con `r.PathValue("name")`:

```go
mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id") // extrae "42" de /users/42
	_, _ = fmt.Fprintf(w, `{"id":"%s"}`, id)
})
```

| Expresión | Tipo | Propósito |
|---|---|---|
| `r.PathValue("id")` | `string` | Segmento de ruta con nombre (`/users/{id}`) |
| `r.URL.Query().Get("q")` | `string` | Parámetro de consulta (`?q=value`) |
| `r.URL.Query()["tag"]` | `[]string` | Múltiples valores para la misma clave (`?tag=a&tag=b`) |

---

## 4. Manual Práctico: ListenAndServe (Fase Cero)

En esta fase, el `README.md` es el manual principal para copiar, pegar y ejecutar.

Nota práctica:
- cada ejemplo usa el puerto `:8080`

![ListenAndServe Flow](docs/diagrams/listenandserve.svg)

- ejecuta un ejemplo a la vez (detén el anterior antes de ejecutar el siguiente)

### Línea de razonamiento de la clase

| Orden | Enfoque | Resultado para el estudiante |
|---|---|---|
| 1 | `HandleFunc` vs `HandlerFunc` | Evita los errores de registro más comunes |
| 2 | Qué acepta `ListenAndServe` | Sabe cómo pasar `nil`, `HandlerFunc`, `ServeMux` o un tipo personalizado |
| 3 | Variaciones base sin `ServeMux` personalizado | Domina el flujo HTTP básico |
| 4 | Leer la solicitud: método, ruta, consulta, cabeceras, cuerpo, `PathValue` | Extrae cualquier dato de la solicitud |
| 5 | Escribir la respuesta: cabeceras, estado, cuerpo | Sabe el orden correcto y evita errores |
| 6 | CRUD completo: GET, POST, PUT, PATCH, DELETE | Cubre los verbos con ejemplos reales |
| 7 | `http.Server` con tiempos de espera | Configura el servidor para producción |
| 8 | Middleware: Logger, Auth, cadena | Compone comportamientos reutilizables |
| 9 | Estandarización de respuestas: motivación para `writeJSON` | Entiende por qué existe la sección 5 |
| 10 | Apagado elegante | Apaga el servidor sin perder solicitudes en curso |
| 11 | Probando manejadores con `httptest` | Prueba manejadores sin iniciar un servidor real |

### 4.1 Diferencia Esencial: `HandleFunc` vs `HandlerFunc`

`HandleFunc` es una función de registro.
`HandlerFunc` es un tipo de adaptador (se convierte en `http.Handler`).

```go
// MAL - HandleFunc no devuelve nada
http.Handle("/route", http.HandleFunc(...))

// BIEN - HandlerFunc es un tipo
http.Handle("/route", http.HandlerFunc(...))

// BIEN - HandleFunc registra directamente
http.HandleFunc("/route", ...)
```

### 4.2 Qué Acepta `ListenAndServe`

Firma:

```go
func ListenAndServe(addr string, handler Handler) error
```

El segundo argumento acepta cualquier cosa que implemente la interfaz `http.Handler`, es decir, cualquier tipo que tenga el método `ServeHTTP(w, r)`.

![What ListenAndServe accepts](docs/diagrams/listenandserve-handler.svg)

| Opción | Cuándo usar |
|---|---|
| `nil` | Usa el `DefaultServeMux` global. Simple para ejemplos, pero evítalo en producción |
| `http.HandlerFunc(fn)` | Adapta una función directamente como manejador. Útil para servidor de una sola ruta |
| `http.NewServeMux()` | Enrutador dedicado y aislado. Recomendado para cualquier API real |
| Tipo personalizado con `ServeHTTP` | Cuando necesitas estado interno, composición o cadena de middleware |

**Sobre `DefaultServeMux` y por qué evitarlo en producción:**

`DefaultServeMux` es un `*ServeMux` global creado automáticamente por el paquete `net/http`. Al pasar `nil`, el servidor usa este mux implícitamente. El problema: cualquier paquete importado puede registrar rutas en él a través de `init()`, creando rutas expuestas invisibles y potencialmente no deseadas.

```go
// PRECAUCIÓN: cualquier importación podría haber hecho esto en init()
import _ "some/package" // podría haber registrado /debug/pprof, /metrics, etc.

http.ListenAndServe(":8080", nil) // expone estas rutas sin tu conocimiento
```

Prefiere siempre tu propio `ServeMux`:

```go
mux := http.NewServeMux()
mux.HandleFunc("GET /ping", pingHandler)
http.ListenAndServe(":8080", mux) // solo tus rutas
```

**Tipo personalizado con `ServeHTTP` en la práctica:**

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

**`http.Server` como la alternativa recomendada para producción:**

`http.ListenAndServe` es conveniente pero no permite configurar tiempos de espera. Para producción usa siempre `http.Server`:

```go
srv := &http.Server{
	Addr:              ":8080",
	Handler:           mux,           // tu ServeMux dedicado
	ReadHeaderTimeout: 5 * time.Second,
	ReadTimeout:       15 * time.Second,
	WriteTimeout:      15 * time.Second,
	IdleTimeout:       60 * time.Second,
}
log.Fatal(srv.ListenAndServe())
```

### 4.3 Variaciones Base (sin `ServeMux` personalizado)

#### Ejemplo 4.3.1 - `DefaultServeMux` con `HandleFunc`

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

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/
curl -i localhost:8080/api
```

#### Ejemplo 4.3.2 - `DefaultServeMux` con `Handle`

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

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/
curl -i localhost:8080/api
```

#### Ejemplo 4.3.3 - Manejador único directo (enrutamiento manual)

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

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/
curl -i localhost:8080/api
curl -i localhost:8080/x
```

#### Ejemplo 4.3.4 - Extraer `HandlerFunc` a una variable

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

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/
```

#### Ejemplo 4.3.5 - Convertir a `HandlerFunc`

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

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/
```

### 4.4 Algunas Posibilidades

#### Ejemplo 4.4.1 - Parámetros de URL (`r.URL.Query`)

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

Ejecutar:

```bash
go run main.go
curl -i "localhost:8080/hello?name=jeff"
curl -i localhost:8080/hello
```

#### Ejemplo 4.4.2 - Respuesta JSON

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

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/api/user
```

#### Ejemplo 4.4.3 - Cabeceras personalizadas + código de estado

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

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/
```


#### Ejemplo 4.4.4 - Parámetros de consulta con múltiples valores

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

Ejecutar:

```bash
go run main.go
curl -i "localhost:8080/search?q=golang&tag=go&tag=api&tag=http"
```

#### Ejemplo 4.4.5 - Leer cabeceras personalizadas de la solicitud

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

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/info \
	-H "Authorization: Bearer token123" \
	-H "X-Trace-Id: abc-456" \
	-H "X-Client-Version: 2.1.0"
```

#### Ejemplo 4.4.6 - `r.PathValue` (Go 1.22+)

`r.PathValue` extrae segmentos con nombre directamente del patrón de ruta.

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

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/users/42
curl -i localhost:8080/posts/10/comments/99
```

#### Ejemplo 4.4.7 - PATCH `/api/v1/users/{id}` (actualización parcial)

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

Ejecutar:

```bash
go run main.go
curl -i -X PATCH localhost:8080/api/v1/users/42 \
	-H "Content-Type: application/json" \
	-d '{"email":"new@email.com"}'
```

#### Ejemplo 4.4.8 - DELETE `/api/v1/users/{id}`

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
	// simulación: el id "0" no existe
	if id == "0" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(fmt.Sprintf(`{"error":"user_not_found","id":"%s"}`, id)))
		return
	}
	w.WriteHeader(http.StatusNoContent) // 204: sin cuerpo
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("DELETE /api/v1/users/{id}", deleteUser)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Ejecutar:

```bash
go run main.go
curl -i -X DELETE localhost:8080/api/v1/users/42
curl -i -X DELETE localhost:8080/api/v1/users/0
```

#### Ejemplo 4.4.9 - `http.Redirect`

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// Redirección permanente (301): la URL se movió definitivamente
	mux.HandleFunc("GET /old-path", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/new-path", http.StatusMovedPermanently)
	})
	// Redirección temporal (302): la URL podría volver a cambiar
	mux.HandleFunc("GET /temp", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/new-path", http.StatusFound)
	})
	mux.HandleFunc("GET /new-path", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("has llegado al nuevo destino"))
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/old-path
curl -iL localhost:8080/old-path
curl -i localhost:8080/temp
```

### 4.5 `ServeMux` con patrón de método + `http.Server`

Sí, esta sintaxis existe y es oficial:

```go
mux.HandleFunc("POST /api/v1/user", handler)
```

Nota:
- el patrón de método (`"GET /x"`, `"POST /x"`) requiere Go 1.22+

**Keep-Alive** (importante para API):
- en **HTTP/1.1**, **keep-alive** es el valor por defecto; el cliente tiende a reutilizar la conexión
- el servidor no "activa el keep-alive manualmente", sino que controla el tiempo de inactividad
- en Go, `IdleTimeout` es una configuración clave para las conexiones persistentes
- los proxies/equilibradores de carga en el camino también pueden cerrar conexiones

#### Principales propiedades de `http.Server`

| Campo | Qué hace | Ejemplo |
|---|---|---|
| `Addr` | Dirección/puerto donde el servidor escuchará | `":8080"` |
| `Handler` | Quién procesará las rutas (`mux`, manejador personalizado, etc.) | `mux` |
| `IdleTimeout` | Tiempo de conexión inactiva esperando la siguiente solicitud (keep-alive) | `60 * time.Second` |
| `ReadTimeout` | Tiempo máximo para leer la solicitud completa (cabeceras + cuerpo) | `15 * time.Second` |
| `ReadHeaderTimeout` | Tiempo máximo para leer solo las cabeceras (anti-slowloris) | `5 * time.Second` |
| `WriteTimeout` | Tiempo máximo para escribir la respuesta | `15 * time.Second` |
| `MaxHeaderBytes` | Tamaño máximo de las cabeceras recibidas | `1 << 20` (1MB) |

Ejemplo completo y educativo:

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Solicitud %s procesada\n", r.URL.Path)
	fmt.Printf("Conexión desde: %s\n", r.RemoteAddr)
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

	fmt.Println("Servidor escuchando en :8080")
	fmt.Println("IdleTimeout:", srv.IdleTimeout)
	log.Fatal(srv.ListenAndServe())
}
```

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/
```

Prueba rápida de reutilización de conexión (HTTP/1.1):

```bash
curl -v --http1.1 http://localhost:8080/ http://localhost:8080/
```

Puntos críticos a explicar en clase:
- `ReadTimeout` cubre la lectura completa (cabeceras + cuerpo)
- `ReadHeaderTimeout` protege contra el envío lento de cabeceras (slowloris)
- `IdleTimeout` controla cuánto tiempo permanece abierta la conexión keep-alive sin una nueva solicitud
- `MaxHeaderBytes: 1 << 20` usa el operador de desplazamiento de bits para definir el límite de 1MB
- el keep-alive depende de que el cliente/proxy reutilice la conexión; el servidor define los límites y políticas

*Nota: **Slowloris** es un tipo de ataque DoS (Denegación de Servicio) que mantiene abiertas varias conexiones HTTP enviando datos muy lentamente, sin finalizar la solicitud.*

#### POST `/api/v1/user` con `json.NewDecoder`

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

const maxBodyBytes = 1 << 20 // 1MB -> Operador de Desplazamiento de Bits
// bit shift left (desplazamiento a la izquierda) que significa -> "Desplaza el número 1 a la izquierda 20 posiciones"
// Matemáticas detrás:
// 1 << n  =  1 × 2^n  =  2^n
// Ejemplos:
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
		`{"message":"usuario creado (decoder)","name":"%s","email":"%s"}`,
		in.Name, in.Email,
	)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/user", postUserWithDecoder)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Ejecutar:

```bash
go run main.go
curl -i -X POST localhost:8080/api/v1/user \
  -H "Content-Type: application/json" \
  -d '{"name":"Jeff","email":"jeff@email.com"}'
```

#### POST `/api/v1/user` leyendo el `Body` + `json.Unmarshal`

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
		`{"message":"usuario creado (unmarshal)","name":"%s","email":"%s"}`,
		in.Name, in.Email,
	)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/v1/user", postUserWithUnmarshal)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Ejecutar:

```bash
go run main.go
curl -i -X POST localhost:8080/api/v1/user \
  -H "Content-Type: application/json" \
  -d '{"name":"Jeff","email":"jeff@email.com"}'
```

#### ¿Cuándo usar `Decoder` vs `Unmarshal`?

| Opción | Cuándo usar | Ventaja | Atención |
|---|---|---|---|
| `json.NewDecoder(r.Body).Decode(&v)` | Flujo HTTP estándar leyendo directamente del cuerpo | Simple y directo en el manejador | Menos control sobre el `[]byte` en bruto |
| `io.ReadAll(r.Body)` + `json.Unmarshal(raw, &v)` | Cuando necesitas el cuerpo en bruto antes de convertir | Permite log, auditoría, firma, validación previa | Más verboso y usa memoria para guardar todo el cuerpo |

Lista de verificación mínima para un endpoint de API:
- limitar el tamaño del cuerpo (`http.MaxBytesReader`)
- validar `Content-Type: application/json` para endpoints JSON
- validar el JSON de entrada y devolver un error claro
- responder con un `Content-Type` consistente
- mantener los manejadores con nombre fuera de `main` a medida que el flujo crece

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

Ejecutar:

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
		`{"message":"usuario actualizado","name":"%s","email":"%s"}`,
		in.Name, in.Email,
	)))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("PUT /api/v1/user", putUser)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Ejecutar:

```bash
go run main.go
curl -i -X PUT localhost:8080/api/v1/user \
  -H "Content-Type: application/json" \
  -d '{"name":"Jeff Otoni","email":"new@email.com"}'
```

### 4.6 ¿Cuándo usar `http.Handler`?

![Middleware Chain](docs/diagrams/middleware-chain.svg)


Regla de oro:
- usa `http.Handler` cuando quieras **componer comportamientos** (middleware, cadena, reutilización)
- usa `http.HandlerFunc` cuando quieras **responder a una ruta directamente** (simple y rápido)

`http.Handler` es la base de todo:

```go
type Handler interface {
    ServeHTTP(http.ResponseWriter, *http.Request)
}
```

No es una función.
Es un comportamiento: cualquier tipo que implemente `ServeHTTP` se convierte en un manejador HTTP.

#### Cuándo `HandlerFunc` es suficiente

Usa la función directa cuando:
- el código es simple
- no hay estado interno
- no hay composición de middleware

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

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/ping
```

#### Ejemplo con `ServeHTTP` (tipo personalizado)

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
	_, _ = w.Write([]byte(fmt.Sprintf("%s | ruta=%s", h.msg, r.URL.Path)))
}

func main() {
	h := helloHandler{msg: "manejador personalizado implementando ServeHTTP"}
	http.Handle("/hello", h)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/hello
```

#### Ejemplo con `http.Handler` para composición (struct + `finalHandler`)

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

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/api/v1/user
```

#### Misma composición, otra forma (middleware func + `finalHandler`)

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

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/api/v1/user
```


### 4.7 Apagado Elegante (Graceful Shutdown)

Al llamar a `log.Fatal(srv.ListenAndServe())`, el servidor se apaga inmediatamente al recibir una señal del SO (SIGINT, SIGTERM). Esto corta abruptamente las conexiones abiertas. `Shutdown` permite que el servidor finalice las solicitudes en curso antes de cerrarse.

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

	// Canal que recibe las señales del SO
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Inicia el servidor en una goroutine separada
	go func() {
		fmt.Println("Servidor escuchando en :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	// Bloquea hasta recibir la señal
	<-quit
	fmt.Println("Apagando el servidor...")

	// Contexto con tiempo de espera para el apagado
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Servidor forzado a apagarse: %v", err)
	}

	fmt.Println("Servidor finalizado limpiamente")
}
```

Ejecutar:

```bash
go run main.go
# En otra terminal:
curl -i localhost:8080/ping
# Para salir con apagado elegante:
Ctrl+C
```

Puntos importantes:
- `signal.Notify` captura `SIGINT` (Ctrl+C) y `SIGTERM` (usado por Docker/Kubernetes)
- `srv.ListenAndServe()` devuelve `http.ErrServerClosed` cuando se llama a `Shutdown`; esto es lo esperado, no un error
- `context.WithTimeout(10s)` garantiza que el servidor no espere indefinidamente
- En Kubernetes, se envía `SIGTERM` antes de retirar el pod del equilibrador de carga; el apagado elegante da tiempo para que las solicitudes en curso finalicen

### 4.8 Probando manejadores con `httptest`

El paquete `net/http/httptest` permite probar manejadores sin iniciar un servidor real en un puerto. Es el estándar de la comunidad de Go para pruebas de API.

```go
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Manejador que queremos probar
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
			name:        "éxito",
			body:        `{"name":"Jeff","email":"jeff@email.com"}`,
			contentType: "application/json",
			wantStatus:  http.StatusCreated,
		},
		{
			name:        "content-type incorrecto",
			body:        `{"name":"Jeff","email":"jeff@email.com"}`,
			contentType: "text/plain",
			wantStatus:  http.StatusUnsupportedMediaType,
		},
		{
			name:        "campos faltantes",
			body:        `{"name":"Jeff"}`,
			contentType: "application/json",
			wantStatus:  http.StatusUnprocessableEntity,
		},
		{
			name:        "json inválido",
			body:        `{inválido}`,
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
				t.Errorf("estado = %d, esperado %d", rec.Code, tt.wantStatus)
			}
		})
	}
}
```

Ejecutar:

```bash
go test -v ./...
```

Componentes de `httptest`:
- `httptest.NewRequest(method, target, body)`: crea un `*http.Request` para pruebas sin conexión real
- `httptest.NewRecorder()`: implementa `http.ResponseWriter` y captura el estado, las cabeceras y el cuerpo
- `rec.Code`: código de estado escrito por el manejador
- `rec.Body.String()`: cuerpo de la respuesta como cadena
- `rec.Header()`: cabeceras de la respuesta

### 4.9 Dos Servidores en Puertos Distintos con Goroutines

Usa un `ServeMux` por puerto e inicia cada servidor en su propia goroutine.
Para todas las variaciones de parada/bloqueo, consulta:
- [`examples/10-dual-listenandserve-goroutines/README.md`](examples/10-dual-listenandserve-goroutines/README.md)
- [`examples/10-dual-listenandserve-goroutines/README_pt.md`](examples/10-dual-listenandserve-goroutines/README_pt.md)

Endpoints:
- `:8080` -> `GET /api/v1/user`, `POST /api/v1/user`
- `:3000` -> `GET /api/v1/mock/user`, `POST /api/v1/mock/user`

Ejemplo práctico en el repositorio:
- [`examples/10-dual-listenandserve-goroutines/main.go`](examples/10-dual-listenandserve-goroutines/main.go)
- [`examples/10-dual-listenandserve-goroutines/main_test.go`](examples/10-dual-listenandserve-goroutines/main_test.go)
- [`examples/10-dual-listenandserve-goroutines/README.md`](examples/10-dual-listenandserve-goroutines/README.md) (todos los enlaces a escenarios)
- [`examples/10-dual-listenandserve-goroutines/README_pt.md`](examples/10-dual-listenandserve-goroutines/README_pt.md) (versión en portugués)
- [`examples/10-1-dual-listenandserve-basic/main.go`](examples/10-1-dual-listenandserve-basic/main.go) (versión compacta 10.1)
- [`examples/10-1-dual-listenandserve-basic/main_test.go`](examples/10-1-dual-listenandserve-basic/main_test.go)

```go
package main

import (
	"log"
	"net/http"
)

// newMainMux registra rutas para el puerto 8080.
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

// newMockMux registra rutas para el puerto 3000.
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

// startServer inicia un servidor HTTP en una goroutine.
func startServer(srv *http.Server) {
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server %s error: %v", srv.Addr, err)
		}
	}()
}

// main inicia ambos servidores y bloquea el proceso.
func main() {
	mainServer := &http.Server{Addr: ":8080", Handler: newMainMux()}
	mockServer := &http.Server{Addr: ":3000", Handler: newMockMux()}

	startServer(mainServer)
	startServer(mockServer)
	select {} // 01) Simple, directo, menos controlable.
	// done := make(chan struct{}); <-done // 02) Canal de bloqueo explícito.
	// var wg sync.WaitGroup; wg.Add(2); ...; wg.Wait() // 03) WaitGroup.
	// stop := make(chan os.Signal, 1); signal.Notify(stop, os.Interrupt, syscall.SIGTERM); <-stop // 04) Canal de señal.
	// time.Sleep(10 * time.Minute) // 05) Ventana de sueño fija.
	// ctx, cancel := context.WithCancel(context.Background()); <-ctx.Done(); cancel() // 06) Cancelación de contexto.
	// for { time.Sleep(time.Hour) } // 07) Bucle infinito.
	// errCh := make(chan error, 2); if err := <-errCh; err != nil { ... } // 08) Grupo de canal de error.
	// ctx, cancel := context.WithCancel(context.Background()); select { case err := <-errCh: cancel() } // 09) Canal de error + contexto.
	// var mu sync.Mutex; mu.Lock(); mu.Lock() // 10) Interbloqueo de Mutex (didáctico).
	// runtime.Goexit() // 11) Finaliza solo la goroutine principal.
	// var mu sync.Mutex; cond := sync.NewCond(&mu); mu.Lock(); cond.Wait() // 12) Espera sync.Cond.
	// ch := make(chan struct{}); for range ch {} // 13) Bloqueo de rango de canal.
	// fmt.Scanln() // 14) Bloqueo en stdin.
	// go runMock(); runMain() // 15) Un servidor bloqueante + una goroutine.
	// ticker := time.NewTicker(time.Hour); defer ticker.Stop(); for range ticker.C {} // 16) Ticker infinito.
	// stop := make(chan struct{}); select { case <-stop: case err := <-errCh: _ = err } // 17) Select con múltiples canales.
}
```

Ejecutar:

```bash
go run ./examples/10-dual-listenandserve-goroutines
curl -i localhost:8080/api/v1/user
curl -i -X POST localhost:8080/api/v1/user
curl -i localhost:3000/api/v1/mock/user
curl -i -X POST localhost:3000/api/v1/mock/user
```


---

## 5. Servidor API

En los ejemplos de la sección 4, cada manejador escribía su respuesta directamente con `w.Header().Set(...)`, `w.WriteHeader(...)` y `w.Write(...)`. Funciona, pero a medida que la API crece esto se vuelve repetitivo e inconsistente: un manejador devuelve `{"error":"..."}`, otro devuelve `{"message":"..."}`, un tercero olvida el `Content-Type`.

La sección 5 resuelve esto con un patrón de respuesta centralizado (`writeJSON` y `writeError`), organización de rutas, validación de entrada, endpoints de salud y autenticación básica. Es la evolución natural de lo construido en la sección 4.

**Enfoque de este bloque:**
- solo lado del servidor
- ejemplos pequeños para copiar, pegar y evolucionar
- estandarizar la API antes de pasar a una estructura mayor

**Ítems seleccionados:**

| Ítem |
|---|
| 5.0 helpers.go: funciones compartidas para la sección |
| 5.1 Estandarización de Respuestas |
| 5.2 Mapa de errores y estados por escenario |
| 5.3 Organización de Rutas |
| 5.4 Validación de entrada del servidor |
| 5.5 Endpoints de Salud (Health) |
| 5.6 Middleware de Auth Básica |
| 5.7 Variables de Entorno |
| 5.8 Middleware de CORS |
| 5.9 Documentación con OpenAPI/Swagger |

### 5.0 helpers.go: funciones compartidas

Los siguientes ejemplos reutilizan `writeJSON` y `writeError`. En lugar de volver a declararlos en cada archivo, extráelos a un `helpers.go` en el mismo paquete:

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

A partir de aquí, todos los ejemplos asumen que `writeJSON` y `writeError` están disponibles vía `helpers.go` en el mismo paquete.

### 5.1 Estandarización de Respuestas

**Objetivo:**
- centralizar la escritura de respuestas en un solo punto
- evitar la repetición de `Header`, `WriteHeader`, `Write`
- mantener un formato consistente de éxito y error

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

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/ok
curl -i localhost:8080/bad
```

### 5.2 Mapa de errores y estados por escenario

| Escenario | Estado | Cuándo usar |
|---|---:|---|
| JSON inválido | `400` | Cuerpo malformado |
| Campo inválido / faltante | `422` | Cuerpo válido, pero regla de negocio inválida |
| Content-Type incorrecto | `415` | Se esperaba `application/json` |
| Recurso no encontrado | `404` | El ID/ruta no existe |
| Método no permitido | `405` | El endpoint existe, el método no |
| Conflicto de estado | `409` | Ej.: correo ya registrado |
| Error interno | `500` | Fallo inesperado del servidor |

Ejemplo de retorno de error estandarizado:

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

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/error/bad-json
curl -i localhost:8080/error/validation
curl -i localhost:8080/error/not-found
```

### 5.3 Organización de Rutas

Objetivo:
- mantener las rutas predecibles
- separar por contexto (`health`, `users`, etc.)
- versionar la API (`/api/v1`)

```go
package main

import (
	"log"
	"net/http"
)

func registerRoutes(mux *http.ServeMux) {
	// Salud (Health)
	mux.HandleFunc("GET /healthz", healthz)
	mux.HandleFunc("GET /readyz", readyz)
	mux.HandleFunc("GET /livez", livez)

	// API v1 - usuarios
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

Consejo:
- si usas `{id}` en el patrón (Go 1.22+), lee con `r.PathValue("id")`

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/healthz
curl -i localhost:8080/api/v1/users/123
```

### 5.4 Validación de Entrada del Servidor

Lista de verificación corta para `POST/PUT`:
1. validar `Content-Type`
2. limitar el tamaño del cuerpo
3. decodificar JSON con `DisallowUnknownFields`
4. validar campos obligatorios
5. devolver el estado correcto (`400`, `415`, `422`)

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

Ejecutar:

```bash
go run main.go
curl -i -X POST localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Jeff","email":"jeff@email.com"}'
curl -i -X POST localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Jeff","email":"jeff@email.com","extra":"x"}'
```

### 5.5 Endpoints de Salud (Health)

![Health Endpoints](docs/diagrams/health-endpoints.svg)


Patrón simple:
- `GET /healthz`: el servidor responde (arriba)
- `GET /livez`: el proceso está vivo
- `GET /readyz`: listo para recibir tráfico (dependencias OK)

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

Ejecutar:

```bash
go run main.go
curl -i localhost:8080/healthz
curl -i localhost:8080/readyz
curl -i localhost:8080/livez
```

### 5.6 Middleware de Auth Básica

Cuándo usar (simple y educativo):
- proteger un endpoint interno de laboratorio/homologación
- probar la autenticación HTTP básica antes de adoptar JWT/OAuth2

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
	// Lee las credenciales de las variables de entorno
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

Ejecutar:

```bash
API_USER=admin API_PASS=s3cr3t go run main.go
curl -i localhost:8080/api/v1/user
curl -i -u admin:s3cr3t localhost:8080/api/v1/user
```

### 5.7 Variables de Entorno

En producción, las configuraciones sensibles (credenciales, puertos, URLs de bases de datos) nunca deben estar codificadas en el código (hardcoded). Go ofrece `os.Getenv` y `os.LookupEnv` para esto.

| Función | Comportamiento |
|---|---|
| `os.Getenv("KEY")` | Devuelve el valor o `""` si no existe |
| `os.LookupEnv("KEY")` | Devuelve valor + bool `ok`, distingue "no definida" de "definida vacía" |

```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Puerto con fallback
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Variable obligatoria: se detiene si no está definida
	dbURL, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Fatal("DATABASE_URL no definida")
	}

	// Variable opcional con fallback
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	fmt.Printf("Iniciando en :%s (env=%s db=%s)\n", port, env, dbURL)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("pong"))
	})

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
```

Ejecutar:

```bash
PORT=9090 APP_ENV=production DATABASE_URL=postgres://localhost/mydb go run main.go
curl -i localhost:9090/ping
```

En Docker:

```dockerfile
ENV PORT=8080
ENV APP_ENV=production
```

O vía `docker run`:

```bash
docker run -e PORT=9090 -e DATABASE_URL=postgres://... -p 9090:9090 nethttp-server:local
```

Mejores prácticas:
- usa `LookupEnv` para variables obligatorias: el servidor falla inmediatamente si está mal configurado
- usa `Getenv` con fallback para variables opcionales con un valor predeterminado sensato
- nunca registres en el log el valor de variables sensibles (contraseñas, tokens)
- en desarrollo, usa un archivo `.env` con una librería como `godotenv`; nunca hagas commit de este archivo


### 5.8 Middleware de CORS

CORS (Cross-Origin Resource Sharing) es un mecanismo de seguridad del navegador que bloquea las solicitudes realizadas desde un origen distinto al de la API. Toda API consumida por un frontend web necesita gestionar esto.

**Cómo funciona:**

1. El navegador envía una solicitud `OPTIONS` (preflight) antes de un `POST/PUT/DELETE` con cabeceras personalizadas
2. El servidor responde con cabeceras `Access-Control-*` indicando qué está permitido
3. El navegador libera (o bloquea) la solicitud real basándose en esa respuesta

```go
package main

import (
	"log"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Permite cualquier origen (restringir en producción)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Trace-Id")

		// Responde al preflight y sale
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

	// CORS aplicado a toda la API
	finalHandler := corsMiddleware(mux)
	log.Fatal(http.ListenAndServe(":8080", finalHandler))
}
```

Ejecutar:

```bash
go run main.go

# Simula el preflight del navegador
curl -i -X OPTIONS localhost:8080/api/v1/users \
  -H "Origin: http://localhost:3000" \
  -H "Access-Control-Request-Method: POST" \
  -H "Access-Control-Request-Headers: Content-Type"

# Solicitud normal
curl -i localhost:8080/api/v1/users -H "Origin: http://localhost:3000"
```

| Cabecera | Qué controla |
|---|---|
| `Access-Control-Allow-Origin` | Qué orígenes pueden acceder (`*` = cualquiera, o `https://misitio.com`) |
| `Access-Control-Allow-Methods` | Qué verbos HTTP están permitidos |
| `Access-Control-Allow-Headers` | Qué cabeceras puede enviar el cliente |
| `Access-Control-Max-Age` | Cuánto tiempo puede el navegador cachear el preflight (segundos) |

**En producción:** sustituye `"*"` por el origen real de tu frontend. Usar `"*"` con credenciales (`Authorization`) no funciona, ya que el navegador exige un origen explícito en ese caso.

### 5.9 Documentación con OpenAPI/Swagger

#### Un poco de historia

En 2010, Tony Tam creó Swagger mientras trabajaba en Wordnik para resolver un problema simple: cómo describir una API REST para que humanos y máquinas la entendieran. La solución fue un archivo de especificación en JSON/YAML que describía endpoints, parámetros, esquemas y respuestas.

En 2015, la especificación fue donada a la OpenAPI Initiative (parte de la Linux Foundation) y pasó a llamarse OpenAPI Specification. El nombre Swagger permaneció popular, pero técnicamente se refiere a las herramientas (Swagger UI, Swagger Editor), no a la especificación.

| Versión | Año | Aspectos destacados |
|---|---:|---|
| Swagger 1.x | 2010 | Origen en Wordnik, formato propietario |
| Swagger 2.0 | 2014 | Gran estandarización, adopción masiva |
| OpenAPI 3.0 | 2017 | Reestructurado, soporte para webhooks y enlaces |
| OpenAPI 3.1 | 2021 | Alineación total con JSON Schema |

Hoy en día, OpenAPI 3.1 es el estándar. Cuando alguien dice "genera el Swagger", normalmente se refiere a "generar el archivo OpenAPI".

**Por qué es importante:**
- otros equipos pueden consumir tu API sin pedir explicaciones
- se pueden generar clientes SDK automáticamente
- sirve como contrato vivo entre frontend y backend
- facilita las pruebas manuales vía interfaz web

**Con `swaggo/swag` en Go:**

```bash
# Instala el generador
go install github.com/swaggo/swag/cmd/swag@latest

# Añade dependencias al proyecto
go get github.com/swaggo/http-swagger
go get github.com/swaggo/swag
```

Anota los manejadores con comentarios especiales:

```go
package main

import (
	"log"
	"net/http"

	_ "myproject/docs" // generado por swag
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           User API
// @version         1.0
// @description     Ejemplo de API documentada con Swagger
// @host            localhost:8080
// @BasePath        /api/v1

// getUsers godoc
// @Summary      Listar usuarios
// @Description  Devuelve todos los usuarios registrados
// @Tags         usuarios
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

	// Sirve la interfaz Swagger UI en /swagger/
	mux.Handle("GET /swagger/", httpSwagger.WrapHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Generar documentación:

```bash
# Genera la carpeta docs/ con el JSON/YAML de OpenAPI
swag init

# Inicia el servidor
go run main.go

# Accede a la interfaz visual
open http://localhost:8080/swagger/index.html
```

**Flujo de trabajo recomendado:**

| Etapa | Acción |
|---|---|
| 1 | Escribe el manejador con anotaciones `// @...` |
| 2 | Ejecuta `swag init` para regenerar `docs/` |
| 3 | Hace commit de `docs/` junto con el código |
| 4 | El CI/CD puede validar que `docs/` esté actualizado |

**Alternativas a swaggo:**

| Herramienta | Enfoque | ¿Ensucia el código? |
|---|---|:---:|
| `swaggo/swag` | Anotaciones en comentarios de Go, genera OpenAPI 2.0/3.0 | Sí |
| `deepmap/oapi-codegen` | Contract-first: genera código Go desde YAML de OpenAPI | No |
| `huma` | Framework que genera OpenAPI automáticamente vía tipos de Go | No |
| Manual YAML + manejador estático | Tú escribes el contrato, Go sirve el archivo | No |
| Postman → OpenAPI | Exporta la colección de Postman y la convierte a YAML | No |

---

#### Alternativa 1: YAML Manual + Swagger UI sin dependencia

El enfoque más limpio. Escribes el `openapi.yaml` a mano (o vía editor online), haces commit en el repositorio y lo sirves con un solo manejador de Go. Cero anotaciones en el código, cero dependencias extra.

**Paso 1: Escribe el contrato en el Swagger Editor online:**

Accede a [editor.swagger.io](https://editor.swagger.io), escribe o pega tu YAML y valida en tiempo real. Cuando estés listo, descarga el archivo.

```yaml
# openapi.yaml: guarda en la raíz o en docs/openapi.yaml
openapi: "3.0.3"
info:
  title: User API
  version: "1.0.0"
paths:
  /api/v1/users:
    get:
      summary: Listar usuarios
      responses:
        "200":
          description: Lista de usuarios
          content:
            application/json:
              schema:
                type: array
                items:
                  $ref: "#/components/schemas/User"
  /api/v1/users/{id}:
    get:
      summary: Buscar usuario por ID
      parameters:
        - name: id
          in: path
          required: true
          schema:
            type: string
      responses:
        "200":
          description: Usuario encontrado
        "404":
          description: No encontrado
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

**Paso 2: Sirve el YAML y la UI vía Go (cero dependencias):**

```go
package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Rutas de la API
	mux.HandleFunc("GET /api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`[{"name":"Jeff","email":"jeff@email.com"}]`))
	})

	// Sirve el archivo openapi.yaml directamente
	mux.HandleFunc("GET /openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		http.ServeFile(w, r, "docs/openapi.yaml")
	})

	// Sirve Swagger UI vía CDN (HTML puro, sin dependencia de Go)
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

	log.Println("Docs en http://localhost:8080/docs")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

Ejecutar:

```bash
go run main.go
open http://localhost:8080/docs
```

O sirve vía Redoc (visual más limpio para documentación pública):

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

#### Alternativa 2: Colección de Postman a OpenAPI YAML

Si ya tienes una colección de Postman, puedes convertirla directamente a OpenAPI sin escribir el YAML desde cero.

**Opción A: A través del propio Postman (interfaz):**

1. Abre la colección en Postman
2. Haz clic en los tres puntos de la colección
3. Selecciona `Export` → elige formato `Collection v2.1`
4. Guarda como `collection.json`

**Opción B: Convertir con script de Node (postman-to-openapi):**

```bash
# Instala la herramienta
npm install -g postman-to-openapi

# Convierte collection.json a openapi.yaml
p2o collection.json -f docs/openapi.yaml
```

**Opción C: Convertir con script de Python:**

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
                # normaliza parámetros de ruta {param} -> {param}
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

    print(f"OpenAPI generado en {output_path}")

if __name__ == "__main__":
    postman_to_openapi(sys.argv[1], sys.argv[2])
```

Ejecutar:

```bash
python3 convert_postman.py collection.json docs/openapi.yaml
```

Después sirve el `openapi.yaml` con el manejador de la Alternativa 1. El resultado es una documentación visual completa sin ninguna anotación en tu código Go.

---

#### Alternativa 3: Contract-first con oapi-codegen

Escribe primero el YAML, genera automáticamente las interfaces y tipos de Go. El código queda limpio porque la especificación vive fuera del código.

```bash
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
oapi-codegen -package api docs/openapi.yaml > api/api.gen.go
```

Go genera las interfaces que tú implementas, sin ninguna anotación y sin ningún comentario especial.

---

**Qué enfoque elegir:**

| Situación | Recomendación |
|---|---|
| El equipo ya usa Postman | Exportar colección y convertir con `postman-to-openapi` o script de Python |
| Quiere cero dependencia de Go | YAML manual + `http.ServeFile` + Swagger UI vía CDN |
| Quiere contrato riguroso y generación de código | `oapi-codegen` (contract-first) |
| Quiere iterar rápido con anotaciones | `swaggo/swag` (acepta la suciedad como compromiso) |
| Documentación pública bonita | Redoc vía CDN sobre el mismo `openapi.yaml` |

---

## 6. Docker: Construir y Ejecutar Local

Objetivo:
- compilar el servidor Go en una imagen ligera (multi-etapa)
- ejecutar localmente en el puerto `8080`
- tener comandos básicos de operación y depuración

Prerrequisito:
- tener un `main.go` funcional en la raíz (puedes usar cualquier ejemplo de este README)
- usar `.dockerignore` para no enviar archivos innecesarios al contexto de construcción

### 6.1 Dockerfile multi-etapa (Alpine + zona horaria de Brasil)

Archivo `Dockerfile` (en la raíz del proyecto):

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

Notas:
- la imagen final es ligera (no tiene el toolchain de Go)
- zona horaria configurada para Brasil (`America/Sao_Paulo`)
- si tu entrypoint es otro archivo, usa `--build-arg APP_FILE=tu_archivo.go`
- la imagen builder debe ser compatible con la versión de `go.mod` (ej.: `go 1.25.6`)

### 6.2 Comandos Básicos de Docker

Construcción de la imagen:

```bash
docker build -t nethttp-server:local .
```

Construcción de la imagen sin usar caché (fuerza la reconstrucción):

```bash
docker build --no-cache -t nethttp-server:local .
```

Construcción eligiendo otro archivo Go:

```bash
docker build -t nethttp-server:local --build-arg APP_FILE=cmd/api/main.go .
```

Limpiar caché de construcción (caché del builder):

```bash
docker builder prune -f
```

Ejecutar contenedor local:

```bash
docker run -d --name nethttp-server -p 8080:8080 nethttp-server:local
```

Listar contenedores:

```bash
docker ps -a
```

Ver logs:

```bash
docker logs -f nethttp-server
```

Validar rutas:

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
  -d '{"name":"Jeff Actualizado","email":"jeff.updated@email.com"}'
```

Comprobar zona horaria en el contenedor:

```bash
docker exec -it nethttp-server date
```

Detener y eliminar contenedor:

```bash
docker stop nethttp-server
docker rm nethttp-server
```

---

<br>

<sub>Hecho con dedicación por <a href="https://github.com/jeffotoni">@jeffotoni</a> · Go es el lenguaje, enseñar es la misión.</sub>
