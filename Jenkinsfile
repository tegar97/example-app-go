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
                sh "docker run -d -p 3001:3001 hello-world-app:${BUILD_NUMBER}"
            }
        }

          stage('Test') {
            steps {
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
