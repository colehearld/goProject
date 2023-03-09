# goProject

## This is a project I am currently building in order to expand my knowledege in a few different areas. The project utilizes Go, shell scripting, and docker.

### The project is a simple container orchestration system that allows a user to work efficiently with docker container by allowing them to store the containers into "packets".

### In order to run the program in its current state, docker and go (1.16.0 or later) need to be installed. In the "goProject" repository, run "go run .", and to run either of the current shell scripts, run ./runPacket.sh or ./stopPacket.sh

### Current Progress:
Currently, I am working on running the containers from the packet until the user stops it, it runs for about a second as of now. Once this is figured out, I will write proper commands for running the packets and stopping them.

### In The Future:
Latar, I plan to properly store the container information. As of now, container ID's are stored in a txt file for convenience of building more urgent properties of the program.
