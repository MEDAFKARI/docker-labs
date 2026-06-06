# Docker deep dive book notes

First docker was a monolith

All the components in one 

docker daomen was doing the whole job building images, executing containers was doing the high level and low level runtimes 

But it was refactored to delegate these processes to 2 new component : Containerd and runc


## Whats Containerd and runc 

containerd = container management platform

containerd responsible for pulling images managing lifecycle managing network plugins


runc = container execution tool

responsible for creating namespaces cgroups and start processes

# Whats the influence of OCI On docker

OCI made sure that all the images and containers build and run in the same way 

They put a template to respect while building the image and running the container

using the image-scpec and run-spec 

image spec = whats inside the box (what the image should have as layers)

run-spec = How to run the box (Defining the config.json on how to run the container by runc*)


