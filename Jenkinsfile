pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Build') {
            steps {
                sh "docker build -t hello-world-app:${BUILD_NUMBER} ."
            }
        }

        stage('Deploy') {
            steps {
                sh "docker ps -q --filter publish=3001 | xargs -r docker stop"
                sh "docker ps -a -q --filter publish=3001 | xargs -r docker rm"
                sh "docker run -d -p 3001:3001 hello-world-app:${BUILD_NUMBER}"
            }
        }

        stage('Test') {
            steps {
                sh "sleep 5"
                sh "curl http://localhost:3001"
            }
        }
    }

    post {
        always {
            sh "docker rmi hello-world-app:${BUILD_NUMBER} || true"
        }
        success {
            echo 'Build completed successfully!'
        }
        failure {
            echo 'Build failed!'
        }
    }
}
