---
title: Royalty API v1.0
toc_footers: []
includes: []
search: true
highlight_theme: darkula
headingLevel: 2

---

<!-- Generator: Widdershins v4.0.1 -->

<h1 id="royalty-api">Royalty API v1.0</h1>

RESTFul API to distribute and redeem the voucher

Base URLs:

* <a href="localhost:8080/api">localhost:8080/api</a>

Email: <a href="mailto:jayapermanarama@gmail.com">Rama Jayapermana</a> 
License: <a href="http://www.apache.org/licenses/LICENSE-2.0.html">Apache 2.0</a>

<h1 id="royalty-api-transaction">Transaction</h1>

## post__transaction

`POST /transaction`

*Create Transaction for each User*

Create Transaction for each User

> Body parameter

```json
{
  "id": 0,
  "invoice_number": "string",
  "tenant_id": 0,
  "total_amount": 0,
  "transaction_date": "string",
  "user_id": 0,
  "voucherGenerated": true
}
```

<h3 id="post__transaction-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[domain.TransactionRequest](#schemadomain.transactionrequest)|true|Transaction information|

> Example responses

> 200 Response

```json
{
  "data": null,
  "message": "string",
  "status_code": 0
}
```

<h3 id="post__transaction-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[domain.Response](#schemadomain.response)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[domain.ErrorMessage](#schemadomain.errormessage)|

<aside class="success">
This operation does not require authentication
</aside>

## get__transaction_:userId

`GET /transaction/:userId`

*List of Transaction for each User*

List of Transaction for each User

<h3 id="get__transaction_:userid-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|userId|path|string|true|User ID|

> Example responses

> 200 Response

```json
{
  "data": null,
  "message": "string",
  "status_code": 0
}
```

<h3 id="get__transaction_:userid-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[domain.Response](#schemadomain.response)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[domain.ErrorMessage](#schemadomain.errormessage)|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="royalty-api-user">User</h1>

## get__user

`GET /user`

*Get list of users*

Retrieve list of users

> Example responses

> 200 Response

```json
{
  "data": null,
  "message": "string",
  "status_code": 0
}
```

<h3 id="get__user-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[domain.Response](#schemadomain.response)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[domain.ErrorMessage](#schemadomain.errormessage)|

<aside class="success">
This operation does not require authentication
</aside>

## post__user

`POST /user`

*Register User*

Register User to the App

> Body parameter

```json
{
  "address": "string",
  "email": "string",
  "id": 0,
  "name": "string",
  "password": "string",
  "phone": "string",
  "username": "string"
}
```

<h3 id="post__user-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|[domain.UserRequest](#schemadomain.userrequest)|true|User information|

> Example responses

> 200 Response

```json
{
  "data": null,
  "message": "string",
  "status_code": 0
}
```

<h3 id="post__user-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[domain.Response](#schemadomain.response)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[domain.ErrorMessage](#schemadomain.errormessage)|

<aside class="success">
This operation does not require authentication
</aside>

## get__user_:id

`GET /user/:id`

*Get user detail*

Retrieve detail of users

<h3 id="get__user_:id-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|User ID|

> Example responses

> 200 Response

```json
{
  "data": null,
  "message": "string",
  "status_code": 0
}
```

<h3 id="get__user_:id-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[domain.Response](#schemadomain.response)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[domain.ErrorMessage](#schemadomain.errormessage)|

<aside class="success">
This operation does not require authentication
</aside>

## delete__user_:id

`DELETE /user/:id`

*Delete user*

Delete user from database

<h3 id="delete__user_:id-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|User ID|

> Example responses

> 200 Response

```json
{
  "data": null,
  "message": "string",
  "status_code": 0
}
```

<h3 id="delete__user_:id-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[domain.Response](#schemadomain.response)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[domain.ErrorMessage](#schemadomain.errormessage)|

<aside class="success">
This operation does not require authentication
</aside>

## patch__user_:id

`PATCH /user/:id`

*Update User*

Update Exist User in the Application

> Body parameter

```json
{
  "address": "string",
  "email": "string",
  "id": 0,
  "name": "string",
  "password": "string",
  "phone": "string",
  "username": "string"
}
```

<h3 id="patch__user_:id-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|id|path|string|true|User ID|
|body|body|[domain.UserRequest](#schemadomain.userrequest)|true|User information|

> Example responses

> 200 Response

```json
{
  "data": null,
  "message": "string",
  "status_code": 0
}
```

<h3 id="patch__user_:id-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[domain.Response](#schemadomain.response)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[domain.ErrorMessage](#schemadomain.errormessage)|

<aside class="success">
This operation does not require authentication
</aside>

<h1 id="royalty-api-voucher">Voucher</h1>

## get__voucher_:userId

`GET /voucher/:userId`

*Get list of vouchers*

Retrieve list of vouchers based on User ID

<h3 id="get__voucher_:userid-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|userId|path|string|true|User ID|

> Example responses

> 200 Response

```json
{
  "data": null,
  "message": "string",
  "status_code": 0
}
```

<h3 id="get__voucher_:userid-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[domain.Response](#schemadomain.response)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[domain.ErrorMessage](#schemadomain.errormessage)|

<aside class="success">
This operation does not require authentication
</aside>

## get__voucher_generate_:invoice

`GET /voucher/generate/:invoice`

*Generate Voucher for User*

Generate Voucher for User based on Invoice

<h3 id="get__voucher_generate_:invoice-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|invoice|path|string|true|Invoice Number|

> Example responses

> 200 Response

```json
{
  "data": null,
  "message": "string",
  "status_code": 0
}
```

<h3 id="get__voucher_generate_:invoice-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[domain.Response](#schemadomain.response)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[domain.ErrorMessage](#schemadomain.errormessage)|

<aside class="success">
This operation does not require authentication
</aside>

## get__voucher_redeem_:code

`GET /voucher/redeem/:code`

*Redeem Voucher for User*

Redeem Voucher for User

<h3 id="get__voucher_redeem_:code-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|code|path|string|true|Voucher Code|

> Example responses

> 200 Response

```json
{
  "data": null,
  "message": "string",
  "status_code": 0
}
```

<h3 id="get__voucher_redeem_:code-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|OK|[domain.Response](#schemadomain.response)|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|Bad Request|[domain.ErrorMessage](#schemadomain.errormessage)|

<aside class="success">
This operation does not require authentication
</aside>

# Schemas

<h2 id="tocS_domain.ErrorMessage">domain.ErrorMessage</h2>
<!-- backwards compatibility -->
<a id="schemadomain.errormessage"></a>
<a id="schema_domain.ErrorMessage"></a>
<a id="tocSdomain.errormessage"></a>
<a id="tocsdomain.errormessage"></a>

```json
{
  "error": true,
  "message": null,
  "status_code": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|error|boolean|false|none|none|
|message|any|false|none|none|
|status_code|integer|false|none|none|

<h2 id="tocS_domain.Response">domain.Response</h2>
<!-- backwards compatibility -->
<a id="schemadomain.response"></a>
<a id="schema_domain.Response"></a>
<a id="tocSdomain.response"></a>
<a id="tocsdomain.response"></a>

```json
{
  "data": null,
  "message": "string",
  "status_code": 0
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|data|any|false|none|none|
|message|string|false|none|none|
|status_code|integer|false|none|none|

<h2 id="tocS_domain.TransactionRequest">domain.TransactionRequest</h2>
<!-- backwards compatibility -->
<a id="schemadomain.transactionrequest"></a>
<a id="schema_domain.TransactionRequest"></a>
<a id="tocSdomain.transactionrequest"></a>
<a id="tocsdomain.transactionrequest"></a>

```json
{
  "id": 0,
  "invoice_number": "string",
  "tenant_id": 0,
  "total_amount": 0,
  "transaction_date": "string",
  "user_id": 0,
  "voucherGenerated": true
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|id|integer|false|none|none|
|invoice_number|string|false|none|none|
|tenant_id|integer|false|none|none|
|total_amount|integer|false|none|none|
|transaction_date|string|false|none|none|
|user_id|integer|false|none|none|
|voucherGenerated|boolean|false|none|none|

<h2 id="tocS_domain.UserRequest">domain.UserRequest</h2>
<!-- backwards compatibility -->
<a id="schemadomain.userrequest"></a>
<a id="schema_domain.UserRequest"></a>
<a id="tocSdomain.userrequest"></a>
<a id="tocsdomain.userrequest"></a>

```json
{
  "address": "string",
  "email": "string",
  "id": 0,
  "name": "string",
  "password": "string",
  "phone": "string",
  "username": "string"
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|address|string|false|none|none|
|email|string|false|none|none|
|id|integer|false|none|none|
|name|string|false|none|none|
|password|string|false|none|none|
|phone|string|false|none|none|
|username|string|false|none|none|

