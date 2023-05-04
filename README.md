# GTD app

## Functional Requirements

### The five important use cases

The users of the app:
1. can capture stuff into Inbox
2. can clarify captured stuff quickly
3. can organize clarified items into the right places:
  - Trash
  - Someday/maybe
  - Reference
  - Project
  - Waiting for
  - Calendar
  - Next action
4. can reflect their actual situation to the app by performing Daily, Weekly, and Monthly reviews
5. can engage in what they should do using:
  - "Next action" list
  - Calendar
  - Levels of Focus

### Other functions

The user of the app
- can use the app from web browsers and as a mobile app on various devices (desktop, tablet, and mobile)
- can log in to the app using Social Auth (e.g., Google, Facebook, Twitter)
- can integrate the app's calendar with their Google Calendar
- can use the app from web browsers on various devices (desktop, tablet, and mobile)

## Non-functional requirements

### Performance

- This app can handle XX simultaneous requests (provide a specific number)
- This app responds to each request within 0.5 seconds

### Security

- Access control: Users authenticated with social login can use all functions of the app
- Data encryption:
  - This app does not store any authentication information because authentication is provided via social login
  - Sensitive user data is encrypted at rest and in transit

### Data backup:

This app backs up every user's data daily.

### Scalability

This app can scale out components that restrict the overall throughput.

### Usability

- Supported browsers:
  - Google Chrome v112.0 or later
- Misc functions:
  - Undo/redo functionality
  - Sorting by specified columns in each list

## Technical Stack

- Cloud Provider: Oracle Cloud
- Application Runtime Platform: Kubernetes
- CI/CD: GitHub Actions
- Programming Languages (Frameworks):
  - Backend: Go
  - Frontend: TypeScript (React)
- Middlewares:
  - RabbitMQ
  - PosgreSQL
  - BackBalancer
- Networking:
  - Between backend services: gRPC
  - Between backend and frondend services: GraphQL

## Architecture

The architecture of this application is Microservice Architecture

```mermaid
graph TB
    User[User]-->Frontend
    Frontend[Web and Mobile Frontend]-->RouteMaster
    RouteMaster --> IDP
    IDP --> RouteMaster
    RouteMaster --> JWT
    JWT

    RouteMaster --> BackBalancer
    BackBalancer --> Checklist
    BackBalancer --> Clarify
    Clarify --> Checklist
    Clarify --> Inbox
    Clarify --> Reference
    Clarify --> Project
    Clarify --> Calendar
    Project --> Task
    BackBalancer --> Task
    BackBalancer --> Project
    BackBalancer --> Reference
    BackBalancer --> Review
    Review --> Focus
    BackBalancer --> Focus
    BackBalancer --> Inbox
    Project --> Focus

    Clarify --> ClarifyBroker
    ClarifyBroker --> Inbox
    ClarifyBroker --> Project
    ClarifyBroker --> Task

    classDef services fill:#2288f9,stroke:#333,stroke-width:1px;
    class RouteMaster,BackBalancer,IDP,Frontend,Checklist,Clarify,ClarifyBroker,Inbox,Reference,Project,Task,Review,Focus,Calendar services;
    classDef tools fill:#858585,stroke:#333,stroke-width:1px;
    class JWT tools
```


| Name of service | Responsibility |
| --- | --- |
| Inbox | CRUD of Inbox items |
| Clarify | Assignment of each Inbox item to single category (e.g.: 'Projects', 'References', 'Someday/maybe') |
| Checklist | CRUD of user-defined checklists |
| Reference | CRUD of reference items |
| Calendar | Integration with Google Calendar |
| Project | CRUD of projects and management of their child tasks; Collaboration management (e.g.: user invitation, access management) |
| Task | CRUD of tasks; Collaboration management (e.g.: task assignment) |
| Review | CRUD of reviews |
| Focus | CRUD of Levels of Focus and management of relation between each level and project |
| IDP | Management of signup, login and access control |

