Download SSM Parameters in Docker with minimal footprint
========================================================
This demonstrates a simple Go application that downloads parameters from AWS SSM
and instead of installing entire AWS CLI, it uses 9MB of space in Alpine.

Put `test.conf` into Parameter Store for testing.

Refer to this blog post: https://pabis.eu/blog/2023-09-11-Slim-SSM-Parameter-Store-Go.html