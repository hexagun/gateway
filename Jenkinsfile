def dev_repository_tag=""

pipeline {
    agent any

    environment {
        // You must set the following environment variables
        // ORGANIZATION_NAME
        YOUR_DOCKERHUB_USERNAME = "morettimathieu"
        
        SERVICE_NAME = "hexagun-gateway"
        DOCKER_REPOSITORY = "${YOUR_DOCKERHUB_USERNAME}/${SERVICE_NAME}"
        PROD_REPOSITORY_TAG = "${DOCKER_REPOSITORY}:${BUILD_ID}-prod"
    }
    stages {
        
        stage('Build') {
            agent {
                kubernetes {
                  yaml '''
                    apiVersion: v1
                    kind: Pod
                    metadata:
                        labels:
                        some-label: some-label-value
                    spec:
                        containers:
                        - name: go-builder
                          image: golang:1.22
                          command:
                          - cat
                          tty: true                       
                    '''
                }
            }
            steps {   
                container('go-builder') {

                    // Output will be something like "go version go1.22 darwin/arm64"
                    sh 'ls -la'
                    sh 'go version'
                    sh 'go env -w GOFLAGS="-buildvcs=false"'
                    sh 'go mod download'
                    sh 'go build -o gateway'
                }                
            }
        }    
    }
}
