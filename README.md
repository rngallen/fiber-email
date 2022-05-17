# FIBER MAIL
Send Emails (news letters) using fiber


-   [x] Send Emails [API](http://localhost:3003/api/v1/emails)
-   [x] Configuration sample file [Config](sample.yaml)


Sample Schema to send emails 
```
[
  {
    "to": {
      "name": "Sample User",
      "address": "sample@gmail.com"
    },
    "subject": "Hello",
    "body": "Hello we're having new stock"
  },
  {
    "to": {
      "name": "Test User",
      "address": "testuser@gmail.com"
    },
    "subject": "Hello",
    "body": "<p>Hello  <b>Test User</b> we're having new stock</p>"
  },
  {
    "to": {
      "name": "Best User",
      "address": "bestuser@gmail.com"
    },
    "subject": "Hello",
    "body": "<p>Hello we're having new stock</p>"
  }
]

```