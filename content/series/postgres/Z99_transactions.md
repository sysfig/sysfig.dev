---
title: Transactions
slug: transactions
date: 2022-04-30T12:11:00+01:00
chapter: z
order: 99
tags:
  - postgres
draft: true
---

A transaction block is an ordered set of SQL commands that are grouped into a single operation. This means if the database crashes whilst executing one of the SQL commands in the transaction, then the whole transaction is deemed to have failed and none of the SQL commands within it (even those that have already been executed) takes effect. Because of this all-or-nothing property, transactions are said to be _atomic_.

In PostgreSQL, you can define a transaction block by enclosing the set of SQL commands between the `BEGIN` and `COMMIT` commands.

```sql
BEGIN;
UPDATE accounts SET balance = balance - 100.00
    WHERE name = 'Alice';
-- etc etc
COMMIT;
```

## Savepoints


`SAVEPOINT`
`ROLLBACK TO`
