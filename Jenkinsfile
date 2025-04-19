pipeline {
    agent any

    environment {
        IMAGE_NAME = 'my-golang-app'
        CONTAINER_NAME = 'golang-container'
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    docker.build("${IMAGE_NAME}")
                }
            }
        }

        stage('Run Container') {
            steps {
                script {
                    sh "docker run -d --name ${CONTAINER_NAME} ${IMAGE_NAME} tail -f /dev/null"
                }
            }
        }

        stage('Build Go Project') {
            steps {
                script {
                    sh "docker exec ${CONTAINER_NAME} go build ./..."
                }
            }
        }

        stage('Run Tests') {
            steps {
                script {
                    sh "docker exec ${CONTAINER_NAME} go test ./..."
                }
            }
        }

        stage('Cleanup') {
            steps {
                script {
                    sh "docker stop ${CONTAINER_NAME}"
                    sh "docker rm ${CONTAINER_NAME}"
                }
            }
        }
    }
}
