# SSL Labs TLS Checker (Nebula Challenge)

Este proyecto es un programa en **Go** que verifica la seguridad **TLS** de un dominio utilizando la **API de SSL Labs**.

Fue desarrollado como parte de una prueba técnica y se enfoca en:
- Invocar la API de evaluación de SSL Labs
- Consultar periódicamente el estado del análisis hasta su finalización
- Procesar y mostrar información relevante sobre la seguridad TLS del dominio

---

## Funcionalidades

- Inicia un análisis de seguridad TLS para un dominio dado
- Consulta el progreso del análisis hasta que finaliza
- Muestra resultados por endpoint (dirección IP)
- Presenta:
  - Calificación TLS (A+, A, etc.)
  - Protocolos TLS soportados
  - Vulnerabilidades TLS conocidas:
    - Heartbleed
    - POODLE
    - FREAK
    - LOGJAM

---

## Requisitos

- Go **1.18 o superior**
- Conexión a Internet (la API de SSL Labs es externa)

---

## Uso

Clona el repositorio:

```bash
git clone https://github.com/<tu-usuario>/ssllabs-tls-checker.git
cd ssllabs-tls-checker
```
Ejecuta el programa indicando el dominio a analizar:

```bash
go run . example.com
```

Ejemplo

```bash
go run . truora.com
```

Salida de ejemplo:

```yaml
Analizando truora.com...
Status: DNS
Status: IN_PROGRESS
Status: READY

Resultados para truora.com
----------------------------
IP: 199.60.103.31
Estado: Ready
Grade: A+
Protocolos soportados:
  - TLS 1.2
Vulnerabilidades:
  Heartbleed: false
  Poodle: false
  Freak: false
  Logjam: false
```

## ¿Cómo funciona?

1. El programa llama al endpoint analyze de SSL Labs con el parámetro startNew=on

2. SSL Labs ejecuta un análisis completo de seguridad TLS (puede tardar varios minutos)

3. El programa consulta periódicamente el estado del análisis

4. Cuando el estado es READY, se procesan y muestran los resultados
