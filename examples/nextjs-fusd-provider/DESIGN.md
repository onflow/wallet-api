# Wallet API Design

## Current design

### Create an account

#### Asynchronous

1. POST `/accounts` -> `Job`
2. GET `/jobs/${Job.id}` -> `(status, address)`
3. Repeat until `status == "Complete"`
4. GET `/accounts/${address}` -> `Account`

#### Synchronous

1. POST `/accounts` -> `Account`

### Send a transaction

#### Asynchronous

1. POST `/transactions` -> `Job`
2. GET `/jobs/${Job.id}` -> `(status, txID)`
3. Repeat until `status == "Complete"`
4. GET `/transactions/${txID}` -> `Transaction`

#### Synchronous

1. POST `/transactions` -> `Transaction`

### Models

#### Job

```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000",
  "result": "ADDRESS|TXID",
  "status": "PENDING",
  "createdAt": "10-06-26 02:31:29",
  "updatedAt": "10-06-26 02:31:29"
}
```

## Proposed design

### Create an account

#### Asynchronous

1. POST `/accounts` -> `Account`
2. GET `/accounts/${UUID}` -> `Account`
3. Repeat until `account.status == "SUCCESS|ERROR"`

#### Synchronous

1. POST `/accounts` -> `Account`

### Send a transaction

#### Asynchronous

1. POST `/transactions` -> `Transaction`
2. GET `/transactions/${UUID}` -> `Transaction`
3. Repeat until `tx.status == "SUCCESS|ERROR"`

#### Synchronous

1. POST `/transactions` -> `Transaction`

### Models

#### Status

- `IN_QUEUE`
- `SIGNED`
- `SUBMITTED`
- `SUCCESS`
- `ERROR`

#### Account

```json
{
  "uuid": "123e4567-e89b-12d3-a456-426614174000",
  "address": "0xa7b30d9833758f1a",
  "status": "PENDING",
  "createdAt": "10-06-26 02:31:29",
  "updatedAt": "10-06-26 02:31:29"
}
```

#### Transaction

```json
{
  "uuid": "123e4567-e89b-12d3-a456-426614174000",
  "transactionId": "83c78e25df65f2003173afc56d3f57866e7240508ee02146f0d6ca584b5579f7",
  "error": null,
  "events": [],
  "status": "SUCCESS",
  "createdAt": "10-06-26 02:31:29",
  "updatedAt": "10-06-26 02:31:29"
}
```
