This folder contains HTTP request files for testing the Cloud Run Zipcode Weather Service.

Use a REST client (e.g., VS Code REST Client extension) or `curl` to run the `.http` files against a locally running server at `http://localhost:8080`.

Examples:

- Successful request: `GET /?zipcode=22021001`
- Invalid zipcode format: `GET /?zipcode=22021-001`
- Empty zipcode: `GET /?zipcode=`
- Not found zipcode: `GET /?zipcode=00000000`

You can run individual requests directly from the `.http` file or copy commands into a terminal.