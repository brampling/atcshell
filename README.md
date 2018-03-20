# atcshell

Go program that provides a CLI shell for manipulating ActiveTrust Cloud configuration

usage: ./atcshell

The ATCKEY environment variable must be set to your ATC API key.

The ATCENV environment variable can be set to preprod to access the test environment. The production environment is used by default.

Docker:

Create a Docker image of this program by running this from within this directory:

docker build -t yourname/atcshell .

You can then run the image:

docker run -e ATCKEY=key -e ATCENV=<prod | preprod> -it yourname/atcshell
