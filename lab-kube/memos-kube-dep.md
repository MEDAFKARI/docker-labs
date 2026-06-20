# Progress

made a deployment for image that contains 2 replicas

I tested to delete a replicas and see if kube gonna create another one 

Made a service to expose the port to use it from my browser

Used forward port command but it wasnt a good practice instead 

i changed the settings inside the yaml file for the service 

gave it a type nodePort 

and gave a port so i can access it 

after that used hte commande minikube service memos-service 

to check the running service and open the address of it in the web 


kubectl cluster-info to get where its running either in localhost or in 192.169.X.X
