# GTD app

## Functional Requirements

### The five important use cases

The users of the app:
- can capture stuff into Inbox
- can clarify captured stuff quickly
- can organize clarified items into the right places:
  - Trash
  - Someday/maybe
  - Reference
  - Project
  - Waiting for
  - Calendar
  - Next action
- can reflect their actual situation to the app by performing Daily, Weekly, and Monthly reviews
- can engage in what they should do using:
  - "Next action" list
  - Calendar
  - Levels of Focus

### Other functions

The user of the app
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

Link to Miro board: https://miro.com/app/board/uXjVPYQe-Ls=/?share_link_id=17090475310
