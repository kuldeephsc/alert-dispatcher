# Dispatcher

A Go-based notification dispatcher service that sends alerts through multiple channels (Email, Slack, Webhook) with built-in retry logic.

## Architecture

### High-Level Design (HLD)

The system follows a clean architecture pattern with clear separation of concerns:

```
Client → Notification Service → Dispatcher → Channels (Email/Slack/Webhook)
```

**Components:**
- **Client**: Entry point that initiates notifications
- **Notification Service**: Builds alert objects and orchestrates the dispatch
- **Dispatcher**: Core orchestration logic - fans out alerts to channels and retries failed sends
- **Channels**: Individual implementations for different notification mediums

Each channel implements the same two-method interface (`Send()` and `Name()`), allowing the Dispatcher to remain agnostic of specific channel implementations. This enables adding new channels without modifying the Dispatcher code.

```mermaid
graph TD
    Client["Client"]
    NS["Notification Service<br/>builds the Alert"]
    D["Dispatcher<br/>fans out, retries once,<br/>collects results"]
    Email["Email Channel<br/>Send(alert)"]
    Slack["Slack Channel<br/>Send(alert)"]
    Webhook["Webhook Channel<br/>Send(alert)"]
    
    Client --> NS
    NS --> D
    D --> Email
    D --> Slack
    D --> Webhook
    
    classDef note fill:#f9f9f9,stroke:#999
    class Email,Slack,Webhook note
```

### Low-Level Design (LLD)

The detailed component architecture showing data flow and interface contracts:

```mermaid
graph TD
    Alert["Alert<br/>data contract, input"]
    D["Dispatcher<br/>fan-out + retry orchestration"]
    Result["Result<br/>status per channel, output"]
    CI["interface<br/>Channel"]
    Email["EmailChannel<br/>SMTP"]
    Slack["SlackChannel<br/>HTTP POST"]
    Webhook["WebhookChannel<br/>HTTP POST"]
    SMS["SMSChannel<br/>added later, no edits"]
    
    Alert --> D
    D --> Result
    D --> CI
    CI --> Email
    CI --> Slack
    CI --> Webhook
    CI --> SMS
    
    style Email fill:#c8e6c9
    style Slack fill:#c8e6c9
    style Webhook fill:#c8e6c9
    style SMS fill:#ffcccc
    style CI fill:#bbdefb
    style D fill:#e1bee7
```

**Key Points:**
- The Dispatcher depends on the `Channel` interface, never on concrete implementations
- Each channel implementation is independent and can be added/modified without affecting the Dispatcher
- New channels (e.g., SMSChannel) can be added by simply implementing the `Channel` interface
- The retry logic is centralized in the Dispatcher for consistency
- Every box in the bottom row implements Channel — Dispatcher's code never mentions any of them by name

## Project Structure

```
dispatcher/
├── main.go                 # Entry point
├── go.mod                  # Module definition
├── README.md              # This file
├── dispatcher/
│   └── dispatcher.go      # Core dispatcher logic
├── channel/
│   ├── channel.go         # Channel interface
│   ├── email.go           # Email channel implementation
│   ├── slack.go           # Slack channel implementation
│   └── webhook.go         # Webhook channel implementation
├── model/
│   └── alert.go           # Alert data model
└── service/
    └── notification.go    # Notification service
```

## How It Works

1. **Client** calls `NotificationService.Notify()` with title and message
2. **Notification Service** creates an `Alert` object and passes it to the Dispatcher
3. **Dispatcher** iterates through all registered channels and attempts to send the alert
4. **Channels** implement their specific send logic (SMTP for Email, HTTP POST for Slack/Webhook)
5. **Retry Logic**: If a channel fails, the Dispatcher retries once before marking as failed
6. **Results** are collected and returned as a map showing success/failure status for each channel

## Extensibility

To add a new channel:

1. Create a new file in the `channel/` package (e.g., `sms.go`)
2. Define a struct that implements the `Channel` interface:
   ```go
   type SMSChannel struct{}
   
   func (s *SMSChannel) Send(alert model.Alert) error {
       // Implementation
   }
   
   func (s *SMSChannel) Name() string {
       return "SMS"
   }
   ```
3. Register it in main: `d.Register(&channel.SMSChannel{})`

The Dispatcher code remains unchanged - this is the power of interface-based design.
