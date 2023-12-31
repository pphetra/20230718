# Dependency

```
┌────────────────────────────────────────────────────────────────────────┐  ┌────────────────────────────────────────────────────────────────────────┐
│                                 member                                 │  │                                 infra                                  │
│                                                                        │  │                                                                        │
│                                                                        │  │                                                                        │
│                                                                        │  │                                                                        │
│       ┌───────────────────────────────────────────────────────┐        │  │                                                                        │
│       │                         infra                         │        │  │                                                                        │
│       │                                                       │        │  │                     ┌────────────┐                                     │
│       │                                                       │        │  │                     │     DB     │                                     │
│       │   ┌─────────────────────────────────┐                 │◀───────┼──┼─────────────────────│ UnitOfWork │                                     │
│       │   │    PostgresMemberRepository     │                 │        │  │                     └────────────┘                                     │
│       │   └─────────────────────────────────┘                 │        │  │                            │                                           │
│       │                                                       │        │  │                            │                                           │
│       └───────────────────────────────────────────────────────┘        │  │                            │                                           │
│                                                                        │  └────────────────────────────┼───────────────────────────────────────────┘
│                    │                              │                    │  ┌────────────────────────────┼───────────────────────────────────────────┐
│                    │                              │                    │  │                            │    shared                                 │
│                    ▼                              │                    │  │                            │                                           │
│       ┌───────────────────────────────────────────┼───────────┐        │  │         ┌──────────────────┼────────────────────────────────────┐      │
│       │                          app              │           │        │  │         │                  │       app                          │      │
│       │                                           │           │        │  │         │                  ▼                                    │      │
│       │  ┌─────────────────────────────────┐      │           │        │  │         │           ┌────────────┐                              │      │
│       │  │ RegisterIndividualMemberUsecase │      │           │────────┼──┼────────▶│           │ UnitOfWork │                              │      │
│       │  └─────────────────────────────────┘      │           │        │  │         │           │ Interface  │                              │      │
│       │                                           │           │        │  │         │           └────────────┘                              │      │
│       │                                           │           │        │  │         │                                                       │      │
│       └───────────────────────────────────────────┼───────────┘        │  │         └───────────────────────────────────────────────────────┘      │
│                    │                              │                    │  │                                                                        │
│                    │                              │                    │  │                                                                        │
│                    ▼                              ▼                    │  │                                                                        │
│       ┌───────────────────────────────────────────────────────┐        │  │         ┌───────────────────────────────────────────────────────┐      │
│       │                        domain                         │        │  │         │                     value_object                      │      │
│       │                                                       │        │  │         │                                                       │      │
│       │         ┌──────┐    ┌───────────────────┐             │        │  │         │           ┌────────┐                                  │      │
│       │         │Member│    │ MemberRepository  │             ├────────┼──┼────────▶│           │Address │                                  │      │
│       │         └──────┘    └───────────────────┘             │        │  │         │           └────────┘                                  │      │
│       │                                                       │        │  │         │                                                       │      │
│       │                                                       │        │  │         │                                                       │      │
│       └───────────────────────────────────────────────────────┘        │  │         └───────────────────────────────────────────────────────┘      │
│                                                                        │  │                                                                        │
│                                                                        │  │                                                                        │
│                                                                        │  │                                                                        │
└────────────────────────────────────────────────────────────────────────┘  └────────────────────────────────────────────────────────────────────────┘

```

v2
```
+------------------------------------------------------------------------+  +------------------------------------------------------------------------+                       
|                                 member                                 |  |                                 infra                                  |                       
|                                                                        |  |                                                                        |                       
|                                                                        |  |                                                                        |                       
|                                                                        |  |                                                                        |                       
|                                                                        |  |                                                                        |                       
|       +-------------------------------------------------------+        |  |                                                                        |                       
|       |                         infra                         |        |  |                                                                        |                       
|       |   +---------------------------------+                 |        |  |                    +------------+                                      |         +------------+
|       |   |    PostgresMemberRepository     |                 |        |  |                    |     DB     |                                      |         |            |
|       |   +---------------------------------+                 |◀-------+--+------------------- | UnitOfWork |--------------------------------------+--------▶|   sql.DB   |
|       |   +---------------------------------+                 |        |  |                    +------------+                                      |         |            |
|       |   |        GoChannelEventBus        |                 |        |  |                           |                                            |         +------------+
|       |   +---------------------------------+                 |        |  |                           |                                            |                       
|       +-------------------------------------------------------+        |  +---------------------------+--------------------------------------------+                       
|                    |                              |                    |  +---------------------------+--------------------------------------------+                       
|                    |                              |                    |  |                           ++    shared                                 |                       
|                    ▼                              |                    |  |                            |                                           |                       
|       +-------------------------------------------+-----------+        |  |         +------------------+------------------------------------+      |                       
|       |                          app              |           |        |  |         |                  |       app                          |      |                       
|       |                                           |           |        |  |         |                  ▼                                    |      |                       
|       |  +---------------------------------+      |           |        |  |         |           +------------+    +----------+              |      |                       
|       |  | RegisterIndividualMemberUsecase |      |           |--------+--+--------▶|           | UnitOfWork |    | EventBus |              |      |                       
|       |  +---------------------------------+      |           |        |  |         |           | Interface  |    |          |              |      |                       
|       |                                           |           |        |  |         |           +------------+    +----------+              |      |                       
|       |                                           |           |        |  |         |                                                       |      |                       
|       +-------------------------------------------+-----------+        |  |         +-------------------------------------------------------+      |                       
|                    |                              |                    |  |                                                                        |                       
|                    |                              |                    |  |                                                                        |                       
|                    ▼                              ▼                    |  |                                                                        |                       
|       +-------------------------------------------------------+        |  |         +-------------------------------------------------------+      |                       
|       |                        domain                         |        |  |         |                     value_object                      |      |                       
|       |    +------+   +-------------------+                   |        |  |         |                                                       |      |                       
|       |    |Member|   | MemberRepository  |                   |        |  |         |           +--------+   +------------+                 |      |                       
|       |    +------+   +-------------------+                   +--------+--+--------▶|           |Address |   |DomainEvent |                 |      |                       
|       |               +--------------------------------+      |        |  |         |           +--------+   +------------+                 |      |                       
|       |               |IndividualMemberRegisteredEvent |      |        |  |         |                                                       |      |                       
|       |               +--------------------------------+      |        |  |         |                                                       |      |                       
|       +-------------------------------------------------------+        |  |         +-------------------------------------------------------+      |                       
|                                                                        |  |                                                                        |                       
|                                                                        |  |                                                                        |                       
|                                                                        |  |                                                                        |                       
+------------------------------------------------------------------------+  +------------------------------------------------------------------------+                       
```