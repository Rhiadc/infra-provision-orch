## Infra provision orchestrator
![arch](https://i.ibb.co/K7MJ1Xh/fdfg.png)

The purpose of this project is to offer an easy way for developers to create, build and deploy their microservices. Choosing a type os service, you create a repo based on that main service (ex: choosing the "Go microservice" as service type, a new git repository is generated with all the microservice template code. Alongside with the microservice repo, a repo with helm chart values is created as well). <br>
This project communicates with: [ms-base-service](https://github.com/Rhiadc/ms-base-go), [infra-ms-base](https://github.com/Rhiadc/infra-ms-base) and [ms-charts](https://github.com/Rhiadc/ms-charts)
### How does it work?
The architecture behaves as the following:
- A user sends a request to the POST endpoint ```/create-service``` with the body: <br>
```
{
    "kind": "ms",
    "name": "ms-my-microservice",
    "description": "some cool description"
}
```
- The orchestor validates the request and the service kind
- After that, it creates a repository based on that template (until now, the only service template is [ms-base-service](https://github.com/Rhiadc/ms-base-go). Which is a template based on a Golang microservice0
- It also generates a infra repository (with Helm chart values) for the ms repository. The infra repo is based on the following template [infra-ms-base](https://github.com/Rhiadc/infra-ms-base)
- The orchestrator also creates a reference for the ms in the argo-gitdir-services repo, for creating the new app reference on the K8S cluster with ArgoCD
- After generating both repositories, the user can tailor and change the config and the code inside the generated code. The user can also change the values on the infra repository, based on his needs.
- The helm chart referenced on that repository is stored on that repo (that acts like a chart-museum) [ms-charts](https://github.com/Rhiadc/ms-charts)
- In the ms repo, after merging into main, a new github actions CI is trigger, building a new image version, and bumping that that on the infra repository.
- After that, ArgoCD will sync and push the changes to the respective namespace in the K8S cluster