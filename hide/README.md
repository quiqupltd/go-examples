# Hide IDs

[Blog Post](https://medium.com/emvi/golang-transforming-ids-to-a-userfriendly-representation-in-web-applications-85bf2f7d71c5)

Most Golang web applications use persitance in some way or another. Usually, the connection between your application and the persistent layer is a technical identification value (ID), a number in most cases. IDs are useful to identify, connect and distinguish data records, often we don't want to expose sequential numberic ids, hide allows it to fix that during marshal / unmarshal

Although this method works, the code of the hide library can serve as an example of how to implement this pattern with own logic.
